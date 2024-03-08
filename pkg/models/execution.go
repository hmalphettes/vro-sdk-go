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

// Execution execution
//
// swagger:model execution
type Execution struct {

	// description
	Description string `json:"description,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// object id
	ObjectID string `json:"object-id,omitempty" xml:"object-id,attr,omitempty"`

	// parameters
	Parameters []*Parameter `json:"parameters" xml:"output-parameters"`

	// relations
	Relations *Relations `json:"relations,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"start-date,omitempty" xml:"start-date,attr,omitempty"`

	// started by
	StartedBy string `json:"started-by,omitempty" xml:"started-by,attr,omitempty"`

	// steps
	Steps []*StepInfo `json:"steps" xml:"steps"`

	// valid
	Valid *bool `json:"valid,omitempty" xml:"valid,attr,omitempty"`
}

// Validate validates this execution
func (m *Execution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Execution) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Execution) validateRelations(formats strfmt.Registry) error {
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

func (m *Execution) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start-date", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Execution) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this execution based on the context it is used
func (m *Execution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Execution) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {

			if swag.IsZero(m.Parameters[i]) { // not required
				return nil
			}

			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Execution) contextValidateRelations(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Execution) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {

			if swag.IsZero(m.Steps[i]) { // not required
				return nil
			}

			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Execution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Execution) UnmarshalBinary(b []byte) error {
	var res Execution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
