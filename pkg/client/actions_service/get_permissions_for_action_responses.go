// Code generated by go-swagger; DO NOT EDIT.

package actions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPermissionsForActionReader is a Reader for the GetPermissionsForAction structure.
type GetPermissionsForActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsForActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsForActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPermissionsForActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPermissionsForActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /actions/{id}/permissions] getPermissionsForAction", response, response.Code())
	}
}

// NewGetPermissionsForActionOK creates a GetPermissionsForActionOK with default headers values
func NewGetPermissionsForActionOK() *GetPermissionsForActionOK {
	return &GetPermissionsForActionOK{}
}

/*
GetPermissionsForActionOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPermissionsForActionOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this get permissions for action o k response has a 2xx status code
func (o *GetPermissionsForActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permissions for action o k response has a 3xx status code
func (o *GetPermissionsForActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for action o k response has a 4xx status code
func (o *GetPermissionsForActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions for action o k response has a 5xx status code
func (o *GetPermissionsForActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for action o k response a status code equal to that given
func (o *GetPermissionsForActionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get permissions for action o k response
func (o *GetPermissionsForActionOK) Code() int {
	return 200
}

func (o *GetPermissionsForActionOK) Error() string {
	return fmt.Sprintf("[GET /actions/{id}/permissions][%d] getPermissionsForActionOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForActionOK) String() string {
	return fmt.Sprintf("[GET /actions/{id}/permissions][%d] getPermissionsForActionOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForActionOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *GetPermissionsForActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsForActionUnauthorized creates a GetPermissionsForActionUnauthorized with default headers values
func NewGetPermissionsForActionUnauthorized() *GetPermissionsForActionUnauthorized {
	return &GetPermissionsForActionUnauthorized{}
}

/*
GetPermissionsForActionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetPermissionsForActionUnauthorized struct {
}

// IsSuccess returns true when this get permissions for action unauthorized response has a 2xx status code
func (o *GetPermissionsForActionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for action unauthorized response has a 3xx status code
func (o *GetPermissionsForActionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for action unauthorized response has a 4xx status code
func (o *GetPermissionsForActionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for action unauthorized response has a 5xx status code
func (o *GetPermissionsForActionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for action unauthorized response a status code equal to that given
func (o *GetPermissionsForActionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get permissions for action unauthorized response
func (o *GetPermissionsForActionUnauthorized) Code() int {
	return 401
}

func (o *GetPermissionsForActionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions/{id}/permissions][%d] getPermissionsForActionUnauthorized ", 401)
}

func (o *GetPermissionsForActionUnauthorized) String() string {
	return fmt.Sprintf("[GET /actions/{id}/permissions][%d] getPermissionsForActionUnauthorized ", 401)
}

func (o *GetPermissionsForActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPermissionsForActionNotFound creates a GetPermissionsForActionNotFound with default headers values
func NewGetPermissionsForActionNotFound() *GetPermissionsForActionNotFound {
	return &GetPermissionsForActionNotFound{}
}

/*
GetPermissionsForActionNotFound describes a response with status code 404, with default header values.

Can not find an action with the specified id
*/
type GetPermissionsForActionNotFound struct {
}

// IsSuccess returns true when this get permissions for action not found response has a 2xx status code
func (o *GetPermissionsForActionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for action not found response has a 3xx status code
func (o *GetPermissionsForActionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for action not found response has a 4xx status code
func (o *GetPermissionsForActionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for action not found response has a 5xx status code
func (o *GetPermissionsForActionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for action not found response a status code equal to that given
func (o *GetPermissionsForActionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get permissions for action not found response
func (o *GetPermissionsForActionNotFound) Code() int {
	return 404
}

func (o *GetPermissionsForActionNotFound) Error() string {
	return fmt.Sprintf("[GET /actions/{id}/permissions][%d] getPermissionsForActionNotFound ", 404)
}

func (o *GetPermissionsForActionNotFound) String() string {
	return fmt.Sprintf("[GET /actions/{id}/permissions][%d] getPermissionsForActionNotFound ", 404)
}

func (o *GetPermissionsForActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
