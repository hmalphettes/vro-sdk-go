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

// GetAllExecutionsForReader is a Reader for the GetAllExecutionsFor structure.
type GetAllExecutionsForReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllExecutionsForReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllExecutionsForOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllExecutionsForUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllExecutionsForNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/executions] getAllExecutionsFor", response, response.Code())
	}
}

// NewGetAllExecutionsForOK creates a GetAllExecutionsForOK with default headers values
func NewGetAllExecutionsForOK() *GetAllExecutionsForOK {
	return &GetAllExecutionsForOK{}
}

/*
GetAllExecutionsForOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetAllExecutionsForOK struct {
	Payload *models.Executions
}

// IsSuccess returns true when this get all executions for o k response has a 2xx status code
func (o *GetAllExecutionsForOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all executions for o k response has a 3xx status code
func (o *GetAllExecutionsForOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions for o k response has a 4xx status code
func (o *GetAllExecutionsForOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all executions for o k response has a 5xx status code
func (o *GetAllExecutionsForOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions for o k response a status code equal to that given
func (o *GetAllExecutionsForOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all executions for o k response
func (o *GetAllExecutionsForOK) Code() int {
	return 200
}

func (o *GetAllExecutionsForOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions][%d] getAllExecutionsForOK  %+v", 200, o.Payload)
}

func (o *GetAllExecutionsForOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions][%d] getAllExecutionsForOK  %+v", 200, o.Payload)
}

func (o *GetAllExecutionsForOK) GetPayload() *models.Executions {
	return o.Payload
}

func (o *GetAllExecutionsForOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Executions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllExecutionsForUnauthorized creates a GetAllExecutionsForUnauthorized with default headers values
func NewGetAllExecutionsForUnauthorized() *GetAllExecutionsForUnauthorized {
	return &GetAllExecutionsForUnauthorized{}
}

/*
GetAllExecutionsForUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetAllExecutionsForUnauthorized struct {
}

// IsSuccess returns true when this get all executions for unauthorized response has a 2xx status code
func (o *GetAllExecutionsForUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all executions for unauthorized response has a 3xx status code
func (o *GetAllExecutionsForUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions for unauthorized response has a 4xx status code
func (o *GetAllExecutionsForUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all executions for unauthorized response has a 5xx status code
func (o *GetAllExecutionsForUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions for unauthorized response a status code equal to that given
func (o *GetAllExecutionsForUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all executions for unauthorized response
func (o *GetAllExecutionsForUnauthorized) Code() int {
	return 401
}

func (o *GetAllExecutionsForUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions][%d] getAllExecutionsForUnauthorized ", 401)
}

func (o *GetAllExecutionsForUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions][%d] getAllExecutionsForUnauthorized ", 401)
}

func (o *GetAllExecutionsForUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllExecutionsForNotFound creates a GetAllExecutionsForNotFound with default headers values
func NewGetAllExecutionsForNotFound() *GetAllExecutionsForNotFound {
	return &GetAllExecutionsForNotFound{}
}

/*
GetAllExecutionsForNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetAllExecutionsForNotFound struct {
}

// IsSuccess returns true when this get all executions for not found response has a 2xx status code
func (o *GetAllExecutionsForNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all executions for not found response has a 3xx status code
func (o *GetAllExecutionsForNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions for not found response has a 4xx status code
func (o *GetAllExecutionsForNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all executions for not found response has a 5xx status code
func (o *GetAllExecutionsForNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions for not found response a status code equal to that given
func (o *GetAllExecutionsForNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all executions for not found response
func (o *GetAllExecutionsForNotFound) Code() int {
	return 404
}

func (o *GetAllExecutionsForNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions][%d] getAllExecutionsForNotFound ", 404)
}

func (o *GetAllExecutionsForNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions][%d] getAllExecutionsForNotFound ", 404)
}

func (o *GetAllExecutionsForNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
