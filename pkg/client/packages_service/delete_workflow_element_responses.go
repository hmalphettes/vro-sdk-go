// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteWorkflowElementReader is a Reader for the DeleteWorkflowElement structure.
type DeleteWorkflowElementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowElementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkflowElementNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteWorkflowElementUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkflowElementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /packages/{packageName}/workflow/{id}] deleteWorkflowElement", response, response.Code())
	}
}

// NewDeleteWorkflowElementNoContent creates a DeleteWorkflowElementNoContent with default headers values
func NewDeleteWorkflowElementNoContent() *DeleteWorkflowElementNoContent {
	return &DeleteWorkflowElementNoContent{}
}

/*
DeleteWorkflowElementNoContent describes a response with status code 204, with default header values.

The request is successful
*/
type DeleteWorkflowElementNoContent struct {
}

// IsSuccess returns true when this delete workflow element no content response has a 2xx status code
func (o *DeleteWorkflowElementNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete workflow element no content response has a 3xx status code
func (o *DeleteWorkflowElementNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow element no content response has a 4xx status code
func (o *DeleteWorkflowElementNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workflow element no content response has a 5xx status code
func (o *DeleteWorkflowElementNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow element no content response a status code equal to that given
func (o *DeleteWorkflowElementNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete workflow element no content response
func (o *DeleteWorkflowElementNoContent) Code() int {
	return 204
}

func (o *DeleteWorkflowElementNoContent) Error() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/workflow/{id}][%d] deleteWorkflowElementNoContent ", 204)
}

func (o *DeleteWorkflowElementNoContent) String() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/workflow/{id}][%d] deleteWorkflowElementNoContent ", 204)
}

func (o *DeleteWorkflowElementNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowElementUnauthorized creates a DeleteWorkflowElementUnauthorized with default headers values
func NewDeleteWorkflowElementUnauthorized() *DeleteWorkflowElementUnauthorized {
	return &DeleteWorkflowElementUnauthorized{}
}

/*
DeleteWorkflowElementUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type DeleteWorkflowElementUnauthorized struct {
}

// IsSuccess returns true when this delete workflow element unauthorized response has a 2xx status code
func (o *DeleteWorkflowElementUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow element unauthorized response has a 3xx status code
func (o *DeleteWorkflowElementUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow element unauthorized response has a 4xx status code
func (o *DeleteWorkflowElementUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow element unauthorized response has a 5xx status code
func (o *DeleteWorkflowElementUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow element unauthorized response a status code equal to that given
func (o *DeleteWorkflowElementUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete workflow element unauthorized response
func (o *DeleteWorkflowElementUnauthorized) Code() int {
	return 401
}

func (o *DeleteWorkflowElementUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/workflow/{id}][%d] deleteWorkflowElementUnauthorized ", 401)
}

func (o *DeleteWorkflowElementUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/workflow/{id}][%d] deleteWorkflowElementUnauthorized ", 401)
}

func (o *DeleteWorkflowElementUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowElementNotFound creates a DeleteWorkflowElementNotFound with default headers values
func NewDeleteWorkflowElementNotFound() *DeleteWorkflowElementNotFound {
	return &DeleteWorkflowElementNotFound{}
}

/*
DeleteWorkflowElementNotFound describes a response with status code 404, with default header values.

Package or workflow element is missing
*/
type DeleteWorkflowElementNotFound struct {
}

// IsSuccess returns true when this delete workflow element not found response has a 2xx status code
func (o *DeleteWorkflowElementNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow element not found response has a 3xx status code
func (o *DeleteWorkflowElementNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow element not found response has a 4xx status code
func (o *DeleteWorkflowElementNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow element not found response has a 5xx status code
func (o *DeleteWorkflowElementNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow element not found response a status code equal to that given
func (o *DeleteWorkflowElementNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete workflow element not found response
func (o *DeleteWorkflowElementNotFound) Code() int {
	return 404
}

func (o *DeleteWorkflowElementNotFound) Error() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/workflow/{id}][%d] deleteWorkflowElementNotFound ", 404)
}

func (o *DeleteWorkflowElementNotFound) String() string {
	return fmt.Sprintf("[DELETE /packages/{packageName}/workflow/{id}][%d] deleteWorkflowElementNotFound ", 404)
}

func (o *DeleteWorkflowElementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
