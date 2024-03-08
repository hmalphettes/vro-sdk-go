// Code generated by go-swagger; DO NOT EDIT.

package workflow_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllTasksForWorkflowReader is a Reader for the GetAllTasksForWorkflow structure.
type GetAllTasksForWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTasksForWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTasksForWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetAllTasksForWorkflowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllTasksForWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllTasksForWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllTasksForWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/tasks] getAllTasksForWorkflow", response, response.Code())
	}
}

// NewGetAllTasksForWorkflowOK creates a GetAllTasksForWorkflowOK with default headers values
func NewGetAllTasksForWorkflowOK() *GetAllTasksForWorkflowOK {
	return &GetAllTasksForWorkflowOK{}
}

/*
GetAllTasksForWorkflowOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetAllTasksForWorkflowOK struct {
	Payload *models.Tasks
}

// IsSuccess returns true when this get all tasks for workflow o k response has a 2xx status code
func (o *GetAllTasksForWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all tasks for workflow o k response has a 3xx status code
func (o *GetAllTasksForWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tasks for workflow o k response has a 4xx status code
func (o *GetAllTasksForWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tasks for workflow o k response has a 5xx status code
func (o *GetAllTasksForWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tasks for workflow o k response a status code equal to that given
func (o *GetAllTasksForWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all tasks for workflow o k response
func (o *GetAllTasksForWorkflowOK) Code() int {
	return 200
}

func (o *GetAllTasksForWorkflowOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetAllTasksForWorkflowOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetAllTasksForWorkflowOK) GetPayload() *models.Tasks {
	return o.Payload
}

func (o *GetAllTasksForWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tasks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTasksForWorkflowNoContent creates a GetAllTasksForWorkflowNoContent with default headers values
func NewGetAllTasksForWorkflowNoContent() *GetAllTasksForWorkflowNoContent {
	return &GetAllTasksForWorkflowNoContent{}
}

/*
GetAllTasksForWorkflowNoContent describes a response with status code 204, with default header values.

No content
*/
type GetAllTasksForWorkflowNoContent struct {
}

// IsSuccess returns true when this get all tasks for workflow no content response has a 2xx status code
func (o *GetAllTasksForWorkflowNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all tasks for workflow no content response has a 3xx status code
func (o *GetAllTasksForWorkflowNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tasks for workflow no content response has a 4xx status code
func (o *GetAllTasksForWorkflowNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tasks for workflow no content response has a 5xx status code
func (o *GetAllTasksForWorkflowNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tasks for workflow no content response a status code equal to that given
func (o *GetAllTasksForWorkflowNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get all tasks for workflow no content response
func (o *GetAllTasksForWorkflowNoContent) Code() int {
	return 204
}

func (o *GetAllTasksForWorkflowNoContent) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowNoContent ", 204)
}

func (o *GetAllTasksForWorkflowNoContent) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowNoContent ", 204)
}

func (o *GetAllTasksForWorkflowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTasksForWorkflowBadRequest creates a GetAllTasksForWorkflowBadRequest with default headers values
func NewGetAllTasksForWorkflowBadRequest() *GetAllTasksForWorkflowBadRequest {
	return &GetAllTasksForWorkflowBadRequest{}
}

/*
GetAllTasksForWorkflowBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type GetAllTasksForWorkflowBadRequest struct {
}

// IsSuccess returns true when this get all tasks for workflow bad request response has a 2xx status code
func (o *GetAllTasksForWorkflowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tasks for workflow bad request response has a 3xx status code
func (o *GetAllTasksForWorkflowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tasks for workflow bad request response has a 4xx status code
func (o *GetAllTasksForWorkflowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all tasks for workflow bad request response has a 5xx status code
func (o *GetAllTasksForWorkflowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tasks for workflow bad request response a status code equal to that given
func (o *GetAllTasksForWorkflowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get all tasks for workflow bad request response
func (o *GetAllTasksForWorkflowBadRequest) Code() int {
	return 400
}

func (o *GetAllTasksForWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowBadRequest ", 400)
}

func (o *GetAllTasksForWorkflowBadRequest) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowBadRequest ", 400)
}

func (o *GetAllTasksForWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTasksForWorkflowUnauthorized creates a GetAllTasksForWorkflowUnauthorized with default headers values
func NewGetAllTasksForWorkflowUnauthorized() *GetAllTasksForWorkflowUnauthorized {
	return &GetAllTasksForWorkflowUnauthorized{}
}

/*
GetAllTasksForWorkflowUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetAllTasksForWorkflowUnauthorized struct {
}

// IsSuccess returns true when this get all tasks for workflow unauthorized response has a 2xx status code
func (o *GetAllTasksForWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tasks for workflow unauthorized response has a 3xx status code
func (o *GetAllTasksForWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tasks for workflow unauthorized response has a 4xx status code
func (o *GetAllTasksForWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all tasks for workflow unauthorized response has a 5xx status code
func (o *GetAllTasksForWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tasks for workflow unauthorized response a status code equal to that given
func (o *GetAllTasksForWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all tasks for workflow unauthorized response
func (o *GetAllTasksForWorkflowUnauthorized) Code() int {
	return 401
}

func (o *GetAllTasksForWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowUnauthorized ", 401)
}

func (o *GetAllTasksForWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowUnauthorized ", 401)
}

func (o *GetAllTasksForWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTasksForWorkflowNotFound creates a GetAllTasksForWorkflowNotFound with default headers values
func NewGetAllTasksForWorkflowNotFound() *GetAllTasksForWorkflowNotFound {
	return &GetAllTasksForWorkflowNotFound{}
}

/*
GetAllTasksForWorkflowNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetAllTasksForWorkflowNotFound struct {
}

// IsSuccess returns true when this get all tasks for workflow not found response has a 2xx status code
func (o *GetAllTasksForWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tasks for workflow not found response has a 3xx status code
func (o *GetAllTasksForWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tasks for workflow not found response has a 4xx status code
func (o *GetAllTasksForWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all tasks for workflow not found response has a 5xx status code
func (o *GetAllTasksForWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tasks for workflow not found response a status code equal to that given
func (o *GetAllTasksForWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all tasks for workflow not found response
func (o *GetAllTasksForWorkflowNotFound) Code() int {
	return 404
}

func (o *GetAllTasksForWorkflowNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowNotFound ", 404)
}

func (o *GetAllTasksForWorkflowNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/tasks][%d] getAllTasksForWorkflowNotFound ", 404)
}

func (o *GetAllTasksForWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
