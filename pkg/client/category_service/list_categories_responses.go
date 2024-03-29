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

// ListCategoriesReader is a Reader for the ListCategories structure.
type ListCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCategoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /categories] listCategories", response, response.Code())
	}
}

// NewListCategoriesOK creates a ListCategoriesOK with default headers values
func NewListCategoriesOK() *ListCategoriesOK {
	return &ListCategoriesOK{}
}

/*
ListCategoriesOK describes a response with status code 200, with default header values.

The request is successful
*/
type ListCategoriesOK struct {
	Payload *models.Categories
}

// IsSuccess returns true when this list categories o k response has a 2xx status code
func (o *ListCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list categories o k response has a 3xx status code
func (o *ListCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories o k response has a 4xx status code
func (o *ListCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list categories o k response has a 5xx status code
func (o *ListCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories o k response a status code equal to that given
func (o *ListCategoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list categories o k response
func (o *ListCategoriesOK) Code() int {
	return 200
}

func (o *ListCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /categories][%d] listCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesOK) String() string {
	return fmt.Sprintf("[GET /categories][%d] listCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesOK) GetPayload() *models.Categories {
	return o.Payload
}

func (o *ListCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Categories)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCategoriesUnauthorized creates a ListCategoriesUnauthorized with default headers values
func NewListCategoriesUnauthorized() *ListCategoriesUnauthorized {
	return &ListCategoriesUnauthorized{}
}

/*
ListCategoriesUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type ListCategoriesUnauthorized struct {
}

// IsSuccess returns true when this list categories unauthorized response has a 2xx status code
func (o *ListCategoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list categories unauthorized response has a 3xx status code
func (o *ListCategoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories unauthorized response has a 4xx status code
func (o *ListCategoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list categories unauthorized response has a 5xx status code
func (o *ListCategoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories unauthorized response a status code equal to that given
func (o *ListCategoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list categories unauthorized response
func (o *ListCategoriesUnauthorized) Code() int {
	return 401
}

func (o *ListCategoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /categories][%d] listCategoriesUnauthorized ", 401)
}

func (o *ListCategoriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /categories][%d] listCategoriesUnauthorized ", 401)
}

func (o *ListCategoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
