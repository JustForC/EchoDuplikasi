// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/predicate"
	"Kynesia/ent/register"
	"Kynesia/ent/socialmedia"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SocialMediaUpdate is the builder for updating SocialMedia entities.
type SocialMediaUpdate struct {
	config
	hooks    []Hook
	mutation *SocialMediaMutation
}

// Where appends a list predicates to the SocialMediaUpdate builder.
func (smu *SocialMediaUpdate) Where(ps ...predicate.SocialMedia) *SocialMediaUpdate {
	smu.mutation.Where(ps...)
	return smu
}

// SetInstagram sets the "instagram" field.
func (smu *SocialMediaUpdate) SetInstagram(s string) *SocialMediaUpdate {
	smu.mutation.SetInstagram(s)
	return smu
}

// SetFacebook sets the "facebook" field.
func (smu *SocialMediaUpdate) SetFacebook(s string) *SocialMediaUpdate {
	smu.mutation.SetFacebook(s)
	return smu
}

// SetTiktok sets the "tiktok" field.
func (smu *SocialMediaUpdate) SetTiktok(s string) *SocialMediaUpdate {
	smu.mutation.SetTiktok(s)
	return smu
}

// SetTwitter sets the "twitter" field.
func (smu *SocialMediaUpdate) SetTwitter(s string) *SocialMediaUpdate {
	smu.mutation.SetTwitter(s)
	return smu
}

// AddRegisterIDs adds the "register" edge to the Register entity by IDs.
func (smu *SocialMediaUpdate) AddRegisterIDs(ids ...int) *SocialMediaUpdate {
	smu.mutation.AddRegisterIDs(ids...)
	return smu
}

// AddRegister adds the "register" edges to the Register entity.
func (smu *SocialMediaUpdate) AddRegister(r ...*Register) *SocialMediaUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return smu.AddRegisterIDs(ids...)
}

// Mutation returns the SocialMediaMutation object of the builder.
func (smu *SocialMediaUpdate) Mutation() *SocialMediaMutation {
	return smu.mutation
}

// ClearRegister clears all "register" edges to the Register entity.
func (smu *SocialMediaUpdate) ClearRegister() *SocialMediaUpdate {
	smu.mutation.ClearRegister()
	return smu
}

// RemoveRegisterIDs removes the "register" edge to Register entities by IDs.
func (smu *SocialMediaUpdate) RemoveRegisterIDs(ids ...int) *SocialMediaUpdate {
	smu.mutation.RemoveRegisterIDs(ids...)
	return smu
}

// RemoveRegister removes "register" edges to Register entities.
func (smu *SocialMediaUpdate) RemoveRegister(r ...*Register) *SocialMediaUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return smu.RemoveRegisterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *SocialMediaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smu.hooks) == 0 {
		affected, err = smu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SocialMediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smu.mutation = mutation
			affected, err = smu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smu.hooks) - 1; i >= 0; i-- {
			if smu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (smu *SocialMediaUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *SocialMediaUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *SocialMediaUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (smu *SocialMediaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   socialmedia.Table,
			Columns: socialmedia.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: socialmedia.FieldID,
			},
		},
	}
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smu.mutation.Instagram(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldInstagram,
		})
	}
	if value, ok := smu.mutation.Facebook(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldFacebook,
		})
	}
	if value, ok := smu.mutation.Tiktok(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldTiktok,
		})
	}
	if value, ok := smu.mutation.Twitter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldTwitter,
		})
	}
	if smu.mutation.RegisterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   socialmedia.RegisterTable,
			Columns: socialmedia.RegisterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: register.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smu.mutation.RemovedRegisterIDs(); len(nodes) > 0 && !smu.mutation.RegisterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   socialmedia.RegisterTable,
			Columns: socialmedia.RegisterPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smu.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   socialmedia.RegisterTable,
			Columns: socialmedia.RegisterPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{socialmedia.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SocialMediaUpdateOne is the builder for updating a single SocialMedia entity.
type SocialMediaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SocialMediaMutation
}

// SetInstagram sets the "instagram" field.
func (smuo *SocialMediaUpdateOne) SetInstagram(s string) *SocialMediaUpdateOne {
	smuo.mutation.SetInstagram(s)
	return smuo
}

// SetFacebook sets the "facebook" field.
func (smuo *SocialMediaUpdateOne) SetFacebook(s string) *SocialMediaUpdateOne {
	smuo.mutation.SetFacebook(s)
	return smuo
}

// SetTiktok sets the "tiktok" field.
func (smuo *SocialMediaUpdateOne) SetTiktok(s string) *SocialMediaUpdateOne {
	smuo.mutation.SetTiktok(s)
	return smuo
}

// SetTwitter sets the "twitter" field.
func (smuo *SocialMediaUpdateOne) SetTwitter(s string) *SocialMediaUpdateOne {
	smuo.mutation.SetTwitter(s)
	return smuo
}

// AddRegisterIDs adds the "register" edge to the Register entity by IDs.
func (smuo *SocialMediaUpdateOne) AddRegisterIDs(ids ...int) *SocialMediaUpdateOne {
	smuo.mutation.AddRegisterIDs(ids...)
	return smuo
}

// AddRegister adds the "register" edges to the Register entity.
func (smuo *SocialMediaUpdateOne) AddRegister(r ...*Register) *SocialMediaUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return smuo.AddRegisterIDs(ids...)
}

// Mutation returns the SocialMediaMutation object of the builder.
func (smuo *SocialMediaUpdateOne) Mutation() *SocialMediaMutation {
	return smuo.mutation
}

// ClearRegister clears all "register" edges to the Register entity.
func (smuo *SocialMediaUpdateOne) ClearRegister() *SocialMediaUpdateOne {
	smuo.mutation.ClearRegister()
	return smuo
}

// RemoveRegisterIDs removes the "register" edge to Register entities by IDs.
func (smuo *SocialMediaUpdateOne) RemoveRegisterIDs(ids ...int) *SocialMediaUpdateOne {
	smuo.mutation.RemoveRegisterIDs(ids...)
	return smuo
}

// RemoveRegister removes "register" edges to Register entities.
func (smuo *SocialMediaUpdateOne) RemoveRegister(r ...*Register) *SocialMediaUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return smuo.RemoveRegisterIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *SocialMediaUpdateOne) Select(field string, fields ...string) *SocialMediaUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated SocialMedia entity.
func (smuo *SocialMediaUpdateOne) Save(ctx context.Context) (*SocialMedia, error) {
	var (
		err  error
		node *SocialMedia
	)
	if len(smuo.hooks) == 0 {
		node, err = smuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SocialMediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smuo.mutation = mutation
			node, err = smuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smuo.hooks) - 1; i >= 0; i-- {
			if smuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *SocialMediaUpdateOne) SaveX(ctx context.Context) *SocialMedia {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *SocialMediaUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *SocialMediaUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (smuo *SocialMediaUpdateOne) sqlSave(ctx context.Context) (_node *SocialMedia, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   socialmedia.Table,
			Columns: socialmedia.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: socialmedia.FieldID,
			},
		},
	}
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SocialMedia.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, socialmedia.FieldID)
		for _, f := range fields {
			if !socialmedia.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != socialmedia.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.Instagram(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldInstagram,
		})
	}
	if value, ok := smuo.mutation.Facebook(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldFacebook,
		})
	}
	if value, ok := smuo.mutation.Tiktok(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldTiktok,
		})
	}
	if value, ok := smuo.mutation.Twitter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: socialmedia.FieldTwitter,
		})
	}
	if smuo.mutation.RegisterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   socialmedia.RegisterTable,
			Columns: socialmedia.RegisterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: register.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smuo.mutation.RemovedRegisterIDs(); len(nodes) > 0 && !smuo.mutation.RegisterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   socialmedia.RegisterTable,
			Columns: socialmedia.RegisterPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smuo.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   socialmedia.RegisterTable,
			Columns: socialmedia.RegisterPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SocialMedia{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{socialmedia.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
