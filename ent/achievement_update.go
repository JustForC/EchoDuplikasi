// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/achievement"
	"Kynesia/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AchievementUpdate is the builder for updating Achievement entities.
type AchievementUpdate struct {
	config
	hooks    []Hook
	mutation *AchievementMutation
}

// Where appends a list predicates to the AchievementUpdate builder.
func (au *AchievementUpdate) Where(ps ...predicate.Achievement) *AchievementUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AchievementUpdate) SetName(s string) *AchievementUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AchievementUpdate) SetNillableName(s *string) *AchievementUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// ClearName clears the value of the "name" field.
func (au *AchievementUpdate) ClearName() *AchievementUpdate {
	au.mutation.ClearName()
	return au
}

// SetOrganizer sets the "organizer" field.
func (au *AchievementUpdate) SetOrganizer(s string) *AchievementUpdate {
	au.mutation.SetOrganizer(s)
	return au
}

// SetNillableOrganizer sets the "organizer" field if the given value is not nil.
func (au *AchievementUpdate) SetNillableOrganizer(s *string) *AchievementUpdate {
	if s != nil {
		au.SetOrganizer(*s)
	}
	return au
}

// ClearOrganizer clears the value of the "organizer" field.
func (au *AchievementUpdate) ClearOrganizer() *AchievementUpdate {
	au.mutation.ClearOrganizer()
	return au
}

// SetLevel sets the "level" field.
func (au *AchievementUpdate) SetLevel(s string) *AchievementUpdate {
	au.mutation.SetLevel(s)
	return au
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (au *AchievementUpdate) SetNillableLevel(s *string) *AchievementUpdate {
	if s != nil {
		au.SetLevel(*s)
	}
	return au
}

// ClearLevel clears the value of the "level" field.
func (au *AchievementUpdate) ClearLevel() *AchievementUpdate {
	au.mutation.ClearLevel()
	return au
}

// Mutation returns the AchievementMutation object of the builder.
func (au *AchievementUpdate) Mutation() *AchievementMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AchievementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AchievementUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AchievementUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AchievementUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AchievementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   achievement.Table,
			Columns: achievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: achievement.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldName,
		})
	}
	if au.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: achievement.FieldName,
		})
	}
	if value, ok := au.mutation.Organizer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldOrganizer,
		})
	}
	if au.mutation.OrganizerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: achievement.FieldOrganizer,
		})
	}
	if value, ok := au.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldLevel,
		})
	}
	if au.mutation.LevelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: achievement.FieldLevel,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AchievementUpdateOne is the builder for updating a single Achievement entity.
type AchievementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AchievementMutation
}

// SetName sets the "name" field.
func (auo *AchievementUpdateOne) SetName(s string) *AchievementUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableName(s *string) *AchievementUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// ClearName clears the value of the "name" field.
func (auo *AchievementUpdateOne) ClearName() *AchievementUpdateOne {
	auo.mutation.ClearName()
	return auo
}

// SetOrganizer sets the "organizer" field.
func (auo *AchievementUpdateOne) SetOrganizer(s string) *AchievementUpdateOne {
	auo.mutation.SetOrganizer(s)
	return auo
}

// SetNillableOrganizer sets the "organizer" field if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableOrganizer(s *string) *AchievementUpdateOne {
	if s != nil {
		auo.SetOrganizer(*s)
	}
	return auo
}

// ClearOrganizer clears the value of the "organizer" field.
func (auo *AchievementUpdateOne) ClearOrganizer() *AchievementUpdateOne {
	auo.mutation.ClearOrganizer()
	return auo
}

// SetLevel sets the "level" field.
func (auo *AchievementUpdateOne) SetLevel(s string) *AchievementUpdateOne {
	auo.mutation.SetLevel(s)
	return auo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (auo *AchievementUpdateOne) SetNillableLevel(s *string) *AchievementUpdateOne {
	if s != nil {
		auo.SetLevel(*s)
	}
	return auo
}

// ClearLevel clears the value of the "level" field.
func (auo *AchievementUpdateOne) ClearLevel() *AchievementUpdateOne {
	auo.mutation.ClearLevel()
	return auo
}

// Mutation returns the AchievementMutation object of the builder.
func (auo *AchievementUpdateOne) Mutation() *AchievementMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AchievementUpdateOne) Select(field string, fields ...string) *AchievementUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Achievement entity.
func (auo *AchievementUpdateOne) Save(ctx context.Context) (*Achievement, error) {
	var (
		err  error
		node *Achievement
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AchievementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AchievementUpdateOne) SaveX(ctx context.Context) *Achievement {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AchievementUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AchievementUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AchievementUpdateOne) sqlSave(ctx context.Context) (_node *Achievement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   achievement.Table,
			Columns: achievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: achievement.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Achievement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievement.FieldID)
		for _, f := range fields {
			if !achievement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != achievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldName,
		})
	}
	if auo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: achievement.FieldName,
		})
	}
	if value, ok := auo.mutation.Organizer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldOrganizer,
		})
	}
	if auo.mutation.OrganizerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: achievement.FieldOrganizer,
		})
	}
	if value, ok := auo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: achievement.FieldLevel,
		})
	}
	if auo.mutation.LevelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: achievement.FieldLevel,
		})
	}
	_node = &Achievement{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
