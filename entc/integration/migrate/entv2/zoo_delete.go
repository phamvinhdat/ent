// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	context "context"

	"entgo.io/ent/dialect/sql"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	predicate "entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/entc/integration/migrate/entv2/zoo"
	field "entgo.io/ent/schema/field"
)

// ZooDelete is the builder for deleting a Zoo entity.
type ZooDelete struct {
	config
	hooks    []Hook
	mutation *ZooMutation
}

// Where appends a list predicates to the ZooDelete builder.
func (zd *ZooDelete) Where(ps ...predicate.Zoo) *ZooDelete {
	zd.mutation.Where(ps...)
	return zd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (zd *ZooDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ZooMutation](ctx, zd.sqlExec, zd.mutation, zd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (zd *ZooDelete) ExecX(ctx context.Context) int {
	n, err := zd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (zd *ZooDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: zoo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zoo.FieldID,
			},
		},
	}
	if ps := zd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, zd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	zd.mutation.done = true
	return affected, err
}

// ZooDeleteOne is the builder for deleting a single Zoo entity.
type ZooDeleteOne struct {
	zd *ZooDelete
}

// Exec executes the deletion query.
func (zdo *ZooDeleteOne) Exec(ctx context.Context) error {
	n, err := zdo.zd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{zoo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (zdo *ZooDeleteOne) ExecX(ctx context.Context) {
	zdo.zd.ExecX(ctx)
}
