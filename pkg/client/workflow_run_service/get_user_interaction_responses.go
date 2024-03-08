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

// GetUserInteractionReader is a Reader for the GetUserInteraction structure.
type GetUserInteractionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserInteractionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserInteractionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserInteractionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserInteractionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/executions/{executionId}/interaction] getUserInteraction", response, response.Code())
	}
}

// NewGetUserInteractionOK creates a GetUserInteractionOK with default headers values
func NewGetUserInteractionOK() *GetUserInteractionOK {
	return &GetUserInteractionOK{}
}

/*
GetUserInteractionOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetUserInteractionOK struct {
	Payload *models.UserInteraction
}

// IsSuccess returns true when this get user interaction o k response has a 2xx status code
func (o *GetUserInteractionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user interaction o k response has a 3xx status code
func (o *GetUserInteractionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user interaction o k response has a 4xx status code
func (o *GetUserInteractionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user interaction o k response has a 5xx status code
func (o *GetUserInteractionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user interaction o k response a status code equal to that given
func (o *GetUserInteractionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user interaction o k response
func (o *GetUserInteractionOK) Code() int {
	return 200
}

func (o *GetUserInteractionOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction][%d] getUserInteractionOK  %+v", 200, o.Payload)
}

func (o *GetUserInteractionOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction][%d] getUserInteractionOK  %+v", 200, o.Payload)
}

func (o *GetUserInteractionOK) GetPayload() *models.UserInteraction {
	return o.Payload
}

func (o *GetUserInteractionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserInteraction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserInteractionUnauthorized creates a GetUserInteractionUnauthorized with default headers values
func NewGetUserInteractionUnauthorized() *GetUserInteractionUnauthorized {
	return &GetUserInteractionUnauthorized{}
}

/*
GetUserInteractionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetUserInteractionUnauthorized struct {
}

// IsSuccess returns true when this get user interaction unauthorized response has a 2xx status code
func (o *GetUserInteractionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user interaction unauthorized response has a 3xx status code
func (o *GetUserInteractionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user interaction unauthorized response has a 4xx status code
func (o *GetUserInteractionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user interaction unauthorized response has a 5xx status code
func (o *GetUserInteractionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user interaction unauthorized response a status code equal to that given
func (o *GetUserInteractionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get user interaction unauthorized response
func (o *GetUserInteractionUnauthorized) Code() int {
	return 401
}

func (o *GetUserInteractionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction][%d] getUserInteractionUnauthorized ", 401)
}

func (o *GetUserInteractionUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction][%d] getUserInteractionUnauthorized ", 401)
}

func (o *GetUserInteractionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserInteractionNotFound creates a GetUserInteractionNotFound with default headers values
func NewGetUserInteractionNotFound() *GetUserInteractionNotFound {
	return &GetUserInteractionNotFound{}
}

/*
GetUserInteractionNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetUserInteractionNotFound struct {
}

// IsSuccess returns true when this get user interaction not found response has a 2xx status code
func (o *GetUserInteractionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user interaction not found response has a 3xx status code
func (o *GetUserInteractionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user interaction not found response has a 4xx status code
func (o *GetUserInteractionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user interaction not found response has a 5xx status code
func (o *GetUserInteractionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user interaction not found response a status code equal to that given
func (o *GetUserInteractionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user interaction not found response
func (o *GetUserInteractionNotFound) Code() int {
	return 404
}

func (o *GetUserInteractionNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction][%d] getUserInteractionNotFound ", 404)
}

func (o *GetUserInteractionNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction][%d] getUserInteractionNotFound ", 404)
}

func (o *GetUserInteractionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
