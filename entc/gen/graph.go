// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package gen is the interface for generating loaded schemas into a Go package.
package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"text/template/parse"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"

	"golang.org/x/tools/imports"
)

type (
	// The Config holds the global codegen configuration to be
	// shared between all generated nodes.
	Config struct {
		// Schema holds the Go package path for the user ent/schema.
		// For example, "<project>/ent/schema".
		Schema string

		// Target defines the filepath for the target directory that
		// holds the generated code. For example, "./project/ent".
		//
		// By default, 'ent generate ./ent/schema' uses './ent' as a
		// target directory.
		Target string

		// Package defines the Go package path of the target directory
		// mentioned above. For example, "github.com/org/project/ent".
		//
		// By default, for schema package named "<project>/ent/schema",
		// 'ent generate' uses "<project>/ent" as a default package.
		Package string

		// Header allows users to provide an optional header signature for
		// the generated files. It defaults to the standard 'go generate'
		// format: '// Code generated by ent, DO NOT EDIT.'.
		Header string

		// Storage configuration for the codegen. Defaults to sql.
		Storage *Storage

		// IDType specifies the type of the id field in the codegen.
		// The supported types are string and int, which also the default.
		IDType *field.TypeInfo

		// Templates specifies a list of alternative templates to execute or
		// to override the default. If nil, the default template is used.
		//
		// Note that, additional templates are executed on the Graph object and
		// the execution output is stored in a file derived by the template name.
		Templates []*Template

		// Features defines a list of additional features to add to the codegen phase.
		// For example, the PrivacyFeature.
		Features []Feature

		// Hooks holds an optional list of Hooks to apply on the graph before/after the code-generation.
		Hooks []Hook

		// Annotations that are injected to the Config object can be accessed
		// globally in all templates. In order to access an annotation from a
		// graph template, do the following:
		//
		//	{{- with $.Annotations.GQL }}
		//		{{/* Annotation usage goes here. */}}
		//	{{- end }}
		//
		// For type templates, we access the Config field to access the global
		// annotations, and not the type-specific annotation.
		//
		//	{{- with $.Config.Annotations.GQL }}
		//		{{/* Annotation usage goes here. */}}
		//	{{- end }}
		//
		// Note that the mapping is from the annotation-name (e.g. "GQL") to a JSON decoded object.
		Annotations Annotations
	}

	// Graph holds the nodes/entities of the loaded graph schema. Note that, it doesn't
	// hold the edges of the graph. Instead, each Type holds the edges for other Types.
	Graph struct {
		*Config
		// Nodes are list of Go types that mapped to the types in the loaded schema.
		Nodes []*Type
		nodes map[string]*Type
		// Schemas holds the raw interfaces for the loaded schemas.
		Schemas []*load.Schema
	}

	// Generator is the interface that wraps the Generate method.
	Generator interface {
		// Generate generates the ent artifacts for the given graph.
		Generate(*Graph) error
	}

	// The GenerateFunc type is an adapter to allow the use of ordinary
	// function as Generator. If f is a function with the appropriate signature,
	// GenerateFunc(f) is a Generator that calls f.
	GenerateFunc func(*Graph) error

	// Hook defines the "generate middleware". A function that gets a Generator
	// and returns a Generator. For example:
	//
	//	hook := func(next gen.Generator) gen.Generator {
	//		return gen.GenerateFunc(func(g *Graph) error {
	//			fmt.Println("Graph:", g)
	//			return next.Generate(g)
	//		})
	//	}
	//
	Hook func(Generator) Generator

	// Annotations defines code generation annotations to be passed to the templates.
	// It can be defined on most elements in the schema (node, field, edge), or globally
	// on the Config object.
	// The mapping is from the annotation name (e.g. "EntGQL") to the annotation itself.
	// Note that, annotations that are defined in the schema must be JSON encoded/decoded.
	Annotations map[string]any
)

// Generate calls f(g).
func (f GenerateFunc) Generate(g *Graph) error {
	return f(g)
}

// NewGraph creates a new Graph for the code generation from the given schema definitions.
// It fails if one of the schemas is invalid.
func NewGraph(c *Config, schemas ...*load.Schema) (g *Graph, err error) {
	defer catch(&err)
	g = &Graph{Config: c, Nodes: make([]*Type, 0, len(schemas)), Schemas: schemas}
	for i := range schemas {
		g.addNode(schemas[i])
	}
	for i := range schemas {
		g.addEdges(schemas[i])
	}
	for _, t := range g.Nodes {
		check(g.resolve(t), "resolve %q relations", t.Name)
	}
	for _, t := range g.Nodes {
		check(t.setupFKs(), "set %q foreign-keys", t.Name)
	}
	check(g.edgeSchemas(), "resolving edges")
	for i := range schemas {
		g.addIndexes(schemas[i])
	}
	aliases(g)
	g.defaults()
	return
}

// defaultIDType holds the default value for IDType.
var defaultIDType = &field.TypeInfo{Type: field.TypeInt}

// defaults sets the default value of the IDType. The IDType field is used
// by multiple templates. If the IDType was not provided, it falls back to
// int, or the one used in the schema (if all schemas share the same IDType).
func (g *Graph) defaults() {
	if g.IDType != nil {
		return
	}
	g.IDType = defaultIDType
	if len(g.Nodes) == 0 {
		return
	}
	idTypes := make([]*field.TypeInfo, 0, len(g.Nodes))
	for _, n := range g.Nodes {
		if n.HasOneFieldID() {
			idTypes = append(idTypes, n.ID.Type)
		}
	}
	// Check that all nodes have the same type for the ID field.
	for i := 0; i < len(idTypes)-1; i++ {
		if idTypes[i].Type != idTypes[i+1].Type {
			return
		}
	}
	g.IDType = idTypes[0]
}

// Gen generates the artifacts for the graph.
func (g *Graph) Gen() error {
	var gen Generator = GenerateFunc(generate)
	for i := len(g.Hooks) - 1; i >= 0; i-- {
		gen = g.Hooks[i](gen)
	}
	return gen.Generate(g)
}

// generate is the default Generator implementation.
func generate(g *Graph) error {
	var (
		assets   assets
		external []GraphTemplate
	)
	templates, external = g.templates()
	for _, n := range g.Nodes {
		assets.addDir(filepath.Join(g.Config.Target, n.PackageDir()))
		for _, tmpl := range Templates {
			b := bytes.NewBuffer(nil)
			if err := templates.ExecuteTemplate(b, tmpl.Name, n); err != nil {
				return fmt.Errorf("execute template %q: %w", tmpl.Name, err)
			}
			assets.add(filepath.Join(g.Config.Target, tmpl.Format(n)), b.Bytes())
		}
	}
	for _, tmpl := range append(GraphTemplates, external...) {
		if tmpl.Skip != nil && tmpl.Skip(g) {
			continue
		}
		if dir := filepath.Dir(tmpl.Format); dir != "." {
			assets.addDir(filepath.Join(g.Config.Target, dir))
		}
		b := bytes.NewBuffer(nil)
		if err := templates.ExecuteTemplate(b, tmpl.Name, g); err != nil {
			return fmt.Errorf("execute template %q: %w", tmpl.Name, err)
		}
		assets.add(filepath.Join(g.Config.Target, tmpl.Format), b.Bytes())
	}
	for _, f := range AllFeatures {
		if f.cleanup == nil || g.featureEnabled(f) {
			continue
		}
		if err := f.cleanup(g.Config); err != nil {
			return fmt.Errorf("cleanup %q feature assets: %w", f.Name, err)
		}
	}
	// Write and format assets only if template execution
	// finished successfully.
	if err := assets.write(); err != nil {
		return err
	}
	// cleanup assets that are not needed anymore.
	cleanOldNodes(assets, g.Config.Target)
	// We can't run "imports" on files when the state is not completed.
	// Because, "goimports" will drop undefined package. Therefore, it
	// is suspended to the end of the writing.
	return assets.format()
}

// addNode creates a new Type/Node/Ent to the graph.
func (g *Graph) addNode(schema *load.Schema) {
	t, err := NewType(g.Config, schema)
	check(err, "create type %s", schema.Name)
	g.Nodes = append(g.Nodes, t)
}

// addIndexes adds the indexes for the schema type.
func (g *Graph) addIndexes(schema *load.Schema) {
	typ, _ := g.typ(schema.Name)
	for _, idx := range schema.Indexes {
		check(typ.AddIndex(idx), "invalid index for schema %q", schema.Name)
	}
}

// addEdges adds the node edges to the graph.
func (g *Graph) addEdges(schema *load.Schema) {
	t, _ := g.typ(schema.Name)
	seen := make(map[string]struct{}, len(schema.Edges))
	for _, e := range schema.Edges {
		typ, ok := g.typ(e.Type)
		expect(ok, "type %q does not exist for edge", e.Type)
		_, ok = t.fields[e.Name]
		expect(!ok, "%s schema cannot contain field and edge with the same name %q", schema.Name, e.Name)
		_, ok = seen[e.Name]
		expect(!ok, "%s schema contains multiple %q edges", schema.Name, e.Name)
		seen[e.Name] = struct{}{}
		switch {
		// Assoc only.
		case !e.Inverse:
			t.Edges = append(t.Edges, &Edge{
				def:         e,
				Type:        typ,
				Name:        e.Name,
				Owner:       t,
				Unique:      e.Unique,
				Optional:    !e.Required,
				StructTag:   structTag(e.Name, e.Tag),
				Annotations: e.Annotations,
			})
		// Inverse only.
		case e.Inverse && e.Ref == nil:
			expect(e.RefName != "", "back-reference edge %s.%s is missing the Ref attribute", t.Name, e.Name)
			t.Edges = append(t.Edges, &Edge{
				def:         e,
				Type:        typ,
				Name:        e.Name,
				Owner:       typ,
				Inverse:     e.RefName,
				Unique:      e.Unique,
				Optional:    !e.Required,
				StructTag:   structTag(e.Name, e.Tag),
				Annotations: e.Annotations,
			})
		// Inverse and assoc.
		case e.Inverse:
			ref := e.Ref
			expect(e.RefName == "", "reference name is derived from the assoc name: %s.%s <-> %s.%s", t.Name, ref.Name, t.Name, e.Name)
			expect(ref.Type == t.Name, "assoc-inverse edge allowed only as o2o relation of the same type")
			from := &Edge{
				def:         e,
				Type:        typ,
				Name:        e.Name,
				Owner:       t,
				Inverse:     ref.Name,
				Unique:      e.Unique,
				Optional:    !e.Required,
				StructTag:   structTag(e.Name, e.Tag),
				Annotations: e.Annotations,
			}
			to := &Edge{
				def:         ref,
				Ref:         from,
				Type:        typ,
				Owner:       t,
				Name:        ref.Name,
				Unique:      ref.Unique,
				Optional:    !ref.Required,
				StructTag:   structTag(ref.Name, ref.Tag),
				Annotations: ref.Annotations,
			}
			from.Ref = to
			t.Edges = append(t.Edges, from, to)
		default:
			panic(graphError{"edge must be either an assoc or inverse edge"})
		}
	}
}

// resolve resolves the type reference and relation of edges.
// It fails if one of the references is missing or invalid.
//
// relation definitions between A and B, where A is the owner of
// the edge and B uses this edge as a back-reference:
//
// 	O2O
// 	 - A have a unique edge (E) to B, and B have a back-reference unique edge (E') for E.
// 	 - A have a unique edge (E) to A.
//
// 	O2M (The "Many" side, keeps a reference to the "One" side).
// 	 - A have an edge (E) to B (not unique), and B doesn't have a back-reference edge for E.
// 	 - A have an edge (E) to B (not unique), and B have a back-reference unique edge (E') for E.
//
// 	M2O (The "Many" side, holds the reference to the "One" side).
// 	 - A have a unique edge (E) to B, and B doesn't have a back-reference edge for E.
// 	 - A have a unique edge (E) to B, and B have a back-reference non-unique edge (E') for E.
//
// 	M2M
// 	 - A have an edge (E) to B (not unique), and B have a back-reference non-unique edge (E') for E.
// 	 - A have an edge (E) to A (not unique).
//
func (g *Graph) resolve(t *Type) error {
	for _, e := range t.Edges {
		switch {
		case e.IsInverse():
			ref, ok := e.Type.HasAssoc(e.Inverse)
			if !ok {
				return fmt.Errorf("edge %q is missing for inverse edge: %s.%s(%s)", e.Inverse, t.Name, e.Name, e.Type.Name)
			}
			if !e.Optional && !ref.Optional {
				return fmt.Errorf("edges cannot be required in both directions: %s.%s <-> %s.%s", t.Name, e.Name, e.Type.Name, ref.Name)
			}
			if ref.Type != t {
				return fmt.Errorf("mismatch type for back-ref %q of %s.%s <-> %s.%s", e.Inverse, t.Name, e.Name, e.Type.Name, ref.Name)
			}
			e.Ref, ref.Ref = ref, e
			table := t.Table()
			// Name the foreign-key column in a format that wouldn't change even if an inverse
			// edge is dropped (or added). The format is: "<Edge-Owner>_<Edge-Name>".
			column := fmt.Sprintf("%s_%s", e.Type.Label(), snake(ref.Name))
			switch a, b := ref.Unique, e.Unique; {
			// If the relation column is in the inverse side/table. The rule is simple, if assoc is O2M,
			// then inverse is M2O and the relation is in its table.
			case a && b:
				e.Rel.Type, ref.Rel.Type = O2O, O2O
			case !a && b:
				e.Rel.Type, ref.Rel.Type = M2O, O2M

			// If the relation column is in the assoc side.
			case a && !b:
				e.Rel.Type, ref.Rel.Type = O2M, M2O
				table = e.Type.Table()

			case !a && !b:
				e.Rel.Type, ref.Rel.Type = M2M, M2M
				table = e.Type.Label() + "_" + ref.Name
				c1, c2 := ref.Owner.Label()+"_id", ref.Type.Label()+"_id"
				// If the relation is from the same type: User has Friends ([]User),
				// we give the second column a different name (the relation name).
				if c1 == c2 {
					c2 = rules.Singularize(e.Name) + "_id"
				}
				e.Rel.Columns = []string{c1, c2}
				ref.Rel.Columns = []string{c1, c2}
			}
			e.Rel.Table, ref.Rel.Table = table, table
			if !e.M2M() {
				e.Rel.Columns = []string{column}
				ref.Rel.Columns = []string{column}
			}
		// Assoc with uninitialized relation.
		case !e.IsInverse() && e.Rel.Type == Unk:
			switch {
			case !e.Unique && e.Type == t:
				e.Rel.Type = M2M
				e.Bidi = true
				e.Rel.Table = t.Label() + "_" + e.Name
				e.Rel.Columns = []string{e.Owner.Label() + "_id", rules.Singularize(e.Name) + "_id"}
			case e.Unique && e.Type == t:
				e.Rel.Type = O2O
				e.Bidi = true
				e.Rel.Table = t.Table()
			case e.Unique:
				e.Rel.Type = M2O
				e.Rel.Table = t.Table()
			default:
				e.Rel.Type = O2M
				e.Rel.Table = e.Type.Table()
			}
			if !e.M2M() {
				e.Rel.Columns = []string{fmt.Sprintf("%s_%s", t.Label(), snake(e.Name))}
			}
		}
	}
	return nil
}

// edgeSchemas visits all edges in the graph and detects which schemas are used as "edge schemas".
// Note, edge schemas cannot be used by more than one association (edge.To), must define two required
// edges (+ edge-fields) to the types that go through them, and allow adding additional fields with
// optional default values.
func (g *Graph) edgeSchemas() error {
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			if e.def.Through == nil {
				continue
			}
			if !e.M2M() {
				return fmt.Errorf("edge %s.%s Through(%q, %s.Type) is allowed only on M2M edges, but got: %q", n.Name, e.Name, e.def.Through.N, e.def.Through.T, e.Rel.Type)
			}
			typ, ok := g.typ(e.def.Through.T)
			switch {
			case !ok:
				return fmt.Errorf("edge %s.%s defined with Through(%q, %s.Type), but type %[4]s was not found", n.Name, e.Name, e.def.Through.N, e.def.Through.T, e.def.Through.T)
			case typ == n:
				return fmt.Errorf("edge %s.%s defined with Through(%q, %s.Type), but edge cannot go through itself", n.Name, e.Name, e.def.Through.N, e.def.Through.T)
			case e.def.Through.N == "" || n.hasEdge(e.def.Through.N):
				return fmt.Errorf("edge %s.%s defined with Through(%q, %s.Type), but schema %[1]s already has an edge named %[3]s", n.Name, e.Name, e.def.Through.N, e.def.Through.T)
			case e.IsInverse():
				if typ.EdgeSchema.From != nil {
					return fmt.Errorf("type %s is already used as an edge schema by other edge.From: %s.%s", typ.Name, typ.EdgeSchema.From.Name, typ.EdgeSchema.From.Owner.Name)
				}
				e.Through = typ
				typ.EdgeSchema.From = e
				if to, from := typ.EdgeSchema.To, typ.EdgeSchema.From; to != nil && from.Ref != to {
					return fmt.Errorf("mismtached edge.From(%q, %s.Type) and edge.To(%q, %s.Type) for edge schema %s", from.Name, from.Type.Name, to.Name, to.Type.Name, typ.Name)
				}
			default: // Assoc.
				if typ.EdgeSchema.To != nil {
					return fmt.Errorf("type %s is already used as an edge schema by other edge.To: %s.%s", typ.Name, typ.EdgeSchema.From.Name, typ.EdgeSchema.From.Owner.Name)
				}
				e.Through = typ
				typ.EdgeSchema.To = e
				if to, from := typ.EdgeSchema.To, typ.EdgeSchema.From; from != nil && from.Ref != to {
					return fmt.Errorf("mismtached edge.To(%q, %s.Type) and edge.From(%q, %s.Type) for edge schema %s", from.Name, from.Type.Name, to.Name, to.Type.Name, typ.Name)
				}
			}
			e.Rel.Table = typ.Table()
			var ref *Edge
			for i, c := range e.Rel.Columns {
				r, ok := func() (*Edge, bool) {
					for _, fk := range typ.ForeignKeys {
						if fk.Field.Name == c {
							return fk.Edge, true
						}
					}
					return nil, false
				}()
				if !ok {
					return fmt.Errorf("missing edge-field %s.%s for edge schema used by %s.%s in Through(%q, %s.Type)", typ.Name, c, n.Name, e.Name, e.def.Through.N, typ.Name)
				}
				if r.Optional {
					return fmt.Errorf("edge-schema %s is missing a Required() attribute for its reference edge %q", typ.Name, e.Name)
				}
				if !e.IsInverse() && i == 0 || e.IsInverse() && i == 1 {
					ref = r
				}
			}
			// Edges from src/dest table are always O2M. One row to many
			// rows in the join table. Hence, a many-to-many relationship.
			n.Edges = append(n.Edges, &Edge{
				def:       &load.Edge{},
				Name:      e.def.Through.N,
				Type:      typ,
				Inverse:   ref.Name,
				Ref:       ref,
				Owner:     n,
				Optional:  true,
				StructTag: structTag(e.def.Through.N, ""),
				Rel: Relation{
					Type:    O2M,
					fk:      ref.Rel.fk,
					Table:   ref.Rel.Table,
					Columns: ref.Rel.Columns,
				},
			})
			// Edge schema contains a composite primary key, and it was not resolved in previous iterations.
			if ant := fieldAnnotate(typ.Annotations); ant != nil && len(ant.ID) > 0 && len(typ.EdgeSchema.ID) == 0 {
				r1, r2 := e.Rel.Columns[0], e.Rel.Columns[1]
				if len(ant.ID) != 2 || ant.ID[0] != r1 || ant.ID[1] != r2 {
					return fmt.Errorf(`edge schema primary key can only be defined on "id" or (%q, %q) in the same order`, r1, r2)
				}
				typ.ID = nil
				for _, f := range ant.ID {
					typ.EdgeSchema.ID = append(typ.EdgeSchema.ID, typ.fields[f])
				}
			}
			if typ.HasCompositeID() {
				continue
			}
			hasI := func() bool {
				for _, idx := range typ.Indexes {
					if !idx.Unique && len(idx.Columns) != 2 {
						continue
					}
					c1, c2 := idx.Columns[0], idx.Columns[1]
					r1, r2 := e.Rel.Columns[0], e.Rel.Columns[1]
					if c1 == r1 && c2 == r2 || c1 == r2 && c2 == r1 {
						return true
					}
				}
				return false
			}()
			if !hasI {
				if err := typ.AddIndex(&load.Index{Unique: true, Fields: e.Rel.Columns}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Tables returns the schema definitions of SQL tables for the graph.
func (g *Graph) Tables() (all []*schema.Table, err error) {
	tables := make(map[string]*schema.Table)
	for _, n := range g.Nodes {
		table := schema.NewTable(n.Table())
		if n.HasOneFieldID() {
			table.AddPrimary(n.ID.PK())
		}
		table.SetAnnotation(n.EntSQL())
		for _, f := range n.Fields {
			if !f.IsEdgeField() {
				table.AddColumn(f.Column())
			}
		}
		tables[table.Name] = table
		all = append(all, table)
	}
	for _, n := range g.Nodes {
		// Foreign key and its reference, or a join table.
		for _, e := range n.Edges {
			if e.IsInverse() {
				continue
			}
			switch e.Rel.Type {
			case O2O, O2M:
				// The "owner" is the table that owns the relation (we set
				// the foreign-key on) and "ref" is the referenced table.
				owner, ref := tables[e.Rel.Table], tables[n.Table()]
				pk := ref.PrimaryKey[0]
				column := &schema.Column{Name: e.Rel.Column(), Size: pk.Size, Type: pk.Type, Unique: e.Rel.Type == O2O, SchemaType: pk.SchemaType, Nullable: true}
				// If it's not a circular reference (self-referencing table),
				// and the inverse edge is required, make it non-nullable.
				if n != e.Type && e.Ref != nil && !e.Ref.Optional {
					column.Nullable = false
				}
				mayAddColumn(owner, column)
				owner.AddForeignKey(&schema.ForeignKey{
					RefTable:   ref,
					OnDelete:   deleteAction(e, column),
					Columns:    []*schema.Column{column},
					RefColumns: []*schema.Column{ref.PrimaryKey[0]},
					Symbol:     fkSymbol(e, owner, ref),
				})
			case M2O:
				ref, owner := tables[e.Type.Table()], tables[e.Rel.Table]
				pk := ref.PrimaryKey[0]
				column := &schema.Column{Name: e.Rel.Column(), Size: pk.Size, Type: pk.Type, SchemaType: pk.SchemaType, Nullable: true}
				// If it's not a circular reference (self-referencing table),
				// and the edge is non-optional (required), make it non-nullable.
				if n != e.Type && !e.Optional {
					column.Nullable = false
				}
				mayAddColumn(owner, column)
				owner.AddForeignKey(&schema.ForeignKey{
					RefTable:   ref,
					OnDelete:   deleteAction(e, column),
					Columns:    []*schema.Column{column},
					RefColumns: []*schema.Column{ref.PrimaryKey[0]},
					Symbol:     fkSymbol(e, owner, ref),
				})
			case M2M:
				// If there is an edge schema for the association (i.e. edge.Through).
				if e.Through != nil || e.Ref != nil && e.Ref.Through != nil {
					continue
				}
				t1, t2 := tables[n.Table()], tables[e.Type.Table()]
				c1 := &schema.Column{Name: e.Rel.Columns[0], Type: field.TypeInt, SchemaType: n.ID.def.SchemaType}
				if ref := n.ID; ref.UserDefined {
					c1.Type = ref.Type.Type
					c1.Size = ref.size()
				}
				c2 := &schema.Column{Name: e.Rel.Columns[1], Type: field.TypeInt, SchemaType: e.Type.ID.def.SchemaType}
				if ref := e.Type.ID; ref.UserDefined {
					c2.Type = ref.Type.Type
					c2.Size = ref.size()
				}
				s1, s2 := fkSymbols(e, c1, c2)
				all = append(all, &schema.Table{
					Name:       e.Rel.Table,
					Columns:    []*schema.Column{c1, c2},
					PrimaryKey: []*schema.Column{c1, c2},
					ForeignKeys: []*schema.ForeignKey{
						{
							RefTable:   t1,
							OnDelete:   schema.Cascade,
							Columns:    []*schema.Column{c1},
							RefColumns: []*schema.Column{t1.PrimaryKey[0]},
							Symbol:     s1,
						},
						{
							RefTable:   t2,
							OnDelete:   schema.Cascade,
							Columns:    []*schema.Column{c2},
							RefColumns: []*schema.Column{t2.PrimaryKey[0]},
							Symbol:     s2,
						},
					},
				})
			}
		}
		if n.HasCompositeID() {
			if err := addCompositePK(tables[n.Table()], n); err != nil {
				return nil, err
			}
		}
	}
	// Append indexes to tables after all columns were added (including relation columns).
	for _, n := range g.Nodes {
		table := tables[n.Table()]
		for _, idx := range n.Indexes {
			table.AddIndex(idx.Name, idx.Unique, idx.Columns)
			// Set the entsql.IndexAnnotation from the schema if exists.
			index, _ := table.Index(idx.Name)
			index.Annotation = entsqlIndexAnnotate(idx.Annotations)
		}
	}
	return
}

// mayAddColumn adds the given column if it does not already exist in the table.
func mayAddColumn(t *schema.Table, c *schema.Column) {
	if !t.HasColumn(c.Name) {
		t.AddColumn(c)
	}
}

func addCompositePK(t *schema.Table, n *Type) error {
	columns := make([]*schema.Column, 0, len(n.EdgeSchema.ID))
	for _, id := range n.EdgeSchema.ID {
		for _, f := range n.Fields {
			if !f.IsEdgeField() || id != f {
				continue
			}
			c, ok := t.Column(f.StorageKey())
			if !ok {
				return fmt.Errorf("missing column %q for edge field %q.%q", f.StorageKey(), n.Name, f.Name)
			}
			columns = append(columns, c)
		}
	}
	t.PrimaryKey = columns
	return nil
}

// fkSymbol returns the symbol of the foreign-key constraint for edges of type O2M, M2O and O2O.
// It returns the symbol of the storage-key if it was provided, and generate custom one otherwise.
func fkSymbol(e *Edge, ownerT, refT *schema.Table) string {
	if k, _ := e.StorageKey(); k != nil && len(k.Symbols) == 1 {
		return k.Symbols[0]
	}
	return fmt.Sprintf("%s_%s_%s", ownerT.Name, refT.Name, e.Name)
}

// fkSymbols is like fkSymbol but for M2M edges.
func fkSymbols(e *Edge, c1, c2 *schema.Column) (string, string) {
	s1 := fmt.Sprintf("%s_%s", e.Rel.Table, c1.Name)
	s2 := fmt.Sprintf("%s_%s", e.Rel.Table, c2.Name)
	if k, _ := e.StorageKey(); k != nil {
		if len(k.Symbols) > 0 {
			s1 = k.Symbols[0]
		}
		if len(k.Symbols) > 1 {
			s2 = k.Symbols[1]
		}
	}
	return s1, s2
}

// deleteAction returns the referential action for DELETE operations of the given edge.
func deleteAction(e *Edge, c *schema.Column) schema.ReferenceOption {
	action := schema.NoAction
	if c.Nullable {
		action = schema.SetNull
	}
	if ant := e.EntSQL(); ant != nil && ant.OnDelete != "" {
		action = schema.ReferenceOption(ant.OnDelete)
	}
	return action
}

// SupportMigrate reports if the codegen supports schema migration.
func (g *Graph) SupportMigrate() bool {
	return g.Storage.SchemaMode.Support(Migrate)
}

// Snapshot holds the information for storing the schema snapshot.
type Snapshot struct {
	Schema   string
	Package  string
	Schemas  []*load.Schema
	Features []string
}

// SchemaSnapshot returns a JSON string represents the graph schema in loadable format.
func (g *Graph) SchemaSnapshot() (string, error) {
	schemas := make([]*load.Schema, len(g.Nodes))
	for i := range g.Nodes {
		schemas[i] = g.Nodes[i].schema
	}
	snap := Snapshot{
		Schema:  g.Schema,
		Package: g.Package,
		Schemas: schemas,
	}
	for _, feat := range g.Features {
		snap.Features = append(snap.Features, feat.Name)
	}
	out, err := json.Marshal(snap)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (g *Graph) typ(name string) (*Type, bool) {
	if g.nodes == nil {
		g.nodes = make(map[string]*Type, len(g.Nodes))
		for _, n := range g.Nodes {
			g.nodes[n.Name] = n
		}
	}
	n, ok := g.nodes[name]
	return n, ok
}

// templates returns the Template to execute on the Graph,
// and a list of optional external templates if provided.
func (g *Graph) templates() (*Template, []GraphTemplate) {
	initTemplates()
	var (
		roots    = make(map[string]struct{})
		helpers  = make(map[string]struct{})
		external = make([]GraphTemplate, 0, len(g.Templates))
	)
	for _, rootT := range g.Templates {
		templates.Funcs(rootT.FuncMap)
		for _, tmpl := range rootT.Templates() {
			if parse.IsEmptyTree(tmpl.Root) {
				continue
			}
			name := tmpl.Name()
			switch {
			// Helper templates can be either global (prefixed with "helper/"),
			// or local, where their names follow the format: "<root-tmpl>/helper/.+").
			case strings.HasPrefix(name, "helper/"):
			case strings.Contains(name, "/helper/"):
				helpers[name] = struct{}{}
			case templates.Lookup(name) == nil && !extendExisting(name):
				// If the template does not override or extend one of
				// the builtin templates, generate it in a new file.
				external = append(external, GraphTemplate{
					Name:   name,
					Format: snake(name) + ".go",
					Skip:   rootT.condition,
				})
				roots[name] = struct{}{}
			}
			templates = MustParse(templates.AddParseTree(name, tmpl.Tree))
		}
	}
	for name := range helpers {
		root := name[:strings.Index(name, "/helper/")]
		// If the name is prefixed with a name of a root
		// template, we treat it as a local helper template.
		if _, ok := roots[root]; ok {
			continue
		}
		external = append(external, GraphTemplate{
			Name:   name,
			Format: snake(name) + ".go",
		})
	}
	for _, f := range g.Features {
		external = append(external, f.GraphTemplates...)
	}
	return templates, external
}

// ModuleInfo returns the entgo.io/ent version.
func (Config) ModuleInfo() (m debug.Module) {
	const pkg = "entgo.io/ent"
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	// Was running as a CLI (ent/cmd/ent).
	if info.Main.Path == pkg {
		return info.Main
	}
	// Or, as a main package (ent/entc).
	for _, dep := range info.Deps {
		if dep.Path == pkg {
			return *dep
		}
	}
	return
}

// FeatureEnabled reports if the given feature name is enabled.
// It's exported to be used by the template engine as follows:
//
//	{{ with $.FeatureEnabled "privacy" }}
//		...
//	{{ end }}
//
func (c Config) FeatureEnabled(name string) (bool, error) {
	for _, f := range AllFeatures {
		if name == f.Name {
			return c.featureEnabled(f), nil
		}
	}
	return false, fmt.Errorf("unexpected feature name %q", name)
}

// featureEnabled reports if the given feature-flag is enabled.
func (c Config) featureEnabled(f Feature) bool {
	for i := range c.Features {
		if f.Name == c.Features[i].Name {
			return true
		}
	}
	return false
}

// PrepareEnv makes sure the generated directory (environment)
// is suitable for loading the `ent` package (avoid cyclic imports).
func PrepareEnv(c *Config) (undo func() error, err error) {
	var (
		nop  = func() error { return nil }
		path = filepath.Join(c.Target, "runtime.go")
	)
	out, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nop, nil
		}
		return nil, err
	}
	fi, err := parser.ParseFile(token.NewFileSet(), path, out, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}
	// Targeted package doesn't import the schema.
	if len(fi.Imports) == 0 {
		return nop, nil
	}
	if err := os.WriteFile(path, append([]byte("// +build tools\n"), out...), 0644); err != nil {
		return nil, err
	}
	return func() error { return os.WriteFile(path, out, 0644) }, nil
}

// cleanOldNodes removes all files that were generated
// for nodes that were removed from the schema.
func cleanOldNodes(assets assets, target string) {
	d, err := os.ReadDir(target)
	if err != nil {
		return
	}
	// Find deleted nodes by selecting one generated
	// file from standard templates (<T>_query.go).
	var deleted []*Type
	for _, f := range d {
		if !strings.HasSuffix(f.Name(), "_query.go") {
			continue
		}
		typ := &Type{Name: strings.TrimSuffix(f.Name(), "_query.go")}
		path := filepath.Join(target, typ.PackageDir())
		if _, ok := assets.dirs[path]; ok {
			continue
		}
		// If it is a node, it must have a model file and a dir (e.g. ent/t.go, ent/t).
		_, err1 := os.Stat(path + ".go")
		f2, err2 := os.Stat(path)
		if err1 == nil && err2 == nil && f2.IsDir() {
			deleted = append(deleted, typ)
		}
	}
	for _, typ := range deleted {
		for _, t := range Templates {
			err := os.Remove(filepath.Join(target, t.Format(typ)))
			if err != nil && !os.IsNotExist(err) {
				log.Printf("remove old file %s: %s\n", filepath.Join(target, t.Format(typ)), err)
			}
		}
		err := os.Remove(filepath.Join(target, typ.PackageDir()))
		if err != nil && !os.IsNotExist(err) {
			log.Printf("remove old dir %s: %s\n", filepath.Join(target, typ.PackageDir()), err)
		}
	}
}

type assets struct {
	dirs  map[string]struct{}
	files map[string][]byte
}

func (a *assets) add(path string, content []byte) {
	if a.files == nil {
		a.files = make(map[string][]byte)
	}
	a.files[path] = content
}

func (a *assets) addDir(path string) {
	if a.dirs == nil {
		a.dirs = make(map[string]struct{})
	}
	a.dirs[path] = struct{}{}
}

// write files and dirs in the assets.
func (a assets) write() error {
	for dir := range a.dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("create dir %q: %w", dir, err)
		}
	}
	for path, content := range a.files {
		if err := os.WriteFile(path, content, 0644); err != nil {
			return fmt.Errorf("write file %q: %w", path, err)
		}
	}
	return nil
}

// format runs "goimports" on all assets.
func (a assets) format() error {
	for path, content := range a.files {
		src, err := imports.Process(path, content, nil)
		if err != nil {
			return fmt.Errorf("format file %s: %w", path, err)
		}
		if err := os.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %w", path, err)
		}
	}
	return nil
}

// expect panics if the condition is false.
func expect(cond bool, msg string, args ...any) {
	if !cond {
		panic(graphError{fmt.Sprintf(msg, args...)})
	}
}

// check panics if the error is not nil.
func check(err error, msg string, args ...any) {
	if err != nil {
		args = append(args, err)
		panic(graphError{fmt.Sprintf(msg+": %s", args...)})
	}
}

type graphError struct {
	msg string
}

func (p graphError) Error() string { return fmt.Sprintf("entc/gen: %s", p.msg) }

func catch(err *error) {
	if e := recover(); e != nil {
		gerr, ok := e.(graphError)
		if !ok {
			panic(e)
		}
		*err = gerr
	}
}

func extendExisting(name string) bool {
	if match(partialPatterns[:], name) {
		return true
	}
	for _, t := range Templates {
		if match(t.ExtendPatterns, name) {
			return true
		}
	}
	for _, t := range GraphTemplates {
		if match(t.ExtendPatterns, name) {
			return true
		}
	}
	return false
}
