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

// ImportPackageDetails import package details
//
// swagger:model ImportPackageDetails
type ImportPackageDetails struct {

	// certificate info
	CertificateInfo *PackageCertificateInfo `json:"certificateInfo,omitempty"`

	// certificate trusted
	CertificateTrusted *bool `json:"certificateTrusted,omitempty" xml:"certificateTrusted,attr,omitempty"`

	// certificate unknown
	CertificateUnknown *bool `json:"certificateUnknown,omitempty" xml:"certificateUnknown,attr,omitempty"`

	// certificate valid
	CertificateValid *bool `json:"certificateValid,omitempty" xml:"certificateValid,attr,omitempty"`

	// content verified
	ContentVerified *bool `json:"contentVerified,omitempty" xml:"contentVerified,attr,omitempty"`

	// import element details
	ImportElementDetails []*ImportElementDetail `json:"importElementDetails"`

	// package already exists
	PackageAlreadyExists *bool `json:"packageAlreadyExists,omitempty" xml:"packageAlreadyExists,attr,omitempty"`

	// package name
	PackageName string `json:"packageName,omitempty" xml:"packageName,attr,omitempty"`
}

// Validate validates this import package details
func (m *ImportPackageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportElementDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportPackageDetails) validateCertificateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateInfo) { // not required
		return nil
	}

	if m.CertificateInfo != nil {
		if err := m.CertificateInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ImportPackageDetails) validateImportElementDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportElementDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ImportElementDetails); i++ {
		if swag.IsZero(m.ImportElementDetails[i]) { // not required
			continue
		}

		if m.ImportElementDetails[i] != nil {
			if err := m.ImportElementDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("importElementDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("importElementDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this import package details based on the context it is used
func (m *ImportPackageDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImportElementDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportPackageDetails) contextValidateCertificateInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CertificateInfo != nil {

		if swag.IsZero(m.CertificateInfo) { // not required
			return nil
		}

		if err := m.CertificateInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ImportPackageDetails) contextValidateImportElementDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImportElementDetails); i++ {

		if m.ImportElementDetails[i] != nil {

			if swag.IsZero(m.ImportElementDetails[i]) { // not required
				return nil
			}

			if err := m.ImportElementDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("importElementDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("importElementDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImportPackageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportPackageDetails) UnmarshalBinary(b []byte) error {
	var res ImportPackageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}