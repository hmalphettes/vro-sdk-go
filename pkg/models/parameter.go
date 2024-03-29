// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Parameter parameter
//
// swagger:model parameter
type Parameter struct {

	// description
	Description string `json:"description,omitempty" xml:"description,attr,omitempty"`

	// encrypt value
	EncryptValue *bool `json:"encrypt-value,omitempty" xml:"encrypt-value,attr,omitempty"`

	// name
	// Required: true
	// Pattern: ^[\p{L}_$][\p{L}0-9_$]*$
	Name *string `json:"name" xml:"name,attr"`

	// scope
	// Enum: [LOCAL TOKEN]
	Scope string `json:"scope,omitempty" xml:"scope,attr,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`

	// updated
	Updated *bool `json:"updated,omitempty" xml:"updated,attr,omitempty"`

	// value
	Value *Value `json:"value,omitempty"`
}

// Validate validates this parameter
func (m *Parameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Parameter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[\p{L}_$][\p{L}0-9_$]*$`); err != nil {
		return err
	}

	return nil
}

var parameterTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LOCAL","TOKEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		parameterTypeScopePropEnum = append(parameterTypeScopePropEnum, v)
	}
}

const (

	// ParameterScopeLOCAL captures enum value "LOCAL"
	ParameterScopeLOCAL string = "LOCAL"

	// ParameterScopeTOKEN captures enum value "TOKEN"
	ParameterScopeTOKEN string = "TOKEN"
)

// prop value enum
func (m *Parameter) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, parameterTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Parameter) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *Parameter) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this parameter based on the context it is used
func (m *Parameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Parameter) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Parameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Parameter) UnmarshalBinary(b []byte) error {
	var res Parameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
