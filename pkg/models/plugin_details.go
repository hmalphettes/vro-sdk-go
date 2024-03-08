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

// PluginDetails plugin details
//
// swagger:model plugin-details
type PluginDetails struct {

	// build
	Build int64 `json:"build,omitempty" xml:"build,attr,omitempty"`

	// description
	Description string `json:"description,omitempty" xml:"description,attr,omitempty"`

	// enums
	Enums []*EnumType `json:"enums"`

	// external Url
	ExternalURL string `json:"externalUrl,omitempty" xml:"externalUrl,attr,omitempty"`

	// icon
	Icon string `json:"icon,omitempty" xml:"icon,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// objects
	Objects []*ObjectType `json:"objects"`

	// state
	// Enum: [RELEASE ALPHA BETA DEPRECATED]
	State string `json:"state,omitempty" xml:"state,attr,omitempty"`

	// types
	Types []*SdkType `json:"types"`

	// version
	Version string `json:"version,omitempty" xml:"version,attr,omitempty"`
}

// Validate validates this plugin details
func (m *PluginDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnums(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginDetails) validateEnums(formats strfmt.Registry) error {
	if swag.IsZero(m.Enums) { // not required
		return nil
	}

	for i := 0; i < len(m.Enums); i++ {
		if swag.IsZero(m.Enums[i]) { // not required
			continue
		}

		if m.Enums[i] != nil {
			if err := m.Enums[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enums" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enums" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginDetails) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var pluginDetailsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RELEASE","ALPHA","BETA","DEPRECATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pluginDetailsTypeStatePropEnum = append(pluginDetailsTypeStatePropEnum, v)
	}
}

const (

	// PluginDetailsStateRELEASE captures enum value "RELEASE"
	PluginDetailsStateRELEASE string = "RELEASE"

	// PluginDetailsStateALPHA captures enum value "ALPHA"
	PluginDetailsStateALPHA string = "ALPHA"

	// PluginDetailsStateBETA captures enum value "BETA"
	PluginDetailsStateBETA string = "BETA"

	// PluginDetailsStateDEPRECATED captures enum value "DEPRECATED"
	PluginDetailsStateDEPRECATED string = "DEPRECATED"
)

// prop value enum
func (m *PluginDetails) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pluginDetailsTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PluginDetails) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *PluginDetails) validateTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.Types) { // not required
		return nil
	}

	for i := 0; i < len(m.Types); i++ {
		if swag.IsZero(m.Types[i]) { // not required
			continue
		}

		if m.Types[i] != nil {
			if err := m.Types[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugin details based on the context it is used
func (m *PluginDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnums(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginDetails) contextValidateEnums(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Enums); i++ {

		if m.Enums[i] != nil {

			if swag.IsZero(m.Enums[i]) { // not required
				return nil
			}

			if err := m.Enums[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enums" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enums" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginDetails) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginDetails) contextValidateTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Types); i++ {

		if m.Types[i] != nil {

			if swag.IsZero(m.Types[i]) { // not required
				return nil
			}

			if err := m.Types[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginDetails) UnmarshalBinary(b []byte) error {
	var res PluginDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}