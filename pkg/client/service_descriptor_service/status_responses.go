// Code generated by go-swagger; DO NOT EDIT.

package service_descriptor_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// StatusReader is a Reader for the Status structure.
type StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /status] status", response, response.Code())
	}
}

// NewStatusOK creates a StatusOK with default headers values
func NewStatusOK() *StatusOK {
	return &StatusOK{}
}

/*
StatusOK describes a response with status code 200, with default header values.

The request is successful.
*/
type StatusOK struct {
	Payload *models.ServiceRegistryStatus
}

// IsSuccess returns true when this status o k response has a 2xx status code
func (o *StatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status o k response has a 3xx status code
func (o *StatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status o k response has a 4xx status code
func (o *StatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status o k response has a 5xx status code
func (o *StatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status o k response a status code equal to that given
func (o *StatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status o k response
func (o *StatusOK) Code() int {
	return 200
}

func (o *StatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] statusOK  %+v", 200, o.Payload)
}

func (o *StatusOK) String() string {
	return fmt.Sprintf("[GET /status][%d] statusOK  %+v", 200, o.Payload)
}

func (o *StatusOK) GetPayload() *models.ServiceRegistryStatus {
	return o.Payload
}

func (o *StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceRegistryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusUnauthorized creates a StatusUnauthorized with default headers values
func NewStatusUnauthorized() *StatusUnauthorized {
	return &StatusUnauthorized{}
}

/*
StatusUnauthorized describes a response with status code 401, with default header values.

User is not authorized.
*/
type StatusUnauthorized struct {
}

// IsSuccess returns true when this status unauthorized response has a 2xx status code
func (o *StatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status unauthorized response has a 3xx status code
func (o *StatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status unauthorized response has a 4xx status code
func (o *StatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status unauthorized response has a 5xx status code
func (o *StatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status unauthorized response a status code equal to that given
func (o *StatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status unauthorized response
func (o *StatusUnauthorized) Code() int {
	return 401
}

func (o *StatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /status][%d] statusUnauthorized ", 401)
}

func (o *StatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /status][%d] statusUnauthorized ", 401)
}

func (o *StatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
