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

// PStep p step
//
// swagger:model PStep
type PStep struct {

	// desc
	Desc string `json:"desc,omitempty"`

	// p group
	// Required: true
	PGroup []*PGroup `json:"p-group"`

	// p param
	PParam []*PParam `json:"p-param"`

	// p qual
	PQual []*PQual `json:"p-qual"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this p step
func (m *PStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePQual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PStep) validatePGroup(formats strfmt.Registry) error {

	if err := validate.Required("p-group", "body", m.PGroup); err != nil {
		return err
	}

	for i := 0; i < len(m.PGroup); i++ {
		if swag.IsZero(m.PGroup[i]) { // not required
			continue
		}

		if m.PGroup[i] != nil {
			if err := m.PGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("p-group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("p-group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PStep) validatePParam(formats strfmt.Registry) error {
	if swag.IsZero(m.PParam) { // not required
		return nil
	}

	for i := 0; i < len(m.PParam); i++ {
		if swag.IsZero(m.PParam[i]) { // not required
			continue
		}

		if m.PParam[i] != nil {
			if err := m.PParam[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("p-param" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("p-param" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PStep) validatePQual(formats strfmt.Registry) error {
	if swag.IsZero(m.PQual) { // not required
		return nil
	}

	for i := 0; i < len(m.PQual); i++ {
		if swag.IsZero(m.PQual[i]) { // not required
			continue
		}

		if m.PQual[i] != nil {
			if err := m.PQual[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("p-qual" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("p-qual" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PStep) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p step based on the context it is used
func (m *PStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePParam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePQual(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PStep) contextValidatePGroup(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PGroup); i++ {

		if m.PGroup[i] != nil {

			if swag.IsZero(m.PGroup[i]) { // not required
				return nil
			}

			if err := m.PGroup[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("p-group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("p-group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PStep) contextValidatePParam(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PParam); i++ {

		if m.PParam[i] != nil {

			if swag.IsZero(m.PParam[i]) { // not required
				return nil
			}

			if err := m.PParam[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("p-param" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("p-param" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PStep) contextValidatePQual(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PQual); i++ {

		if m.PQual[i] != nil {

			if swag.IsZero(m.PQual[i]) { // not required
				return nil
			}

			if err := m.PQual[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("p-qual" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("p-qual" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PStep) UnmarshalBinary(b []byte) error {
	var res PStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
