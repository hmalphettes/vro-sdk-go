// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddActionsToPackageReader is a Reader for the AddActionsToPackage structure.
type AddActionsToPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddActionsToPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddActionsToPackageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddActionsToPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddActionsToPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /packages/{packageName}/action/{categoryName}] addActionsToPackage", response, response.Code())
	}
}

// NewAddActionsToPackageNoContent creates a AddActionsToPackageNoContent with default headers values
func NewAddActionsToPackageNoContent() *AddActionsToPackageNoContent {
	return &AddActionsToPackageNoContent{}
}

/*
AddActionsToPackageNoContent describes a response with status code 204, with default header values.

The request is successful
*/
type AddActionsToPackageNoContent struct {
}

// IsSuccess returns true when this add actions to package no content response has a 2xx status code
func (o *AddActionsToPackageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add actions to package no content response has a 3xx status code
func (o *AddActionsToPackageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add actions to package no content response has a 4xx status code
func (o *AddActionsToPackageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add actions to package no content response has a 5xx status code
func (o *AddActionsToPackageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add actions to package no content response a status code equal to that given
func (o *AddActionsToPackageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add actions to package no content response
func (o *AddActionsToPackageNoContent) Code() int {
	return 204
}

func (o *AddActionsToPackageNoContent) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/action/{categoryName}][%d] addActionsToPackageNoContent ", 204)
}

func (o *AddActionsToPackageNoContent) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/action/{categoryName}][%d] addActionsToPackageNoContent ", 204)
}

func (o *AddActionsToPackageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddActionsToPackageUnauthorized creates a AddActionsToPackageUnauthorized with default headers values
func NewAddActionsToPackageUnauthorized() *AddActionsToPackageUnauthorized {
	return &AddActionsToPackageUnauthorized{}
}

/*
AddActionsToPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type AddActionsToPackageUnauthorized struct {
}

// IsSuccess returns true when this add actions to package unauthorized response has a 2xx status code
func (o *AddActionsToPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add actions to package unauthorized response has a 3xx status code
func (o *AddActionsToPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add actions to package unauthorized response has a 4xx status code
func (o *AddActionsToPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add actions to package unauthorized response has a 5xx status code
func (o *AddActionsToPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add actions to package unauthorized response a status code equal to that given
func (o *AddActionsToPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add actions to package unauthorized response
func (o *AddActionsToPackageUnauthorized) Code() int {
	return 401
}

func (o *AddActionsToPackageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/action/{categoryName}][%d] addActionsToPackageUnauthorized ", 401)
}

func (o *AddActionsToPackageUnauthorized) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/action/{categoryName}][%d] addActionsToPackageUnauthorized ", 401)
}

func (o *AddActionsToPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddActionsToPackageNotFound creates a AddActionsToPackageNotFound with default headers values
func NewAddActionsToPackageNotFound() *AddActionsToPackageNotFound {
	return &AddActionsToPackageNotFound{}
}

/*
AddActionsToPackageNotFound describes a response with status code 404, with default header values.

Package or workflow is missing
*/
type AddActionsToPackageNotFound struct {
}

// IsSuccess returns true when this add actions to package not found response has a 2xx status code
func (o *AddActionsToPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add actions to package not found response has a 3xx status code
func (o *AddActionsToPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add actions to package not found response has a 4xx status code
func (o *AddActionsToPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add actions to package not found response has a 5xx status code
func (o *AddActionsToPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add actions to package not found response a status code equal to that given
func (o *AddActionsToPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add actions to package not found response
func (o *AddActionsToPackageNotFound) Code() int {
	return 404
}

func (o *AddActionsToPackageNotFound) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/action/{categoryName}][%d] addActionsToPackageNotFound ", 404)
}

func (o *AddActionsToPackageNotFound) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/action/{categoryName}][%d] addActionsToPackageNotFound ", 404)
}

func (o *AddActionsToPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
