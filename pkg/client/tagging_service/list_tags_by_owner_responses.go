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

// ListTagsByOwnerReader is a Reader for the ListTagsByOwner structure.
type ListTagsByOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTagsByOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTagsByOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListTagsByOwnerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tags/{owner}] listTagsByOwner", response, response.Code())
	}
}

// NewListTagsByOwnerOK creates a ListTagsByOwnerOK with default headers values
func NewListTagsByOwnerOK() *ListTagsByOwnerOK {
	return &ListTagsByOwnerOK{}
}

/*
ListTagsByOwnerOK describes a response with status code 200, with default header values.

The request is successful
*/
type ListTagsByOwnerOK struct {
	Payload *models.Tags
}

// IsSuccess returns true when this list tags by owner o k response has a 2xx status code
func (o *ListTagsByOwnerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tags by owner o k response has a 3xx status code
func (o *ListTagsByOwnerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tags by owner o k response has a 4xx status code
func (o *ListTagsByOwnerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tags by owner o k response has a 5xx status code
func (o *ListTagsByOwnerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list tags by owner o k response a status code equal to that given
func (o *ListTagsByOwnerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list tags by owner o k response
func (o *ListTagsByOwnerOK) Code() int {
	return 200
}

func (o *ListTagsByOwnerOK) Error() string {
	return fmt.Sprintf("[GET /tags/{owner}][%d] listTagsByOwnerOK  %+v", 200, o.Payload)
}

func (o *ListTagsByOwnerOK) String() string {
	return fmt.Sprintf("[GET /tags/{owner}][%d] listTagsByOwnerOK  %+v", 200, o.Payload)
}

func (o *ListTagsByOwnerOK) GetPayload() *models.Tags {
	return o.Payload
}

func (o *ListTagsByOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTagsByOwnerUnauthorized creates a ListTagsByOwnerUnauthorized with default headers values
func NewListTagsByOwnerUnauthorized() *ListTagsByOwnerUnauthorized {
	return &ListTagsByOwnerUnauthorized{}
}

/*
ListTagsByOwnerUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type ListTagsByOwnerUnauthorized struct {
}

// IsSuccess returns true when this list tags by owner unauthorized response has a 2xx status code
func (o *ListTagsByOwnerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tags by owner unauthorized response has a 3xx status code
func (o *ListTagsByOwnerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tags by owner unauthorized response has a 4xx status code
func (o *ListTagsByOwnerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tags by owner unauthorized response has a 5xx status code
func (o *ListTagsByOwnerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list tags by owner unauthorized response a status code equal to that given
func (o *ListTagsByOwnerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list tags by owner unauthorized response
func (o *ListTagsByOwnerUnauthorized) Code() int {
	return 401
}

func (o *ListTagsByOwnerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tags/{owner}][%d] listTagsByOwnerUnauthorized ", 401)
}

func (o *ListTagsByOwnerUnauthorized) String() string {
	return fmt.Sprintf("[GET /tags/{owner}][%d] listTagsByOwnerUnauthorized ", 401)
}

func (o *ListTagsByOwnerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
