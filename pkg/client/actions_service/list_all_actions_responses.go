// Code generated by go-swagger; DO NOT EDIT.

package actions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// ListAllActionsReader is a Reader for the ListAllActions structure.
type ListAllActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAllActionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /actions] listAllActions", response, response.Code())
	}
}

// NewListAllActionsOK creates a ListAllActionsOK with default headers values
func NewListAllActionsOK() *ListAllActionsOK {
	return &ListAllActionsOK{}
}

/*
ListAllActionsOK describes a response with status code 200, with default header values.

The request is successful
*/
type ListAllActionsOK struct {
	Payload *models.Actions
}

// IsSuccess returns true when this list all actions o k response has a 2xx status code
func (o *ListAllActionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list all actions o k response has a 3xx status code
func (o *ListAllActionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all actions o k response has a 4xx status code
func (o *ListAllActionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list all actions o k response has a 5xx status code
func (o *ListAllActionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list all actions o k response a status code equal to that given
func (o *ListAllActionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list all actions o k response
func (o *ListAllActionsOK) Code() int {
	return 200
}

func (o *ListAllActionsOK) Error() string {
	return fmt.Sprintf("[GET /actions][%d] listAllActionsOK  %+v", 200, o.Payload)
}

func (o *ListAllActionsOK) String() string {
	return fmt.Sprintf("[GET /actions][%d] listAllActionsOK  %+v", 200, o.Payload)
}

func (o *ListAllActionsOK) GetPayload() *models.Actions {
	return o.Payload
}

func (o *ListAllActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Actions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllActionsUnauthorized creates a ListAllActionsUnauthorized with default headers values
func NewListAllActionsUnauthorized() *ListAllActionsUnauthorized {
	return &ListAllActionsUnauthorized{}
}

/*
ListAllActionsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type ListAllActionsUnauthorized struct {
}

// IsSuccess returns true when this list all actions unauthorized response has a 2xx status code
func (o *ListAllActionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list all actions unauthorized response has a 3xx status code
func (o *ListAllActionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all actions unauthorized response has a 4xx status code
func (o *ListAllActionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list all actions unauthorized response has a 5xx status code
func (o *ListAllActionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list all actions unauthorized response a status code equal to that given
func (o *ListAllActionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list all actions unauthorized response
func (o *ListAllActionsUnauthorized) Code() int {
	return 401
}

func (o *ListAllActionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions][%d] listAllActionsUnauthorized ", 401)
}

func (o *ListAllActionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /actions][%d] listAllActionsUnauthorized ", 401)
}

func (o *ListAllActionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
