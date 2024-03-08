// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BaseDecorator base decorator
//
// swagger:model BaseDecorator
type BaseDecorator struct {

	// kind
	Kind string `json:"kind,omitempty" xml:"kind,attr,omitempty"`
}

// Validate validates this base decorator
func (m *BaseDecorator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base decorator based on context it is used
func (m *BaseDecorator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseDecorator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseDecorator) UnmarshalBinary(b []byte) error {
	var res BaseDecorator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
