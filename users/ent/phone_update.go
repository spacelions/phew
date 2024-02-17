// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"users/ent/phone"
	"users/ent/predicate"
	"users/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PhoneUpdate is the builder for updating Phone entities.
type PhoneUpdate struct {
	config
	hooks    []Hook
	mutation *PhoneMutation
}

// Where appends a list predicates to the PhoneUpdate builder.
func (pu *PhoneUpdate) Where(ps ...predicate.Phone) *PhoneUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PhoneUpdate) SetUpdateTime(t time.Time) *PhoneUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetCountryCode sets the "country_code" field.
func (pu *PhoneUpdate) SetCountryCode(s string) *PhoneUpdate {
	pu.mutation.SetCountryCode(s)
	return pu
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (pu *PhoneUpdate) SetNillableCountryCode(s *string) *PhoneUpdate {
	if s != nil {
		pu.SetCountryCode(*s)
	}
	return pu
}

// SetNumber sets the "number" field.
func (pu *PhoneUpdate) SetNumber(s string) *PhoneUpdate {
	pu.mutation.SetNumber(s)
	return pu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (pu *PhoneUpdate) SetNillableNumber(s *string) *PhoneUpdate {
	if s != nil {
		pu.SetNumber(*s)
	}
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PhoneUpdate) SetUserID(id int) *PhoneUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *PhoneUpdate) SetNillableUserID(id *int) *PhoneUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PhoneUpdate) SetUser(u *User) *PhoneUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (pu *PhoneUpdate) Mutation() *PhoneMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PhoneUpdate) ClearUser() *PhoneUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PhoneUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PhoneUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PhoneUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PhoneUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PhoneUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := phone.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PhoneUpdate) check() error {
	if v, ok := pu.mutation.CountryCode(); ok {
		if err := phone.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`ent: validator failed for field "Phone.country_code": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Number(); ok {
		if err := phone.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Phone.number": %w`, err)}
		}
	}
	return nil
}

func (pu *PhoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(phone.Table, phone.Columns, sqlgraph.NewFieldSpec(phone.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(phone.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.CountryCode(); ok {
		_spec.SetField(phone.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.Number(); ok {
		_spec.SetField(phone.FieldNumber, field.TypeString, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.UserTable,
			Columns: []string{phone.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.UserTable,
			Columns: []string{phone.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PhoneUpdateOne is the builder for updating a single Phone entity.
type PhoneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PhoneMutation
}

// SetUpdateTime sets the "update_time" field.
func (puo *PhoneUpdateOne) SetUpdateTime(t time.Time) *PhoneUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetCountryCode sets the "country_code" field.
func (puo *PhoneUpdateOne) SetCountryCode(s string) *PhoneUpdateOne {
	puo.mutation.SetCountryCode(s)
	return puo
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (puo *PhoneUpdateOne) SetNillableCountryCode(s *string) *PhoneUpdateOne {
	if s != nil {
		puo.SetCountryCode(*s)
	}
	return puo
}

// SetNumber sets the "number" field.
func (puo *PhoneUpdateOne) SetNumber(s string) *PhoneUpdateOne {
	puo.mutation.SetNumber(s)
	return puo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (puo *PhoneUpdateOne) SetNillableNumber(s *string) *PhoneUpdateOne {
	if s != nil {
		puo.SetNumber(*s)
	}
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PhoneUpdateOne) SetUserID(id int) *PhoneUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *PhoneUpdateOne) SetNillableUserID(id *int) *PhoneUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PhoneUpdateOne) SetUser(u *User) *PhoneUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (puo *PhoneUpdateOne) Mutation() *PhoneMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PhoneUpdateOne) ClearUser() *PhoneUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PhoneUpdate builder.
func (puo *PhoneUpdateOne) Where(ps ...predicate.Phone) *PhoneUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PhoneUpdateOne) Select(field string, fields ...string) *PhoneUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Phone entity.
func (puo *PhoneUpdateOne) Save(ctx context.Context) (*Phone, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PhoneUpdateOne) SaveX(ctx context.Context) *Phone {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PhoneUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PhoneUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PhoneUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := phone.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PhoneUpdateOne) check() error {
	if v, ok := puo.mutation.CountryCode(); ok {
		if err := phone.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`ent: validator failed for field "Phone.country_code": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Number(); ok {
		if err := phone.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Phone.number": %w`, err)}
		}
	}
	return nil
}

func (puo *PhoneUpdateOne) sqlSave(ctx context.Context) (_node *Phone, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(phone.Table, phone.Columns, sqlgraph.NewFieldSpec(phone.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Phone.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phone.FieldID)
		for _, f := range fields {
			if !phone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != phone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(phone.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.CountryCode(); ok {
		_spec.SetField(phone.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.Number(); ok {
		_spec.SetField(phone.FieldNumber, field.TypeString, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.UserTable,
			Columns: []string{phone.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.UserTable,
			Columns: []string{phone.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Phone{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}