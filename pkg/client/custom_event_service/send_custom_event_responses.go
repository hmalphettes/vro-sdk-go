// Code generated by go-swagger; DO NOT EDIT.

package custom_event_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SendCustomEventReader is a Reader for the SendCustomEvent structure.
type SendCustomEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendCustomEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendCustomEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSendCustomEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /customevent/{eventname}] sendCustomEvent", response, response.Code())
	}
}

// NewSendCustomEventOK creates a SendCustomEventOK with default headers values
func NewSendCustomEventOK() *SendCustomEventOK {
	return &SendCustomEventOK{}
}

/*
SendCustomEventOK describes a response with status code 200, with default header values.

The request is successful
*/
type SendCustomEventOK struct {
}

// IsSuccess returns true when this send custom event o k response has a 2xx status code
func (o *SendCustomEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send custom event o k response has a 3xx status code
func (o *SendCustomEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send custom event o k response has a 4xx status code
func (o *SendCustomEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this send custom event o k response has a 5xx status code
func (o *SendCustomEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this send custom event o k response a status code equal to that given
func (o *SendCustomEventOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the send custom event o k response
func (o *SendCustomEventOK) Code() int {
	return 200
}

func (o *SendCustomEventOK) Error() string {
	return fmt.Sprintf("[POST /customevent/{eventname}][%d] sendCustomEventOK ", 200)
}

func (o *SendCustomEventOK) String() string {
	return fmt.Sprintf("[POST /customevent/{eventname}][%d] sendCustomEventOK ", 200)
}

func (o *SendCustomEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendCustomEventUnauthorized creates a SendCustomEventUnauthorized with default headers values
func NewSendCustomEventUnauthorized() *SendCustomEventUnauthorized {
	return &SendCustomEventUnauthorized{}
}

/*
SendCustomEventUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type SendCustomEventUnauthorized struct {
}

// IsSuccess returns true when this send custom event unauthorized response has a 2xx status code
func (o *SendCustomEventUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send custom event unauthorized response has a 3xx status code
func (o *SendCustomEventUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send custom event unauthorized response has a 4xx status code
func (o *SendCustomEventUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this send custom event unauthorized response has a 5xx status code
func (o *SendCustomEventUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this send custom event unauthorized response a status code equal to that given
func (o *SendCustomEventUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the send custom event unauthorized response
func (o *SendCustomEventUnauthorized) Code() int {
	return 401
}

func (o *SendCustomEventUnauthorized) Error() string {
	return fmt.Sprintf("[POST /customevent/{eventname}][%d] sendCustomEventUnauthorized ", 401)
}

func (o *SendCustomEventUnauthorized) String() string {
	return fmt.Sprintf("[POST /customevent/{eventname}][%d] sendCustomEventUnauthorized ", 401)
}

func (o *SendCustomEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
