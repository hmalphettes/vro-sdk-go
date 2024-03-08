// Code generated by go-swagger; DO NOT EDIT.

package content_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportContentWorkflowReader is a Reader for the ImportContentWorkflow structure.
type ImportContentWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportContentWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportContentWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportContentWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportContentWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportContentWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /content/workflows/{categoryId}] importContentWorkflow", response, response.Code())
	}
}

// NewImportContentWorkflowOK creates a ImportContentWorkflowOK with default headers values
func NewImportContentWorkflowOK() *ImportContentWorkflowOK {
	return &ImportContentWorkflowOK{}
}

/*
ImportContentWorkflowOK describes a response with status code 200, with default header values.

The request is successful
*/
type ImportContentWorkflowOK struct {
}

// IsSuccess returns true when this import content workflow o k response has a 2xx status code
func (o *ImportContentWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import content workflow o k response has a 3xx status code
func (o *ImportContentWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import content workflow o k response has a 4xx status code
func (o *ImportContentWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this import content workflow o k response has a 5xx status code
func (o *ImportContentWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this import content workflow o k response a status code equal to that given
func (o *ImportContentWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the import content workflow o k response
func (o *ImportContentWorkflowOK) Code() int {
	return 200
}

func (o *ImportContentWorkflowOK) Error() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowOK ", 200)
}

func (o *ImportContentWorkflowOK) String() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowOK ", 200)
}

func (o *ImportContentWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportContentWorkflowBadRequest creates a ImportContentWorkflowBadRequest with default headers values
func NewImportContentWorkflowBadRequest() *ImportContentWorkflowBadRequest {
	return &ImportContentWorkflowBadRequest{}
}

/*
ImportContentWorkflowBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type ImportContentWorkflowBadRequest struct {
}

// IsSuccess returns true when this import content workflow bad request response has a 2xx status code
func (o *ImportContentWorkflowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import content workflow bad request response has a 3xx status code
func (o *ImportContentWorkflowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import content workflow bad request response has a 4xx status code
func (o *ImportContentWorkflowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this import content workflow bad request response has a 5xx status code
func (o *ImportContentWorkflowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this import content workflow bad request response a status code equal to that given
func (o *ImportContentWorkflowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the import content workflow bad request response
func (o *ImportContentWorkflowBadRequest) Code() int {
	return 400
}

func (o *ImportContentWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowBadRequest ", 400)
}

func (o *ImportContentWorkflowBadRequest) String() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowBadRequest ", 400)
}

func (o *ImportContentWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportContentWorkflowUnauthorized creates a ImportContentWorkflowUnauthorized with default headers values
func NewImportContentWorkflowUnauthorized() *ImportContentWorkflowUnauthorized {
	return &ImportContentWorkflowUnauthorized{}
}

/*
ImportContentWorkflowUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type ImportContentWorkflowUnauthorized struct {
}

// IsSuccess returns true when this import content workflow unauthorized response has a 2xx status code
func (o *ImportContentWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import content workflow unauthorized response has a 3xx status code
func (o *ImportContentWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import content workflow unauthorized response has a 4xx status code
func (o *ImportContentWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this import content workflow unauthorized response has a 5xx status code
func (o *ImportContentWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this import content workflow unauthorized response a status code equal to that given
func (o *ImportContentWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the import content workflow unauthorized response
func (o *ImportContentWorkflowUnauthorized) Code() int {
	return 401
}

func (o *ImportContentWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowUnauthorized ", 401)
}

func (o *ImportContentWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowUnauthorized ", 401)
}

func (o *ImportContentWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportContentWorkflowNotFound creates a ImportContentWorkflowNotFound with default headers values
func NewImportContentWorkflowNotFound() *ImportContentWorkflowNotFound {
	return &ImportContentWorkflowNotFound{}
}

/*
ImportContentWorkflowNotFound describes a response with status code 404, with default header values.

Cannot find a category with the specified ID.
*/
type ImportContentWorkflowNotFound struct {
}

// IsSuccess returns true when this import content workflow not found response has a 2xx status code
func (o *ImportContentWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import content workflow not found response has a 3xx status code
func (o *ImportContentWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import content workflow not found response has a 4xx status code
func (o *ImportContentWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this import content workflow not found response has a 5xx status code
func (o *ImportContentWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this import content workflow not found response a status code equal to that given
func (o *ImportContentWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the import content workflow not found response
func (o *ImportContentWorkflowNotFound) Code() int {
	return 404
}

func (o *ImportContentWorkflowNotFound) Error() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowNotFound ", 404)
}

func (o *ImportContentWorkflowNotFound) String() string {
	return fmt.Sprintf("[POST /content/workflows/{categoryId}][%d] importContentWorkflowNotFound ", 404)
}

func (o *ImportContentWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
