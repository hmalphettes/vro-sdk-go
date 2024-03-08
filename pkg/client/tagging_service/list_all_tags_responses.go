// Code generated by go-swagger; DO NOT EDIT.

package tagging_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ListAllTagsReader is a Reader for the ListAllTags structure.
type ListAllTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAllTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tags/all] listAllTags", response, response.Code())
	}
}

// NewListAllTagsOK creates a ListAllTagsOK with default headers values
func NewListAllTagsOK() *ListAllTagsOK {
	return &ListAllTagsOK{}
}

/*
ListAllTagsOK describes a response with status code 200, with default header values.

The request is successful
*/
type ListAllTagsOK struct {
	Payload *models.Tags
}

// IsSuccess returns true when this list all tags o k response has a 2xx status code
func (o *ListAllTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list all tags o k response has a 3xx status code
func (o *ListAllTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all tags o k response has a 4xx status code
func (o *ListAllTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list all tags o k response has a 5xx status code
func (o *ListAllTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list all tags o k response a status code equal to that given
func (o *ListAllTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list all tags o k response
func (o *ListAllTagsOK) Code() int {
	return 200
}

func (o *ListAllTagsOK) Error() string {
	return fmt.Sprintf("[GET /tags/all][%d] listAllTagsOK  %+v", 200, o.Payload)
}

func (o *ListAllTagsOK) String() string {
	return fmt.Sprintf("[GET /tags/all][%d] listAllTagsOK  %+v", 200, o.Payload)
}

func (o *ListAllTagsOK) GetPayload() *models.Tags {
	return o.Payload
}

func (o *ListAllTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllTagsUnauthorized creates a ListAllTagsUnauthorized with default headers values
func NewListAllTagsUnauthorized() *ListAllTagsUnauthorized {
	return &ListAllTagsUnauthorized{}
}

/*
ListAllTagsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type ListAllTagsUnauthorized struct {
}

// IsSuccess returns true when this list all tags unauthorized response has a 2xx status code
func (o *ListAllTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list all tags unauthorized response has a 3xx status code
func (o *ListAllTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all tags unauthorized response has a 4xx status code
func (o *ListAllTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list all tags unauthorized response has a 5xx status code
func (o *ListAllTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list all tags unauthorized response a status code equal to that given
func (o *ListAllTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list all tags unauthorized response
func (o *ListAllTagsUnauthorized) Code() int {
	return 401
}

func (o *ListAllTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tags/all][%d] listAllTagsUnauthorized ", 401)
}

func (o *ListAllTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /tags/all][%d] listAllTagsUnauthorized ", 401)
}

func (o *ListAllTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}