// Code generated by go-swagger; DO NOT EDIT.

package service_descriptor_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// EnumerateServicesReader is a Reader for the EnumerateServices structure.
type EnumerateServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumerateServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnumerateServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnumerateServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /] enumerateServices", response, response.Code())
	}
}

// NewEnumerateServicesOK creates a EnumerateServicesOK with default headers values
func NewEnumerateServicesOK() *EnumerateServicesOK {
	return &EnumerateServicesOK{}
}

/*
EnumerateServicesOK describes a response with status code 200, with default header values.

The request is successful.
*/
type EnumerateServicesOK struct {
	Payload *models.ServiceDescriptors
}

// IsSuccess returns true when this enumerate services o k response has a 2xx status code
func (o *EnumerateServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enumerate services o k response has a 3xx status code
func (o *EnumerateServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate services o k response has a 4xx status code
func (o *EnumerateServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enumerate services o k response has a 5xx status code
func (o *EnumerateServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate services o k response a status code equal to that given
func (o *EnumerateServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enumerate services o k response
func (o *EnumerateServicesOK) Code() int {
	return 200
}

func (o *EnumerateServicesOK) Error() string {
	return fmt.Sprintf("[GET /][%d] enumerateServicesOK  %+v", 200, o.Payload)
}

func (o *EnumerateServicesOK) String() string {
	return fmt.Sprintf("[GET /][%d] enumerateServicesOK  %+v", 200, o.Payload)
}

func (o *EnumerateServicesOK) GetPayload() *models.ServiceDescriptors {
	return o.Payload
}

func (o *EnumerateServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDescriptors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerateServicesUnauthorized creates a EnumerateServicesUnauthorized with default headers values
func NewEnumerateServicesUnauthorized() *EnumerateServicesUnauthorized {
	return &EnumerateServicesUnauthorized{}
}

/*
EnumerateServicesUnauthorized describes a response with status code 401, with default header values.

User is not authorized.
*/
type EnumerateServicesUnauthorized struct {
}

// IsSuccess returns true when this enumerate services unauthorized response has a 2xx status code
func (o *EnumerateServicesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate services unauthorized response has a 3xx status code
func (o *EnumerateServicesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate services unauthorized response has a 4xx status code
func (o *EnumerateServicesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate services unauthorized response has a 5xx status code
func (o *EnumerateServicesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate services unauthorized response a status code equal to that given
func (o *EnumerateServicesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enumerate services unauthorized response
func (o *EnumerateServicesUnauthorized) Code() int {
	return 401
}

func (o *EnumerateServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /][%d] enumerateServicesUnauthorized ", 401)
}

func (o *EnumerateServicesUnauthorized) String() string {
	return fmt.Sprintf("[GET /][%d] enumerateServicesUnauthorized ", 401)
}

func (o *EnumerateServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}