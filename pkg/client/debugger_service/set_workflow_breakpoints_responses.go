// Code generated by go-swagger; DO NOT EDIT.

package debugger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// SetWorkflowBreakpointsReader is a Reader for the SetWorkflowBreakpoints structure.
type SetWorkflowBreakpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetWorkflowBreakpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetWorkflowBreakpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetWorkflowBreakpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetWorkflowBreakpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /debugger/{workflowId}/breakpoints] setWorkflowBreakpoints", response, response.Code())
	}
}

// NewSetWorkflowBreakpointsOK creates a SetWorkflowBreakpointsOK with default headers values
func NewSetWorkflowBreakpointsOK() *SetWorkflowBreakpointsOK {
	return &SetWorkflowBreakpointsOK{}
}

/*
SetWorkflowBreakpointsOK describes a response with status code 200, with default header values.

The request is successful
*/
type SetWorkflowBreakpointsOK struct {
	Payload *models.Breakpoints
}

// IsSuccess returns true when this set workflow breakpoints o k response has a 2xx status code
func (o *SetWorkflowBreakpointsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set workflow breakpoints o k response has a 3xx status code
func (o *SetWorkflowBreakpointsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow breakpoints o k response has a 4xx status code
func (o *SetWorkflowBreakpointsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set workflow breakpoints o k response has a 5xx status code
func (o *SetWorkflowBreakpointsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow breakpoints o k response a status code equal to that given
func (o *SetWorkflowBreakpointsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set workflow breakpoints o k response
func (o *SetWorkflowBreakpointsOK) Code() int {
	return 200
}

func (o *SetWorkflowBreakpointsOK) Error() string {
	return fmt.Sprintf("[POST /debugger/{workflowId}/breakpoints][%d] setWorkflowBreakpointsOK  %+v", 200, o.Payload)
}

func (o *SetWorkflowBreakpointsOK) String() string {
	return fmt.Sprintf("[POST /debugger/{workflowId}/breakpoints][%d] setWorkflowBreakpointsOK  %+v", 200, o.Payload)
}

func (o *SetWorkflowBreakpointsOK) GetPayload() *models.Breakpoints {
	return o.Payload
}

func (o *SetWorkflowBreakpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Breakpoints)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowBreakpointsUnauthorized creates a SetWorkflowBreakpointsUnauthorized with default headers values
func NewSetWorkflowBreakpointsUnauthorized() *SetWorkflowBreakpointsUnauthorized {
	return &SetWorkflowBreakpointsUnauthorized{}
}

/*
SetWorkflowBreakpointsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type SetWorkflowBreakpointsUnauthorized struct {
}

// IsSuccess returns true when this set workflow breakpoints unauthorized response has a 2xx status code
func (o *SetWorkflowBreakpointsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow breakpoints unauthorized response has a 3xx status code
func (o *SetWorkflowBreakpointsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow breakpoints unauthorized response has a 4xx status code
func (o *SetWorkflowBreakpointsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this set workflow breakpoints unauthorized response has a 5xx status code
func (o *SetWorkflowBreakpointsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow breakpoints unauthorized response a status code equal to that given
func (o *SetWorkflowBreakpointsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the set workflow breakpoints unauthorized response
func (o *SetWorkflowBreakpointsUnauthorized) Code() int {
	return 401
}

func (o *SetWorkflowBreakpointsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /debugger/{workflowId}/breakpoints][%d] setWorkflowBreakpointsUnauthorized ", 401)
}

func (o *SetWorkflowBreakpointsUnauthorized) String() string {
	return fmt.Sprintf("[POST /debugger/{workflowId}/breakpoints][%d] setWorkflowBreakpointsUnauthorized ", 401)
}

func (o *SetWorkflowBreakpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetWorkflowBreakpointsNotFound creates a SetWorkflowBreakpointsNotFound with default headers values
func NewSetWorkflowBreakpointsNotFound() *SetWorkflowBreakpointsNotFound {
	return &SetWorkflowBreakpointsNotFound{}
}

/*
SetWorkflowBreakpointsNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type SetWorkflowBreakpointsNotFound struct {
}

// IsSuccess returns true when this set workflow breakpoints not found response has a 2xx status code
func (o *SetWorkflowBreakpointsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow breakpoints not found response has a 3xx status code
func (o *SetWorkflowBreakpointsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow breakpoints not found response has a 4xx status code
func (o *SetWorkflowBreakpointsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set workflow breakpoints not found response has a 5xx status code
func (o *SetWorkflowBreakpointsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow breakpoints not found response a status code equal to that given
func (o *SetWorkflowBreakpointsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set workflow breakpoints not found response
func (o *SetWorkflowBreakpointsNotFound) Code() int {
	return 404
}

func (o *SetWorkflowBreakpointsNotFound) Error() string {
	return fmt.Sprintf("[POST /debugger/{workflowId}/breakpoints][%d] setWorkflowBreakpointsNotFound ", 404)
}

func (o *SetWorkflowBreakpointsNotFound) String() string {
	return fmt.Sprintf("[POST /debugger/{workflowId}/breakpoints][%d] setWorkflowBreakpointsNotFound ", 404)
}

func (o *SetWorkflowBreakpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
