// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	context "context"
	fmt "fmt"
	math "math"

	"entgo.io/ent/dialect/sql"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/conversion"
	predicate "entgo.io/ent/entc/integration/migrate/entv2/predicate"
	field "entgo.io/ent/schema/field"
)

// ConversionQuery is the builder for querying Conversion entities.
type ConversionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Conversion
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConversionQuery builder.
func (cq *ConversionQuery) Where(ps ...predicate.Conversion) *ConversionQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConversionQuery) Limit(limit int) *ConversionQuery {
	cq.limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConversionQuery) Offset(offset int) *ConversionQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConversionQuery) Unique(unique bool) *ConversionQuery {
	cq.unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConversionQuery) Order(o ...OrderFunc) *ConversionQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Conversion entity from the query.
// Returns a *NotFoundError when no Conversion was found.
func (cq *ConversionQuery) First(ctx context.Context) (*Conversion, error) {
	nodes, err := cq.Limit(1).All(newQueryContext(ctx, TypeConversion, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{conversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConversionQuery) FirstX(ctx context.Context) *Conversion {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Conversion ID from the query.
// Returns a *NotFoundError when no Conversion ID was found.
func (cq *ConversionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(newQueryContext(ctx, TypeConversion, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{conversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConversionQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Conversion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Conversion entity is found.
// Returns a *NotFoundError when no Conversion entities are found.
func (cq *ConversionQuery) Only(ctx context.Context) (*Conversion, error) {
	nodes, err := cq.Limit(2).All(newQueryContext(ctx, TypeConversion, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{conversion.Label}
	default:
		return nil, &NotSingularError{conversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConversionQuery) OnlyX(ctx context.Context) *Conversion {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Conversion ID in the query.
// Returns a *NotSingularError when more than one Conversion ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConversionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(newQueryContext(ctx, TypeConversion, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{conversion.Label}
	default:
		err = &NotSingularError{conversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConversionQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Conversions.
func (cq *ConversionQuery) All(ctx context.Context) ([]*Conversion, error) {
	ctx = newQueryContext(ctx, TypeConversion, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Conversion, *ConversionQuery]()
	return withInterceptors[[]*Conversion](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConversionQuery) AllX(ctx context.Context) []*Conversion {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Conversion IDs.
func (cq *ConversionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeConversion, "IDs")
	if err := cq.Select(conversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConversionQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConversionQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeConversion, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConversionQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConversionQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConversionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeConversion, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entv2: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConversionQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConversionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConversionQuery) Clone() *ConversionQuery {
	if cq == nil {
		return nil
	}
	return &ConversionQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Conversion{}, cq.predicates...),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Conversion.Query().
//		GroupBy(conversion.FieldName).
//		Aggregate(entv2.Count()).
//		Scan(ctx, &v)
func (cq *ConversionQuery) GroupBy(field string, fields ...string) *ConversionGroupBy {
	cq.fields = append([]string{field}, fields...)
	grbuild := &ConversionGroupBy{build: cq}
	grbuild.flds = &cq.fields
	grbuild.label = conversion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Conversion.Query().
//		Select(conversion.FieldName).
//		Scan(ctx, &v)
func (cq *ConversionQuery) Select(fields ...string) *ConversionSelect {
	cq.fields = append(cq.fields, fields...)
	sbuild := &ConversionSelect{ConversionQuery: cq}
	sbuild.label = conversion.Label
	sbuild.flds, sbuild.scan = &cq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConversionSelect configured with the given aggregations.
func (cq *ConversionQuery) Aggregate(fns ...AggregateFunc) *ConversionSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConversionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("entv2: uninitialized interceptor (forgotten import entv2/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.fields {
		if !conversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConversionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Conversion, error) {
	var (
		nodes = []*Conversion{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Conversion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Conversion{config: cq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *ConversionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConversionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   conversion.Table,
			Columns: conversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversion.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversion.FieldID)
		for i := range fields {
			if fields[i] != conversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConversionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(conversion.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = conversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConversionGroupBy is the group-by builder for Conversion entities.
type ConversionGroupBy struct {
	selector
	build *ConversionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConversionGroupBy) Aggregate(fns ...AggregateFunc) *ConversionGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConversionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeConversion, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConversionQuery, *ConversionGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConversionGroupBy) sqlScan(ctx context.Context, root *ConversionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConversionSelect is the builder for selecting fields of Conversion entities.
type ConversionSelect struct {
	*ConversionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConversionSelect) Aggregate(fns ...AggregateFunc) *ConversionSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConversionSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeConversion, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConversionQuery, *ConversionSelect](ctx, cs.ConversionQuery, cs, cs.inters, v)
}

func (cs *ConversionSelect) sqlScan(ctx context.Context, root *ConversionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
