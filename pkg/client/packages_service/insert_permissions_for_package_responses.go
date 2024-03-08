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

// InsertPermissionsForPackageReader is a Reader for the InsertPermissionsForPackage structure.
type InsertPermissionsForPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertPermissionsForPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInsertPermissionsForPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewInsertPermissionsForPackageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInsertPermissionsForPackageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInsertPermissionsForPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /packages/{packageName}/permissions] insertPermissionsForPackage", response, response.Code())
	}
}

// NewInsertPermissionsForPackageOK creates a InsertPermissionsForPackageOK with default headers values
func NewInsertPermissionsForPackageOK() *InsertPermissionsForPackageOK {
	return &InsertPermissionsForPackageOK{}
}

/*
InsertPermissionsForPackageOK describes a response with status code 200, with default header values.

successful operation
*/
type InsertPermissionsForPackageOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this insert permissions for package o k response has a 2xx status code
func (o *InsertPermissionsForPackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert permissions for package o k response has a 3xx status code
func (o *InsertPermissionsForPackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for package o k response has a 4xx status code
func (o *InsertPermissionsForPackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert permissions for package o k response has a 5xx status code
func (o *InsertPermissionsForPackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for package o k response a status code equal to that given
func (o *InsertPermissionsForPackageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the insert permissions for package o k response
func (o *InsertPermissionsForPackageOK) Code() int {
	return 200
}

func (o *InsertPermissionsForPackageOK) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageOK  %+v", 200, o.Payload)
}

func (o *InsertPermissionsForPackageOK) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageOK  %+v", 200, o.Payload)
}

func (o *InsertPermissionsForPackageOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *InsertPermissionsForPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertPermissionsForPackageCreated creates a InsertPermissionsForPackageCreated with default headers values
func NewInsertPermissionsForPackageCreated() *InsertPermissionsForPackageCreated {
	return &InsertPermissionsForPackageCreated{}
}

/*
InsertPermissionsForPackageCreated describes a response with status code 201, with default header values.

The request is successful
*/
type InsertPermissionsForPackageCreated struct {
}

// IsSuccess returns true when this insert permissions for package created response has a 2xx status code
func (o *InsertPermissionsForPackageCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert permissions for package created response has a 3xx status code
func (o *InsertPermissionsForPackageCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for package created response has a 4xx status code
func (o *InsertPermissionsForPackageCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert permissions for package created response has a 5xx status code
func (o *InsertPermissionsForPackageCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for package created response a status code equal to that given
func (o *InsertPermissionsForPackageCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the insert permissions for package created response
func (o *InsertPermissionsForPackageCreated) Code() int {
	return 201
}

func (o *InsertPermissionsForPackageCreated) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageCreated ", 201)
}

func (o *InsertPermissionsForPackageCreated) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageCreated ", 201)
}

func (o *InsertPermissionsForPackageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForPackageBadRequest creates a InsertPermissionsForPackageBadRequest with default headers values
func NewInsertPermissionsForPackageBadRequest() *InsertPermissionsForPackageBadRequest {
	return &InsertPermissionsForPackageBadRequest{}
}

/*
InsertPermissionsForPackageBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type InsertPermissionsForPackageBadRequest struct {
}

// IsSuccess returns true when this insert permissions for package bad request response has a 2xx status code
func (o *InsertPermissionsForPackageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for package bad request response has a 3xx status code
func (o *InsertPermissionsForPackageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for package bad request response has a 4xx status code
func (o *InsertPermissionsForPackageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for package bad request response has a 5xx status code
func (o *InsertPermissionsForPackageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for package bad request response a status code equal to that given
func (o *InsertPermissionsForPackageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the insert permissions for package bad request response
func (o *InsertPermissionsForPackageBadRequest) Code() int {
	return 400
}

func (o *InsertPermissionsForPackageBadRequest) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageBadRequest ", 400)
}

func (o *InsertPermissionsForPackageBadRequest) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageBadRequest ", 400)
}

func (o *InsertPermissionsForPackageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForPackageUnauthorized creates a InsertPermissionsForPackageUnauthorized with default headers values
func NewInsertPermissionsForPackageUnauthorized() *InsertPermissionsForPackageUnauthorized {
	return &InsertPermissionsForPackageUnauthorized{}
}

/*
InsertPermissionsForPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type InsertPermissionsForPackageUnauthorized struct {
}

// IsSuccess returns true when this insert permissions for package unauthorized response has a 2xx status code
func (o *InsertPermissionsForPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for package unauthorized response has a 3xx status code
func (o *InsertPermissionsForPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for package unauthorized response has a 4xx status code
func (o *InsertPermissionsForPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for package unauthorized response has a 5xx status code
func (o *InsertPermissionsForPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for package unauthorized response a status code equal to that given
func (o *InsertPermissionsForPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the insert permissions for package unauthorized response
func (o *InsertPermissionsForPackageUnauthorized) Code() int {
	return 401
}

func (o *InsertPermissionsForPackageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageUnauthorized ", 401)
}

func (o *InsertPermissionsForPackageUnauthorized) String() string {
	return fmt.Sprintf("[POST /packages/{packageName}/permissions][%d] insertPermissionsForPackageUnauthorized ", 401)
}

func (o *InsertPermissionsForPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
