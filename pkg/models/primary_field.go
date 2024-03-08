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

// PrimaryField primary field
//
// swagger:model PrimaryField
type PrimaryField struct {

	// affected fields
	AffectedFields []*AffectedField `json:"affected-fields" xml:"affected-fields"`

	// constraints
	// Read Only: true
	Constraints []BaseConstraint `json:"constraints" xml:"constraints"`

	// decorators
	// Read Only: true
	Decorators []*BaseDecorator `json:"decorators" xml:"decorators"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display-name,omitempty"`

	// fields
	Fields []*Field `json:"fields" xml:"fields"`

	// hidden
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// messages
	Messages []*PresentationMessageInfo `json:"messages" xml:"messages"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`

	// value
	Value *Value `json:"value,omitempty"`
}

// Validate validates this primary field
func (m *PrimaryField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffectedFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecorators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrimaryField) validateAffectedFields(formats strfmt.Registry) error {
	if swag.IsZero(m.AffectedFields) { // not required
		return nil
	}

	for i := 0; i < len(m.AffectedFields); i++ {
		if swag.IsZero(m.AffectedFields[i]) { // not required
			continue
		}

		if m.AffectedFields[i] != nil {
			if err := m.AffectedFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affected-fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("affected-fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) validateDecorators(formats strfmt.Registry) error {
	if swag.IsZero(m.Decorators) { // not required
		return nil
	}

	for i := 0; i < len(m.Decorators); i++ {
		if swag.IsZero(m.Decorators[i]) { // not required
			continue
		}

		if m.Decorators[i] != nil {
			if err := m.Decorators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("decorators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("decorators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this primary field based on the context it is used
func (m *PrimaryField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAffectedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDecorators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrimaryField) contextValidateAffectedFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AffectedFields); i++ {

		if m.AffectedFields[i] != nil {

			if swag.IsZero(m.AffectedFields[i]) { // not required
				return nil
			}

			if err := m.AffectedFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affected-fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("affected-fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "constraints", "body", []BaseConstraint(m.Constraints)); err != nil {
		return err
	}

	return nil
}

func (m *PrimaryField) contextValidateDecorators(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "decorators", "body", []*BaseDecorator(m.Decorators)); err != nil {
		return err
	}

	for i := 0; i < len(m.Decorators); i++ {

		if m.Decorators[i] != nil {

			if swag.IsZero(m.Decorators[i]) { // not required
				return nil
			}

			if err := m.Decorators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("decorators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("decorators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fields); i++ {

		if m.Fields[i] != nil {

			if swag.IsZero(m.Fields[i]) { // not required
				return nil
			}

			if err := m.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {

			if swag.IsZero(m.Messages[i]) { // not required
				return nil
			}

			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrimaryField) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrimaryField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrimaryField) UnmarshalBinary(b []byte) error {
	var res PrimaryField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
