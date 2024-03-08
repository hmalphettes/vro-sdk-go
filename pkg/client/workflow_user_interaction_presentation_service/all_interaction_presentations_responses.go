// Code generated by go-swagger; DO NOT EDIT.

package workflow_user_interaction_presentation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// AllInteractionPresentationsReader is a Reader for the AllInteractionPresentations structure.
type AllInteractionPresentationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllInteractionPresentationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllInteractionPresentationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAllInteractionPresentationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllInteractionPresentationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances] allInteractionPresentations", response, response.Code())
	}
}

// NewAllInteractionPresentationsOK creates a AllInteractionPresentationsOK with default headers values
func NewAllInteractionPresentationsOK() *AllInteractionPresentationsOK {
	return &AllInteractionPresentationsOK{}
}

/*
AllInteractionPresentationsOK describes a response with status code 200, with default header values.

The request is successful
*/
type AllInteractionPresentationsOK struct {
	Payload *models.Executions
}

// IsSuccess returns true when this all interaction presentations o k response has a 2xx status code
func (o *AllInteractionPresentationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this all interaction presentations o k response has a 3xx status code
func (o *AllInteractionPresentationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this all interaction presentations o k response has a 4xx status code
func (o *AllInteractionPresentationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this all interaction presentations o k response has a 5xx status code
func (o *AllInteractionPresentationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this all interaction presentations o k response a status code equal to that given
func (o *AllInteractionPresentationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the all interaction presentations o k response
func (o *AllInteractionPresentationsOK) Code() int {
	return 200
}

func (o *AllInteractionPresentationsOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] allInteractionPresentationsOK  %+v", 200, o.Payload)
}

func (o *AllInteractionPresentationsOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] allInteractionPresentationsOK  %+v", 200, o.Payload)
}

func (o *AllInteractionPresentationsOK) GetPayload() *models.Executions {
	return o.Payload
}

func (o *AllInteractionPresentationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Executions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllInteractionPresentationsUnauthorized creates a AllInteractionPresentationsUnauthorized with default headers values
func NewAllInteractionPresentationsUnauthorized() *AllInteractionPresentationsUnauthorized {
	return &AllInteractionPresentationsUnauthorized{}
}

/*
AllInteractionPresentationsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type AllInteractionPresentationsUnauthorized struct {
}

// IsSuccess returns true when this all interaction presentations unauthorized response has a 2xx status code
func (o *AllInteractionPresentationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this all interaction presentations unauthorized response has a 3xx status code
func (o *AllInteractionPresentationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this all interaction presentations unauthorized response has a 4xx status code
func (o *AllInteractionPresentationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this all interaction presentations unauthorized response has a 5xx status code
func (o *AllInteractionPresentationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this all interaction presentations unauthorized response a status code equal to that given
func (o *AllInteractionPresentationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the all interaction presentations unauthorized response
func (o *AllInteractionPresentationsUnauthorized) Code() int {
	return 401
}

func (o *AllInteractionPresentationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] allInteractionPresentationsUnauthorized ", 401)
}

func (o *AllInteractionPresentationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] allInteractionPresentationsUnauthorized ", 401)
}

func (o *AllInteractionPresentationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAllInteractionPresentationsNotFound creates a AllInteractionPresentationsNotFound with default headers values
func NewAllInteractionPresentationsNotFound() *AllInteractionPresentationsNotFound {
	return &AllInteractionPresentationsNotFound{}
}

/*
AllInteractionPresentationsNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type AllInteractionPresentationsNotFound struct {
}

// IsSuccess returns true when this all interaction presentations not found response has a 2xx status code
func (o *AllInteractionPresentationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this all interaction presentations not found response has a 3xx status code
func (o *AllInteractionPresentationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this all interaction presentations not found response has a 4xx status code
func (o *AllInteractionPresentationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this all interaction presentations not found response has a 5xx status code
func (o *AllInteractionPresentationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this all interaction presentations not found response a status code equal to that given
func (o *AllInteractionPresentationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the all interaction presentations not found response
func (o *AllInteractionPresentationsNotFound) Code() int {
	return 404
}

func (o *AllInteractionPresentationsNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] allInteractionPresentationsNotFound ", 404)
}

func (o *AllInteractionPresentationsNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] allInteractionPresentationsNotFound ", 404)
}

func (o *AllInteractionPresentationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}