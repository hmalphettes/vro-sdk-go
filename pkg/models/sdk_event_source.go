// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SdkEventSource sdk event source
//
// swagger:model sdk-event-source
type SdkEventSource struct {
	WsPolicyEventSource

	// id
	ID string `json:"id,omitempty"`

	// sdk type
	SdkType string `json:"sdk-type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SdkEventSource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WsPolicyEventSource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WsPolicyEventSource = aO0

	// AO1
	var dataAO1 struct {
		ID string `json:"id,omitempty"`

		SdkType string `json:"sdk-type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.SdkType = dataAO1.SdkType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SdkEventSource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WsPolicyEventSource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID string `json:"id,omitempty"`

		SdkType string `json:"sdk-type,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.SdkType = m.SdkType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sdk event source
func (m *SdkEventSource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WsPolicyEventSource

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sdk event source based on context it is used
func (m *SdkEventSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SdkEventSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdkEventSource) UnmarshalBinary(b []byte) error {
	var res SdkEventSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}