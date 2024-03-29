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

// PackageDetails package details
//
// swagger:model package-details
type PackageDetails struct {

	// actions
	Actions []*Link `json:"actions" xml:"actions"`

	// configurations
	Configurations []*Link `json:"configurations" xml:"configurations"`

	// description
	Description string `json:"description,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// policy templates
	PolicyTemplates []*Link `json:"policyTemplates" xml:"policy-templates"`

	// resources
	Resources []*Link `json:"resources" xml:"resources"`

	// used plugins
	UsedPlugins []*UsedPlugin `json:"usedPlugins" xml:"used-plugins"`

	// workflows
	Workflows []*Link `json:"workflows" xml:"workflows"`
}

// Validate validates this package details
func (m *PackageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedPlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageDetails) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) validateConfigurations(formats strfmt.Registry) error {
	if swag.IsZero(m.Configurations) { // not required
		return nil
	}

	for i := 0; i < len(m.Configurations); i++ {
		if swag.IsZero(m.Configurations[i]) { // not required
			continue
		}

		if m.Configurations[i] != nil {
			if err := m.Configurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) validatePolicyTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyTemplates); i++ {
		if swag.IsZero(m.PolicyTemplates[i]) { // not required
			continue
		}

		if m.PolicyTemplates[i] != nil {
			if err := m.PolicyTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) validateUsedPlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.UsedPlugins) { // not required
		return nil
	}

	for i := 0; i < len(m.UsedPlugins); i++ {
		if swag.IsZero(m.UsedPlugins[i]) { // not required
			continue
		}

		if m.UsedPlugins[i] != nil {
			if err := m.UsedPlugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usedPlugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usedPlugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) validateWorkflows(formats strfmt.Registry) error {
	if swag.IsZero(m.Workflows) { // not required
		return nil
	}

	for i := 0; i < len(m.Workflows); i++ {
		if swag.IsZero(m.Workflows[i]) { // not required
			continue
		}

		if m.Workflows[i] != nil {
			if err := m.Workflows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workflows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this package details based on the context it is used
func (m *PackageDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedPlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkflows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageDetails) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if m.Actions[i] != nil {

			if swag.IsZero(m.Actions[i]) { // not required
				return nil
			}

			if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) contextValidateConfigurations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Configurations); i++ {

		if m.Configurations[i] != nil {

			if swag.IsZero(m.Configurations[i]) { // not required
				return nil
			}

			if err := m.Configurations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) contextValidatePolicyTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyTemplates); i++ {

		if m.PolicyTemplates[i] != nil {

			if swag.IsZero(m.PolicyTemplates[i]) { // not required
				return nil
			}

			if err := m.PolicyTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {

			if swag.IsZero(m.Resources[i]) { // not required
				return nil
			}

			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) contextValidateUsedPlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UsedPlugins); i++ {

		if m.UsedPlugins[i] != nil {

			if swag.IsZero(m.UsedPlugins[i]) { // not required
				return nil
			}

			if err := m.UsedPlugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usedPlugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usedPlugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackageDetails) contextValidateWorkflows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workflows); i++ {

		if m.Workflows[i] != nil {

			if swag.IsZero(m.Workflows[i]) { // not required
				return nil
			}

			if err := m.Workflows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workflows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDetails) UnmarshalBinary(b []byte) error {
	var res PackageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
