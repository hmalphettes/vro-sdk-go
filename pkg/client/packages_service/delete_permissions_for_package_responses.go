// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePermissionsForPackageReader is a Reader for the DeletePermissionsForPackage structure.
type DeletePermissionsForPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePermissionsForPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePermissionsForPackageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePermissionsForPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /packages/{packageName}/permissions] deletePermissionsForPackage", response, response.Code())
	}
}

// NewDeletePermissionsForPackageNoContent creates a DeletePermissionsForPackageNoContent with default headers values
func NewDeletePermissionsForPackageNoContent() *DeletePermissionsForPackageNoContent {
	return &DeletePermissionsForPackageNoContent{}
}

/*
DeletePermissionsForPackageNoContent describes a response with status code 204, with default header values.

No content
*/
type DeletePermissionsForPackageNoContent struct {
}

// IsSuccess returns true when this delete permissions for package no content response has a 2xx status code
func (o *DeletePermissionsForPackageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete permissions for package no content response has a 3xx status code
func (o *DeletePermissionsForPackageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete permissions for package no content response has a 4xx status code
func (o *DeletePermissionsForPackageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete permissions for package no content response has a 5xx status code
func (o *DeletePermissionsForPackageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete permissions for package no content response a status code equal to that given
func (o *DeletePermissionsForPackageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete permissions for package no content response
func (o *DeletePermissionsForPackageNoContent) Code() int {
	return 204
}

func (o *DeletePermissionsForPackageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/permissions][%d] deletePermissionsForPackageNoContent ", 204)
}

func (o *DeletePermissionsForPackageNoContent) String() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/permissions][%d] deletePermissionsForPackageNoContent ", 204)
}

func (o *DeletePermissionsForPackageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePermissionsForPackageUnauthorized creates a DeletePermissionsForPackageUnauthorized with default headers values
func NewDeletePermissionsForPackageUnauthorized() *DeletePermissionsForPackageUnauthorized {
	return &DeletePermissionsForPackageUnauthorized{}
}

/*
DeletePermissionsForPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type DeletePermissionsForPackageUnauthorized struct {
}

// IsSuccess returns true when this delete permissions for package unauthorized response has a 2xx status code
func (o *DeletePermissionsForPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete permissions for package unauthorized response has a 3xx status code
func (o *DeletePermissionsForPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete permissions for package unauthorized response has a 4xx status code
func (o *DeletePermissionsForPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete permissions for package unauthorized response has a 5xx status code
func (o *DeletePermissionsForPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete permissions for package unauthorized response a status code equal to that given
func (o *DeletePermissionsForPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete permissions for package unauthorized response
func (o *DeletePermissionsForPackageUnauthorized) Code() int {
	return 401
}

func (o *DeletePermissionsForPackageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/permissions][%d] deletePermissionsForPackageUnauthorized ", 401)
}

func (o *DeletePermissionsForPackageUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/permissions][%d] deletePermissionsForPackageUnauthorized ", 401)
}

func (o *DeletePermissionsForPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
