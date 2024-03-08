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

// Tags tags
//
// swagger:model tags
type Tags struct {

	// link
	Link []*Link `json:"link"`

	// start
	Start int32 `json:"start,omitempty" xml:"start,attr,omitempty"`

	// total
	Total int32 `json:"total,omitempty" xml:"total,attr,omitempty"`
}

// Validate validates this tags
func (m *Tags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tags) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	for i := 0; i < len(m.Link); i++ {
		if swag.IsZero(m.Link[i]) { // not required
			continue
		}

		if m.Link[i] != nil {
			if err := m.Link[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("link" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("link" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tags based on the context it is used
func (m *Tags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tags) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Link); i++ {

		if m.Link[i] != nil {

			if swag.IsZero(m.Link[i]) { // not required
				return nil
			}

			if err := m.Link[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("link" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("link" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tags) UnmarshalBinary(b []byte) error {
	var res Tags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
