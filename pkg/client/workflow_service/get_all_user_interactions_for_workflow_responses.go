// Code generated by go-swagger; DO NOT EDIT.

package workflow_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetAllUserInteractionsForWorkflowReader is a Reader for the GetAllUserInteractionsForWorkflow structure.
type GetAllUserInteractionsForWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUserInteractionsForWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUserInteractionsForWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUserInteractionsForWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUserInteractionsForWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/interactions] getAllUserInteractionsForWorkflow", response, response.Code())
	}
}

// NewGetAllUserInteractionsForWorkflowOK creates a GetAllUserInteractionsForWorkflowOK with default headers values
func NewGetAllUserInteractionsForWorkflowOK() *GetAllUserInteractionsForWorkflowOK {
	return &GetAllUserInteractionsForWorkflowOK{}
}

/*
GetAllUserInteractionsForWorkflowOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetAllUserInteractionsForWorkflowOK struct {
	Payload *models.Interactions
}

// IsSuccess returns true when this get all user interactions for workflow o k response has a 2xx status code
func (o *GetAllUserInteractionsForWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all user interactions for workflow o k response has a 3xx status code
func (o *GetAllUserInteractionsForWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user interactions for workflow o k response has a 4xx status code
func (o *GetAllUserInteractionsForWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all user interactions for workflow o k response has a 5xx status code
func (o *GetAllUserInteractionsForWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user interactions for workflow o k response a status code equal to that given
func (o *GetAllUserInteractionsForWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all user interactions for workflow o k response
func (o *GetAllUserInteractionsForWorkflowOK) Code() int {
	return 200
}

func (o *GetAllUserInteractionsForWorkflowOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/interactions][%d] getAllUserInteractionsForWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetAllUserInteractionsForWorkflowOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/interactions][%d] getAllUserInteractionsForWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetAllUserInteractionsForWorkflowOK) GetPayload() *models.Interactions {
	return o.Payload
}

func (o *GetAllUserInteractionsForWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interactions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUserInteractionsForWorkflowUnauthorized creates a GetAllUserInteractionsForWorkflowUnauthorized with default headers values
func NewGetAllUserInteractionsForWorkflowUnauthorized() *GetAllUserInteractionsForWorkflowUnauthorized {
	return &GetAllUserInteractionsForWorkflowUnauthorized{}
}

/*
GetAllUserInteractionsForWorkflowUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetAllUserInteractionsForWorkflowUnauthorized struct {
}

// IsSuccess returns true when this get all user interactions for workflow unauthorized response has a 2xx status code
func (o *GetAllUserInteractionsForWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all user interactions for workflow unauthorized response has a 3xx status code
func (o *GetAllUserInteractionsForWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user interactions for workflow unauthorized response has a 4xx status code
func (o *GetAllUserInteractionsForWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all user interactions for workflow unauthorized response has a 5xx status code
func (o *GetAllUserInteractionsForWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user interactions for workflow unauthorized response a status code equal to that given
func (o *GetAllUserInteractionsForWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all user interactions for workflow unauthorized response
func (o *GetAllUserInteractionsForWorkflowUnauthorized) Code() int {
	return 401
}

func (o *GetAllUserInteractionsForWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/interactions][%d] getAllUserInteractionsForWorkflowUnauthorized ", 401)
}

func (o *GetAllUserInteractionsForWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/interactions][%d] getAllUserInteractionsForWorkflowUnauthorized ", 401)
}

func (o *GetAllUserInteractionsForWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUserInteractionsForWorkflowNotFound creates a GetAllUserInteractionsForWorkflowNotFound with default headers values
func NewGetAllUserInteractionsForWorkflowNotFound() *GetAllUserInteractionsForWorkflowNotFound {
	return &GetAllUserInteractionsForWorkflowNotFound{}
}

/*
GetAllUserInteractionsForWorkflowNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetAllUserInteractionsForWorkflowNotFound struct {
}

// IsSuccess returns true when this get all user interactions for workflow not found response has a 2xx status code
func (o *GetAllUserInteractionsForWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all user interactions for workflow not found response has a 3xx status code
func (o *GetAllUserInteractionsForWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user interactions for workflow not found response has a 4xx status code
func (o *GetAllUserInteractionsForWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all user interactions for workflow not found response has a 5xx status code
func (o *GetAllUserInteractionsForWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user interactions for workflow not found response a status code equal to that given
func (o *GetAllUserInteractionsForWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all user interactions for workflow not found response
func (o *GetAllUserInteractionsForWorkflowNotFound) Code() int {
	return 404
}

func (o *GetAllUserInteractionsForWorkflowNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/interactions][%d] getAllUserInteractionsForWorkflowNotFound ", 404)
}

func (o *GetAllUserInteractionsForWorkflowNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/interactions][%d] getAllUserInteractionsForWorkflowNotFound ", 404)
}

func (o *GetAllUserInteractionsForWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
