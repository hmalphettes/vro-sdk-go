// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Action action
//
// swagger:model action
type Action struct {

	// description
	Description string `json:"description,omitempty"`

	// fqn
	Fqn string `json:"fqn,omitempty" xml:"fqn,attr,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// module
	// Required: true
	// Pattern: ^[\p{L}_]+(\.[\p{L}0-9_]+)*$
	Module *string `json:"module"`

	// name
	// Required: true
	// Pattern: ^[\p{L}_$][\p{L}0-9_$]*$
	Name *string `json:"name"`

	// output type
	OutputType string `json:"output-type,omitempty" xml:"output-type,attr,omitempty"`

	// parameter
	Parameter []*Parameter `json:"parameter" xml:"input-parameters"`

	// relations
	Relations *Relations `json:"relations,omitempty"`

	// script
	Script string `json:"script,omitempty" xml:"script,attr,omitempty"`

	// version
	Version string `json:"version,omitempty" xml:"version,attr,omitempty"`
}

// Validate validates this action
func (m *Action) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Action) validateModule(formats strfmt.Registry) error {

	if err := validate.Required("module", "body", m.Module); err != nil {
		return err
	}

	if err := validate.Pattern("module", "body", *m.Module, `^[\p{L}_]+(\.[\p{L}0-9_]+)*$`); err != nil {
		return err
	}

	return nil
}

func (m *Action) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[\p{L}_$][\p{L}0-9_$]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Action) validateParameter(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameter) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameter); i++ {
		if swag.IsZero(m.Parameter[i]) { // not required
			continue
		}

		if m.Parameter[i] != nil {
			if err := m.Parameter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameter" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Action) validateRelations(formats strfmt.Registry) error {
	if swag.IsZero(m.Relations) { // not required
		return nil
	}

	if m.Relations != nil {
		if err := m.Relations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relations")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this action based on the context it is used
func (m *Action) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Action) contextValidateParameter(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameter); i++ {

		if m.Parameter[i] != nil {

			if swag.IsZero(m.Parameter[i]) { // not required
				return nil
			}

			if err := m.Parameter[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameter" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Action) contextValidateRelations(ctx context.Context, formats strfmt.Registry) error {

	if m.Relations != nil {

		if swag.IsZero(m.Relations) { // not required
			return nil
		}

		if err := m.Relations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Action) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Action) UnmarshalBinary(b []byte) error {
	var res Action
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
