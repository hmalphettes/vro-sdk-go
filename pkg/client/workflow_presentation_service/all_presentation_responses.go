// Code generated by go-swagger; DO NOT EDIT.

package workflow_presentation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// AllPresentationReader is a Reader for the AllPresentation structure.
type AllPresentationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllPresentationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllPresentationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAllPresentationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllPresentationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/presentation/instances] allPresentation", response, response.Code())
	}
}

// NewAllPresentationOK creates a AllPresentationOK with default headers values
func NewAllPresentationOK() *AllPresentationOK {
	return &AllPresentationOK{}
}

/*
AllPresentationOK describes a response with status code 200, with default header values.

The request is successful
*/
type AllPresentationOK struct {
	Payload *models.Executions
}

// IsSuccess returns true when this all presentation o k response has a 2xx status code
func (o *AllPresentationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this all presentation o k response has a 3xx status code
func (o *AllPresentationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this all presentation o k response has a 4xx status code
func (o *AllPresentationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this all presentation o k response has a 5xx status code
func (o *AllPresentationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this all presentation o k response a status code equal to that given
func (o *AllPresentationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the all presentation o k response
func (o *AllPresentationOK) Code() int {
	return 200
}

func (o *AllPresentationOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation/instances][%d] allPresentationOK  %+v", 200, o.Payload)
}

func (o *AllPresentationOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation/instances][%d] allPresentationOK  %+v", 200, o.Payload)
}

func (o *AllPresentationOK) GetPayload() *models.Executions {
	return o.Payload
}

func (o *AllPresentationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Executions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllPresentationUnauthorized creates a AllPresentationUnauthorized with default headers values
func NewAllPresentationUnauthorized() *AllPresentationUnauthorized {
	return &AllPresentationUnauthorized{}
}

/*
AllPresentationUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type AllPresentationUnauthorized struct {
}

// IsSuccess returns true when this all presentation unauthorized response has a 2xx status code
func (o *AllPresentationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this all presentation unauthorized response has a 3xx status code
func (o *AllPresentationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this all presentation unauthorized response has a 4xx status code
func (o *AllPresentationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this all presentation unauthorized response has a 5xx status code
func (o *AllPresentationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this all presentation unauthorized response a status code equal to that given
func (o *AllPresentationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the all presentation unauthorized response
func (o *AllPresentationUnauthorized) Code() int {
	return 401
}

func (o *AllPresentationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation/instances][%d] allPresentationUnauthorized ", 401)
}

func (o *AllPresentationUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation/instances][%d] allPresentationUnauthorized ", 401)
}

func (o *AllPresentationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAllPresentationNotFound creates a AllPresentationNotFound with default headers values
func NewAllPresentationNotFound() *AllPresentationNotFound {
	return &AllPresentationNotFound{}
}

/*
AllPresentationNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type AllPresentationNotFound struct {
}

// IsSuccess returns true when this all presentation not found response has a 2xx status code
func (o *AllPresentationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this all presentation not found response has a 3xx status code
func (o *AllPresentationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this all presentation not found response has a 4xx status code
func (o *AllPresentationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this all presentation not found response has a 5xx status code
func (o *AllPresentationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this all presentation not found response a status code equal to that given
func (o *AllPresentationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the all presentation not found response
func (o *AllPresentationNotFound) Code() int {
	return 404
}

func (o *AllPresentationNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation/instances][%d] allPresentationNotFound ", 404)
}

func (o *AllPresentationNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation/instances][%d] allPresentationNotFound ", 404)
}

func (o *AllPresentationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
