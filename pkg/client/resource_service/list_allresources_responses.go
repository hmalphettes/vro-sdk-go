// Code generated by go-swagger; DO NOT EDIT.

package resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// ListAllresourcesReader is a Reader for the ListAllresources structure.
type ListAllresourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllresourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllresourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAllresourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /resources] listAllresources", response, response.Code())
	}
}

// NewListAllresourcesOK creates a ListAllresourcesOK with default headers values
func NewListAllresourcesOK() *ListAllresourcesOK {
	return &ListAllresourcesOK{}
}

/*
ListAllresourcesOK describes a response with status code 200, with default header values.

The request is successful
*/
type ListAllresourcesOK struct {
	Payload *models.Resources
}

// IsSuccess returns true when this list allresources o k response has a 2xx status code
func (o *ListAllresourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list allresources o k response has a 3xx status code
func (o *ListAllresourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list allresources o k response has a 4xx status code
func (o *ListAllresourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list allresources o k response has a 5xx status code
func (o *ListAllresourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list allresources o k response a status code equal to that given
func (o *ListAllresourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list allresources o k response
func (o *ListAllresourcesOK) Code() int {
	return 200
}

func (o *ListAllresourcesOK) Error() string {
	return fmt.Sprintf("[GET /resources][%d] listAllresourcesOK  %+v", 200, o.Payload)
}

func (o *ListAllresourcesOK) String() string {
	return fmt.Sprintf("[GET /resources][%d] listAllresourcesOK  %+v", 200, o.Payload)
}

func (o *ListAllresourcesOK) GetPayload() *models.Resources {
	return o.Payload
}

func (o *ListAllresourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Resources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllresourcesUnauthorized creates a ListAllresourcesUnauthorized with default headers values
func NewListAllresourcesUnauthorized() *ListAllresourcesUnauthorized {
	return &ListAllresourcesUnauthorized{}
}

/*
ListAllresourcesUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type ListAllresourcesUnauthorized struct {
}

// IsSuccess returns true when this list allresources unauthorized response has a 2xx status code
func (o *ListAllresourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list allresources unauthorized response has a 3xx status code
func (o *ListAllresourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list allresources unauthorized response has a 4xx status code
func (o *ListAllresourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list allresources unauthorized response has a 5xx status code
func (o *ListAllresourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list allresources unauthorized response a status code equal to that given
func (o *ListAllresourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list allresources unauthorized response
func (o *ListAllresourcesUnauthorized) Code() int {
	return 401
}

func (o *ListAllresourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /resources][%d] listAllresourcesUnauthorized ", 401)
}

func (o *ListAllresourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /resources][%d] listAllresourcesUnauthorized ", 401)
}

func (o *ListAllresourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
