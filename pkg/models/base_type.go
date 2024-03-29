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

// BaseType base type
//
// swagger:model BaseType
type BaseType struct {

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
}

// Validate validates this base type
func (m *BaseType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var baseTypeTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RELEASE","ALPHA","BETA","DEPRECATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		baseTypeTypeStatePropEnum = append(baseTypeTypeStatePropEnum, v)
	}
}

const (

	// BaseTypeStateRELEASE captures enum value "RELEASE"
	BaseTypeStateRELEASE string = "RELEASE"

	// BaseTypeStateALPHA captures enum value "ALPHA"
	BaseTypeStateALPHA string = "ALPHA"

	// BaseTypeStateBETA captures enum value "BETA"
	BaseTypeStateBETA string = "BETA"

	// BaseTypeStateDEPRECATED captures enum value "DEPRECATED"
	BaseTypeStateDEPRECATED string = "DEPRECATED"
)

// prop value enum
func (m *BaseType) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, baseTypeTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BaseType) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this base type based on context it is used
func (m *BaseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseType) UnmarshalBinary(b []byte) error {
	var res BaseType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
