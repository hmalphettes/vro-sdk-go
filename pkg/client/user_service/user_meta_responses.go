// Code generated by go-swagger; DO NOT EDIT.

package user_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// UserMetaReader is a Reader for the UserMeta structure.
type UserMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUserMetaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users] userMeta", response, response.Code())
	}
}

// NewUserMetaOK creates a UserMetaOK with default headers values
func NewUserMetaOK() *UserMetaOK {
	return &UserMetaOK{}
}

/*
UserMetaOK describes a response with status code 200, with default header values.

The request is successful
*/
type UserMetaOK struct {
	Payload *models.User
}

// IsSuccess returns true when this user meta o k response has a 2xx status code
func (o *UserMetaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user meta o k response has a 3xx status code
func (o *UserMetaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user meta o k response has a 4xx status code
func (o *UserMetaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user meta o k response has a 5xx status code
func (o *UserMetaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user meta o k response a status code equal to that given
func (o *UserMetaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user meta o k response
func (o *UserMetaOK) Code() int {
	return 200
}

func (o *UserMetaOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] userMetaOK  %+v", 200, o.Payload)
}

func (o *UserMetaOK) String() string {
	return fmt.Sprintf("[GET /users][%d] userMetaOK  %+v", 200, o.Payload)
}

func (o *UserMetaOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UserMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMetaUnauthorized creates a UserMetaUnauthorized with default headers values
func NewUserMetaUnauthorized() *UserMetaUnauthorized {
	return &UserMetaUnauthorized{}
}

/*
UserMetaUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type UserMetaUnauthorized struct {
}

// IsSuccess returns true when this user meta unauthorized response has a 2xx status code
func (o *UserMetaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user meta unauthorized response has a 3xx status code
func (o *UserMetaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user meta unauthorized response has a 4xx status code
func (o *UserMetaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user meta unauthorized response has a 5xx status code
func (o *UserMetaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user meta unauthorized response a status code equal to that given
func (o *UserMetaUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user meta unauthorized response
func (o *UserMetaUnauthorized) Code() int {
	return 401
}

func (o *UserMetaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] userMetaUnauthorized ", 401)
}

func (o *UserMetaUnauthorized) String() string {
	return fmt.Sprintf("[GET /users][%d] userMetaUnauthorized ", 401)
}

func (o *UserMetaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
