// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/education"
	"Kynesia/ent/register"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EducationCreate is the builder for creating a Education entity.
type EducationCreate struct {
	config
	mutation *EducationMutation
	hooks    []Hook
}

// SetGrade sets the "grade" field.
func (ec *EducationCreate) SetGrade(s string) *EducationCreate {
	ec.mutation.SetGrade(s)
	return ec
}

// SetName sets the "name" field.
func (ec *EducationCreate) SetName(s string) *EducationCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetProvince sets the "province" field.
func (ec *EducationCreate) SetProvince(s string) *EducationCreate {
	ec.mutation.SetProvince(s)
	return ec
}

// SetCity sets the "city" field.
func (ec *EducationCreate) SetCity(s string) *EducationCreate {
	ec.mutation.SetCity(s)
	return ec
}

// SetEnter sets the "enter" field.
func (ec *EducationCreate) SetEnter(s string) *EducationCreate {
	ec.mutation.SetEnter(s)
	return ec
}

// SetGraduate sets the "graduate" field.
func (ec *EducationCreate) SetGraduate(s string) *EducationCreate {
	ec.mutation.SetGraduate(s)
	return ec
}

// AddRegisterIDs adds the "register" edge to the Register entity by IDs.
func (ec *EducationCreate) AddRegisterIDs(ids ...int) *EducationCreate {
	ec.mutation.AddRegisterIDs(ids...)
	return ec
}

// AddRegister adds the "register" edges to the Register entity.
func (ec *EducationCreate) AddRegister(r ...*Register) *EducationCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddRegisterIDs(ids...)
}

// Mutation returns the EducationMutation object of the builder.
func (ec *EducationCreate) Mutation() *EducationMutation {
	return ec.mutation
}

// Save creates the Education in the database.
func (ec *EducationCreate) Save(ctx context.Context) (*Education, error) {
	var (
		err  error
		node *Education
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EducationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EducationCreate) SaveX(ctx context.Context) *Education {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EducationCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EducationCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EducationCreate) check() error {
	if _, ok := ec.mutation.Grade(); !ok {
		return &ValidationError{Name: "grade", err: errors.New(`ent: missing required field "Education.grade"`)}
	}
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Education.name"`)}
	}
	if _, ok := ec.mutation.Province(); !ok {
		return &ValidationError{Name: "province", err: errors.New(`ent: missing required field "Education.province"`)}
	}
	if _, ok := ec.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "Education.city"`)}
	}
	if _, ok := ec.mutation.Enter(); !ok {
		return &ValidationError{Name: "enter", err: errors.New(`ent: missing required field "Education.enter"`)}
	}
	if _, ok := ec.mutation.Graduate(); !ok {
		return &ValidationError{Name: "graduate", err: errors.New(`ent: missing required field "Education.graduate"`)}
	}
	return nil
}

func (ec *EducationCreate) sqlSave(ctx context.Context) (*Education, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EducationCreate) createSpec() (*Education, *sqlgraph.CreateSpec) {
	var (
		_node = &Education{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: education.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: education.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Grade(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldGrade,
		})
		_node.Grade = value
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ec.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldProvince,
		})
		_node.Province = value
	}
	if value, ok := ec.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldCity,
		})
		_node.City = value
	}
	if value, ok := ec.mutation.Enter(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldEnter,
		})
		_node.Enter = value
	}
	if value, ok := ec.mutation.Graduate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldGraduate,
		})
		_node.Graduate = value
	}
	if nodes := ec.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   education.RegisterTable,
			Columns: education.RegisterPrimaryKey,
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

// EducationCreateBulk is the builder for creating many Education entities in bulk.
type EducationCreateBulk struct {
	config
	builders []*EducationCreate
}

// Save creates the Education entities in the database.
func (ecb *EducationCreateBulk) Save(ctx context.Context) ([]*Education, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Education, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EducationMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EducationCreateBulk) SaveX(ctx context.Context) []*Education {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EducationCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EducationCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
