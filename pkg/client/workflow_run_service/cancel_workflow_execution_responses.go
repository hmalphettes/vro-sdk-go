// Code generated by go-swagger; DO NOT EDIT.

package workflow_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CancelWorkflowExecutionReader is a Reader for the CancelWorkflowExecution structure.
type CancelWorkflowExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelWorkflowExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCancelWorkflowExecutionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelWorkflowExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelWorkflowExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelWorkflowExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workflows/{workflowId}/executions/{executionId}/state] cancelWorkflowExecution", response, response.Code())
	}
}

// NewCancelWorkflowExecutionNoContent creates a CancelWorkflowExecutionNoContent with default headers values
func NewCancelWorkflowExecutionNoContent() *CancelWorkflowExecutionNoContent {
	return &CancelWorkflowExecutionNoContent{}
}

/*
CancelWorkflowExecutionNoContent describes a response with status code 204, with default header values.

No content
*/
type CancelWorkflowExecutionNoContent struct {
}

// IsSuccess returns true when this cancel workflow execution no content response has a 2xx status code
func (o *CancelWorkflowExecutionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel workflow execution no content response has a 3xx status code
func (o *CancelWorkflowExecutionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel workflow execution no content response has a 4xx status code
func (o *CancelWorkflowExecutionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel workflow execution no content response has a 5xx status code
func (o *CancelWorkflowExecutionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel workflow execution no content response a status code equal to that given
func (o *CancelWorkflowExecutionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the cancel workflow execution no content response
func (o *CancelWorkflowExecutionNoContent) Code() int {
	return 204
}

func (o *CancelWorkflowExecutionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionNoContent ", 204)
}

func (o *CancelWorkflowExecutionNoContent) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionNoContent ", 204)
}

func (o *CancelWorkflowExecutionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelWorkflowExecutionBadRequest creates a CancelWorkflowExecutionBadRequest with default headers values
func NewCancelWorkflowExecutionBadRequest() *CancelWorkflowExecutionBadRequest {
	return &CancelWorkflowExecutionBadRequest{}
}

/*
CancelWorkflowExecutionBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type CancelWorkflowExecutionBadRequest struct {
}

// IsSuccess returns true when this cancel workflow execution bad request response has a 2xx status code
func (o *CancelWorkflowExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel workflow execution bad request response has a 3xx status code
func (o *CancelWorkflowExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel workflow execution bad request response has a 4xx status code
func (o *CancelWorkflowExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel workflow execution bad request response has a 5xx status code
func (o *CancelWorkflowExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel workflow execution bad request response a status code equal to that given
func (o *CancelWorkflowExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cancel workflow execution bad request response
func (o *CancelWorkflowExecutionBadRequest) Code() int {
	return 400
}

func (o *CancelWorkflowExecutionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionBadRequest ", 400)
}

func (o *CancelWorkflowExecutionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionBadRequest ", 400)
}

func (o *CancelWorkflowExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelWorkflowExecutionUnauthorized creates a CancelWorkflowExecutionUnauthorized with default headers values
func NewCancelWorkflowExecutionUnauthorized() *CancelWorkflowExecutionUnauthorized {
	return &CancelWorkflowExecutionUnauthorized{}
}

/*
CancelWorkflowExecutionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type CancelWorkflowExecutionUnauthorized struct {
}

// IsSuccess returns true when this cancel workflow execution unauthorized response has a 2xx status code
func (o *CancelWorkflowExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel workflow execution unauthorized response has a 3xx status code
func (o *CancelWorkflowExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel workflow execution unauthorized response has a 4xx status code
func (o *CancelWorkflowExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel workflow execution unauthorized response has a 5xx status code
func (o *CancelWorkflowExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel workflow execution unauthorized response a status code equal to that given
func (o *CancelWorkflowExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cancel workflow execution unauthorized response
func (o *CancelWorkflowExecutionUnauthorized) Code() int {
	return 401
}

func (o *CancelWorkflowExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionUnauthorized ", 401)
}

func (o *CancelWorkflowExecutionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionUnauthorized ", 401)
}

func (o *CancelWorkflowExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelWorkflowExecutionNotFound creates a CancelWorkflowExecutionNotFound with default headers values
func NewCancelWorkflowExecutionNotFound() *CancelWorkflowExecutionNotFound {
	return &CancelWorkflowExecutionNotFound{}
}

/*
CancelWorkflowExecutionNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type CancelWorkflowExecutionNotFound struct {
}

// IsSuccess returns true when this cancel workflow execution not found response has a 2xx status code
func (o *CancelWorkflowExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel workflow execution not found response has a 3xx status code
func (o *CancelWorkflowExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel workflow execution not found response has a 4xx status code
func (o *CancelWorkflowExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel workflow execution not found response has a 5xx status code
func (o *CancelWorkflowExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel workflow execution not found response a status code equal to that given
func (o *CancelWorkflowExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cancel workflow execution not found response
func (o *CancelWorkflowExecutionNotFound) Code() int {
	return 404
}

func (o *CancelWorkflowExecutionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionNotFound ", 404)
}

func (o *CancelWorkflowExecutionNotFound) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/state][%d] cancelWorkflowExecutionNotFound ", 404)
}

func (o *CancelWorkflowExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
