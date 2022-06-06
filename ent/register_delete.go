// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/predicate"
	"Kynesia/ent/register"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RegisterDelete is the builder for deleting a Register entity.
type RegisterDelete struct {
	config
	hooks    []Hook
	mutation *RegisterMutation
}

// Where appends a list predicates to the RegisterDelete builder.
func (rd *RegisterDelete) Where(ps ...predicate.Register) *RegisterDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RegisterDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegisterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			if rd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RegisterDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RegisterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: register.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: register.FieldID,
			},
		},
	}
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// RegisterDeleteOne is the builder for deleting a single Register entity.
type RegisterDeleteOne struct {
	rd *RegisterDelete
}

// Exec executes the deletion query.
func (rdo *RegisterDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{register.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RegisterDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}