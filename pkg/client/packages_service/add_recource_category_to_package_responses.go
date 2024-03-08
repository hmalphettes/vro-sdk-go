// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddRecourceCategoryToPackageReader is a Reader for the AddRecourceCategoryToPackage structure.
type AddRecourceCategoryToPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRecourceCategoryToPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddRecourceCategoryToPackageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddRecourceCategoryToPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddRecourceCategoryToPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /packages/{packageName}/resource_category/{categoryName}] addRecourceCategoryToPackage", response, response.Code())
	}
}

// NewAddRecourceCategoryToPackageNoContent creates a AddRecourceCategoryToPackageNoContent with default headers values
func NewAddRecourceCategoryToPackageNoContent() *AddRecourceCategoryToPackageNoContent {
	return &AddRecourceCategoryToPackageNoContent{}
}

/*
AddRecourceCategoryToPackageNoContent describes a response with status code 204, with default header values.

The request is successful
*/
type AddRecourceCategoryToPackageNoContent struct {
}

// IsSuccess returns true when this add recource category to package no content response has a 2xx status code
func (o *AddRecourceCategoryToPackageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add recource category to package no content response has a 3xx status code
func (o *AddRecourceCategoryToPackageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add recource category to package no content response has a 4xx status code
func (o *AddRecourceCategoryToPackageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add recource category to package no content response has a 5xx status code
func (o *AddRecourceCategoryToPackageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add recource category to package no content response a status code equal to that given
func (o *AddRecourceCategoryToPackageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the add recource category to package no content response
func (o *AddRecourceCategoryToPackageNoContent) Code() int {
	return 204
}

func (o *AddRecourceCategoryToPackageNoContent) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/resource_category/{categoryName}][%d] addRecourceCategoryToPackageNoContent ", 204)
}

func (o *AddRecourceCategoryToPackageNoContent) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/resource_category/{categoryName}][%d] addRecourceCategoryToPackageNoContent ", 204)
}

func (o *AddRecourceCategoryToPackageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddRecourceCategoryToPackageUnauthorized creates a AddRecourceCategoryToPackageUnauthorized with default headers values
func NewAddRecourceCategoryToPackageUnauthorized() *AddRecourceCategoryToPackageUnauthorized {
	return &AddRecourceCategoryToPackageUnauthorized{}
}

/*
AddRecourceCategoryToPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type AddRecourceCategoryToPackageUnauthorized struct {
}

// IsSuccess returns true when this add recource category to package unauthorized response has a 2xx status code
func (o *AddRecourceCategoryToPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add recource category to package unauthorized response has a 3xx status code
func (o *AddRecourceCategoryToPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add recource category to package unauthorized response has a 4xx status code
func (o *AddRecourceCategoryToPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add recource category to package unauthorized response has a 5xx status code
func (o *AddRecourceCategoryToPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add recource category to package unauthorized response a status code equal to that given
func (o *AddRecourceCategoryToPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add recource category to package unauthorized response
func (o *AddRecourceCategoryToPackageUnauthorized) Code() int {
	return 401
}

func (o *AddRecourceCategoryToPackageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/resource_category/{categoryName}][%d] addRecourceCategoryToPackageUnauthorized ", 401)
}

func (o *AddRecourceCategoryToPackageUnauthorized) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/resource_category/{categoryName}][%d] addRecourceCategoryToPackageUnauthorized ", 401)
}

func (o *AddRecourceCategoryToPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddRecourceCategoryToPackageNotFound creates a AddRecourceCategoryToPackageNotFound with default headers values
func NewAddRecourceCategoryToPackageNotFound() *AddRecourceCategoryToPackageNotFound {
	return &AddRecourceCategoryToPackageNotFound{}
}

/*
AddRecourceCategoryToPackageNotFound describes a response with status code 404, with default header values.

Package or workflow is missing
*/
type AddRecourceCategoryToPackageNotFound struct {
}

// IsSuccess returns true when this add recource category to package not found response has a 2xx status code
func (o *AddRecourceCategoryToPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add recource category to package not found response has a 3xx status code
func (o *AddRecourceCategoryToPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add recource category to package not found response has a 4xx status code
func (o *AddRecourceCategoryToPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add recource category to package not found response has a 5xx status code
func (o *AddRecourceCategoryToPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add recource category to package not found response a status code equal to that given
func (o *AddRecourceCategoryToPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add recource category to package not found response
func (o *AddRecourceCategoryToPackageNotFound) Code() int {
	return 404
}

func (o *AddRecourceCategoryToPackageNotFound) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/resource_category/{categoryName}][%d] addRecourceCategoryToPackageNotFound ", 404)
}

func (o *AddRecourceCategoryToPackageNotFound) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/resource_category/{categoryName}][%d] addRecourceCategoryToPackageNotFound ", 404)
}

func (o *AddRecourceCategoryToPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}