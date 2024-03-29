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

// HealthStatus health status
//
// swagger:model health-status
type HealthStatus struct {

	// childs
	Childs []*HealthStatusProvider `json:"childs" xml:"health-status-providers"`

	// description
	Description string `json:"description,omitempty"`

	// state
	// Enum: [UNKNOWN OK WARN ERROR]
	State string `json:"state,omitempty" xml:"state,attr,omitempty"`

	// time
	Time int64 `json:"time,omitempty" xml:"time,attr,omitempty"`
}

// Validate validates this health status
func (m *HealthStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChilds(formats); err != nil {
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

func (m *HealthStatus) validateChilds(formats strfmt.Registry) error {
	if swag.IsZero(m.Childs) { // not required
		return nil
	}

	for i := 0; i < len(m.Childs); i++ {
		if swag.IsZero(m.Childs[i]) { // not required
			continue
		}

		if m.Childs[i] != nil {
			if err := m.Childs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var healthStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","OK","WARN","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		healthStatusTypeStatePropEnum = append(healthStatusTypeStatePropEnum, v)
	}
}

const (

	// HealthStatusStateUNKNOWN captures enum value "UNKNOWN"
	HealthStatusStateUNKNOWN string = "UNKNOWN"

	// HealthStatusStateOK captures enum value "OK"
	HealthStatusStateOK string = "OK"

	// HealthStatusStateWARN captures enum value "WARN"
	HealthStatusStateWARN string = "WARN"

	// HealthStatusStateERROR captures enum value "ERROR"
	HealthStatusStateERROR string = "ERROR"
)

// prop value enum
func (m *HealthStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, healthStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HealthStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this health status based on the context it is used
func (m *HealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChilds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthStatus) contextValidateChilds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Childs); i++ {

		if m.Childs[i] != nil {

			if swag.IsZero(m.Childs[i]) { // not required
				return nil
			}

			if err := m.Childs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthStatus) UnmarshalBinary(b []byte) error {
	var res HealthStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
