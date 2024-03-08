// Code generated by go-swagger; DO NOT EDIT.

package resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPermissionsForResourceReader is a Reader for the GetPermissionsForResource structure.
type GetPermissionsForResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsForResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsForResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPermissionsForResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /resources/{id}/permissions] getPermissionsForResource", response, response.Code())
	}
}

// NewGetPermissionsForResourceOK creates a GetPermissionsForResourceOK with default headers values
func NewGetPermissionsForResourceOK() *GetPermissionsForResourceOK {
	return &GetPermissionsForResourceOK{}
}

/*
GetPermissionsForResourceOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPermissionsForResourceOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this get permissions for resource o k response has a 2xx status code
func (o *GetPermissionsForResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permissions for resource o k response has a 3xx status code
func (o *GetPermissionsForResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for resource o k response has a 4xx status code
func (o *GetPermissionsForResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions for resource o k response has a 5xx status code
func (o *GetPermissionsForResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for resource o k response a status code equal to that given
func (o *GetPermissionsForResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get permissions for resource o k response
func (o *GetPermissionsForResourceOK) Code() int {
	return 200
}

func (o *GetPermissionsForResourceOK) Error() string {
	return fmt.Sprintf("[GET /resources/{id}/permissions][%d] getPermissionsForResourceOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForResourceOK) String() string {
	return fmt.Sprintf("[GET /resources/{id}/permissions][%d] getPermissionsForResourceOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForResourceOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *GetPermissionsForResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsForResourceUnauthorized creates a GetPermissionsForResourceUnauthorized with default headers values
func NewGetPermissionsForResourceUnauthorized() *GetPermissionsForResourceUnauthorized {
	return &GetPermissionsForResourceUnauthorized{}
}

/*
GetPermissionsForResourceUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type GetPermissionsForResourceUnauthorized struct {
}

// IsSuccess returns true when this get permissions for resource unauthorized response has a 2xx status code
func (o *GetPermissionsForResourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for resource unauthorized response has a 3xx status code
func (o *GetPermissionsForResourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for resource unauthorized response has a 4xx status code
func (o *GetPermissionsForResourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for resource unauthorized response has a 5xx status code
func (o *GetPermissionsForResourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for resource unauthorized response a status code equal to that given
func (o *GetPermissionsForResourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get permissions for resource unauthorized response
func (o *GetPermissionsForResourceUnauthorized) Code() int {
	return 401
}

func (o *GetPermissionsForResourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /resources/{id}/permissions][%d] getPermissionsForResourceUnauthorized ", 401)
}

func (o *GetPermissionsForResourceUnauthorized) String() string {
	return fmt.Sprintf("[GET /resources/{id}/permissions][%d] getPermissionsForResourceUnauthorized ", 401)
}

func (o *GetPermissionsForResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
