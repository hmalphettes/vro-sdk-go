// Code generated by go-swagger; DO NOT EDIT.

package content_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ListWorkflowsReader is a Reader for the ListWorkflows structure.
type ListWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListWorkflowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /content/workflows] listWorkflows", response, response.Code())
	}
}

// NewListWorkflowsOK creates a ListWorkflowsOK with default headers values
func NewListWorkflowsOK() *ListWorkflowsOK {
	return &ListWorkflowsOK{}
}

/*
ListWorkflowsOK describes a response with status code 200, with default header values.

The request is successful
*/
type ListWorkflowsOK struct {
	Payload *models.InventoryItems
}

// IsSuccess returns true when this list workflows o k response has a 2xx status code
func (o *ListWorkflowsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list workflows o k response has a 3xx status code
func (o *ListWorkflowsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list workflows o k response has a 4xx status code
func (o *ListWorkflowsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list workflows o k response has a 5xx status code
func (o *ListWorkflowsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list workflows o k response a status code equal to that given
func (o *ListWorkflowsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list workflows o k response
func (o *ListWorkflowsOK) Code() int {
	return 200
}

func (o *ListWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /content/workflows][%d] listWorkflowsOK  %+v", 200, o.Payload)
}

func (o *ListWorkflowsOK) String() string {
	return fmt.Sprintf("[GET /content/workflows][%d] listWorkflowsOK  %+v", 200, o.Payload)
}

func (o *ListWorkflowsOK) GetPayload() *models.InventoryItems {
	return o.Payload
}

func (o *ListWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItems)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListWorkflowsUnauthorized creates a ListWorkflowsUnauthorized with default headers values
func NewListWorkflowsUnauthorized() *ListWorkflowsUnauthorized {
	return &ListWorkflowsUnauthorized{}
}

/*
ListWorkflowsUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type ListWorkflowsUnauthorized struct {
}

// IsSuccess returns true when this list workflows unauthorized response has a 2xx status code
func (o *ListWorkflowsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list workflows unauthorized response has a 3xx status code
func (o *ListWorkflowsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list workflows unauthorized response has a 4xx status code
func (o *ListWorkflowsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list workflows unauthorized response has a 5xx status code
func (o *ListWorkflowsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list workflows unauthorized response a status code equal to that given
func (o *ListWorkflowsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list workflows unauthorized response
func (o *ListWorkflowsUnauthorized) Code() int {
	return 401
}

func (o *ListWorkflowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /content/workflows][%d] listWorkflowsUnauthorized ", 401)
}

func (o *ListWorkflowsUnauthorized) String() string {
	return fmt.Sprintf("[GET /content/workflows][%d] listWorkflowsUnauthorized ", 401)
}

func (o *ListWorkflowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}