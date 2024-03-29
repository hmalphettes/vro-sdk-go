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

// StartWorkflowExecutionReader is a Reader for the StartWorkflowExecution structure.
type StartWorkflowExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartWorkflowExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartWorkflowExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStartWorkflowExecutionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartWorkflowExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStartWorkflowExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartWorkflowExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workflows/{workflowId}/executions] startWorkflowExecution", response, response.Code())
	}
}

// NewStartWorkflowExecutionOK creates a StartWorkflowExecutionOK with default headers values
func NewStartWorkflowExecutionOK() *StartWorkflowExecutionOK {
	return &StartWorkflowExecutionOK{}
}

/*
StartWorkflowExecutionOK describes a response with status code 200, with default header values.

successful operation
*/
type StartWorkflowExecutionOK struct {
	Payload *models.WorkflowExecution
}

// IsSuccess returns true when this start workflow execution o k response has a 2xx status code
func (o *StartWorkflowExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start workflow execution o k response has a 3xx status code
func (o *StartWorkflowExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start workflow execution o k response has a 4xx status code
func (o *StartWorkflowExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start workflow execution o k response has a 5xx status code
func (o *StartWorkflowExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start workflow execution o k response a status code equal to that given
func (o *StartWorkflowExecutionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start workflow execution o k response
func (o *StartWorkflowExecutionOK) Code() int {
	return 200
}

func (o *StartWorkflowExecutionOK) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionOK  %+v", 200, o.Payload)
}

func (o *StartWorkflowExecutionOK) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionOK  %+v", 200, o.Payload)
}

func (o *StartWorkflowExecutionOK) GetPayload() *models.WorkflowExecution {
	return o.Payload
}

func (o *StartWorkflowExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartWorkflowExecutionAccepted creates a StartWorkflowExecutionAccepted with default headers values
func NewStartWorkflowExecutionAccepted() *StartWorkflowExecutionAccepted {
	return &StartWorkflowExecutionAccepted{}
}

/*
StartWorkflowExecutionAccepted describes a response with status code 202, with default header values.

The request is successful
*/
type StartWorkflowExecutionAccepted struct {
}

// IsSuccess returns true when this start workflow execution accepted response has a 2xx status code
func (o *StartWorkflowExecutionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start workflow execution accepted response has a 3xx status code
func (o *StartWorkflowExecutionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start workflow execution accepted response has a 4xx status code
func (o *StartWorkflowExecutionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start workflow execution accepted response has a 5xx status code
func (o *StartWorkflowExecutionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start workflow execution accepted response a status code equal to that given
func (o *StartWorkflowExecutionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start workflow execution accepted response
func (o *StartWorkflowExecutionAccepted) Code() int {
	return 202
}

func (o *StartWorkflowExecutionAccepted) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionAccepted ", 202)
}

func (o *StartWorkflowExecutionAccepted) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionAccepted ", 202)
}

func (o *StartWorkflowExecutionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartWorkflowExecutionBadRequest creates a StartWorkflowExecutionBadRequest with default headers values
func NewStartWorkflowExecutionBadRequest() *StartWorkflowExecutionBadRequest {
	return &StartWorkflowExecutionBadRequest{}
}

/*
StartWorkflowExecutionBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type StartWorkflowExecutionBadRequest struct {
}

// IsSuccess returns true when this start workflow execution bad request response has a 2xx status code
func (o *StartWorkflowExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start workflow execution bad request response has a 3xx status code
func (o *StartWorkflowExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start workflow execution bad request response has a 4xx status code
func (o *StartWorkflowExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start workflow execution bad request response has a 5xx status code
func (o *StartWorkflowExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start workflow execution bad request response a status code equal to that given
func (o *StartWorkflowExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start workflow execution bad request response
func (o *StartWorkflowExecutionBadRequest) Code() int {
	return 400
}

func (o *StartWorkflowExecutionBadRequest) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionBadRequest ", 400)
}

func (o *StartWorkflowExecutionBadRequest) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionBadRequest ", 400)
}

func (o *StartWorkflowExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartWorkflowExecutionUnauthorized creates a StartWorkflowExecutionUnauthorized with default headers values
func NewStartWorkflowExecutionUnauthorized() *StartWorkflowExecutionUnauthorized {
	return &StartWorkflowExecutionUnauthorized{}
}

/*
StartWorkflowExecutionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type StartWorkflowExecutionUnauthorized struct {
}

// IsSuccess returns true when this start workflow execution unauthorized response has a 2xx status code
func (o *StartWorkflowExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start workflow execution unauthorized response has a 3xx status code
func (o *StartWorkflowExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start workflow execution unauthorized response has a 4xx status code
func (o *StartWorkflowExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this start workflow execution unauthorized response has a 5xx status code
func (o *StartWorkflowExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this start workflow execution unauthorized response a status code equal to that given
func (o *StartWorkflowExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the start workflow execution unauthorized response
func (o *StartWorkflowExecutionUnauthorized) Code() int {
	return 401
}

func (o *StartWorkflowExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionUnauthorized ", 401)
}

func (o *StartWorkflowExecutionUnauthorized) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionUnauthorized ", 401)
}

func (o *StartWorkflowExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartWorkflowExecutionNotFound creates a StartWorkflowExecutionNotFound with default headers values
func NewStartWorkflowExecutionNotFound() *StartWorkflowExecutionNotFound {
	return &StartWorkflowExecutionNotFound{}
}

/*
StartWorkflowExecutionNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type StartWorkflowExecutionNotFound struct {
}

// IsSuccess returns true when this start workflow execution not found response has a 2xx status code
func (o *StartWorkflowExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start workflow execution not found response has a 3xx status code
func (o *StartWorkflowExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start workflow execution not found response has a 4xx status code
func (o *StartWorkflowExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this start workflow execution not found response has a 5xx status code
func (o *StartWorkflowExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this start workflow execution not found response a status code equal to that given
func (o *StartWorkflowExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the start workflow execution not found response
func (o *StartWorkflowExecutionNotFound) Code() int {
	return 404
}

func (o *StartWorkflowExecutionNotFound) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionNotFound ", 404)
}

func (o *StartWorkflowExecutionNotFound) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions][%d] startWorkflowExecutionNotFound ", 404)
}

func (o *StartWorkflowExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
