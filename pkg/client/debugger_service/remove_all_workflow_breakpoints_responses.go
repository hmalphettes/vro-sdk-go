// Code generated by go-swagger; DO NOT EDIT.

package debugger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveAllWorkflowBreakpointsReader is a Reader for the RemoveAllWorkflowBreakpoints structure.
type RemoveAllWorkflowBreakpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAllWorkflowBreakpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveAllWorkflowBreakpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveAllWorkflowBreakpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveAllWorkflowBreakpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /debugger/{workflowId}/breakpoints] removeAllWorkflowBreakpoints", response, response.Code())
	}
}

// NewRemoveAllWorkflowBreakpointsOK creates a RemoveAllWorkflowBreakpointsOK with default headers values
func NewRemoveAllWorkflowBreakpointsOK() *RemoveAllWorkflowBreakpointsOK {
	return &RemoveAllWorkflowBreakpointsOK{}
}

/*
RemoveAllWorkflowBreakpointsOK describes a response with status code 200, with default header values.

The request is successful
*/
type RemoveAllWorkflowBreakpointsOK struct {
}

// IsSuccess returns true when this remove all workflow breakpoints o k response has a 2xx status code
func (o *RemoveAllWorkflowBreakpointsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove all workflow breakpoints o k response has a 3xx status code
func (o *RemoveAllWorkflowBreakpointsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove all workflow breakpoints o k response has a 4xx status code
func (o *RemoveAllWorkflowBreakpointsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove all workflow breakpoints o k response has a 5xx status code
func (o *RemoveAllWorkflowBreakpointsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove all workflow breakpoints o k response a status code equal to that given
func (o *RemoveAllWorkflowBreakpointsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove all workflow breakpoints o k response
func (o *RemoveAllWorkflowBreakpointsOK) Code() int {
	return 200
}

func (o *RemoveAllWorkflowBreakpointsOK) Error() string {
	return fmt.Sprintf("[DELETE /debugger/{workflowId}/breakpoints][%d] removeAllWorkflowBreakpointsOK ", 200)
}

func (o *RemoveAllWorkflowBreakpointsOK) String() string {
	return fmt.Sprintf("[DELETE /debugger/{workflowId}/breakpoints][%d] removeAllWorkflowBreakpointsOK ", 200)
}

func (o *RemoveAllWorkflowBreakpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAllWorkflowBreakpointsUnauthorized creates a RemoveAllWorkflowBreakpointsUnauthorized with default headers values
func NewRemoveAllWorkflowBreakpointsUnauthorized() *RemoveAllWorkflowBreakpointsUnauthorized {
	return &RemoveAllWorkflowBreakpointsUnauthorized{}
}

/*
RemoveAllWorkflowBreakpointsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type RemoveAllWorkflowBreakpointsUnauthorized struct {
}

// IsSuccess returns true when this remove all workflow breakpoints unauthorized response has a 2xx status code
func (o *RemoveAllWorkflowBreakpointsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove all workflow breakpoints unauthorized response has a 3xx status code
func (o *RemoveAllWorkflowBreakpointsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove all workflow breakpoints unauthorized response has a 4xx status code
func (o *RemoveAllWorkflowBreakpointsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove all workflow breakpoints unauthorized response has a 5xx status code
func (o *RemoveAllWorkflowBreakpointsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove all workflow breakpoints unauthorized response a status code equal to that given
func (o *RemoveAllWorkflowBreakpointsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove all workflow breakpoints unauthorized response
func (o *RemoveAllWorkflowBreakpointsUnauthorized) Code() int {
	return 401
}

func (o *RemoveAllWorkflowBreakpointsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /debugger/{workflowId}/breakpoints][%d] removeAllWorkflowBreakpointsUnauthorized ", 401)
}

func (o *RemoveAllWorkflowBreakpointsUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /debugger/{workflowId}/breakpoints][%d] removeAllWorkflowBreakpointsUnauthorized ", 401)
}

func (o *RemoveAllWorkflowBreakpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAllWorkflowBreakpointsNotFound creates a RemoveAllWorkflowBreakpointsNotFound with default headers values
func NewRemoveAllWorkflowBreakpointsNotFound() *RemoveAllWorkflowBreakpointsNotFound {
	return &RemoveAllWorkflowBreakpointsNotFound{}
}

/*
RemoveAllWorkflowBreakpointsNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type RemoveAllWorkflowBreakpointsNotFound struct {
}

// IsSuccess returns true when this remove all workflow breakpoints not found response has a 2xx status code
func (o *RemoveAllWorkflowBreakpointsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove all workflow breakpoints not found response has a 3xx status code
func (o *RemoveAllWorkflowBreakpointsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove all workflow breakpoints not found response has a 4xx status code
func (o *RemoveAllWorkflowBreakpointsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove all workflow breakpoints not found response has a 5xx status code
func (o *RemoveAllWorkflowBreakpointsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove all workflow breakpoints not found response a status code equal to that given
func (o *RemoveAllWorkflowBreakpointsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove all workflow breakpoints not found response
func (o *RemoveAllWorkflowBreakpointsNotFound) Code() int {
	return 404
}

func (o *RemoveAllWorkflowBreakpointsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /debugger/{workflowId}/breakpoints][%d] removeAllWorkflowBreakpointsNotFound ", 404)
}

func (o *RemoveAllWorkflowBreakpointsNotFound) String() string {
	return fmt.Sprintf("[DELETE /debugger/{workflowId}/breakpoints][%d] removeAllWorkflowBreakpointsNotFound ", 404)
}

func (o *RemoveAllWorkflowBreakpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
