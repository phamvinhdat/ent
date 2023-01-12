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
	"entgo.io/ent/entc/integration/migrate/entv2/customtype"
	predicate "entgo.io/ent/entc/integration/migrate/entv2/predicate"
	field "entgo.io/ent/schema/field"
)

// CustomTypeQuery is the builder for querying CustomType entities.
type CustomTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.CustomType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomTypeQuery builder.
func (ctq *CustomTypeQuery) Where(ps ...predicate.CustomType) *CustomTypeQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit the number of records to be returned by this query.
func (ctq *CustomTypeQuery) Limit(limit int) *CustomTypeQuery {
	ctq.limit = &limit
	return ctq
}

// Offset to start from.
func (ctq *CustomTypeQuery) Offset(offset int) *CustomTypeQuery {
	ctq.offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *CustomTypeQuery) Unique(unique bool) *CustomTypeQuery {
	ctq.unique = &unique
	return ctq
}

// Order specifies how the records should be ordered.
func (ctq *CustomTypeQuery) Order(o ...OrderFunc) *CustomTypeQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// First returns the first CustomType entity from the query.
// Returns a *NotFoundError when no CustomType was found.
func (ctq *CustomTypeQuery) First(ctx context.Context) (*CustomType, error) {
	nodes, err := ctq.Limit(1).All(newQueryContext(ctx, TypeCustomType, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CustomTypeQuery) FirstX(ctx context.Context) *CustomType {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomType ID from the query.
// Returns a *NotFoundError when no CustomType ID was found.
func (ctq *CustomTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(newQueryContext(ctx, TypeCustomType, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CustomTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CustomType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CustomType entity is found.
// Returns a *NotFoundError when no CustomType entities are found.
func (ctq *CustomTypeQuery) Only(ctx context.Context) (*CustomType, error) {
	nodes, err := ctq.Limit(2).All(newQueryContext(ctx, TypeCustomType, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customtype.Label}
	default:
		return nil, &NotSingularError{customtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CustomTypeQuery) OnlyX(ctx context.Context) *CustomType {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CustomType ID in the query.
// Returns a *NotSingularError when more than one CustomType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *CustomTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(newQueryContext(ctx, TypeCustomType, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = &NotSingularError{customtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CustomTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomTypes.
func (ctq *CustomTypeQuery) All(ctx context.Context) ([]*CustomType, error) {
	ctx = newQueryContext(ctx, TypeCustomType, "All")
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CustomType, *CustomTypeQuery]()
	return withInterceptors[[]*CustomType](ctx, ctq, qr, ctq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CustomTypeQuery) AllX(ctx context.Context) []*CustomType {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomType IDs.
func (ctq *CustomTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeCustomType, "IDs")
	if err := ctq.Select(customtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CustomTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CustomTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeCustomType, "Count")
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ctq, querierCount[*CustomTypeQuery](), ctq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CustomTypeQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CustomTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeCustomType, "Exist")
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entv2: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CustomTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CustomTypeQuery) Clone() *CustomTypeQuery {
	if ctq == nil {
		return nil
	}
	return &CustomTypeQuery{
		config:     ctq.config,
		limit:      ctq.limit,
		offset:     ctq.offset,
		order:      append([]OrderFunc{}, ctq.order...),
		inters:     append([]Interceptor{}, ctq.inters...),
		predicates: append([]predicate.CustomType{}, ctq.predicates...),
		// clone intermediate query.
		sql:    ctq.sql.Clone(),
		path:   ctq.path,
		unique: ctq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Custom string `json:"custom,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomType.Query().
//		GroupBy(customtype.FieldCustom).
//		Aggregate(entv2.Count()).
//		Scan(ctx, &v)
func (ctq *CustomTypeQuery) GroupBy(field string, fields ...string) *CustomTypeGroupBy {
	ctq.fields = append([]string{field}, fields...)
	grbuild := &CustomTypeGroupBy{build: ctq}
	grbuild.flds = &ctq.fields
	grbuild.label = customtype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Custom string `json:"custom,omitempty"`
//	}
//
//	client.CustomType.Query().
//		Select(customtype.FieldCustom).
//		Scan(ctx, &v)
func (ctq *CustomTypeQuery) Select(fields ...string) *CustomTypeSelect {
	ctq.fields = append(ctq.fields, fields...)
	sbuild := &CustomTypeSelect{CustomTypeQuery: ctq}
	sbuild.label = customtype.Label
	sbuild.flds, sbuild.scan = &ctq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomTypeSelect configured with the given aggregations.
func (ctq *CustomTypeQuery) Aggregate(fns ...AggregateFunc) *CustomTypeSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *CustomTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ctq.inters {
		if inter == nil {
			return fmt.Errorf("entv2: uninitialized interceptor (forgotten import entv2/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ctq); err != nil {
				return err
			}
		}
	}
	for _, f := range ctq.fields {
		if !customtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CustomTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CustomType, error) {
	var (
		nodes = []*CustomType{}
		_spec = ctq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CustomType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CustomType{config: ctq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ctq *CustomTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	_spec.Node.Columns = ctq.fields
	if len(ctq.fields) > 0 {
		_spec.Unique = ctq.unique != nil && *ctq.unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CustomTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customtype.Table,
			Columns: customtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		},
		From:   ctq.sql,
		Unique: true,
	}
	if unique := ctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customtype.FieldID)
		for i := range fields {
			if fields[i] != customtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *CustomTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(customtype.Table)
	columns := ctq.fields
	if len(columns) == 0 {
		columns = customtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.unique != nil && *ctq.unique {
		selector.Distinct()
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CustomTypeGroupBy is the group-by builder for CustomType entities.
type CustomTypeGroupBy struct {
	selector
	build *CustomTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CustomTypeGroupBy) Aggregate(fns ...AggregateFunc) *CustomTypeGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the selector query and scans the result into the given value.
func (ctgb *CustomTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeCustomType, "GroupBy")
	if err := ctgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomTypeQuery, *CustomTypeGroupBy](ctx, ctgb.build, ctgb, ctgb.build.inters, v)
}

func (ctgb *CustomTypeGroupBy) sqlScan(ctx context.Context, root *CustomTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ctgb.flds)+len(ctgb.fns))
		for _, f := range *ctgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ctgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomTypeSelect is the builder for selecting fields of CustomType entities.
type CustomTypeSelect struct {
	*CustomTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *CustomTypeSelect) Aggregate(fns ...AggregateFunc) *CustomTypeSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *CustomTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeCustomType, "Select")
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomTypeQuery, *CustomTypeSelect](ctx, cts.CustomTypeQuery, cts, cts.inters, v)
}

func (cts *CustomTypeSelect) sqlScan(ctx context.Context, root *CustomTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
