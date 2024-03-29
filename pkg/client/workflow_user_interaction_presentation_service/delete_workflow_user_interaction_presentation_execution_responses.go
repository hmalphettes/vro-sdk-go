// Code generated by go-swagger; DO NOT EDIT.

package workflow_user_interaction_presentation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteWorkflowUserInteractionPresentationExecutionReader is a Reader for the DeleteWorkflowUserInteractionPresentationExecution structure.
type DeleteWorkflowUserInteractionPresentationExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowUserInteractionPresentationExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkflowUserInteractionPresentationExecutionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteWorkflowUserInteractionPresentationExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkflowUserInteractionPresentationExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}] deleteWorkflowUserInteractionPresentationExecution", response, response.Code())
	}
}

// NewDeleteWorkflowUserInteractionPresentationExecutionNoContent creates a DeleteWorkflowUserInteractionPresentationExecutionNoContent with default headers values
func NewDeleteWorkflowUserInteractionPresentationExecutionNoContent() *DeleteWorkflowUserInteractionPresentationExecutionNoContent {
	return &DeleteWorkflowUserInteractionPresentationExecutionNoContent{}
}

/*
DeleteWorkflowUserInteractionPresentationExecutionNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteWorkflowUserInteractionPresentationExecutionNoContent struct {
}

// IsSuccess returns true when this delete workflow user interaction presentation execution no content response has a 2xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete workflow user interaction presentation execution no content response has a 3xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow user interaction presentation execution no content response has a 4xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete workflow user interaction presentation execution no content response has a 5xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow user interaction presentation execution no content response a status code equal to that given
func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete workflow user interaction presentation execution no content response
func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) Code() int {
	return 204
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}][%d] deleteWorkflowUserInteractionPresentationExecutionNoContent ", 204)
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}][%d] deleteWorkflowUserInteractionPresentationExecutionNoContent ", 204)
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowUserInteractionPresentationExecutionUnauthorized creates a DeleteWorkflowUserInteractionPresentationExecutionUnauthorized with default headers values
func NewDeleteWorkflowUserInteractionPresentationExecutionUnauthorized() *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized {
	return &DeleteWorkflowUserInteractionPresentationExecutionUnauthorized{}
}

/*
DeleteWorkflowUserInteractionPresentationExecutionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized.
*/
type DeleteWorkflowUserInteractionPresentationExecutionUnauthorized struct {
}

// IsSuccess returns true when this delete workflow user interaction presentation execution unauthorized response has a 2xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow user interaction presentation execution unauthorized response has a 3xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow user interaction presentation execution unauthorized response has a 4xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow user interaction presentation execution unauthorized response has a 5xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow user interaction presentation execution unauthorized response a status code equal to that given
func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete workflow user interaction presentation execution unauthorized response
func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) Code() int {
	return 401
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}][%d] deleteWorkflowUserInteractionPresentationExecutionUnauthorized ", 401)
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}][%d] deleteWorkflowUserInteractionPresentationExecutionUnauthorized ", 401)
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowUserInteractionPresentationExecutionNotFound creates a DeleteWorkflowUserInteractionPresentationExecutionNotFound with default headers values
func NewDeleteWorkflowUserInteractionPresentationExecutionNotFound() *DeleteWorkflowUserInteractionPresentationExecutionNotFound {
	return &DeleteWorkflowUserInteractionPresentationExecutionNotFound{}
}

/*
DeleteWorkflowUserInteractionPresentationExecutionNotFound describes a response with status code 404, with default header values.

Cannot find presentation instance for executionId
*/
type DeleteWorkflowUserInteractionPresentationExecutionNotFound struct {
}

// IsSuccess returns true when this delete workflow user interaction presentation execution not found response has a 2xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete workflow user interaction presentation execution not found response has a 3xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete workflow user interaction presentation execution not found response has a 4xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete workflow user interaction presentation execution not found response has a 5xx status code
func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete workflow user interaction presentation execution not found response a status code equal to that given
func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete workflow user interaction presentation execution not found response
func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) Code() int {
	return 404
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}][%d] deleteWorkflowUserInteractionPresentationExecutionNotFound ", 404)
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) String() string {
	return fmt.Sprintf("[DELETE /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}][%d] deleteWorkflowUserInteractionPresentationExecutionNotFound ", 404)
}

func (o *DeleteWorkflowUserInteractionPresentationExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
