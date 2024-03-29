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

// PolicyTemplate policy template
//
// swagger:model policy-template
type PolicyTemplate struct {

	// description
	// Max Length: 1024
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// event handlers
	EventHandlers []*EventHandler `json:"eventHandlers" xml:"event-handler"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name" xml:"name,attr"`

	// parameter
	Parameter []*Parameter `json:"parameter" xml:"parameters"`

	// policy item
	PolicyItem []*PolicyItem `json:"policy-item" xml:"policy-items"`

	// relations
	Relations *Relations `json:"relations,omitempty"`

	// version
	// Max Length: 20
	// Min Length: 0
	Version *string `json:"version,omitempty"`
}

// Validate validates this policy template
func (m *PolicyTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventHandlers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 1024); err != nil {
		return err
	}

	return nil
}

func (m *PolicyTemplate) validateEventHandlers(formats strfmt.Registry) error {
	if swag.IsZero(m.EventHandlers) { // not required
		return nil
	}

	for i := 0; i < len(m.EventHandlers); i++ {
		if swag.IsZero(m.EventHandlers[i]) { // not required
			continue
		}

		if m.EventHandlers[i] != nil {
			if err := m.EventHandlers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventHandlers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("eventHandlers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *PolicyTemplate) validateParameter(formats strfmt.Registry) error {
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

func (m *PolicyTemplate) validatePolicyItem(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyItem) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyItem); i++ {
		if swag.IsZero(m.PolicyItem[i]) { // not required
			continue
		}

		if m.PolicyItem[i] != nil {
			if err := m.PolicyItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policy-item" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policy-item" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyTemplate) validateRelations(formats strfmt.Registry) error {
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

func (m *PolicyTemplate) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", *m.Version, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("version", "body", *m.Version, 20); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policy template based on the context it is used
func (m *PolicyTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventHandlers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyItem(ctx, formats); err != nil {
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

func (m *PolicyTemplate) contextValidateEventHandlers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EventHandlers); i++ {

		if m.EventHandlers[i] != nil {

			if swag.IsZero(m.EventHandlers[i]) { // not required
				return nil
			}

			if err := m.EventHandlers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventHandlers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("eventHandlers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyTemplate) contextValidateParameter(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PolicyTemplate) contextValidatePolicyItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyItem); i++ {

		if m.PolicyItem[i] != nil {

			if swag.IsZero(m.PolicyItem[i]) { // not required
				return nil
			}

			if err := m.PolicyItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policy-item" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policy-item" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyTemplate) contextValidateRelations(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PolicyTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplate) UnmarshalBinary(b []byte) error {
	var res PolicyTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
