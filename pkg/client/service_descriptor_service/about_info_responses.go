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

// AboutInfoReader is a Reader for the AboutInfo structure.
type AboutInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AboutInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAboutInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAboutInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /about] aboutInfo", response, response.Code())
	}
}

// NewAboutInfoOK creates a AboutInfoOK with default headers values
func NewAboutInfoOK() *AboutInfoOK {
	return &AboutInfoOK{}
}

/*
AboutInfoOK describes a response with status code 200, with default header values.

The request is successful.
*/
type AboutInfoOK struct {
	Payload *models.AboutInfo
}

// IsSuccess returns true when this about info o k response has a 2xx status code
func (o *AboutInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this about info o k response has a 3xx status code
func (o *AboutInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this about info o k response has a 4xx status code
func (o *AboutInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this about info o k response has a 5xx status code
func (o *AboutInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this about info o k response a status code equal to that given
func (o *AboutInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the about info o k response
func (o *AboutInfoOK) Code() int {
	return 200
}

func (o *AboutInfoOK) Error() string {
	return fmt.Sprintf("[GET /about][%d] aboutInfoOK  %+v", 200, o.Payload)
}

func (o *AboutInfoOK) String() string {
	return fmt.Sprintf("[GET /about][%d] aboutInfoOK  %+v", 200, o.Payload)
}

func (o *AboutInfoOK) GetPayload() *models.AboutInfo {
	return o.Payload
}

func (o *AboutInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AboutInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAboutInfoUnauthorized creates a AboutInfoUnauthorized with default headers values
func NewAboutInfoUnauthorized() *AboutInfoUnauthorized {
	return &AboutInfoUnauthorized{}
}

/*
AboutInfoUnauthorized describes a response with status code 401, with default header values.

User is not authorized.
*/
type AboutInfoUnauthorized struct {
}

// IsSuccess returns true when this about info unauthorized response has a 2xx status code
func (o *AboutInfoUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this about info unauthorized response has a 3xx status code
func (o *AboutInfoUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this about info unauthorized response has a 4xx status code
func (o *AboutInfoUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this about info unauthorized response has a 5xx status code
func (o *AboutInfoUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this about info unauthorized response a status code equal to that given
func (o *AboutInfoUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the about info unauthorized response
func (o *AboutInfoUnauthorized) Code() int {
	return 401
}

func (o *AboutInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /about][%d] aboutInfoUnauthorized ", 401)
}

func (o *AboutInfoUnauthorized) String() string {
	return fmt.Sprintf("[GET /about][%d] aboutInfoUnauthorized ", 401)
}

func (o *AboutInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}