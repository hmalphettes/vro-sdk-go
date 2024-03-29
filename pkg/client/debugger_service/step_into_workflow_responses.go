// Code generated by go-swagger; DO NOT EDIT.

package debugger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StepIntoWorkflowReader is a Reader for the StepIntoWorkflow structure.
type StepIntoWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StepIntoWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStepIntoWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStepIntoWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStepIntoWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /debugger/{executionId}/stepInto] stepIntoWorkflow", response, response.Code())
	}
}

// NewStepIntoWorkflowOK creates a StepIntoWorkflowOK with default headers values
func NewStepIntoWorkflowOK() *StepIntoWorkflowOK {
	return &StepIntoWorkflowOK{}
}

/*
StepIntoWorkflowOK describes a response with status code 200, with default header values.

The request is successful
*/
type StepIntoWorkflowOK struct {
}

// IsSuccess returns true when this step into workflow o k response has a 2xx status code
func (o *StepIntoWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this step into workflow o k response has a 3xx status code
func (o *StepIntoWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this step into workflow o k response has a 4xx status code
func (o *StepIntoWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this step into workflow o k response has a 5xx status code
func (o *StepIntoWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this step into workflow o k response a status code equal to that given
func (o *StepIntoWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the step into workflow o k response
func (o *StepIntoWorkflowOK) Code() int {
	return 200
}

func (o *StepIntoWorkflowOK) Error() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/stepInto][%d] stepIntoWorkflowOK ", 200)
}

func (o *StepIntoWorkflowOK) String() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/stepInto][%d] stepIntoWorkflowOK ", 200)
}

func (o *StepIntoWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStepIntoWorkflowUnauthorized creates a StepIntoWorkflowUnauthorized with default headers values
func NewStepIntoWorkflowUnauthorized() *StepIntoWorkflowUnauthorized {
	return &StepIntoWorkflowUnauthorized{}
}

/*
StepIntoWorkflowUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type StepIntoWorkflowUnauthorized struct {
}

// IsSuccess returns true when this step into workflow unauthorized response has a 2xx status code
func (o *StepIntoWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this step into workflow unauthorized response has a 3xx status code
func (o *StepIntoWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this step into workflow unauthorized response has a 4xx status code
func (o *StepIntoWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this step into workflow unauthorized response has a 5xx status code
func (o *StepIntoWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this step into workflow unauthorized response a status code equal to that given
func (o *StepIntoWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the step into workflow unauthorized response
func (o *StepIntoWorkflowUnauthorized) Code() int {
	return 401
}

func (o *StepIntoWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/stepInto][%d] stepIntoWorkflowUnauthorized ", 401)
}

func (o *StepIntoWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/stepInto][%d] stepIntoWorkflowUnauthorized ", 401)
}

func (o *StepIntoWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStepIntoWorkflowNotFound creates a StepIntoWorkflowNotFound with default headers values
func NewStepIntoWorkflowNotFound() *StepIntoWorkflowNotFound {
	return &StepIntoWorkflowNotFound{}
}

/*
StepIntoWorkflowNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type StepIntoWorkflowNotFound struct {
}

// IsSuccess returns true when this step into workflow not found response has a 2xx status code
func (o *StepIntoWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this step into workflow not found response has a 3xx status code
func (o *StepIntoWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this step into workflow not found response has a 4xx status code
func (o *StepIntoWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this step into workflow not found response has a 5xx status code
func (o *StepIntoWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this step into workflow not found response a status code equal to that given
func (o *StepIntoWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the step into workflow not found response
func (o *StepIntoWorkflowNotFound) Code() int {
	return 404
}

func (o *StepIntoWorkflowNotFound) Error() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/stepInto][%d] stepIntoWorkflowNotFound ", 404)
}

func (o *StepIntoWorkflowNotFound) String() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/stepInto][%d] stepIntoWorkflowNotFound ", 404)
}

func (o *StepIntoWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
