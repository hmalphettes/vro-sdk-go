// Code generated by go-swagger; DO NOT EDIT.

package configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPermissionsForConfigurationReader is a Reader for the GetPermissionsForConfiguration structure.
type GetPermissionsForConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsForConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsForConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPermissionsForConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPermissionsForConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /configurations/{id}/permissions] getPermissionsForConfiguration", response, response.Code())
	}
}

// NewGetPermissionsForConfigurationOK creates a GetPermissionsForConfigurationOK with default headers values
func NewGetPermissionsForConfigurationOK() *GetPermissionsForConfigurationOK {
	return &GetPermissionsForConfigurationOK{}
}

/*
GetPermissionsForConfigurationOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPermissionsForConfigurationOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this get permissions for configuration o k response has a 2xx status code
func (o *GetPermissionsForConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permissions for configuration o k response has a 3xx status code
func (o *GetPermissionsForConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for configuration o k response has a 4xx status code
func (o *GetPermissionsForConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions for configuration o k response has a 5xx status code
func (o *GetPermissionsForConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for configuration o k response a status code equal to that given
func (o *GetPermissionsForConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get permissions for configuration o k response
func (o *GetPermissionsForConfigurationOK) Code() int {
	return 200
}

func (o *GetPermissionsForConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions][%d] getPermissionsForConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForConfigurationOK) String() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions][%d] getPermissionsForConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForConfigurationOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *GetPermissionsForConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsForConfigurationUnauthorized creates a GetPermissionsForConfigurationUnauthorized with default headers values
func NewGetPermissionsForConfigurationUnauthorized() *GetPermissionsForConfigurationUnauthorized {
	return &GetPermissionsForConfigurationUnauthorized{}
}

/*
GetPermissionsForConfigurationUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type GetPermissionsForConfigurationUnauthorized struct {
}

// IsSuccess returns true when this get permissions for configuration unauthorized response has a 2xx status code
func (o *GetPermissionsForConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for configuration unauthorized response has a 3xx status code
func (o *GetPermissionsForConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for configuration unauthorized response has a 4xx status code
func (o *GetPermissionsForConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for configuration unauthorized response has a 5xx status code
func (o *GetPermissionsForConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for configuration unauthorized response a status code equal to that given
func (o *GetPermissionsForConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get permissions for configuration unauthorized response
func (o *GetPermissionsForConfigurationUnauthorized) Code() int {
	return 401
}

func (o *GetPermissionsForConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions][%d] getPermissionsForConfigurationUnauthorized ", 401)
}

func (o *GetPermissionsForConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions][%d] getPermissionsForConfigurationUnauthorized ", 401)
}

func (o *GetPermissionsForConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPermissionsForConfigurationNotFound creates a GetPermissionsForConfigurationNotFound with default headers values
func NewGetPermissionsForConfigurationNotFound() *GetPermissionsForConfigurationNotFound {
	return &GetPermissionsForConfigurationNotFound{}
}

/*
GetPermissionsForConfigurationNotFound describes a response with status code 404, with default header values.

Cannot find configuration with the specified ID
*/
type GetPermissionsForConfigurationNotFound struct {
}

// IsSuccess returns true when this get permissions for configuration not found response has a 2xx status code
func (o *GetPermissionsForConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for configuration not found response has a 3xx status code
func (o *GetPermissionsForConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for configuration not found response has a 4xx status code
func (o *GetPermissionsForConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for configuration not found response has a 5xx status code
func (o *GetPermissionsForConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for configuration not found response a status code equal to that given
func (o *GetPermissionsForConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get permissions for configuration not found response
func (o *GetPermissionsForConfigurationNotFound) Code() int {
	return 404
}

func (o *GetPermissionsForConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions][%d] getPermissionsForConfigurationNotFound ", 404)
}

func (o *GetPermissionsForConfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions][%d] getPermissionsForConfigurationNotFound ", 404)
}

func (o *GetPermissionsForConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
