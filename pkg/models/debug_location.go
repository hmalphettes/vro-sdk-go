// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DebugLocation debug location
//
// swagger:model debug-location
type DebugLocation struct {

	// dunes object Id
	// Required: true
	DunesObjectID *string `json:"dunesObjectId" xml:"dunesObjectId,attr"`

	// dunes object type
	// Required: true
	DunesObjectType *string `json:"dunesObjectType" xml:"dunesObjectType,attr"`

	// element name
	// Required: true
	ElementName *string `json:"elementName" xml:"elementName,attr"`

	// line number
	// Required: true
	LineNumber *int32 `json:"lineNumber" xml:"lineNumber,attr"`

	// suspended
	// Required: true
	Suspended bool `json:"suspended" xml:"suspended,attr"`

	// workflow token Id
	// Required: true
	WorkflowTokenID *string `json:"workflowTokenId" xml:"workflowTokenId,attr"`
}

// Validate validates this debug location
func (m *DebugLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDunesObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDunesObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElementName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuspended(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowTokenID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebugLocation) validateDunesObjectID(formats strfmt.Registry) error {

	if err := validate.Required("dunesObjectId", "body", m.DunesObjectID); err != nil {
		return err
	}

	return nil
}

func (m *DebugLocation) validateDunesObjectType(formats strfmt.Registry) error {

	if err := validate.Required("dunesObjectType", "body", m.DunesObjectType); err != nil {
		return err
	}

	return nil
}

func (m *DebugLocation) validateElementName(formats strfmt.Registry) error {

	if err := validate.Required("elementName", "body", m.ElementName); err != nil {
		return err
	}

	return nil
}

func (m *DebugLocation) validateLineNumber(formats strfmt.Registry) error {

	if err := validate.Required("lineNumber", "body", m.LineNumber); err != nil {
		return err
	}

	return nil
}

func (m *DebugLocation) validateSuspended(formats strfmt.Registry) error {

	if err := validate.Required("suspended", "body", bool(m.Suspended)); err != nil {
		return err
	}

	return nil
}

func (m *DebugLocation) validateWorkflowTokenID(formats strfmt.Registry) error {

	if err := validate.Required("workflowTokenId", "body", m.WorkflowTokenID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this debug location based on context it is used
func (m *DebugLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DebugLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugLocation) UnmarshalBinary(b []byte) error {
	var res DebugLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
