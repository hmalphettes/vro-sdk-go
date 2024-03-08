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

// InsertPermissionsForConfigurationReader is a Reader for the InsertPermissionsForConfiguration structure.
type InsertPermissionsForConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertPermissionsForConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInsertPermissionsForConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewInsertPermissionsForConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInsertPermissionsForConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInsertPermissionsForConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInsertPermissionsForConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /configurations/{id}/permissions] insertPermissionsForConfiguration", response, response.Code())
	}
}

// NewInsertPermissionsForConfigurationOK creates a InsertPermissionsForConfigurationOK with default headers values
func NewInsertPermissionsForConfigurationOK() *InsertPermissionsForConfigurationOK {
	return &InsertPermissionsForConfigurationOK{}
}

/*
InsertPermissionsForConfigurationOK describes a response with status code 200, with default header values.

successful operation
*/
type InsertPermissionsForConfigurationOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this insert permissions for configuration o k response has a 2xx status code
func (o *InsertPermissionsForConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert permissions for configuration o k response has a 3xx status code
func (o *InsertPermissionsForConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for configuration o k response has a 4xx status code
func (o *InsertPermissionsForConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert permissions for configuration o k response has a 5xx status code
func (o *InsertPermissionsForConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for configuration o k response a status code equal to that given
func (o *InsertPermissionsForConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the insert permissions for configuration o k response
func (o *InsertPermissionsForConfigurationOK) Code() int {
	return 200
}

func (o *InsertPermissionsForConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationOK  %+v", 200, o.Payload)
}

func (o *InsertPermissionsForConfigurationOK) String() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationOK  %+v", 200, o.Payload)
}

func (o *InsertPermissionsForConfigurationOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *InsertPermissionsForConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertPermissionsForConfigurationCreated creates a InsertPermissionsForConfigurationCreated with default headers values
func NewInsertPermissionsForConfigurationCreated() *InsertPermissionsForConfigurationCreated {
	return &InsertPermissionsForConfigurationCreated{}
}

/*
InsertPermissionsForConfigurationCreated describes a response with status code 201, with default header values.

The request is successful
*/
type InsertPermissionsForConfigurationCreated struct {
}

// IsSuccess returns true when this insert permissions for configuration created response has a 2xx status code
func (o *InsertPermissionsForConfigurationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert permissions for configuration created response has a 3xx status code
func (o *InsertPermissionsForConfigurationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for configuration created response has a 4xx status code
func (o *InsertPermissionsForConfigurationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert permissions for configuration created response has a 5xx status code
func (o *InsertPermissionsForConfigurationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for configuration created response a status code equal to that given
func (o *InsertPermissionsForConfigurationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the insert permissions for configuration created response
func (o *InsertPermissionsForConfigurationCreated) Code() int {
	return 201
}

func (o *InsertPermissionsForConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationCreated ", 201)
}

func (o *InsertPermissionsForConfigurationCreated) String() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationCreated ", 201)
}

func (o *InsertPermissionsForConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForConfigurationBadRequest creates a InsertPermissionsForConfigurationBadRequest with default headers values
func NewInsertPermissionsForConfigurationBadRequest() *InsertPermissionsForConfigurationBadRequest {
	return &InsertPermissionsForConfigurationBadRequest{}
}

/*
InsertPermissionsForConfigurationBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type InsertPermissionsForConfigurationBadRequest struct {
}

// IsSuccess returns true when this insert permissions for configuration bad request response has a 2xx status code
func (o *InsertPermissionsForConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for configuration bad request response has a 3xx status code
func (o *InsertPermissionsForConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for configuration bad request response has a 4xx status code
func (o *InsertPermissionsForConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for configuration bad request response has a 5xx status code
func (o *InsertPermissionsForConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for configuration bad request response a status code equal to that given
func (o *InsertPermissionsForConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the insert permissions for configuration bad request response
func (o *InsertPermissionsForConfigurationBadRequest) Code() int {
	return 400
}

func (o *InsertPermissionsForConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationBadRequest ", 400)
}

func (o *InsertPermissionsForConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationBadRequest ", 400)
}

func (o *InsertPermissionsForConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForConfigurationUnauthorized creates a InsertPermissionsForConfigurationUnauthorized with default headers values
func NewInsertPermissionsForConfigurationUnauthorized() *InsertPermissionsForConfigurationUnauthorized {
	return &InsertPermissionsForConfigurationUnauthorized{}
}

/*
InsertPermissionsForConfigurationUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type InsertPermissionsForConfigurationUnauthorized struct {
}

// IsSuccess returns true when this insert permissions for configuration unauthorized response has a 2xx status code
func (o *InsertPermissionsForConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for configuration unauthorized response has a 3xx status code
func (o *InsertPermissionsForConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for configuration unauthorized response has a 4xx status code
func (o *InsertPermissionsForConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for configuration unauthorized response has a 5xx status code
func (o *InsertPermissionsForConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for configuration unauthorized response a status code equal to that given
func (o *InsertPermissionsForConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the insert permissions for configuration unauthorized response
func (o *InsertPermissionsForConfigurationUnauthorized) Code() int {
	return 401
}

func (o *InsertPermissionsForConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationUnauthorized ", 401)
}

func (o *InsertPermissionsForConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationUnauthorized ", 401)
}

func (o *InsertPermissionsForConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForConfigurationNotFound creates a InsertPermissionsForConfigurationNotFound with default headers values
func NewInsertPermissionsForConfigurationNotFound() *InsertPermissionsForConfigurationNotFound {
	return &InsertPermissionsForConfigurationNotFound{}
}

/*
InsertPermissionsForConfigurationNotFound describes a response with status code 404, with default header values.

Cannot find configuration with the specified ID
*/
type InsertPermissionsForConfigurationNotFound struct {
}

// IsSuccess returns true when this insert permissions for configuration not found response has a 2xx status code
func (o *InsertPermissionsForConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for configuration not found response has a 3xx status code
func (o *InsertPermissionsForConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for configuration not found response has a 4xx status code
func (o *InsertPermissionsForConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for configuration not found response has a 5xx status code
func (o *InsertPermissionsForConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for configuration not found response a status code equal to that given
func (o *InsertPermissionsForConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the insert permissions for configuration not found response
func (o *InsertPermissionsForConfigurationNotFound) Code() int {
	return 404
}

func (o *InsertPermissionsForConfigurationNotFound) Error() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationNotFound ", 404)
}

func (o *InsertPermissionsForConfigurationNotFound) String() string {
	return fmt.Sprintf("[POST /configurations/{id}/permissions][%d] insertPermissionsForConfigurationNotFound ", 404)
}

func (o *InsertPermissionsForConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}