// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WsValidationMessage ws validation message
//
// swagger:model WsValidationMessage
type WsValidationMessage struct {

	// owner
	Owner string `json:"owner,omitempty" xml:"owner,attr,omitempty"`

	// severity
	Severity string `json:"severity,omitempty" xml:"severity,attr,omitempty"`

	// text
	Text string `json:"text,omitempty" xml:"text,attr,omitempty"`

	// title
	Title string `json:"title,omitempty" xml:"title,attr,omitempty"`
}

// Validate validates this ws validation message
func (m *WsValidationMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ws validation message based on context it is used
func (m *WsValidationMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WsValidationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WsValidationMessage) UnmarshalBinary(b []byte) error {
	var res WsValidationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
