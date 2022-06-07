// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/family"
	"Kynesia/ent/register"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FamilyCreate is the builder for creating a Family entity.
type FamilyCreate struct {
	config
	mutation *FamilyMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (fc *FamilyCreate) SetStatus(s string) *FamilyCreate {
	fc.mutation.SetStatus(s)
	return fc
}

// SetName sets the "name" field.
func (fc *FamilyCreate) SetName(s string) *FamilyCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetGender sets the "gender" field.
func (fc *FamilyCreate) SetGender(s string) *FamilyCreate {
	fc.mutation.SetGender(s)
	return fc
}

// SetBirthplace sets the "birthplace" field.
func (fc *FamilyCreate) SetBirthplace(s string) *FamilyCreate {
	fc.mutation.SetBirthplace(s)
	return fc
}

// SetBirthdate sets the "birthdate" field.
func (fc *FamilyCreate) SetBirthdate(t time.Time) *FamilyCreate {
	fc.mutation.SetBirthdate(t)
	return fc
}

// SetEducation sets the "education" field.
func (fc *FamilyCreate) SetEducation(s string) *FamilyCreate {
	fc.mutation.SetEducation(s)
	return fc
}

// SetJob sets the "job" field.
func (fc *FamilyCreate) SetJob(s string) *FamilyCreate {
	fc.mutation.SetJob(s)
	return fc
}

// AddRegisterIDs adds the "register" edge to the Register entity by IDs.
func (fc *FamilyCreate) AddRegisterIDs(ids ...int) *FamilyCreate {
	fc.mutation.AddRegisterIDs(ids...)
	return fc
}

// AddRegister adds the "register" edges to the Register entity.
func (fc *FamilyCreate) AddRegister(r ...*Register) *FamilyCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return fc.AddRegisterIDs(ids...)
}

// Mutation returns the FamilyMutation object of the builder.
func (fc *FamilyCreate) Mutation() *FamilyMutation {
	return fc.mutation
}

// Save creates the Family in the database.
func (fc *FamilyCreate) Save(ctx context.Context) (*Family, error) {
	var (
		err  error
		node *Family
	)
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FamilyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FamilyCreate) SaveX(ctx context.Context) *Family {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FamilyCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FamilyCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FamilyCreate) check() error {
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Family.status"`)}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Family.name"`)}
	}
	if _, ok := fc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Family.gender"`)}
	}
	if _, ok := fc.mutation.Birthplace(); !ok {
		return &ValidationError{Name: "birthplace", err: errors.New(`ent: missing required field "Family.birthplace"`)}
	}
	if _, ok := fc.mutation.Birthdate(); !ok {
		return &ValidationError{Name: "birthdate", err: errors.New(`ent: missing required field "Family.birthdate"`)}
	}
	if _, ok := fc.mutation.Education(); !ok {
		return &ValidationError{Name: "education", err: errors.New(`ent: missing required field "Family.education"`)}
	}
	if _, ok := fc.mutation.Job(); !ok {
		return &ValidationError{Name: "job", err: errors.New(`ent: missing required field "Family.job"`)}
	}
	return nil
}

func (fc *FamilyCreate) sqlSave(ctx context.Context) (*Family, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FamilyCreate) createSpec() (*Family, *sqlgraph.CreateSpec) {
	var (
		_node = &Family{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: family.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: family.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: family.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: family.FieldName,
		})
		_node.Name = value
	}
	if value, ok := fc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: family.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := fc.mutation.Birthplace(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: family.FieldBirthplace,
		})
		_node.Birthplace = value
	}
	if value, ok := fc.mutation.Birthdate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: family.FieldBirthdate,
		})
		_node.Birthdate = value
	}
	if value, ok := fc.mutation.Education(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: family.FieldEducation,
		})
		_node.Education = value
	}
	if value, ok := fc.mutation.Job(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: family.FieldJob,
		})
		_node.Job = value
	}
	if nodes := fc.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   family.RegisterTable,
			Columns: family.RegisterPrimaryKey,
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

// FamilyCreateBulk is the builder for creating many Family entities in bulk.
type FamilyCreateBulk struct {
	config
	builders []*FamilyCreate
}

// Save creates the Family entities in the database.
func (fcb *FamilyCreateBulk) Save(ctx context.Context) ([]*Family, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Family, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FamilyMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FamilyCreateBulk) SaveX(ctx context.Context) []*Family {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FamilyCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FamilyCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
