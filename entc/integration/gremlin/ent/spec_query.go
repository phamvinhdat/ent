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
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/spec"
)

// SpecQuery is the builder for querying Spec entities.
type SpecQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Spec
	withCard   *CardQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the SpecQuery builder.
func (sq *SpecQuery) Where(ps ...predicate.Spec) *SpecQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SpecQuery) Limit(limit int) *SpecQuery {
	sq.limit = &limit
	return sq
}

// Offset to start from.
func (sq *SpecQuery) Offset(offset int) *SpecQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SpecQuery) Unique(unique bool) *SpecQuery {
	sq.unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SpecQuery) Order(o ...OrderFunc) *SpecQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCard chains the current query on the "card" edge.
func (sq *SpecQuery) QueryCard() *CardQuery {
	query := (&CardClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := sq.gremlinQuery(ctx)
		fromU = gremlin.OutE(spec.CardLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first Spec entity from the query.
// Returns a *NotFoundError when no Spec was found.
func (sq *SpecQuery) First(ctx context.Context) (*Spec, error) {
	nodes, err := sq.Limit(1).All(newQueryContext(ctx, TypeSpec, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{spec.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SpecQuery) FirstX(ctx context.Context) *Spec {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Spec ID from the query.
// Returns a *NotFoundError when no Spec ID was found.
func (sq *SpecQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(newQueryContext(ctx, TypeSpec, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{spec.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SpecQuery) FirstIDX(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Spec entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Spec entity is found.
// Returns a *NotFoundError when no Spec entities are found.
func (sq *SpecQuery) Only(ctx context.Context) (*Spec, error) {
	nodes, err := sq.Limit(2).All(newQueryContext(ctx, TypeSpec, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{spec.Label}
	default:
		return nil, &NotSingularError{spec.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SpecQuery) OnlyX(ctx context.Context) *Spec {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Spec ID in the query.
// Returns a *NotSingularError when more than one Spec ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SpecQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(newQueryContext(ctx, TypeSpec, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{spec.Label}
	default:
		err = &NotSingularError{spec.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SpecQuery) OnlyIDX(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Specs.
func (sq *SpecQuery) All(ctx context.Context) ([]*Spec, error) {
	ctx = newQueryContext(ctx, TypeSpec, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Spec, *SpecQuery]()
	return withInterceptors[[]*Spec](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SpecQuery) AllX(ctx context.Context) []*Spec {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Spec IDs.
func (sq *SpecQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = newQueryContext(ctx, TypeSpec, "IDs")
	if err := sq.Select(spec.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SpecQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SpecQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeSpec, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SpecQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SpecQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SpecQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeSpec, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SpecQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpecQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SpecQuery) Clone() *SpecQuery {
	if sq == nil {
		return nil
	}
	return &SpecQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Spec{}, sq.predicates...),
		withCard:   sq.withCard.Clone(),
		// clone intermediate query.
		gremlin: sq.gremlin.Clone(),
		path:    sq.path,
		unique:  sq.unique,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SpecQuery) WithCard(opts ...func(*CardQuery)) *SpecQuery {
	query := (&CardClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCard = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *SpecQuery) GroupBy(field string, fields ...string) *SpecGroupBy {
	sq.fields = append([]string{field}, fields...)
	grbuild := &SpecGroupBy{build: sq}
	grbuild.flds = &sq.fields
	grbuild.label = spec.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sq *SpecQuery) Select(fields ...string) *SpecSelect {
	sq.fields = append(sq.fields, fields...)
	sbuild := &SpecSelect{SpecQuery: sq}
	sbuild.label = spec.Label
	sbuild.flds, sbuild.scan = &sq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SpecSelect configured with the given aggregations.
func (sq *SpecQuery) Aggregate(fns ...AggregateFunc) *SpecSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SpecQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.gremlin = prev
	}
	return nil
}

func (sq *SpecQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Spec, error) {
	res := &gremlin.Response{}
	traversal := sq.gremlinQuery(ctx)
	if len(sq.fields) > 0 {
		fields := make([]any, len(sq.fields))
		for i, f := range sq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := sq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var sSlice Specs
	if err := sSlice.FromResponse(res); err != nil {
		return nil, err
	}
	sSlice.config(sq.config)
	return sSlice, nil
}

func (sq *SpecQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := sq.gremlinQuery(ctx).Count().Query()
	if err := sq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (sq *SpecQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(spec.Label)
	if sq.gremlin != nil {
		v = sq.gremlin.Clone()
	}
	for _, p := range sq.predicates {
		p(v)
	}
	if len(sq.order) > 0 {
		v.Order()
		for _, p := range sq.order {
			p(v)
		}
	}
	switch limit, offset := sq.limit, sq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := sq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// SpecGroupBy is the group-by builder for Spec entities.
type SpecGroupBy struct {
	selector
	build *SpecQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SpecGroupBy) Aggregate(fns ...AggregateFunc) *SpecGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SpecGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeSpec, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecQuery, *SpecGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SpecGroupBy) gremlinScan(ctx context.Context, root *SpecQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range sgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *sgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*sgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := sgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*sgb.flds)+len(sgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// SpecSelect is the builder for selecting fields of Spec entities.
type SpecSelect struct {
	*SpecQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SpecSelect) Aggregate(fns ...AggregateFunc) *SpecSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SpecSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeSpec, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecQuery, *SpecSelect](ctx, ss.SpecQuery, ss, ss.inters, v)
}

func (ss *SpecSelect) gremlinScan(ctx context.Context, root *SpecQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if len(ss.fields) == 1 {
		if ss.fields[0] != spec.FieldID {
			traversal = traversal.Values(ss.fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ss.fields))
		for i, f := range ss.fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ss.driver.Exec(ctx, query, bindings, res); err != nil {
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
