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

// PackageCertificateInfo package certificate info
//
// swagger:model PackageCertificateInfo
type PackageCertificateInfo struct {

	// common name
	CommonName string `json:"commonName,omitempty" xml:"commonName,attr,omitempty"`

	// country
	Country string `json:"country,omitempty" xml:"country,attr,omitempty"`

	// fingerprint s h a1
	FingerprintSHA1 string `json:"fingerprintSHA1,omitempty" xml:"fingerprintSHA1,attr,omitempty"`

	// organization
	Organization string `json:"organization,omitempty" xml:"organization,attr,omitempty"`

	// organizational unit
	OrganizationalUnit string `json:"organizationalUnit,omitempty" xml:"organizationalUnit,attr,omitempty"`

	// public key algorithm
	PublicKeyAlgorithm string `json:"publicKeyAlgorithm,omitempty" xml:"publicKeyAlgorithm,attr,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty" xml:"serialNumber,attr,omitempty"`

	// valid from date
	// Format: date-time
	ValidFromDate strfmt.DateTime `json:"validFromDate,omitempty" xml:"validFromDate,attr,omitempty"`

	// valid until date
	// Format: date-time
	ValidUntilDate strfmt.DateTime `json:"validUntilDate,omitempty" xml:"validUntilDate,attr,omitempty"`
}

// Validate validates this package certificate info
func (m *PackageCertificateInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidFromDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidUntilDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageCertificateInfo) validateValidFromDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidFromDate) { // not required
		return nil
	}

	if err := validate.FormatOf("validFromDate", "body", "date-time", m.ValidFromDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PackageCertificateInfo) validateValidUntilDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidUntilDate) { // not required
		return nil
	}

	if err := validate.FormatOf("validUntilDate", "body", "date-time", m.ValidUntilDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this package certificate info based on context it is used
func (m *PackageCertificateInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageCertificateInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageCertificateInfo) UnmarshalBinary(b []byte) error {
	var res PackageCertificateInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
