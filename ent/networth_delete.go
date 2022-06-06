// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/networth"
	"Kynesia/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NetworthDelete is the builder for deleting a Networth entity.
type NetworthDelete struct {
	config
	hooks    []Hook
	mutation *NetworthMutation
}

// Where appends a list predicates to the NetworthDelete builder.
func (nd *NetworthDelete) Where(ps ...predicate.Networth) *NetworthDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NetworthDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nd.hooks) == 0 {
		affected, err = nd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetworthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nd.mutation = mutation
			affected, err = nd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nd.hooks) - 1; i >= 0; i-- {
			if nd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NetworthDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NetworthDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: networth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: networth.FieldID,
			},
		},
	}
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
}

// NetworthDeleteOne is the builder for deleting a single Networth entity.
type NetworthDeleteOne struct {
	nd *NetworthDelete
}

// Exec executes the deletion query.
func (ndo *NetworthDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{networth.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NetworthDeleteOne) ExecX(ctx context.Context) {
	ndo.nd.ExecX(ctx)
}
