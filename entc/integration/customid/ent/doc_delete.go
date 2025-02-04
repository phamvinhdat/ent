// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"

	"entgo.io/ent/dialect/sql"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	predicate "entgo.io/ent/entc/integration/customid/ent/predicate"
	field "entgo.io/ent/schema/field"
)

// DocDelete is the builder for deleting a Doc entity.
type DocDelete struct {
	config
	hooks    []Hook
	mutation *DocMutation
}

// Where appends a list predicates to the DocDelete builder.
func (dd *DocDelete) Where(ps ...predicate.Doc) *DocDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DocDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DocMutation](ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DocDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DocDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: doc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: doc.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DocDeleteOne is the builder for deleting a single Doc entity.
type DocDeleteOne struct {
	dd *DocDelete
}

// Exec executes the deletion query.
func (ddo *DocDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{doc.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DocDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
