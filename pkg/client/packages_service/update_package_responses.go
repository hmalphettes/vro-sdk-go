// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePackageReader is a Reader for the UpdatePackage structure.
type UpdatePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdatePackageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdatePackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /packages/{packageName}] updatePackage", response, response.Code())
	}
}

// NewUpdatePackageNoContent creates a UpdatePackageNoContent with default headers values
func NewUpdatePackageNoContent() *UpdatePackageNoContent {
	return &UpdatePackageNoContent{}
}

/*
UpdatePackageNoContent describes a response with status code 204, with default header values.

The request is successful
*/
type UpdatePackageNoContent struct {
}

// IsSuccess returns true when this update package no content response has a 2xx status code
func (o *UpdatePackageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update package no content response has a 3xx status code
func (o *UpdatePackageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update package no content response has a 4xx status code
func (o *UpdatePackageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update package no content response has a 5xx status code
func (o *UpdatePackageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update package no content response a status code equal to that given
func (o *UpdatePackageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update package no content response
func (o *UpdatePackageNoContent) Code() int {
	return 204
}

func (o *UpdatePackageNoContent) Error() string {
	return fmt.Sprintf("[PATCH /packages/{packageName}][%d] updatePackageNoContent ", 204)
}

func (o *UpdatePackageNoContent) String() string {
	return fmt.Sprintf("[PATCH /packages/{packageName}][%d] updatePackageNoContent ", 204)
}

func (o *UpdatePackageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePackageUnauthorized creates a UpdatePackageUnauthorized with default headers values
func NewUpdatePackageUnauthorized() *UpdatePackageUnauthorized {
	return &UpdatePackageUnauthorized{}
}

/*
UpdatePackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type UpdatePackageUnauthorized struct {
}

// IsSuccess returns true when this update package unauthorized response has a 2xx status code
func (o *UpdatePackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update package unauthorized response has a 3xx status code
func (o *UpdatePackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update package unauthorized response has a 4xx status code
func (o *UpdatePackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update package unauthorized response has a 5xx status code
func (o *UpdatePackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update package unauthorized response a status code equal to that given
func (o *UpdatePackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update package unauthorized response
func (o *UpdatePackageUnauthorized) Code() int {
	return 401
}

func (o *UpdatePackageUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /packages/{packageName}][%d] updatePackageUnauthorized ", 401)
}

func (o *UpdatePackageUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /packages/{packageName}][%d] updatePackageUnauthorized ", 401)
}

func (o *UpdatePackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
