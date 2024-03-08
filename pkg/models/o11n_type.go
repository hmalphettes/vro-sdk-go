// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// O11nType o11n type
//
// swagger:model o11nType
type O11nType struct {

	// attr
	Attr *bool `json:"attr,omitempty" xml:"attr,attr,omitempty"`

	// in
	In *bool `json:"in,omitempty" xml:"in,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// out
	Out *bool `json:"out,omitempty" xml:"out,attr,omitempty"`

	// ret
	Ret *bool `json:"ret,omitempty" xml:"ret,attr,omitempty"`
}

// Validate validates this o11n type
func (m *O11nType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o11n type based on context it is used
func (m *O11nType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *O11nType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O11nType) UnmarshalBinary(b []byte) error {
	var res O11nType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
