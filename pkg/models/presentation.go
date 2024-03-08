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
)

// Presentation presentation
//
// swagger:model presentation
type Presentation struct {

	// description
	Description string `json:"description,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// input parameters
	InputParameters []*Parameter `json:"inputParameters" xml:"parameter"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// output parameters
	OutputParameters []*Parameter `json:"outputParameters" xml:"parameter"`

	// relations
	Relations *Relations `json:"relations,omitempty"`

	// steps
	Steps []*StepInfo `json:"steps" xml:"steps"`
}

// Validate validates this presentation
func (m *Presentation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
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

func (m *Presentation) validateInputParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.InputParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.InputParameters); i++ {
		if swag.IsZero(m.InputParameters[i]) { // not required
			continue
		}

		if m.InputParameters[i] != nil {
			if err := m.InputParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Presentation) validateOutputParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.OutputParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.OutputParameters); i++ {
		if swag.IsZero(m.OutputParameters[i]) { // not required
			continue
		}

		if m.OutputParameters[i] != nil {
			if err := m.OutputParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Presentation) validateRelations(formats strfmt.Registry) error {
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

func (m *Presentation) validateSteps(formats strfmt.Registry) error {
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

// ContextValidate validate this presentation based on the context it is used
func (m *Presentation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputParameters(ctx, formats); err != nil {
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

func (m *Presentation) contextValidateInputParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InputParameters); i++ {

		if m.InputParameters[i] != nil {

			if swag.IsZero(m.InputParameters[i]) { // not required
				return nil
			}

			if err := m.InputParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Presentation) contextValidateOutputParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OutputParameters); i++ {

		if m.OutputParameters[i] != nil {

			if swag.IsZero(m.OutputParameters[i]) { // not required
				return nil
			}

			if err := m.OutputParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Presentation) contextValidateRelations(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Presentation) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Presentation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Presentation) UnmarshalBinary(b []byte) error {
	var res Presentation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
