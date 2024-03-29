// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Bind bind
//
// swagger:model Bind
type Bind struct {

	// description
	Description string `json:"description,omitempty"`

	// export name
	ExportName string `json:"export-name,omitempty" xml:"export-name,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`
}

// Validate validates this bind
func (m *Bind) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bind based on context it is used
func (m *Bind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Bind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bind) UnmarshalBinary(b []byte) error {
	var res Bind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
