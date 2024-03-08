// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetPermissionsForPackageReader is a Reader for the GetPermissionsForPackage structure.
type GetPermissionsForPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsForPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsForPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPermissionsForPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /packages/{packageName}/permissions] getPermissionsForPackage", response, response.Code())
	}
}

// NewGetPermissionsForPackageOK creates a GetPermissionsForPackageOK with default headers values
func NewGetPermissionsForPackageOK() *GetPermissionsForPackageOK {
	return &GetPermissionsForPackageOK{}
}

/*
GetPermissionsForPackageOK describes a response with status code 200, with default header values.

The request is successful.
*/
type GetPermissionsForPackageOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this get permissions for package o k response has a 2xx status code
func (o *GetPermissionsForPackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permissions for package o k response has a 3xx status code
func (o *GetPermissionsForPackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for package o k response has a 4xx status code
func (o *GetPermissionsForPackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions for package o k response has a 5xx status code
func (o *GetPermissionsForPackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for package o k response a status code equal to that given
func (o *GetPermissionsForPackageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get permissions for package o k response
func (o *GetPermissionsForPackageOK) Code() int {
	return 200
}

func (o *GetPermissionsForPackageOK) Error() string {
	return fmt.Sprintf("[GET /packages/{packageName}/permissions][%d] getPermissionsForPackageOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForPackageOK) String() string {
	return fmt.Sprintf("[GET /packages/{packageName}/permissions][%d] getPermissionsForPackageOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForPackageOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *GetPermissionsForPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsForPackageUnauthorized creates a GetPermissionsForPackageUnauthorized with default headers values
func NewGetPermissionsForPackageUnauthorized() *GetPermissionsForPackageUnauthorized {
	return &GetPermissionsForPackageUnauthorized{}
}

/*
GetPermissionsForPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type GetPermissionsForPackageUnauthorized struct {
}

// IsSuccess returns true when this get permissions for package unauthorized response has a 2xx status code
func (o *GetPermissionsForPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for package unauthorized response has a 3xx status code
func (o *GetPermissionsForPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for package unauthorized response has a 4xx status code
func (o *GetPermissionsForPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for package unauthorized response has a 5xx status code
func (o *GetPermissionsForPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for package unauthorized response a status code equal to that given
func (o *GetPermissionsForPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get permissions for package unauthorized response
func (o *GetPermissionsForPackageUnauthorized) Code() int {
	return 401
}

func (o *GetPermissionsForPackageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /packages/{packageName}/permissions][%d] getPermissionsForPackageUnauthorized ", 401)
}

func (o *GetPermissionsForPackageUnauthorized) String() string {
	return fmt.Sprintf("[GET /packages/{packageName}/permissions][%d] getPermissionsForPackageUnauthorized ", 401)
}

func (o *GetPermissionsForPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
