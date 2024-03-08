// Code generated by go-swagger; DO NOT EDIT.

package debugger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetWorkflowDebugLocationReader is a Reader for the GetWorkflowDebugLocation structure.
type GetWorkflowDebugLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowDebugLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowDebugLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkflowDebugLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowDebugLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /debugger/{executionId}] getWorkflowDebugLocation", response, response.Code())
	}
}

// NewGetWorkflowDebugLocationOK creates a GetWorkflowDebugLocationOK with default headers values
func NewGetWorkflowDebugLocationOK() *GetWorkflowDebugLocationOK {
	return &GetWorkflowDebugLocationOK{}
}

/*
GetWorkflowDebugLocationOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetWorkflowDebugLocationOK struct {
	Payload *models.DebugLocation
}

// IsSuccess returns true when this get workflow debug location o k response has a 2xx status code
func (o *GetWorkflowDebugLocationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow debug location o k response has a 3xx status code
func (o *GetWorkflowDebugLocationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow debug location o k response has a 4xx status code
func (o *GetWorkflowDebugLocationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow debug location o k response has a 5xx status code
func (o *GetWorkflowDebugLocationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow debug location o k response a status code equal to that given
func (o *GetWorkflowDebugLocationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow debug location o k response
func (o *GetWorkflowDebugLocationOK) Code() int {
	return 200
}

func (o *GetWorkflowDebugLocationOK) Error() string {
	return fmt.Sprintf("[GET /debugger/{executionId}][%d] getWorkflowDebugLocationOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowDebugLocationOK) String() string {
	return fmt.Sprintf("[GET /debugger/{executionId}][%d] getWorkflowDebugLocationOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowDebugLocationOK) GetPayload() *models.DebugLocation {
	return o.Payload
}

func (o *GetWorkflowDebugLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugLocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDebugLocationUnauthorized creates a GetWorkflowDebugLocationUnauthorized with default headers values
func NewGetWorkflowDebugLocationUnauthorized() *GetWorkflowDebugLocationUnauthorized {
	return &GetWorkflowDebugLocationUnauthorized{}
}

/*
GetWorkflowDebugLocationUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetWorkflowDebugLocationUnauthorized struct {
}

// IsSuccess returns true when this get workflow debug location unauthorized response has a 2xx status code
func (o *GetWorkflowDebugLocationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow debug location unauthorized response has a 3xx status code
func (o *GetWorkflowDebugLocationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow debug location unauthorized response has a 4xx status code
func (o *GetWorkflowDebugLocationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow debug location unauthorized response has a 5xx status code
func (o *GetWorkflowDebugLocationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow debug location unauthorized response a status code equal to that given
func (o *GetWorkflowDebugLocationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get workflow debug location unauthorized response
func (o *GetWorkflowDebugLocationUnauthorized) Code() int {
	return 401
}

func (o *GetWorkflowDebugLocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /debugger/{executionId}][%d] getWorkflowDebugLocationUnauthorized ", 401)
}

func (o *GetWorkflowDebugLocationUnauthorized) String() string {
	return fmt.Sprintf("[GET /debugger/{executionId}][%d] getWorkflowDebugLocationUnauthorized ", 401)
}

func (o *GetWorkflowDebugLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowDebugLocationNotFound creates a GetWorkflowDebugLocationNotFound with default headers values
func NewGetWorkflowDebugLocationNotFound() *GetWorkflowDebugLocationNotFound {
	return &GetWorkflowDebugLocationNotFound{}
}

/*
GetWorkflowDebugLocationNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetWorkflowDebugLocationNotFound struct {
}

// IsSuccess returns true when this get workflow debug location not found response has a 2xx status code
func (o *GetWorkflowDebugLocationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow debug location not found response has a 3xx status code
func (o *GetWorkflowDebugLocationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow debug location not found response has a 4xx status code
func (o *GetWorkflowDebugLocationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow debug location not found response has a 5xx status code
func (o *GetWorkflowDebugLocationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow debug location not found response a status code equal to that given
func (o *GetWorkflowDebugLocationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow debug location not found response
func (o *GetWorkflowDebugLocationNotFound) Code() int {
	return 404
}

func (o *GetWorkflowDebugLocationNotFound) Error() string {
	return fmt.Sprintf("[GET /debugger/{executionId}][%d] getWorkflowDebugLocationNotFound ", 404)
}

func (o *GetWorkflowDebugLocationNotFound) String() string {
	return fmt.Sprintf("[GET /debugger/{executionId}][%d] getWorkflowDebugLocationNotFound ", 404)
}

func (o *GetWorkflowDebugLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
