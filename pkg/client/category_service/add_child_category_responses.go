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

// AddChildCategoryReader is a Reader for the AddChildCategory structure.
type AddChildCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddChildCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddChildCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddChildCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddChildCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddChildCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /categories/{id}] addChildCategory", response, response.Code())
	}
}

// NewAddChildCategoryOK creates a AddChildCategoryOK with default headers values
func NewAddChildCategoryOK() *AddChildCategoryOK {
	return &AddChildCategoryOK{}
}

/*
AddChildCategoryOK describes a response with status code 200, with default header values.

The request is successful
*/
type AddChildCategoryOK struct {
	Payload *models.Category
}

// IsSuccess returns true when this add child category o k response has a 2xx status code
func (o *AddChildCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add child category o k response has a 3xx status code
func (o *AddChildCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add child category o k response has a 4xx status code
func (o *AddChildCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add child category o k response has a 5xx status code
func (o *AddChildCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add child category o k response a status code equal to that given
func (o *AddChildCategoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add child category o k response
func (o *AddChildCategoryOK) Code() int {
	return 200
}

func (o *AddChildCategoryOK) Error() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryOK  %+v", 200, o.Payload)
}

func (o *AddChildCategoryOK) String() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryOK  %+v", 200, o.Payload)
}

func (o *AddChildCategoryOK) GetPayload() *models.Category {
	return o.Payload
}

func (o *AddChildCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Category)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddChildCategoryBadRequest creates a AddChildCategoryBadRequest with default headers values
func NewAddChildCategoryBadRequest() *AddChildCategoryBadRequest {
	return &AddChildCategoryBadRequest{}
}

/*
AddChildCategoryBadRequest describes a response with status code 400, with default header values.

Category name must be specified
*/
type AddChildCategoryBadRequest struct {
}

// IsSuccess returns true when this add child category bad request response has a 2xx status code
func (o *AddChildCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add child category bad request response has a 3xx status code
func (o *AddChildCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add child category bad request response has a 4xx status code
func (o *AddChildCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add child category bad request response has a 5xx status code
func (o *AddChildCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add child category bad request response a status code equal to that given
func (o *AddChildCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add child category bad request response
func (o *AddChildCategoryBadRequest) Code() int {
	return 400
}

func (o *AddChildCategoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryBadRequest ", 400)
}

func (o *AddChildCategoryBadRequest) String() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryBadRequest ", 400)
}

func (o *AddChildCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddChildCategoryUnauthorized creates a AddChildCategoryUnauthorized with default headers values
func NewAddChildCategoryUnauthorized() *AddChildCategoryUnauthorized {
	return &AddChildCategoryUnauthorized{}
}

/*
AddChildCategoryUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type AddChildCategoryUnauthorized struct {
}

// IsSuccess returns true when this add child category unauthorized response has a 2xx status code
func (o *AddChildCategoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add child category unauthorized response has a 3xx status code
func (o *AddChildCategoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add child category unauthorized response has a 4xx status code
func (o *AddChildCategoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add child category unauthorized response has a 5xx status code
func (o *AddChildCategoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add child category unauthorized response a status code equal to that given
func (o *AddChildCategoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add child category unauthorized response
func (o *AddChildCategoryUnauthorized) Code() int {
	return 401
}

func (o *AddChildCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryUnauthorized ", 401)
}

func (o *AddChildCategoryUnauthorized) String() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryUnauthorized ", 401)
}

func (o *AddChildCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddChildCategoryNotFound creates a AddChildCategoryNotFound with default headers values
func NewAddChildCategoryNotFound() *AddChildCategoryNotFound {
	return &AddChildCategoryNotFound{}
}

/*
AddChildCategoryNotFound describes a response with status code 404, with default header values.

Cannot find a category with the specified ID
*/
type AddChildCategoryNotFound struct {
}

// IsSuccess returns true when this add child category not found response has a 2xx status code
func (o *AddChildCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add child category not found response has a 3xx status code
func (o *AddChildCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add child category not found response has a 4xx status code
func (o *AddChildCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add child category not found response has a 5xx status code
func (o *AddChildCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add child category not found response a status code equal to that given
func (o *AddChildCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add child category not found response
func (o *AddChildCategoryNotFound) Code() int {
	return 404
}

func (o *AddChildCategoryNotFound) Error() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryNotFound ", 404)
}

func (o *AddChildCategoryNotFound) String() string {
	return fmt.Sprintf("[POST /categories/{id}][%d] addChildCategoryNotFound ", 404)
}

func (o *AddChildCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
