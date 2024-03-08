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

// ObjectType object type
//
// swagger:model ObjectType
type ObjectType struct {

	// attributes
	Attributes []*AttributeType `json:"attributes"`

	// constructors
	Constructors []*MethodType `json:"constructors"`

	// description
	Description string `json:"description,omitempty" xml:"description,attr,omitempty"`

	// events
	Events []*BaseType `json:"events"`

	// external Url
	ExternalURL string `json:"externalUrl,omitempty" xml:"externalUrl,attr,omitempty"`

	// icon
	Icon string `json:"icon,omitempty" xml:"icon,attr,omitempty"`

	// is namespace
	IsNamespace *bool `json:"isNamespace,omitempty" xml:"isNamespace,attr,omitempty"`

	// members
	Members []*MemberType `json:"members"`

	// methods
	Methods []*MethodType `json:"methods"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// state
	// Enum: [RELEASE ALPHA BETA DEPRECATED]
	State string `json:"state,omitempty" xml:"state,attr,omitempty"`
}

// Validate validates this object type
func (m *ObjectType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstructors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectType) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) validateConstructors(formats strfmt.Registry) error {
	if swag.IsZero(m.Constructors) { // not required
		return nil
	}

	for i := 0; i < len(m.Constructors); i++ {
		if swag.IsZero(m.Constructors[i]) { // not required
			continue
		}

		if m.Constructors[i] != nil {
			if err := m.Constructors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constructors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constructors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) validateMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.Methods) { // not required
		return nil
	}

	for i := 0; i < len(m.Methods); i++ {
		if swag.IsZero(m.Methods[i]) { // not required
			continue
		}

		if m.Methods[i] != nil {
			if err := m.Methods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("methods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var objectTypeTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RELEASE","ALPHA","BETA","DEPRECATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectTypeTypeStatePropEnum = append(objectTypeTypeStatePropEnum, v)
	}
}

const (

	// ObjectTypeStateRELEASE captures enum value "RELEASE"
	ObjectTypeStateRELEASE string = "RELEASE"

	// ObjectTypeStateALPHA captures enum value "ALPHA"
	ObjectTypeStateALPHA string = "ALPHA"

	// ObjectTypeStateBETA captures enum value "BETA"
	ObjectTypeStateBETA string = "BETA"

	// ObjectTypeStateDEPRECATED captures enum value "DEPRECATED"
	ObjectTypeStateDEPRECATED string = "DEPRECATED"
)

// prop value enum
func (m *ObjectType) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectTypeTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectType) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this object type based on the context it is used
func (m *ObjectType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstructors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectType) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {

			if swag.IsZero(m.Attributes[i]) { // not required
				return nil
			}

			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) contextValidateConstructors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Constructors); i++ {

		if m.Constructors[i] != nil {

			if swag.IsZero(m.Constructors[i]) { // not required
				return nil
			}

			if err := m.Constructors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constructors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constructors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {

			if swag.IsZero(m.Events[i]) { // not required
				return nil
			}

			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Members); i++ {

		if m.Members[i] != nil {

			if swag.IsZero(m.Members[i]) { // not required
				return nil
			}

			if err := m.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectType) contextValidateMethods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Methods); i++ {

		if m.Methods[i] != nil {

			if swag.IsZero(m.Methods[i]) { // not required
				return nil
			}

			if err := m.Methods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("methods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectType) UnmarshalBinary(b []byte) error {
	var res ObjectType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
