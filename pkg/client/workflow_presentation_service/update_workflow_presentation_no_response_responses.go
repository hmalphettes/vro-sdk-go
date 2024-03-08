// Code generated by go-swagger; DO NOT EDIT.

package workflow_presentation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateWorkflowPresentationNoResponseReader is a Reader for the UpdateWorkflowPresentationNoResponse structure.
type UpdateWorkflowPresentationNoResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWorkflowPresentationNoResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWorkflowPresentationNoResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateWorkflowPresentationNoResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateWorkflowPresentationNoResponseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateWorkflowPresentationNoResponseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateWorkflowPresentationNoResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /workflows/{workflowId}/presentation/instances/{executionId}] updateWorkflowPresentationNoResponse", response, response.Code())
	}
}

// NewUpdateWorkflowPresentationNoResponseOK creates a UpdateWorkflowPresentationNoResponseOK with default headers values
func NewUpdateWorkflowPresentationNoResponseOK() *UpdateWorkflowPresentationNoResponseOK {
	return &UpdateWorkflowPresentationNoResponseOK{}
}

/*
UpdateWorkflowPresentationNoResponseOK describes a response with status code 200, with default header values.

The request is successful
*/
type UpdateWorkflowPresentationNoResponseOK struct {
	Payload *models.Execution
}

// IsSuccess returns true when this update workflow presentation no response o k response has a 2xx status code
func (o *UpdateWorkflowPresentationNoResponseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update workflow presentation no response o k response has a 3xx status code
func (o *UpdateWorkflowPresentationNoResponseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update workflow presentation no response o k response has a 4xx status code
func (o *UpdateWorkflowPresentationNoResponseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update workflow presentation no response o k response has a 5xx status code
func (o *UpdateWorkflowPresentationNoResponseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update workflow presentation no response o k response a status code equal to that given
func (o *UpdateWorkflowPresentationNoResponseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update workflow presentation no response o k response
func (o *UpdateWorkflowPresentationNoResponseOK) Code() int {
	return 200
}

func (o *UpdateWorkflowPresentationNoResponseOK) Error() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseOK  %+v", 200, o.Payload)
}

func (o *UpdateWorkflowPresentationNoResponseOK) String() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseOK  %+v", 200, o.Payload)
}

func (o *UpdateWorkflowPresentationNoResponseOK) GetPayload() *models.Execution {
	return o.Payload
}

func (o *UpdateWorkflowPresentationNoResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Execution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWorkflowPresentationNoResponseBadRequest creates a UpdateWorkflowPresentationNoResponseBadRequest with default headers values
func NewUpdateWorkflowPresentationNoResponseBadRequest() *UpdateWorkflowPresentationNoResponseBadRequest {
	return &UpdateWorkflowPresentationNoResponseBadRequest{}
}

/*
UpdateWorkflowPresentationNoResponseBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type UpdateWorkflowPresentationNoResponseBadRequest struct {
}

// IsSuccess returns true when this update workflow presentation no response bad request response has a 2xx status code
func (o *UpdateWorkflowPresentationNoResponseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update workflow presentation no response bad request response has a 3xx status code
func (o *UpdateWorkflowPresentationNoResponseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update workflow presentation no response bad request response has a 4xx status code
func (o *UpdateWorkflowPresentationNoResponseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update workflow presentation no response bad request response has a 5xx status code
func (o *UpdateWorkflowPresentationNoResponseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update workflow presentation no response bad request response a status code equal to that given
func (o *UpdateWorkflowPresentationNoResponseBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update workflow presentation no response bad request response
func (o *UpdateWorkflowPresentationNoResponseBadRequest) Code() int {
	return 400
}

func (o *UpdateWorkflowPresentationNoResponseBadRequest) Error() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseBadRequest ", 400)
}

func (o *UpdateWorkflowPresentationNoResponseBadRequest) String() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseBadRequest ", 400)
}

func (o *UpdateWorkflowPresentationNoResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWorkflowPresentationNoResponseUnauthorized creates a UpdateWorkflowPresentationNoResponseUnauthorized with default headers values
func NewUpdateWorkflowPresentationNoResponseUnauthorized() *UpdateWorkflowPresentationNoResponseUnauthorized {
	return &UpdateWorkflowPresentationNoResponseUnauthorized{}
}

/*
UpdateWorkflowPresentationNoResponseUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type UpdateWorkflowPresentationNoResponseUnauthorized struct {
}

// IsSuccess returns true when this update workflow presentation no response unauthorized response has a 2xx status code
func (o *UpdateWorkflowPresentationNoResponseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update workflow presentation no response unauthorized response has a 3xx status code
func (o *UpdateWorkflowPresentationNoResponseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update workflow presentation no response unauthorized response has a 4xx status code
func (o *UpdateWorkflowPresentationNoResponseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update workflow presentation no response unauthorized response has a 5xx status code
func (o *UpdateWorkflowPresentationNoResponseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update workflow presentation no response unauthorized response a status code equal to that given
func (o *UpdateWorkflowPresentationNoResponseUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update workflow presentation no response unauthorized response
func (o *UpdateWorkflowPresentationNoResponseUnauthorized) Code() int {
	return 401
}

func (o *UpdateWorkflowPresentationNoResponseUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseUnauthorized ", 401)
}

func (o *UpdateWorkflowPresentationNoResponseUnauthorized) String() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseUnauthorized ", 401)
}

func (o *UpdateWorkflowPresentationNoResponseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWorkflowPresentationNoResponseForbidden creates a UpdateWorkflowPresentationNoResponseForbidden with default headers values
func NewUpdateWorkflowPresentationNoResponseForbidden() *UpdateWorkflowPresentationNoResponseForbidden {
	return &UpdateWorkflowPresentationNoResponseForbidden{}
}

/*
UpdateWorkflowPresentationNoResponseForbidden describes a response with status code 403, with default header values.

The user does not have 'read' access rights for that workflow
*/
type UpdateWorkflowPresentationNoResponseForbidden struct {
}

// IsSuccess returns true when this update workflow presentation no response forbidden response has a 2xx status code
func (o *UpdateWorkflowPresentationNoResponseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update workflow presentation no response forbidden response has a 3xx status code
func (o *UpdateWorkflowPresentationNoResponseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update workflow presentation no response forbidden response has a 4xx status code
func (o *UpdateWorkflowPresentationNoResponseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update workflow presentation no response forbidden response has a 5xx status code
func (o *UpdateWorkflowPresentationNoResponseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update workflow presentation no response forbidden response a status code equal to that given
func (o *UpdateWorkflowPresentationNoResponseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update workflow presentation no response forbidden response
func (o *UpdateWorkflowPresentationNoResponseForbidden) Code() int {
	return 403
}

func (o *UpdateWorkflowPresentationNoResponseForbidden) Error() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseForbidden ", 403)
}

func (o *UpdateWorkflowPresentationNoResponseForbidden) String() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseForbidden ", 403)
}

func (o *UpdateWorkflowPresentationNoResponseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWorkflowPresentationNoResponseNotFound creates a UpdateWorkflowPresentationNoResponseNotFound with default headers values
func NewUpdateWorkflowPresentationNoResponseNotFound() *UpdateWorkflowPresentationNoResponseNotFound {
	return &UpdateWorkflowPresentationNoResponseNotFound{}
}

/*
UpdateWorkflowPresentationNoResponseNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID
*/
type UpdateWorkflowPresentationNoResponseNotFound struct {
}

// IsSuccess returns true when this update workflow presentation no response not found response has a 2xx status code
func (o *UpdateWorkflowPresentationNoResponseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update workflow presentation no response not found response has a 3xx status code
func (o *UpdateWorkflowPresentationNoResponseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update workflow presentation no response not found response has a 4xx status code
func (o *UpdateWorkflowPresentationNoResponseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update workflow presentation no response not found response has a 5xx status code
func (o *UpdateWorkflowPresentationNoResponseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update workflow presentation no response not found response a status code equal to that given
func (o *UpdateWorkflowPresentationNoResponseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update workflow presentation no response not found response
func (o *UpdateWorkflowPresentationNoResponseNotFound) Code() int {
	return 404
}

func (o *UpdateWorkflowPresentationNoResponseNotFound) Error() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseNotFound ", 404)
}

func (o *UpdateWorkflowPresentationNoResponseNotFound) String() string {
	return fmt.Sprintf("[PUT /workflows/{workflowId}/presentation/instances/{executionId}][%d] updateWorkflowPresentationNoResponseNotFound ", 404)
}

func (o *UpdateWorkflowPresentationNoResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
