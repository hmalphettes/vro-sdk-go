// Code generated by go-swagger; DO NOT EDIT.

package authorization_groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetAuthorizationGroupReader is a Reader for the GetAuthorizationGroup structure.
type GetAuthorizationGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAuthorizationGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /authorization-groups/{id}] getAuthorizationGroup", response, response.Code())
	}
}

// NewGetAuthorizationGroupOK creates a GetAuthorizationGroupOK with default headers values
func NewGetAuthorizationGroupOK() *GetAuthorizationGroupOK {
	return &GetAuthorizationGroupOK{}
}

/*
GetAuthorizationGroupOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetAuthorizationGroupOK struct {
	Payload *models.AuthorizationGroup
}

// IsSuccess returns true when this get authorization group o k response has a 2xx status code
func (o *GetAuthorizationGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get authorization group o k response has a 3xx status code
func (o *GetAuthorizationGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization group o k response has a 4xx status code
func (o *GetAuthorizationGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization group o k response has a 5xx status code
func (o *GetAuthorizationGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization group o k response a status code equal to that given
func (o *GetAuthorizationGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get authorization group o k response
func (o *GetAuthorizationGroupOK) Code() int {
	return 200
}

func (o *GetAuthorizationGroupOK) Error() string {
	return fmt.Sprintf("[GET /authorization-groups/{id}][%d] getAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationGroupOK) String() string {
	return fmt.Sprintf("[GET /authorization-groups/{id}][%d] getAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationGroupOK) GetPayload() *models.AuthorizationGroup {
	return o.Payload
}

func (o *GetAuthorizationGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationGroupNotFound creates a GetAuthorizationGroupNotFound with default headers values
func NewGetAuthorizationGroupNotFound() *GetAuthorizationGroupNotFound {
	return &GetAuthorizationGroupNotFound{}
}

/*
GetAuthorizationGroupNotFound describes a response with status code 404, with default header values.

Cannot find a group with the specified id.
*/
type GetAuthorizationGroupNotFound struct {
}

// IsSuccess returns true when this get authorization group not found response has a 2xx status code
func (o *GetAuthorizationGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization group not found response has a 3xx status code
func (o *GetAuthorizationGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization group not found response has a 4xx status code
func (o *GetAuthorizationGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization group not found response has a 5xx status code
func (o *GetAuthorizationGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization group not found response a status code equal to that given
func (o *GetAuthorizationGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get authorization group not found response
func (o *GetAuthorizationGroupNotFound) Code() int {
	return 404
}

func (o *GetAuthorizationGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /authorization-groups/{id}][%d] getAuthorizationGroupNotFound ", 404)
}

func (o *GetAuthorizationGroupNotFound) String() string {
	return fmt.Sprintf("[GET /authorization-groups/{id}][%d] getAuthorizationGroupNotFound ", 404)
}

func (o *GetAuthorizationGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
