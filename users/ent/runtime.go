// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"users/ent/phone"
	"users/ent/schema"
	"users/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	phoneMixin := schema.Phone{}.Mixin()
	phoneMixinFields0 := phoneMixin[0].Fields()
	_ = phoneMixinFields0
	phoneFields := schema.Phone{}.Fields()
	_ = phoneFields
	// phoneDescCreateTime is the schema descriptor for create_time field.
	phoneDescCreateTime := phoneMixinFields0[0].Descriptor()
	// phone.DefaultCreateTime holds the default value on creation for the create_time field.
	phone.DefaultCreateTime = phoneDescCreateTime.Default.(func() time.Time)
	// phoneDescUpdateTime is the schema descriptor for update_time field.
	phoneDescUpdateTime := phoneMixinFields0[1].Descriptor()
	// phone.DefaultUpdateTime holds the default value on creation for the update_time field.
	phone.DefaultUpdateTime = phoneDescUpdateTime.Default.(func() time.Time)
	// phone.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	phone.UpdateDefaultUpdateTime = phoneDescUpdateTime.UpdateDefault.(func() time.Time)
	// phoneDescCountryCode is the schema descriptor for country_code field.
	phoneDescCountryCode := phoneFields[0].Descriptor()
	// phone.CountryCodeValidator is a validator for the "country_code" field. It is called by the builders before save.
	phone.CountryCodeValidator = phoneDescCountryCode.Validators[0].(func(string) error)
	// phoneDescNumber is the schema descriptor for number field.
	phoneDescNumber := phoneFields[1].Descriptor()
	// phone.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	phone.NumberValidator = phoneDescNumber.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
}