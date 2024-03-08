// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigEntry config entry
//
// swagger:model configEntry
type ConfigEntry struct {

	// data
	Data string `json:"data,omitempty"`

	// dependencies
	Dependencies string `json:"dependencies,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,attr,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// plugin
	Plugin *bool `json:"plugin,omitempty" xml:"plugin,attr,omitempty"`

	// remote
	Remote *bool `json:"remote,omitempty" xml:"remote,attr,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// tooltip
	Tooltip string `json:"tooltip,omitempty"`
}

// Validate validates this config entry
func (m *ConfigEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config entry based on context it is used
func (m *ConfigEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigEntry) UnmarshalBinary(b []byte) error {
	var res ConfigEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
