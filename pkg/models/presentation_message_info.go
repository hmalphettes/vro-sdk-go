// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PresentationMessageInfo presentation message info
//
// swagger:model PresentationMessageInfo
type PresentationMessageInfo struct {

	// details
	Details string `json:"Details,omitempty"`

	// summary
	// Required: true
	Summary *string `json:"Summary"`

	// code
	Code string `json:"code,omitempty" xml:"code,attr,omitempty"`

	// severity
	// Enum: [INFO WARN ERROR]
	Severity string `json:"severity,omitempty" xml:"severity,attr,omitempty"`
}

// Validate validates this presentation message info
func (m *PresentationMessageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PresentationMessageInfo) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("Summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

var presentationMessageInfoTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INFO","WARN","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		presentationMessageInfoTypeSeverityPropEnum = append(presentationMessageInfoTypeSeverityPropEnum, v)
	}
}

const (

	// PresentationMessageInfoSeverityINFO captures enum value "INFO"
	PresentationMessageInfoSeverityINFO string = "INFO"

	// PresentationMessageInfoSeverityWARN captures enum value "WARN"
	PresentationMessageInfoSeverityWARN string = "WARN"

	// PresentationMessageInfoSeverityERROR captures enum value "ERROR"
	PresentationMessageInfoSeverityERROR string = "ERROR"
)

// prop value enum
func (m *PresentationMessageInfo) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, presentationMessageInfoTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PresentationMessageInfo) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this presentation message info based on context it is used
func (m *PresentationMessageInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PresentationMessageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PresentationMessageInfo) UnmarshalBinary(b []byte) error {
	var res PresentationMessageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
