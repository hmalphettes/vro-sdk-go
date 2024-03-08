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

// StartInteractionPresentationReader is a Reader for the StartInteractionPresentation structure.
type StartInteractionPresentationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartInteractionPresentationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartInteractionPresentationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStartInteractionPresentationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartInteractionPresentationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances] startInteractionPresentation", response, response.Code())
	}
}

// NewStartInteractionPresentationOK creates a StartInteractionPresentationOK with default headers values
func NewStartInteractionPresentationOK() *StartInteractionPresentationOK {
	return &StartInteractionPresentationOK{}
}

/*
StartInteractionPresentationOK describes a response with status code 200, with default header values.

The request is successful
*/
type StartInteractionPresentationOK struct {
	Payload *models.Execution
}

// IsSuccess returns true when this start interaction presentation o k response has a 2xx status code
func (o *StartInteractionPresentationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start interaction presentation o k response has a 3xx status code
func (o *StartInteractionPresentationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start interaction presentation o k response has a 4xx status code
func (o *StartInteractionPresentationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start interaction presentation o k response has a 5xx status code
func (o *StartInteractionPresentationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start interaction presentation o k response a status code equal to that given
func (o *StartInteractionPresentationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start interaction presentation o k response
func (o *StartInteractionPresentationOK) Code() int {
	return 200
}

func (o *StartInteractionPresentationOK) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] startInteractionPresentationOK  %+v", 200, o.Payload)
}

func (o *StartInteractionPresentationOK) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] startInteractionPresentationOK  %+v", 200, o.Payload)
}

func (o *StartInteractionPresentationOK) GetPayload() *models.Execution {
	return o.Payload
}

func (o *StartInteractionPresentationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Execution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartInteractionPresentationUnauthorized creates a StartInteractionPresentationUnauthorized with default headers values
func NewStartInteractionPresentationUnauthorized() *StartInteractionPresentationUnauthorized {
	return &StartInteractionPresentationUnauthorized{}
}

/*
StartInteractionPresentationUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type StartInteractionPresentationUnauthorized struct {
}

// IsSuccess returns true when this start interaction presentation unauthorized response has a 2xx status code
func (o *StartInteractionPresentationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start interaction presentation unauthorized response has a 3xx status code
func (o *StartInteractionPresentationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start interaction presentation unauthorized response has a 4xx status code
func (o *StartInteractionPresentationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this start interaction presentation unauthorized response has a 5xx status code
func (o *StartInteractionPresentationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this start interaction presentation unauthorized response a status code equal to that given
func (o *StartInteractionPresentationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the start interaction presentation unauthorized response
func (o *StartInteractionPresentationUnauthorized) Code() int {
	return 401
}

func (o *StartInteractionPresentationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] startInteractionPresentationUnauthorized ", 401)
}

func (o *StartInteractionPresentationUnauthorized) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] startInteractionPresentationUnauthorized ", 401)
}

func (o *StartInteractionPresentationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartInteractionPresentationNotFound creates a StartInteractionPresentationNotFound with default headers values
func NewStartInteractionPresentationNotFound() *StartInteractionPresentationNotFound {
	return &StartInteractionPresentationNotFound{}
}

/*
StartInteractionPresentationNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type StartInteractionPresentationNotFound struct {
}

// IsSuccess returns true when this start interaction presentation not found response has a 2xx status code
func (o *StartInteractionPresentationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start interaction presentation not found response has a 3xx status code
func (o *StartInteractionPresentationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start interaction presentation not found response has a 4xx status code
func (o *StartInteractionPresentationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this start interaction presentation not found response has a 5xx status code
func (o *StartInteractionPresentationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this start interaction presentation not found response a status code equal to that given
func (o *StartInteractionPresentationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the start interaction presentation not found response
func (o *StartInteractionPresentationNotFound) Code() int {
	return 404
}

func (o *StartInteractionPresentationNotFound) Error() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] startInteractionPresentationNotFound ", 404)
}

func (o *StartInteractionPresentationNotFound) String() string {
	return fmt.Sprintf("[POST /workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances][%d] startInteractionPresentationNotFound ", 404)
}

func (o *StartInteractionPresentationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}