// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddWorkflowCategoryToPackageReader is a Reader for the AddWorkflowCategoryToPackage structure.
type AddWorkflowCategoryToPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWorkflowCategoryToPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddWorkflowCategoryToPackageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddWorkflowCategoryToPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddWorkflowCategoryToPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /packages/{packageName}/workflow_category/{categoryId}] addWorkflowCategoryToPackage", response, response.Code())
	}
}

// NewAddWorkflowCategoryToPackageNoContent creates a AddWorkflowCategoryToPackageNoContent with default headers values
func NewAddWorkflowCategoryToPackageNoContent() *AddWorkflowCategoryToPackageNoContent {
	return &AddWorkflowCategoryToPackageNoContent{}
}

/*
AddWorkflowCategoryToPackageNoContent describes a response with status code 204, with default header values.

The request is successful
*/
type AddWorkflowCategoryToPackageNoContent struct {
}

// IsSuccess returns true when this add workflow category to package no content response has a 2xx status code
func (o *AddWorkflowCategoryToPackageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add workflow category to package no content response has a 3xx status code
func (o *AddWorkflowCategoryToPackageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add workflow category to package no content response has a 4xx status code
func (o *AddWorkflowCategoryToPackageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add workflow category to package no content response has a 5xx status code
func (o *AddWorkflowCategoryToPackageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add workflow category to package no content response a status code equal to that given
func (o *AddWorkflowCategoryToPackageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add workflow category to package no content response
func (o *AddWorkflowCategoryToPackageNoContent) Code() int {
	return 204
}

func (o *AddWorkflowCategoryToPackageNoContent) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow_category/{categoryId}][%d] addWorkflowCategoryToPackageNoContent ", 204)
}

func (o *AddWorkflowCategoryToPackageNoContent) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow_category/{categoryId}][%d] addWorkflowCategoryToPackageNoContent ", 204)
}

func (o *AddWorkflowCategoryToPackageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWorkflowCategoryToPackageUnauthorized creates a AddWorkflowCategoryToPackageUnauthorized with default headers values
func NewAddWorkflowCategoryToPackageUnauthorized() *AddWorkflowCategoryToPackageUnauthorized {
	return &AddWorkflowCategoryToPackageUnauthorized{}
}

/*
AddWorkflowCategoryToPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type AddWorkflowCategoryToPackageUnauthorized struct {
}

// IsSuccess returns true when this add workflow category to package unauthorized response has a 2xx status code
func (o *AddWorkflowCategoryToPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add workflow category to package unauthorized response has a 3xx status code
func (o *AddWorkflowCategoryToPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add workflow category to package unauthorized response has a 4xx status code
func (o *AddWorkflowCategoryToPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add workflow category to package unauthorized response has a 5xx status code
func (o *AddWorkflowCategoryToPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add workflow category to package unauthorized response a status code equal to that given
func (o *AddWorkflowCategoryToPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add workflow category to package unauthorized response
func (o *AddWorkflowCategoryToPackageUnauthorized) Code() int {
	return 401
}

func (o *AddWorkflowCategoryToPackageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow_category/{categoryId}][%d] addWorkflowCategoryToPackageUnauthorized ", 401)
}

func (o *AddWorkflowCategoryToPackageUnauthorized) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow_category/{categoryId}][%d] addWorkflowCategoryToPackageUnauthorized ", 401)
}

func (o *AddWorkflowCategoryToPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddWorkflowCategoryToPackageNotFound creates a AddWorkflowCategoryToPackageNotFound with default headers values
func NewAddWorkflowCategoryToPackageNotFound() *AddWorkflowCategoryToPackageNotFound {
	return &AddWorkflowCategoryToPackageNotFound{}
}

/*
AddWorkflowCategoryToPackageNotFound describes a response with status code 404, with default header values.

Package or workflow category is missing
*/
type AddWorkflowCategoryToPackageNotFound struct {
}

// IsSuccess returns true when this add workflow category to package not found response has a 2xx status code
func (o *AddWorkflowCategoryToPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add workflow category to package not found response has a 3xx status code
func (o *AddWorkflowCategoryToPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add workflow category to package not found response has a 4xx status code
func (o *AddWorkflowCategoryToPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add workflow category to package not found response has a 5xx status code
func (o *AddWorkflowCategoryToPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add workflow category to package not found response a status code equal to that given
func (o *AddWorkflowCategoryToPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add workflow category to package not found response
func (o *AddWorkflowCategoryToPackageNotFound) Code() int {
	return 404
}

func (o *AddWorkflowCategoryToPackageNotFound) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow_category/{categoryId}][%d] addWorkflowCategoryToPackageNotFound ", 404)
}

func (o *AddWorkflowCategoryToPackageNotFound) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/workflow_category/{categoryId}][%d] addWorkflowCategoryToPackageNotFound ", 404)
}

func (o *AddWorkflowCategoryToPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
