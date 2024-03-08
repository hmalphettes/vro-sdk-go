// Code generated by go-swagger; DO NOT EDIT.

package workflow_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteWorkflowRunExecutionReader is a Reader for the DeleteWorkflowRunExecution structure.
type DeleteWorkflowRunExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowRunExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkflowRunExecutionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkflowRunExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkflowRunExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkflowRunExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workflows/{workflowId}/executions/{executionId}] deleteWorkflowRunExecution", response, response.Code())
	}
}

// NewDeleteWorkflowRunExecutionNoContent creates a DeleteWorkflowRunExecutionNoContent with default headers values
func NewDeleteWorkflowRunExecutionNoContent() *DeleteWorkflowRunExecutionNoContent {
	return &DeleteWorkflowRunExecutionNoContent{}
}

/*
DeleteWorkflowRunExecutionNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteWorkflowRunExecutionNoContent struct {
}

// IsSuccess returns true when this delete workflow run execution no content response has a 2xx status code
func (o *DeleteWorkflowRunExecutionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete workflow run execution no content response has a 3xx status code
func (o *DeleteWorkflowRunExecutionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow run execution no content response has a 4xx status code
func (o *DeleteWorkflowRunExecutionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workflow run execution no content response has a 5xx status code
func (o *DeleteWorkflowRunExecutionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow run execution no content response a status code equal to that given
func (o *DeleteWorkflowRunExecutionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete workflow run execution no content response
func (o *DeleteWorkflowRunExecutionNoContent) Code() int {
	return 204
}

func (o *DeleteWorkflowRunExecutionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionNoContent ", 204)
}

func (o *DeleteWorkflowRunExecutionNoContent) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionNoContent ", 204)
}

func (o *DeleteWorkflowRunExecutionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowRunExecutionBadRequest creates a DeleteWorkflowRunExecutionBadRequest with default headers values
func NewDeleteWorkflowRunExecutionBadRequest() *DeleteWorkflowRunExecutionBadRequest {
	return &DeleteWorkflowRunExecutionBadRequest{}
}

/*
DeleteWorkflowRunExecutionBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type DeleteWorkflowRunExecutionBadRequest struct {
}

// IsSuccess returns true when this delete workflow run execution bad request response has a 2xx status code
func (o *DeleteWorkflowRunExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow run execution bad request response has a 3xx status code
func (o *DeleteWorkflowRunExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow run execution bad request response has a 4xx status code
func (o *DeleteWorkflowRunExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow run execution bad request response has a 5xx status code
func (o *DeleteWorkflowRunExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow run execution bad request response a status code equal to that given
func (o *DeleteWorkflowRunExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete workflow run execution bad request response
func (o *DeleteWorkflowRunExecutionBadRequest) Code() int {
	return 400
}

func (o *DeleteWorkflowRunExecutionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionBadRequest ", 400)
}

func (o *DeleteWorkflowRunExecutionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionBadRequest ", 400)
}

func (o *DeleteWorkflowRunExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowRunExecutionUnauthorized creates a DeleteWorkflowRunExecutionUnauthorized with default headers values
func NewDeleteWorkflowRunExecutionUnauthorized() *DeleteWorkflowRunExecutionUnauthorized {
	return &DeleteWorkflowRunExecutionUnauthorized{}
}

/*
DeleteWorkflowRunExecutionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type DeleteWorkflowRunExecutionUnauthorized struct {
}

// IsSuccess returns true when this delete workflow run execution unauthorized response has a 2xx status code
func (o *DeleteWorkflowRunExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow run execution unauthorized response has a 3xx status code
func (o *DeleteWorkflowRunExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow run execution unauthorized response has a 4xx status code
func (o *DeleteWorkflowRunExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow run execution unauthorized response has a 5xx status code
func (o *DeleteWorkflowRunExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow run execution unauthorized response a status code equal to that given
func (o *DeleteWorkflowRunExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete workflow run execution unauthorized response
func (o *DeleteWorkflowRunExecutionUnauthorized) Code() int {
	return 401
}

func (o *DeleteWorkflowRunExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionUnauthorized ", 401)
}

func (o *DeleteWorkflowRunExecutionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionUnauthorized ", 401)
}

func (o *DeleteWorkflowRunExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowRunExecutionNotFound creates a DeleteWorkflowRunExecutionNotFound with default headers values
func NewDeleteWorkflowRunExecutionNotFound() *DeleteWorkflowRunExecutionNotFound {
	return &DeleteWorkflowRunExecutionNotFound{}
}

/*
DeleteWorkflowRunExecutionNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type DeleteWorkflowRunExecutionNotFound struct {
}

// IsSuccess returns true when this delete workflow run execution not found response has a 2xx status code
func (o *DeleteWorkflowRunExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow run execution not found response has a 3xx status code
func (o *DeleteWorkflowRunExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow run execution not found response has a 4xx status code
func (o *DeleteWorkflowRunExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow run execution not found response has a 5xx status code
func (o *DeleteWorkflowRunExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow run execution not found response a status code equal to that given
func (o *DeleteWorkflowRunExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete workflow run execution not found response
func (o *DeleteWorkflowRunExecutionNotFound) Code() int {
	return 404
}

func (o *DeleteWorkflowRunExecutionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionNotFound ", 404)
}

func (o *DeleteWorkflowRunExecutionNotFound) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}][%d] deleteWorkflowRunExecutionNotFound ", 404)
}

func (o *DeleteWorkflowRunExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
