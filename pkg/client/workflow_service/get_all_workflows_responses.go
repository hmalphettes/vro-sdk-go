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

// GetAllWorkflowsReader is a Reader for the GetAllWorkflows structure.
type GetAllWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllWorkflowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows] getAllWorkflows", response, response.Code())
	}
}

// NewGetAllWorkflowsOK creates a GetAllWorkflowsOK with default headers values
func NewGetAllWorkflowsOK() *GetAllWorkflowsOK {
	return &GetAllWorkflowsOK{}
}

/*
GetAllWorkflowsOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetAllWorkflowsOK struct {
	Payload *models.InventoryItems
}

// IsSuccess returns true when this get all workflows o k response has a 2xx status code
func (o *GetAllWorkflowsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all workflows o k response has a 3xx status code
func (o *GetAllWorkflowsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all workflows o k response has a 4xx status code
func (o *GetAllWorkflowsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all workflows o k response has a 5xx status code
func (o *GetAllWorkflowsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all workflows o k response a status code equal to that given
func (o *GetAllWorkflowsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all workflows o k response
func (o *GetAllWorkflowsOK) Code() int {
	return 200
}

func (o *GetAllWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /workflows][%d] getAllWorkflowsOK  %+v", 200, o.Payload)
}

func (o *GetAllWorkflowsOK) String() string {
	return fmt.Sprintf("[GET /workflows][%d] getAllWorkflowsOK  %+v", 200, o.Payload)
}

func (o *GetAllWorkflowsOK) GetPayload() *models.InventoryItems {
	return o.Payload
}

func (o *GetAllWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItems)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllWorkflowsUnauthorized creates a GetAllWorkflowsUnauthorized with default headers values
func NewGetAllWorkflowsUnauthorized() *GetAllWorkflowsUnauthorized {
	return &GetAllWorkflowsUnauthorized{}
}

/*
GetAllWorkflowsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetAllWorkflowsUnauthorized struct {
}

// IsSuccess returns true when this get all workflows unauthorized response has a 2xx status code
func (o *GetAllWorkflowsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all workflows unauthorized response has a 3xx status code
func (o *GetAllWorkflowsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all workflows unauthorized response has a 4xx status code
func (o *GetAllWorkflowsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all workflows unauthorized response has a 5xx status code
func (o *GetAllWorkflowsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all workflows unauthorized response a status code equal to that given
func (o *GetAllWorkflowsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all workflows unauthorized response
func (o *GetAllWorkflowsUnauthorized) Code() int {
	return 401
}

func (o *GetAllWorkflowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows][%d] getAllWorkflowsUnauthorized ", 401)
}

func (o *GetAllWorkflowsUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows][%d] getAllWorkflowsUnauthorized ", 401)
}

func (o *GetAllWorkflowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
