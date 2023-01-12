// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"

	gremlin "entgo.io/ent/dialect/gremlin"
	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	g "entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/license"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// LicenseDelete is the builder for deleting a License entity.
type LicenseDelete struct {
	config
	hooks    []Hook
	mutation *LicenseMutation
}

// Where appends a list predicates to the LicenseDelete builder.
func (ld *LicenseDelete) Where(ps ...predicate.License) *LicenseDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LicenseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LicenseMutation](ctx, ld.gremlinExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LicenseDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LicenseDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ld.gremlin().Query()
	if err := ld.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	ld.mutation.done = true
	return res.ReadInt()
}

func (ld *LicenseDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(license.Label)
	for _, p := range ld.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// LicenseDeleteOne is the builder for deleting a single License entity.
type LicenseDeleteOne struct {
	ld *LicenseDelete
}

// Exec executes the deletion query.
func (ldo *LicenseDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{license.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LicenseDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}
