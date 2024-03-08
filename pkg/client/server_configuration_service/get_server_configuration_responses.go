// Code generated by go-swagger; DO NOT EDIT.

package server_configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetServerConfigurationReader is a Reader for the GetServerConfiguration structure.
type GetServerConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetServerConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServerConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /server-configuration] getServerConfiguration", response, response.Code())
	}
}

// NewGetServerConfigurationOK creates a GetServerConfigurationOK with default headers values
func NewGetServerConfigurationOK() *GetServerConfigurationOK {
	return &GetServerConfigurationOK{}
}

/*
GetServerConfigurationOK describes a response with status code 200, with default header values.

The request is successful.
*/
type GetServerConfigurationOK struct {
	Payload *models.ConfigEntries
}

// IsSuccess returns true when this get server configuration o k response has a 2xx status code
func (o *GetServerConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get server configuration o k response has a 3xx status code
func (o *GetServerConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server configuration o k response has a 4xx status code
func (o *GetServerConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get server configuration o k response has a 5xx status code
func (o *GetServerConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get server configuration o k response a status code equal to that given
func (o *GetServerConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get server configuration o k response
func (o *GetServerConfigurationOK) Code() int {
	return 200
}

func (o *GetServerConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /server-configuration][%d] getServerConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetServerConfigurationOK) String() string {
	return fmt.Sprintf("[GET /server-configuration][%d] getServerConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetServerConfigurationOK) GetPayload() *models.ConfigEntries {
	return o.Payload
}

func (o *GetServerConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerConfigurationUnauthorized creates a GetServerConfigurationUnauthorized with default headers values
func NewGetServerConfigurationUnauthorized() *GetServerConfigurationUnauthorized {
	return &GetServerConfigurationUnauthorized{}
}

/*
GetServerConfigurationUnauthorized describes a response with status code 401, with default header values.

User is not authenticated.
*/
type GetServerConfigurationUnauthorized struct {
}

// IsSuccess returns true when this get server configuration unauthorized response has a 2xx status code
func (o *GetServerConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get server configuration unauthorized response has a 3xx status code
func (o *GetServerConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server configuration unauthorized response has a 4xx status code
func (o *GetServerConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get server configuration unauthorized response has a 5xx status code
func (o *GetServerConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get server configuration unauthorized response a status code equal to that given
func (o *GetServerConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get server configuration unauthorized response
func (o *GetServerConfigurationUnauthorized) Code() int {
	return 401
}

func (o *GetServerConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /server-configuration][%d] getServerConfigurationUnauthorized ", 401)
}

func (o *GetServerConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[GET /server-configuration][%d] getServerConfigurationUnauthorized ", 401)
}

func (o *GetServerConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServerConfigurationForbidden creates a GetServerConfigurationForbidden with default headers values
func NewGetServerConfigurationForbidden() *GetServerConfigurationForbidden {
	return &GetServerConfigurationForbidden{}
}

/*
GetServerConfigurationForbidden describes a response with status code 403, with default header values.

User is not authorized.
*/
type GetServerConfigurationForbidden struct {
}

// IsSuccess returns true when this get server configuration forbidden response has a 2xx status code
func (o *GetServerConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get server configuration forbidden response has a 3xx status code
func (o *GetServerConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server configuration forbidden response has a 4xx status code
func (o *GetServerConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get server configuration forbidden response has a 5xx status code
func (o *GetServerConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get server configuration forbidden response a status code equal to that given
func (o *GetServerConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get server configuration forbidden response
func (o *GetServerConfigurationForbidden) Code() int {
	return 403
}

func (o *GetServerConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /server-configuration][%d] getServerConfigurationForbidden ", 403)
}

func (o *GetServerConfigurationForbidden) String() string {
	return fmt.Sprintf("[GET /server-configuration][%d] getServerConfigurationForbidden ", 403)
}

func (o *GetServerConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
