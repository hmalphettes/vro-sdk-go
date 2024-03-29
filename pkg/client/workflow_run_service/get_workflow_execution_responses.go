// Code generated by go-swagger; DO NOT EDIT.

package workflow_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetWorkflowExecutionReader is a Reader for the GetWorkflowExecution structure.
type GetWorkflowExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetWorkflowExecutionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkflowExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkflowExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/executions/{executionId}] getWorkflowExecution", response, response.Code())
	}
}

// NewGetWorkflowExecutionOK creates a GetWorkflowExecutionOK with default headers values
func NewGetWorkflowExecutionOK() *GetWorkflowExecutionOK {
	return &GetWorkflowExecutionOK{}
}

/*
GetWorkflowExecutionOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetWorkflowExecutionOK struct {
	Payload *models.WorkflowExecution
}

// IsSuccess returns true when this get workflow execution o k response has a 2xx status code
func (o *GetWorkflowExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow execution o k response has a 3xx status code
func (o *GetWorkflowExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution o k response has a 4xx status code
func (o *GetWorkflowExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow execution o k response has a 5xx status code
func (o *GetWorkflowExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution o k response a status code equal to that given
func (o *GetWorkflowExecutionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow execution o k response
func (o *GetWorkflowExecutionOK) Code() int {
	return 200
}

func (o *GetWorkflowExecutionOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowExecutionOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowExecutionOK) GetPayload() *models.WorkflowExecution {
	return o.Payload
}

func (o *GetWorkflowExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowExecutionNoContent creates a GetWorkflowExecutionNoContent with default headers values
func NewGetWorkflowExecutionNoContent() *GetWorkflowExecutionNoContent {
	return &GetWorkflowExecutionNoContent{}
}

/*
GetWorkflowExecutionNoContent describes a response with status code 204, with default header values.

No content
*/
type GetWorkflowExecutionNoContent struct {
}

// IsSuccess returns true when this get workflow execution no content response has a 2xx status code
func (o *GetWorkflowExecutionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow execution no content response has a 3xx status code
func (o *GetWorkflowExecutionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution no content response has a 4xx status code
func (o *GetWorkflowExecutionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow execution no content response has a 5xx status code
func (o *GetWorkflowExecutionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution no content response a status code equal to that given
func (o *GetWorkflowExecutionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get workflow execution no content response
func (o *GetWorkflowExecutionNoContent) Code() int {
	return 204
}

func (o *GetWorkflowExecutionNoContent) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionNoContent ", 204)
}

func (o *GetWorkflowExecutionNoContent) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionNoContent ", 204)
}

func (o *GetWorkflowExecutionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowExecutionBadRequest creates a GetWorkflowExecutionBadRequest with default headers values
func NewGetWorkflowExecutionBadRequest() *GetWorkflowExecutionBadRequest {
	return &GetWorkflowExecutionBadRequest{}
}

/*
GetWorkflowExecutionBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type GetWorkflowExecutionBadRequest struct {
}

// IsSuccess returns true when this get workflow execution bad request response has a 2xx status code
func (o *GetWorkflowExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow execution bad request response has a 3xx status code
func (o *GetWorkflowExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution bad request response has a 4xx status code
func (o *GetWorkflowExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow execution bad request response has a 5xx status code
func (o *GetWorkflowExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution bad request response a status code equal to that given
func (o *GetWorkflowExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get workflow execution bad request response
func (o *GetWorkflowExecutionBadRequest) Code() int {
	return 400
}

func (o *GetWorkflowExecutionBadRequest) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionBadRequest ", 400)
}

func (o *GetWorkflowExecutionBadRequest) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionBadRequest ", 400)
}

func (o *GetWorkflowExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowExecutionUnauthorized creates a GetWorkflowExecutionUnauthorized with default headers values
func NewGetWorkflowExecutionUnauthorized() *GetWorkflowExecutionUnauthorized {
	return &GetWorkflowExecutionUnauthorized{}
}

/*
GetWorkflowExecutionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetWorkflowExecutionUnauthorized struct {
}

// IsSuccess returns true when this get workflow execution unauthorized response has a 2xx status code
func (o *GetWorkflowExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow execution unauthorized response has a 3xx status code
func (o *GetWorkflowExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution unauthorized response has a 4xx status code
func (o *GetWorkflowExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow execution unauthorized response has a 5xx status code
func (o *GetWorkflowExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution unauthorized response a status code equal to that given
func (o *GetWorkflowExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get workflow execution unauthorized response
func (o *GetWorkflowExecutionUnauthorized) Code() int {
	return 401
}

func (o *GetWorkflowExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionUnauthorized ", 401)
}

func (o *GetWorkflowExecutionUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionUnauthorized ", 401)
}

func (o *GetWorkflowExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowExecutionNotFound creates a GetWorkflowExecutionNotFound with default headers values
func NewGetWorkflowExecutionNotFound() *GetWorkflowExecutionNotFound {
	return &GetWorkflowExecutionNotFound{}
}

/*
GetWorkflowExecutionNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetWorkflowExecutionNotFound struct {
}

// IsSuccess returns true when this get workflow execution not found response has a 2xx status code
func (o *GetWorkflowExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow execution not found response has a 3xx status code
func (o *GetWorkflowExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution not found response has a 4xx status code
func (o *GetWorkflowExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow execution not found response has a 5xx status code
func (o *GetWorkflowExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution not found response a status code equal to that given
func (o *GetWorkflowExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow execution not found response
func (o *GetWorkflowExecutionNotFound) Code() int {
	return 404
}

func (o *GetWorkflowExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionNotFound ", 404)
}

func (o *GetWorkflowExecutionNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}][%d] getWorkflowExecutionNotFound ", 404)
}

func (o *GetWorkflowExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
