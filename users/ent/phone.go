// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/spacelions/phew/users/ent/phone"
	"github.com/spacelions/phew/users/ent/user"
)

// Phone is the model entity for the Phone schema.
type Phone struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhoneQuery when eager-loading is set.
	Edges        PhoneEdges `json:"edges"`
	phone_user   *int
	selectValues sql.SelectValues
}

// PhoneEdges holds the relations/edges for other nodes in the graph.
type PhoneEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhoneEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Phone) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case phone.FieldID:
			values[i] = new(sql.NullInt64)
		case phone.FieldNumber, phone.FieldCountryCode:
			values[i] = new(sql.NullString)
		case phone.ForeignKeys[0]: // phone_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Phone fields.
func (ph *Phone) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case phone.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ph.ID = int(value.Int64)
		case phone.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				ph.Number = value.String
			}
		case phone.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				ph.CountryCode = value.String
			}
		case phone.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field phone_user", value)
			} else if value.Valid {
				ph.phone_user = new(int)
				*ph.phone_user = int(value.Int64)
			}
		default:
			ph.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Phone.
// This includes values selected through modifiers, order, etc.
func (ph *Phone) Value(name string) (ent.Value, error) {
	return ph.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Phone entity.
func (ph *Phone) QueryUser() *UserQuery {
	return NewPhoneClient(ph.config).QueryUser(ph)
}

// Update returns a builder for updating this Phone.
// Note that you need to call Phone.Unwrap() before calling this method if this Phone
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *Phone) Update() *PhoneUpdateOne {
	return NewPhoneClient(ph.config).UpdateOne(ph)
}

// Unwrap unwraps the Phone entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *Phone) Unwrap() *Phone {
	_tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: Phone is not a transactional entity")
	}
	ph.config.driver = _tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *Phone) String() string {
	var builder strings.Builder
	builder.WriteString("Phone(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ph.ID))
	builder.WriteString("number=")
	builder.WriteString(ph.Number)
	builder.WriteString(", ")
	builder.WriteString("country_code=")
	builder.WriteString(ph.CountryCode)
	builder.WriteByte(')')
	return builder.String()
}

// Phones is a parsable slice of Phone.
type Phones []*Phone
