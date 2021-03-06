// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/language"
	"Kynesia/ent/register"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LanguageCreate is the builder for creating a Language entity.
type LanguageCreate struct {
	config
	mutation *LanguageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (lc *LanguageCreate) SetName(s string) *LanguageCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetTalk sets the "talk" field.
func (lc *LanguageCreate) SetTalk(s string) *LanguageCreate {
	lc.mutation.SetTalk(s)
	return lc
}

// SetWrite sets the "write" field.
func (lc *LanguageCreate) SetWrite(s string) *LanguageCreate {
	lc.mutation.SetWrite(s)
	return lc
}

// SetRead sets the "read" field.
func (lc *LanguageCreate) SetRead(s string) *LanguageCreate {
	lc.mutation.SetRead(s)
	return lc
}

// SetListen sets the "listen" field.
func (lc *LanguageCreate) SetListen(s string) *LanguageCreate {
	lc.mutation.SetListen(s)
	return lc
}

// AddRegisterIDs adds the "register" edge to the Register entity by IDs.
func (lc *LanguageCreate) AddRegisterIDs(ids ...int) *LanguageCreate {
	lc.mutation.AddRegisterIDs(ids...)
	return lc
}

// AddRegister adds the "register" edges to the Register entity.
func (lc *LanguageCreate) AddRegister(r ...*Register) *LanguageCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return lc.AddRegisterIDs(ids...)
}

// Mutation returns the LanguageMutation object of the builder.
func (lc *LanguageCreate) Mutation() *LanguageMutation {
	return lc.mutation
}

// Save creates the Language in the database.
func (lc *LanguageCreate) Save(ctx context.Context) (*Language, error) {
	var (
		err  error
		node *Language
	)
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LanguageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LanguageCreate) SaveX(ctx context.Context) *Language {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LanguageCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LanguageCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LanguageCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Language.name"`)}
	}
	if _, ok := lc.mutation.Talk(); !ok {
		return &ValidationError{Name: "talk", err: errors.New(`ent: missing required field "Language.talk"`)}
	}
	if _, ok := lc.mutation.Write(); !ok {
		return &ValidationError{Name: "write", err: errors.New(`ent: missing required field "Language.write"`)}
	}
	if _, ok := lc.mutation.Read(); !ok {
		return &ValidationError{Name: "read", err: errors.New(`ent: missing required field "Language.read"`)}
	}
	if _, ok := lc.mutation.Listen(); !ok {
		return &ValidationError{Name: "listen", err: errors.New(`ent: missing required field "Language.listen"`)}
	}
	return nil
}

func (lc *LanguageCreate) sqlSave(ctx context.Context) (*Language, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LanguageCreate) createSpec() (*Language, *sqlgraph.CreateSpec) {
	var (
		_node = &Language{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: language.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: language.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lc.mutation.Talk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldTalk,
		})
		_node.Talk = value
	}
	if value, ok := lc.mutation.Write(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldWrite,
		})
		_node.Write = value
	}
	if value, ok := lc.mutation.Read(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldRead,
		})
		_node.Read = value
	}
	if value, ok := lc.mutation.Listen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldListen,
		})
		_node.Listen = value
	}
	if nodes := lc.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   language.RegisterTable,
			Columns: language.RegisterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: register.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LanguageCreateBulk is the builder for creating many Language entities in bulk.
type LanguageCreateBulk struct {
	config
	builders []*LanguageCreate
}

// Save creates the Language entities in the database.
func (lcb *LanguageCreateBulk) Save(ctx context.Context) ([]*Language, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Language, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LanguageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LanguageCreateBulk) SaveX(ctx context.Context) []*Language {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LanguageCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LanguageCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
