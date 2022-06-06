// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/language"
	"Kynesia/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LanguageUpdate is the builder for updating Language entities.
type LanguageUpdate struct {
	config
	hooks    []Hook
	mutation *LanguageMutation
}

// Where appends a list predicates to the LanguageUpdate builder.
func (lu *LanguageUpdate) Where(ps ...predicate.Language) *LanguageUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *LanguageUpdate) SetName(s string) *LanguageUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableName(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// ClearName clears the value of the "name" field.
func (lu *LanguageUpdate) ClearName() *LanguageUpdate {
	lu.mutation.ClearName()
	return lu
}

// SetTalk sets the "talk" field.
func (lu *LanguageUpdate) SetTalk(s string) *LanguageUpdate {
	lu.mutation.SetTalk(s)
	return lu
}

// SetNillableTalk sets the "talk" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableTalk(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetTalk(*s)
	}
	return lu
}

// ClearTalk clears the value of the "talk" field.
func (lu *LanguageUpdate) ClearTalk() *LanguageUpdate {
	lu.mutation.ClearTalk()
	return lu
}

// SetWrite sets the "write" field.
func (lu *LanguageUpdate) SetWrite(s string) *LanguageUpdate {
	lu.mutation.SetWrite(s)
	return lu
}

// SetNillableWrite sets the "write" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableWrite(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetWrite(*s)
	}
	return lu
}

// ClearWrite clears the value of the "write" field.
func (lu *LanguageUpdate) ClearWrite() *LanguageUpdate {
	lu.mutation.ClearWrite()
	return lu
}

// SetRead sets the "read" field.
func (lu *LanguageUpdate) SetRead(s string) *LanguageUpdate {
	lu.mutation.SetRead(s)
	return lu
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableRead(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetRead(*s)
	}
	return lu
}

// ClearRead clears the value of the "read" field.
func (lu *LanguageUpdate) ClearRead() *LanguageUpdate {
	lu.mutation.ClearRead()
	return lu
}

// SetListen sets the "listen" field.
func (lu *LanguageUpdate) SetListen(s string) *LanguageUpdate {
	lu.mutation.SetListen(s)
	return lu
}

// SetNillableListen sets the "listen" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableListen(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetListen(*s)
	}
	return lu
}

// ClearListen clears the value of the "listen" field.
func (lu *LanguageUpdate) ClearListen() *LanguageUpdate {
	lu.mutation.ClearListen()
	return lu
}

// Mutation returns the LanguageMutation object of the builder.
func (lu *LanguageUpdate) Mutation() *LanguageMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LanguageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LanguageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LanguageUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LanguageUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LanguageUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LanguageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   language.Table,
			Columns: language.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: language.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldName,
		})
	}
	if lu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldName,
		})
	}
	if value, ok := lu.mutation.Talk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldTalk,
		})
	}
	if lu.mutation.TalkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldTalk,
		})
	}
	if value, ok := lu.mutation.Write(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldWrite,
		})
	}
	if lu.mutation.WriteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldWrite,
		})
	}
	if value, ok := lu.mutation.Read(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldRead,
		})
	}
	if lu.mutation.ReadCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldRead,
		})
	}
	if value, ok := lu.mutation.Listen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldListen,
		})
	}
	if lu.mutation.ListenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldListen,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LanguageUpdateOne is the builder for updating a single Language entity.
type LanguageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LanguageMutation
}

// SetName sets the "name" field.
func (luo *LanguageUpdateOne) SetName(s string) *LanguageUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableName(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// ClearName clears the value of the "name" field.
func (luo *LanguageUpdateOne) ClearName() *LanguageUpdateOne {
	luo.mutation.ClearName()
	return luo
}

// SetTalk sets the "talk" field.
func (luo *LanguageUpdateOne) SetTalk(s string) *LanguageUpdateOne {
	luo.mutation.SetTalk(s)
	return luo
}

// SetNillableTalk sets the "talk" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableTalk(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetTalk(*s)
	}
	return luo
}

// ClearTalk clears the value of the "talk" field.
func (luo *LanguageUpdateOne) ClearTalk() *LanguageUpdateOne {
	luo.mutation.ClearTalk()
	return luo
}

// SetWrite sets the "write" field.
func (luo *LanguageUpdateOne) SetWrite(s string) *LanguageUpdateOne {
	luo.mutation.SetWrite(s)
	return luo
}

// SetNillableWrite sets the "write" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableWrite(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetWrite(*s)
	}
	return luo
}

// ClearWrite clears the value of the "write" field.
func (luo *LanguageUpdateOne) ClearWrite() *LanguageUpdateOne {
	luo.mutation.ClearWrite()
	return luo
}

// SetRead sets the "read" field.
func (luo *LanguageUpdateOne) SetRead(s string) *LanguageUpdateOne {
	luo.mutation.SetRead(s)
	return luo
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableRead(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetRead(*s)
	}
	return luo
}

// ClearRead clears the value of the "read" field.
func (luo *LanguageUpdateOne) ClearRead() *LanguageUpdateOne {
	luo.mutation.ClearRead()
	return luo
}

// SetListen sets the "listen" field.
func (luo *LanguageUpdateOne) SetListen(s string) *LanguageUpdateOne {
	luo.mutation.SetListen(s)
	return luo
}

// SetNillableListen sets the "listen" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableListen(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetListen(*s)
	}
	return luo
}

// ClearListen clears the value of the "listen" field.
func (luo *LanguageUpdateOne) ClearListen() *LanguageUpdateOne {
	luo.mutation.ClearListen()
	return luo
}

// Mutation returns the LanguageMutation object of the builder.
func (luo *LanguageUpdateOne) Mutation() *LanguageMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LanguageUpdateOne) Select(field string, fields ...string) *LanguageUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Language entity.
func (luo *LanguageUpdateOne) Save(ctx context.Context) (*Language, error) {
	var (
		err  error
		node *Language
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LanguageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LanguageUpdateOne) SaveX(ctx context.Context) *Language {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LanguageUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LanguageUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LanguageUpdateOne) sqlSave(ctx context.Context) (_node *Language, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   language.Table,
			Columns: language.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: language.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Language.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, language.FieldID)
		for _, f := range fields {
			if !language.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != language.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldName,
		})
	}
	if luo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldName,
		})
	}
	if value, ok := luo.mutation.Talk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldTalk,
		})
	}
	if luo.mutation.TalkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldTalk,
		})
	}
	if value, ok := luo.mutation.Write(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldWrite,
		})
	}
	if luo.mutation.WriteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldWrite,
		})
	}
	if value, ok := luo.mutation.Read(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldRead,
		})
	}
	if luo.mutation.ReadCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldRead,
		})
	}
	if value, ok := luo.mutation.Listen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: language.FieldListen,
		})
	}
	if luo.mutation.ListenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: language.FieldListen,
		})
	}
	_node = &Language{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}