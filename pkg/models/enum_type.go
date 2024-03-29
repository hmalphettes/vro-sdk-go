// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EnumType enum type
//
// swagger:model EnumType
type EnumType struct {

	// description
	Description string `json:"description,omitempty" xml:"description,attr,omitempty"`

	// external Url
	ExternalURL string `json:"externalUrl,omitempty" xml:"externalUrl,attr,omitempty"`

	// icon
	Icon string `json:"icon,omitempty" xml:"icon,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// state
	// Enum: [RELEASE ALPHA BETA DEPRECATED]
	State string `json:"state,omitempty" xml:"state,attr,omitempty"`

	// values
	Values []*BaseType `json:"values"`
}

// Validate validates this enum type
func (m *EnumType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var enumTypeTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RELEASE","ALPHA","BETA","DEPRECATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enumTypeTypeStatePropEnum = append(enumTypeTypeStatePropEnum, v)
	}
}

const (

	// EnumTypeStateRELEASE captures enum value "RELEASE"
	EnumTypeStateRELEASE string = "RELEASE"

	// EnumTypeStateALPHA captures enum value "ALPHA"
	EnumTypeStateALPHA string = "ALPHA"

	// EnumTypeStateBETA captures enum value "BETA"
	EnumTypeStateBETA string = "BETA"

	// EnumTypeStateDEPRECATED captures enum value "DEPRECATED"
	EnumTypeStateDEPRECATED string = "DEPRECATED"
)

// prop value enum
func (m *EnumType) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, enumTypeTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EnumType) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *EnumType) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this enum type based on the context it is used
func (m *EnumType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnumType) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {

			if swag.IsZero(m.Values[i]) { // not required
				return nil
			}

			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnumType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnumType) UnmarshalBinary(b []byte) error {
	var res EnumType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
