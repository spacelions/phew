// Code generated by ent, DO NOT EDIT.

package phone

import (
	"time"
	"users/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Phone {
	return predicate.Phone(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Phone {
	return predicate.Phone(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Phone {
	return predicate.Phone(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Phone {
	return predicate.Phone(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Phone {
	return predicate.Phone(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Phone {
	return predicate.Phone(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Phone {
	return predicate.Phone(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldUpdateTime, v))
}

// CountryCode applies equality check predicate on the "country_code" field. It's identical to CountryCodeEQ.
func CountryCode(v string) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldCountryCode, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldNumber, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Phone {
	return predicate.Phone(sql.FieldLTE(FieldUpdateTime, v))
}

// CountryCodeEQ applies the EQ predicate on the "country_code" field.
func CountryCodeEQ(v string) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldCountryCode, v))
}

// CountryCodeNEQ applies the NEQ predicate on the "country_code" field.
func CountryCodeNEQ(v string) predicate.Phone {
	return predicate.Phone(sql.FieldNEQ(FieldCountryCode, v))
}

// CountryCodeIn applies the In predicate on the "country_code" field.
func CountryCodeIn(vs ...string) predicate.Phone {
	return predicate.Phone(sql.FieldIn(FieldCountryCode, vs...))
}

// CountryCodeNotIn applies the NotIn predicate on the "country_code" field.
func CountryCodeNotIn(vs ...string) predicate.Phone {
	return predicate.Phone(sql.FieldNotIn(FieldCountryCode, vs...))
}

// CountryCodeGT applies the GT predicate on the "country_code" field.
func CountryCodeGT(v string) predicate.Phone {
	return predicate.Phone(sql.FieldGT(FieldCountryCode, v))
}

// CountryCodeGTE applies the GTE predicate on the "country_code" field.
func CountryCodeGTE(v string) predicate.Phone {
	return predicate.Phone(sql.FieldGTE(FieldCountryCode, v))
}

// CountryCodeLT applies the LT predicate on the "country_code" field.
func CountryCodeLT(v string) predicate.Phone {
	return predicate.Phone(sql.FieldLT(FieldCountryCode, v))
}

// CountryCodeLTE applies the LTE predicate on the "country_code" field.
func CountryCodeLTE(v string) predicate.Phone {
	return predicate.Phone(sql.FieldLTE(FieldCountryCode, v))
}

// CountryCodeContains applies the Contains predicate on the "country_code" field.
func CountryCodeContains(v string) predicate.Phone {
	return predicate.Phone(sql.FieldContains(FieldCountryCode, v))
}

// CountryCodeHasPrefix applies the HasPrefix predicate on the "country_code" field.
func CountryCodeHasPrefix(v string) predicate.Phone {
	return predicate.Phone(sql.FieldHasPrefix(FieldCountryCode, v))
}

// CountryCodeHasSuffix applies the HasSuffix predicate on the "country_code" field.
func CountryCodeHasSuffix(v string) predicate.Phone {
	return predicate.Phone(sql.FieldHasSuffix(FieldCountryCode, v))
}

// CountryCodeEqualFold applies the EqualFold predicate on the "country_code" field.
func CountryCodeEqualFold(v string) predicate.Phone {
	return predicate.Phone(sql.FieldEqualFold(FieldCountryCode, v))
}

// CountryCodeContainsFold applies the ContainsFold predicate on the "country_code" field.
func CountryCodeContainsFold(v string) predicate.Phone {
	return predicate.Phone(sql.FieldContainsFold(FieldCountryCode, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Phone {
	return predicate.Phone(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Phone {
	return predicate.Phone(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Phone {
	return predicate.Phone(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Phone {
	return predicate.Phone(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Phone {
	return predicate.Phone(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Phone {
	return predicate.Phone(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Phone {
	return predicate.Phone(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Phone {
	return predicate.Phone(sql.FieldLTE(FieldNumber, v))
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Phone {
	return predicate.Phone(sql.FieldContains(FieldNumber, v))
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Phone {
	return predicate.Phone(sql.FieldHasPrefix(FieldNumber, v))
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Phone {
	return predicate.Phone(sql.FieldHasSuffix(FieldNumber, v))
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Phone {
	return predicate.Phone(sql.FieldEqualFold(FieldNumber, v))
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Phone {
	return predicate.Phone(sql.FieldContainsFold(FieldNumber, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Phone {
	return predicate.Phone(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Phone {
	return predicate.Phone(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Phone) predicate.Phone {
	return predicate.Phone(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Phone) predicate.Phone {
	return predicate.Phone(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Phone) predicate.Phone {
	return predicate.Phone(sql.NotPredicates(p))
}
