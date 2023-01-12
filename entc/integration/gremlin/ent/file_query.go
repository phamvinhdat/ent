// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"
	fmt "fmt"
	math "math"

	gremlin "entgo.io/ent/dialect/gremlin"
	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	g "entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/file"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// FileQuery is the builder for querying File entities.
type FileQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.File
	withOwner  *UserQuery
	withType   *FileTypeQuery
	withField  *FieldTypeQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the FileQuery builder.
func (fq *FileQuery) Where(ps ...predicate.File) *FileQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FileQuery) Limit(limit int) *FileQuery {
	fq.limit = &limit
	return fq
}

// Offset to start from.
func (fq *FileQuery) Offset(offset int) *FileQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FileQuery) Unique(unique bool) *FileQuery {
	fq.unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FileQuery) Order(o ...OrderFunc) *FileQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryOwner chains the current query on the "owner" edge.
func (fq *FileQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := fq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.FilesLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryType chains the current query on the "type" edge.
func (fq *FileQuery) QueryType() *FileTypeQuery {
	query := (&FileTypeClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := fq.gremlinQuery(ctx)
		fromU = gremlin.InE(filetype.FilesLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryField chains the current query on the "field" edge.
func (fq *FileQuery) QueryField() *FieldTypeQuery {
	query := (&FieldTypeClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := fq.gremlinQuery(ctx)
		fromU = gremlin.OutE(file.FieldLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first File entity from the query.
// Returns a *NotFoundError when no File was found.
func (fq *FileQuery) First(ctx context.Context) (*File, error) {
	nodes, err := fq.Limit(1).All(newQueryContext(ctx, TypeFile, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{file.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FileQuery) FirstX(ctx context.Context) *File {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first File ID from the query.
// Returns a *NotFoundError when no File ID was found.
func (fq *FileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(1).IDs(newQueryContext(ctx, TypeFile, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{file.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FileQuery) FirstIDX(ctx context.Context) string {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single File entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one File entity is found.
// Returns a *NotFoundError when no File entities are found.
func (fq *FileQuery) Only(ctx context.Context) (*File, error) {
	nodes, err := fq.Limit(2).All(newQueryContext(ctx, TypeFile, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{file.Label}
	default:
		return nil, &NotSingularError{file.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FileQuery) OnlyX(ctx context.Context) *File {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only File ID in the query.
// Returns a *NotSingularError when more than one File ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(2).IDs(newQueryContext(ctx, TypeFile, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{file.Label}
	default:
		err = &NotSingularError{file.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FileQuery) OnlyIDX(ctx context.Context) string {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Files.
func (fq *FileQuery) All(ctx context.Context) ([]*File, error) {
	ctx = newQueryContext(ctx, TypeFile, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*File, *FileQuery]()
	return withInterceptors[[]*File](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FileQuery) AllX(ctx context.Context) []*File {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of File IDs.
func (fq *FileQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = newQueryContext(ctx, TypeFile, "IDs")
	if err := fq.Select(file.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FileQuery) IDsX(ctx context.Context) []string {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FileQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeFile, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FileQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FileQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeFile, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FileQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FileQuery) Clone() *FileQuery {
	if fq == nil {
		return nil
	}
	return &FileQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.File{}, fq.predicates...),
		withOwner:  fq.withOwner.Clone(),
		withType:   fq.withType.Clone(),
		withField:  fq.withField.Clone(),
		// clone intermediate query.
		gremlin: fq.gremlin.Clone(),
		path:    fq.path,
		unique:  fq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithOwner(opts ...func(*UserQuery)) *FileQuery {
	query := (&UserClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withOwner = query
	return fq
}

// WithType tells the query-builder to eager-load the nodes that are connected to
// the "type" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithType(opts ...func(*FileTypeQuery)) *FileQuery {
	query := (&FileTypeClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withType = query
	return fq
}

// WithField tells the query-builder to eager-load the nodes that are connected to
// the "field" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithField(opts ...func(*FieldTypeQuery)) *FileQuery {
	query := (&FieldTypeClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withField = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Size int `json:"size,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.File.Query().
//		GroupBy(file.FieldSize).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FileQuery) GroupBy(field string, fields ...string) *FileGroupBy {
	fq.fields = append([]string{field}, fields...)
	grbuild := &FileGroupBy{build: fq}
	grbuild.flds = &fq.fields
	grbuild.label = file.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Size int `json:"size,omitempty"`
//	}
//
//	client.File.Query().
//		Select(file.FieldSize).
//		Scan(ctx, &v)
func (fq *FileQuery) Select(fields ...string) *FileSelect {
	fq.fields = append(fq.fields, fields...)
	sbuild := &FileSelect{FileQuery: fq}
	sbuild.label = file.Label
	sbuild.flds, sbuild.scan = &fq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileSelect configured with the given aggregations.
func (fq *FileQuery) Aggregate(fns ...AggregateFunc) *FileSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.gremlin = prev
	}
	return nil
}

func (fq *FileQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*File, error) {
	res := &gremlin.Response{}
	traversal := fq.gremlinQuery(ctx)
	if len(fq.fields) > 0 {
		fields := make([]any, len(fq.fields))
		for i, f := range fq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := fq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fs Files
	if err := fs.FromResponse(res); err != nil {
		return nil, err
	}
	fs.config(fq.config)
	return fs, nil
}

func (fq *FileQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := fq.gremlinQuery(ctx).Count().Query()
	if err := fq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (fq *FileQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(file.Label)
	if fq.gremlin != nil {
		v = fq.gremlin.Clone()
	}
	for _, p := range fq.predicates {
		p(v)
	}
	if len(fq.order) > 0 {
		v.Order()
		for _, p := range fq.order {
			p(v)
		}
	}
	switch limit, offset := fq.limit, fq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := fq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// FileGroupBy is the group-by builder for File entities.
type FileGroupBy struct {
	selector
	build *FileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FileGroupBy) Aggregate(fns ...AggregateFunc) *FileGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeFile, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileQuery, *FileGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FileGroupBy) gremlinScan(ctx context.Context, root *FileQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range fgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *fgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*fgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := fgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*fgb.flds)+len(fgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// FileSelect is the builder for selecting fields of File entities.
type FileSelect struct {
	*FileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FileSelect) Aggregate(fns ...AggregateFunc) *FileSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FileSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeFile, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileQuery, *FileSelect](ctx, fs.FileQuery, fs, fs.inters, v)
}

func (fs *FileSelect) gremlinScan(ctx context.Context, root *FileQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if len(fs.fields) == 1 {
		if fs.fields[0] != file.FieldID {
			traversal = traversal.Values(fs.fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(fs.fields))
		for i, f := range fs.fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := fs.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
