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

// GetPresentationForReader is a Reader for the GetPresentationFor structure.
type GetPresentationForReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPresentationForReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPresentationForOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPresentationForUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPresentationForNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/presentation] getPresentationFor", response, response.Code())
	}
}

// NewGetPresentationForOK creates a GetPresentationForOK with default headers values
func NewGetPresentationForOK() *GetPresentationForOK {
	return &GetPresentationForOK{}
}

/*
GetPresentationForOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPresentationForOK struct {
	Payload *models.Presentation
}

// IsSuccess returns true when this get presentation for o k response has a 2xx status code
func (o *GetPresentationForOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get presentation for o k response has a 3xx status code
func (o *GetPresentationForOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presentation for o k response has a 4xx status code
func (o *GetPresentationForOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presentation for o k response has a 5xx status code
func (o *GetPresentationForOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get presentation for o k response a status code equal to that given
func (o *GetPresentationForOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get presentation for o k response
func (o *GetPresentationForOK) Code() int {
	return 200
}

func (o *GetPresentationForOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation][%d] getPresentationForOK  %+v", 200, o.Payload)
}

func (o *GetPresentationForOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation][%d] getPresentationForOK  %+v", 200, o.Payload)
}

func (o *GetPresentationForOK) GetPayload() *models.Presentation {
	return o.Payload
}

func (o *GetPresentationForOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Presentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresentationForUnauthorized creates a GetPresentationForUnauthorized with default headers values
func NewGetPresentationForUnauthorized() *GetPresentationForUnauthorized {
	return &GetPresentationForUnauthorized{}
}

/*
GetPresentationForUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetPresentationForUnauthorized struct {
}

// IsSuccess returns true when this get presentation for unauthorized response has a 2xx status code
func (o *GetPresentationForUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presentation for unauthorized response has a 3xx status code
func (o *GetPresentationForUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presentation for unauthorized response has a 4xx status code
func (o *GetPresentationForUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presentation for unauthorized response has a 5xx status code
func (o *GetPresentationForUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get presentation for unauthorized response a status code equal to that given
func (o *GetPresentationForUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get presentation for unauthorized response
func (o *GetPresentationForUnauthorized) Code() int {
	return 401
}

func (o *GetPresentationForUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation][%d] getPresentationForUnauthorized ", 401)
}

func (o *GetPresentationForUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation][%d] getPresentationForUnauthorized ", 401)
}

func (o *GetPresentationForUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPresentationForNotFound creates a GetPresentationForNotFound with default headers values
func NewGetPresentationForNotFound() *GetPresentationForNotFound {
	return &GetPresentationForNotFound{}
}

/*
GetPresentationForNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetPresentationForNotFound struct {
}

// IsSuccess returns true when this get presentation for not found response has a 2xx status code
func (o *GetPresentationForNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presentation for not found response has a 3xx status code
func (o *GetPresentationForNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presentation for not found response has a 4xx status code
func (o *GetPresentationForNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presentation for not found response has a 5xx status code
func (o *GetPresentationForNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get presentation for not found response a status code equal to that given
func (o *GetPresentationForNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get presentation for not found response
func (o *GetPresentationForNotFound) Code() int {
	return 404
}

func (o *GetPresentationForNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation][%d] getPresentationForNotFound ", 404)
}

func (o *GetPresentationForNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/presentation][%d] getPresentationForNotFound ", 404)
}

func (o *GetPresentationForNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
