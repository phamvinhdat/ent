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
	"entgo.io/ent/entc/integration/gremlin/ent/pet"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// PetDelete is the builder for deleting a Pet entity.
type PetDelete struct {
	config
	hooks    []Hook
	mutation *PetMutation
}

// Where appends a list predicates to the PetDelete builder.
func (pd *PetDelete) Where(ps ...predicate.Pet) *PetDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PetMutation](ctx, pd.gremlinExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PetDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PetDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pd.gremlin().Query()
	if err := pd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	pd.mutation.done = true
	return res.ReadInt()
}

func (pd *PetDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(pet.Label)
	for _, p := range pd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// PetDeleteOne is the builder for deleting a single Pet entity.
type PetDeleteOne struct {
	pd *PetDelete
}

// Exec executes the deletion query.
func (pdo *PetDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PetDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
