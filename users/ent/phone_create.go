// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/spacelions/phew/users/ent/phone"
	"github.com/spacelions/phew/users/ent/user"
)

// PhoneCreate is the builder for creating a Phone entity.
type PhoneCreate struct {
	config
	mutation *PhoneMutation
	hooks    []Hook
}

// SetNumber sets the "number" field.
func (pc *PhoneCreate) SetNumber(s string) *PhoneCreate {
	pc.mutation.SetNumber(s)
	return pc
}

// SetCountryCode sets the "country_code" field.
func (pc *PhoneCreate) SetCountryCode(s string) *PhoneCreate {
	pc.mutation.SetCountryCode(s)
	return pc
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (pc *PhoneCreate) SetNillableCountryCode(s *string) *PhoneCreate {
	if s != nil {
		pc.SetCountryCode(*s)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PhoneCreate) SetUserID(id int) *PhoneCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PhoneCreate) SetNillableUserID(id *int) *PhoneCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PhoneCreate) SetUser(u *User) *PhoneCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (pc *PhoneCreate) Mutation() *PhoneMutation {
	return pc.mutation
}

// Save creates the Phone in the database.
func (pc *PhoneCreate) Save(ctx context.Context) (*Phone, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PhoneCreate) SaveX(ctx context.Context) *Phone {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PhoneCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PhoneCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PhoneCreate) defaults() {
	if _, ok := pc.mutation.CountryCode(); !ok {
		v := phone.DefaultCountryCode
		pc.mutation.SetCountryCode(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PhoneCreate) check() error {
	if _, ok := pc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Phone.number"`)}
	}
	if v, ok := pc.mutation.Number(); ok {
		if err := phone.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Phone.number": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`ent: missing required field "Phone.country_code"`)}
	}
	if v, ok := pc.mutation.CountryCode(); ok {
		if err := phone.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`ent: validator failed for field "Phone.country_code": %w`, err)}
		}
	}
	return nil
}

func (pc *PhoneCreate) sqlSave(ctx context.Context) (*Phone, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PhoneCreate) createSpec() (*Phone, *sqlgraph.CreateSpec) {
	var (
		_node = &Phone{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(phone.Table, sqlgraph.NewFieldSpec(phone.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Number(); ok {
		_spec.SetField(phone.FieldNumber, field.TypeString, value)
		_node.Number = value
	}
	if value, ok := pc.mutation.CountryCode(); ok {
		_spec.SetField(phone.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.phone_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PhoneCreateBulk is the builder for creating many Phone entities in bulk.
type PhoneCreateBulk struct {
	config
	err      error
	builders []*PhoneCreate
}

// Save creates the Phone entities in the database.
func (pcb *PhoneCreateBulk) Save(ctx context.Context) ([]*Phone, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Phone, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PhoneMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PhoneCreateBulk) SaveX(ctx context.Context) []*Phone {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PhoneCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PhoneCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
