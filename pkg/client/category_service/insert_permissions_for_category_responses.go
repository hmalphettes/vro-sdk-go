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

// InsertPermissionsForCategoryReader is a Reader for the InsertPermissionsForCategory structure.
type InsertPermissionsForCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertPermissionsForCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInsertPermissionsForCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewInsertPermissionsForCategoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInsertPermissionsForCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInsertPermissionsForCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInsertPermissionsForCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /categories/{id}/permissions] insertPermissionsForCategory", response, response.Code())
	}
}

// NewInsertPermissionsForCategoryOK creates a InsertPermissionsForCategoryOK with default headers values
func NewInsertPermissionsForCategoryOK() *InsertPermissionsForCategoryOK {
	return &InsertPermissionsForCategoryOK{}
}

/*
InsertPermissionsForCategoryOK describes a response with status code 200, with default header values.

successful operation
*/
type InsertPermissionsForCategoryOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this insert permissions for category o k response has a 2xx status code
func (o *InsertPermissionsForCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert permissions for category o k response has a 3xx status code
func (o *InsertPermissionsForCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for category o k response has a 4xx status code
func (o *InsertPermissionsForCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert permissions for category o k response has a 5xx status code
func (o *InsertPermissionsForCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for category o k response a status code equal to that given
func (o *InsertPermissionsForCategoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the insert permissions for category o k response
func (o *InsertPermissionsForCategoryOK) Code() int {
	return 200
}

func (o *InsertPermissionsForCategoryOK) Error() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryOK  %+v", 200, o.Payload)
}

func (o *InsertPermissionsForCategoryOK) String() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryOK  %+v", 200, o.Payload)
}

func (o *InsertPermissionsForCategoryOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *InsertPermissionsForCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertPermissionsForCategoryCreated creates a InsertPermissionsForCategoryCreated with default headers values
func NewInsertPermissionsForCategoryCreated() *InsertPermissionsForCategoryCreated {
	return &InsertPermissionsForCategoryCreated{}
}

/*
InsertPermissionsForCategoryCreated describes a response with status code 201, with default header values.

The request is successful
*/
type InsertPermissionsForCategoryCreated struct {
}

// IsSuccess returns true when this insert permissions for category created response has a 2xx status code
func (o *InsertPermissionsForCategoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this insert permissions for category created response has a 3xx status code
func (o *InsertPermissionsForCategoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for category created response has a 4xx status code
func (o *InsertPermissionsForCategoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this insert permissions for category created response has a 5xx status code
func (o *InsertPermissionsForCategoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for category created response a status code equal to that given
func (o *InsertPermissionsForCategoryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the insert permissions for category created response
func (o *InsertPermissionsForCategoryCreated) Code() int {
	return 201
}

func (o *InsertPermissionsForCategoryCreated) Error() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryCreated ", 201)
}

func (o *InsertPermissionsForCategoryCreated) String() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryCreated ", 201)
}

func (o *InsertPermissionsForCategoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForCategoryBadRequest creates a InsertPermissionsForCategoryBadRequest with default headers values
func NewInsertPermissionsForCategoryBadRequest() *InsertPermissionsForCategoryBadRequest {
	return &InsertPermissionsForCategoryBadRequest{}
}

/*
InsertPermissionsForCategoryBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type InsertPermissionsForCategoryBadRequest struct {
}

// IsSuccess returns true when this insert permissions for category bad request response has a 2xx status code
func (o *InsertPermissionsForCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for category bad request response has a 3xx status code
func (o *InsertPermissionsForCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for category bad request response has a 4xx status code
func (o *InsertPermissionsForCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for category bad request response has a 5xx status code
func (o *InsertPermissionsForCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for category bad request response a status code equal to that given
func (o *InsertPermissionsForCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the insert permissions for category bad request response
func (o *InsertPermissionsForCategoryBadRequest) Code() int {
	return 400
}

func (o *InsertPermissionsForCategoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryBadRequest ", 400)
}

func (o *InsertPermissionsForCategoryBadRequest) String() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryBadRequest ", 400)
}

func (o *InsertPermissionsForCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForCategoryUnauthorized creates a InsertPermissionsForCategoryUnauthorized with default headers values
func NewInsertPermissionsForCategoryUnauthorized() *InsertPermissionsForCategoryUnauthorized {
	return &InsertPermissionsForCategoryUnauthorized{}
}

/*
InsertPermissionsForCategoryUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type InsertPermissionsForCategoryUnauthorized struct {
}

// IsSuccess returns true when this insert permissions for category unauthorized response has a 2xx status code
func (o *InsertPermissionsForCategoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for category unauthorized response has a 3xx status code
func (o *InsertPermissionsForCategoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for category unauthorized response has a 4xx status code
func (o *InsertPermissionsForCategoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for category unauthorized response has a 5xx status code
func (o *InsertPermissionsForCategoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for category unauthorized response a status code equal to that given
func (o *InsertPermissionsForCategoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the insert permissions for category unauthorized response
func (o *InsertPermissionsForCategoryUnauthorized) Code() int {
	return 401
}

func (o *InsertPermissionsForCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryUnauthorized ", 401)
}

func (o *InsertPermissionsForCategoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryUnauthorized ", 401)
}

func (o *InsertPermissionsForCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInsertPermissionsForCategoryNotFound creates a InsertPermissionsForCategoryNotFound with default headers values
func NewInsertPermissionsForCategoryNotFound() *InsertPermissionsForCategoryNotFound {
	return &InsertPermissionsForCategoryNotFound{}
}

/*
InsertPermissionsForCategoryNotFound describes a response with status code 404, with default header values.

Cannot find a category with the specified ID
*/
type InsertPermissionsForCategoryNotFound struct {
}

// IsSuccess returns true when this insert permissions for category not found response has a 2xx status code
func (o *InsertPermissionsForCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this insert permissions for category not found response has a 3xx status code
func (o *InsertPermissionsForCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this insert permissions for category not found response has a 4xx status code
func (o *InsertPermissionsForCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this insert permissions for category not found response has a 5xx status code
func (o *InsertPermissionsForCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this insert permissions for category not found response a status code equal to that given
func (o *InsertPermissionsForCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the insert permissions for category not found response
func (o *InsertPermissionsForCategoryNotFound) Code() int {
	return 404
}

func (o *InsertPermissionsForCategoryNotFound) Error() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryNotFound ", 404)
}

func (o *InsertPermissionsForCategoryNotFound) String() string {
	return fmt.Sprintf("[POST /categories/{id}/permissions][%d] insertPermissionsForCategoryNotFound ", 404)
}

func (o *InsertPermissionsForCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
