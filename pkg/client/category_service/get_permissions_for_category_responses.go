// Code generated by go-swagger; DO NOT EDIT.

package category_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetPermissionsForCategoryReader is a Reader for the GetPermissionsForCategory structure.
type GetPermissionsForCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsForCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsForCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPermissionsForCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPermissionsForCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /categories/{id}/permissions] getPermissionsForCategory", response, response.Code())
	}
}

// NewGetPermissionsForCategoryOK creates a GetPermissionsForCategoryOK with default headers values
func NewGetPermissionsForCategoryOK() *GetPermissionsForCategoryOK {
	return &GetPermissionsForCategoryOK{}
}

/*
GetPermissionsForCategoryOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPermissionsForCategoryOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this get permissions for category o k response has a 2xx status code
func (o *GetPermissionsForCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permissions for category o k response has a 3xx status code
func (o *GetPermissionsForCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for category o k response has a 4xx status code
func (o *GetPermissionsForCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions for category o k response has a 5xx status code
func (o *GetPermissionsForCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for category o k response a status code equal to that given
func (o *GetPermissionsForCategoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get permissions for category o k response
func (o *GetPermissionsForCategoryOK) Code() int {
	return 200
}

func (o *GetPermissionsForCategoryOK) Error() string {
	return fmt.Sprintf("[GET /categories/{id}/permissions][%d] getPermissionsForCategoryOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForCategoryOK) String() string {
	return fmt.Sprintf("[GET /categories/{id}/permissions][%d] getPermissionsForCategoryOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsForCategoryOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *GetPermissionsForCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsForCategoryUnauthorized creates a GetPermissionsForCategoryUnauthorized with default headers values
func NewGetPermissionsForCategoryUnauthorized() *GetPermissionsForCategoryUnauthorized {
	return &GetPermissionsForCategoryUnauthorized{}
}

/*
GetPermissionsForCategoryUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type GetPermissionsForCategoryUnauthorized struct {
}

// IsSuccess returns true when this get permissions for category unauthorized response has a 2xx status code
func (o *GetPermissionsForCategoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for category unauthorized response has a 3xx status code
func (o *GetPermissionsForCategoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for category unauthorized response has a 4xx status code
func (o *GetPermissionsForCategoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for category unauthorized response has a 5xx status code
func (o *GetPermissionsForCategoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for category unauthorized response a status code equal to that given
func (o *GetPermissionsForCategoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get permissions for category unauthorized response
func (o *GetPermissionsForCategoryUnauthorized) Code() int {
	return 401
}

func (o *GetPermissionsForCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /categories/{id}/permissions][%d] getPermissionsForCategoryUnauthorized ", 401)
}

func (o *GetPermissionsForCategoryUnauthorized) String() string {
	return fmt.Sprintf("[GET /categories/{id}/permissions][%d] getPermissionsForCategoryUnauthorized ", 401)
}

func (o *GetPermissionsForCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPermissionsForCategoryNotFound creates a GetPermissionsForCategoryNotFound with default headers values
func NewGetPermissionsForCategoryNotFound() *GetPermissionsForCategoryNotFound {
	return &GetPermissionsForCategoryNotFound{}
}

/*
GetPermissionsForCategoryNotFound describes a response with status code 404, with default header values.

Cannot find a category with the specified ID,
*/
type GetPermissionsForCategoryNotFound struct {
}

// IsSuccess returns true when this get permissions for category not found response has a 2xx status code
func (o *GetPermissionsForCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions for category not found response has a 3xx status code
func (o *GetPermissionsForCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions for category not found response has a 4xx status code
func (o *GetPermissionsForCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions for category not found response has a 5xx status code
func (o *GetPermissionsForCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions for category not found response a status code equal to that given
func (o *GetPermissionsForCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get permissions for category not found response
func (o *GetPermissionsForCategoryNotFound) Code() int {
	return 404
}

func (o *GetPermissionsForCategoryNotFound) Error() string {
	return fmt.Sprintf("[GET /categories/{id}/permissions][%d] getPermissionsForCategoryNotFound ", 404)
}

func (o *GetPermissionsForCategoryNotFound) String() string {
	return fmt.Sprintf("[GET /categories/{id}/permissions][%d] getPermissionsForCategoryNotFound ", 404)
}

func (o *GetPermissionsForCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
