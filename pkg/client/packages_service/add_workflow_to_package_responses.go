// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddWorkflowToPackageReader is a Reader for the AddWorkflowToPackage structure.
type AddWorkflowToPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWorkflowToPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddWorkflowToPackageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddWorkflowToPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddWorkflowToPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /packages/{packageName}/workflow/{workflowId}] addWorkflowToPackage", response, response.Code())
	}
}

// NewAddWorkflowToPackageNoContent creates a AddWorkflowToPackageNoContent with default headers values
func NewAddWorkflowToPackageNoContent() *AddWorkflowToPackageNoContent {
	return &AddWorkflowToPackageNoContent{}
}

/*
AddWorkflowToPackageNoContent describes a response with status code 204, with default header values.

The request is successful
*/
type AddWorkflowToPackageNoContent struct {
}

// IsSuccess returns true when this add workflow to package no content response has a 2xx status code
func (o *AddWorkflowToPackageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add workflow to package no content response has a 3xx status code
func (o *AddWorkflowToPackageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add workflow to package no content response has a 4xx status code
func (o *AddWorkflowToPackageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add workflow to package no content response has a 5xx status code
func (o *AddWorkflowToPackageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add workflow to package no content response a status code equal to that given
func (o *AddWorkflowToPackageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add workflow to package no content response
func (o *AddWorkflowToPackageNoContent) Code() int {
	return 204
}

func (o *AddWorkflowToPackageNoContent) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow/{workflowId}][%d] addWorkflowToPackageNoContent ", 204)
}

func (o *AddWorkflowToPackageNoContent) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow/{workflowId}][%d] addWorkflowToPackageNoContent ", 204)
}

func (o *AddWorkflowToPackageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWorkflowToPackageUnauthorized creates a AddWorkflowToPackageUnauthorized with default headers values
func NewAddWorkflowToPackageUnauthorized() *AddWorkflowToPackageUnauthorized {
	return &AddWorkflowToPackageUnauthorized{}
}

/*
AddWorkflowToPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type AddWorkflowToPackageUnauthorized struct {
}

// IsSuccess returns true when this add workflow to package unauthorized response has a 2xx status code
func (o *AddWorkflowToPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add workflow to package unauthorized response has a 3xx status code
func (o *AddWorkflowToPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add workflow to package unauthorized response has a 4xx status code
func (o *AddWorkflowToPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add workflow to package unauthorized response has a 5xx status code
func (o *AddWorkflowToPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add workflow to package unauthorized response a status code equal to that given
func (o *AddWorkflowToPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add workflow to package unauthorized response
func (o *AddWorkflowToPackageUnauthorized) Code() int {
	return 401
}

func (o *AddWorkflowToPackageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow/{workflowId}][%d] addWorkflowToPackageUnauthorized ", 401)
}

func (o *AddWorkflowToPackageUnauthorized) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow/{workflowId}][%d] addWorkflowToPackageUnauthorized ", 401)
}

func (o *AddWorkflowToPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWorkflowToPackageNotFound creates a AddWorkflowToPackageNotFound with default headers values
func NewAddWorkflowToPackageNotFound() *AddWorkflowToPackageNotFound {
	return &AddWorkflowToPackageNotFound{}
}

/*
AddWorkflowToPackageNotFound describes a response with status code 404, with default header values.

Package or workflow is missing
*/
type AddWorkflowToPackageNotFound struct {
}

// IsSuccess returns true when this add workflow to package not found response has a 2xx status code
func (o *AddWorkflowToPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add workflow to package not found response has a 3xx status code
func (o *AddWorkflowToPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add workflow to package not found response has a 4xx status code
func (o *AddWorkflowToPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add workflow to package not found response has a 5xx status code
func (o *AddWorkflowToPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add workflow to package not found response a status code equal to that given
func (o *AddWorkflowToPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add workflow to package not found response
func (o *AddWorkflowToPackageNotFound) Code() int {
	return 404
}

func (o *AddWorkflowToPackageNotFound) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow/{workflowId}][%d] addWorkflowToPackageNotFound ", 404)
}

func (o *AddWorkflowToPackageNotFound) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow/{workflowId}][%d] addWorkflowToPackageNotFound ", 404)
}

func (o *AddWorkflowToPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}