// Code generated by go-swagger; DO NOT EDIT.

package actions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetActionReader is a Reader for the GetAction structure.
type GetActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /actions/{categoryName}/{actionName}] getAction", response, response.Code())
	}
}

// NewGetActionOK creates a GetActionOK with default headers values
func NewGetActionOK() *GetActionOK {
	return &GetActionOK{}
}

/*
GetActionOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetActionOK struct {
	Payload *models.Action
}

// IsSuccess returns true when this get action o k response has a 2xx status code
func (o *GetActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get action o k response has a 3xx status code
func (o *GetActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action o k response has a 4xx status code
func (o *GetActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get action o k response has a 5xx status code
func (o *GetActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get action o k response a status code equal to that given
func (o *GetActionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get action o k response
func (o *GetActionOK) Code() int {
	return 200
}

func (o *GetActionOK) Error() string {
	return fmt.Sprintf("[GET /actions/{categoryName}/{actionName}][%d] getActionOK  %+v", 200, o.Payload)
}

func (o *GetActionOK) String() string {
	return fmt.Sprintf("[GET /actions/{categoryName}/{actionName}][%d] getActionOK  %+v", 200, o.Payload)
}

func (o *GetActionOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *GetActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionUnauthorized creates a GetActionUnauthorized with default headers values
func NewGetActionUnauthorized() *GetActionUnauthorized {
	return &GetActionUnauthorized{}
}

/*
GetActionUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetActionUnauthorized struct {
}

// IsSuccess returns true when this get action unauthorized response has a 2xx status code
func (o *GetActionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action unauthorized response has a 3xx status code
func (o *GetActionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action unauthorized response has a 4xx status code
func (o *GetActionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get action unauthorized response has a 5xx status code
func (o *GetActionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get action unauthorized response a status code equal to that given
func (o *GetActionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get action unauthorized response
func (o *GetActionUnauthorized) Code() int {
	return 401
}

func (o *GetActionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions/{categoryName}/{actionName}][%d] getActionUnauthorized ", 401)
}

func (o *GetActionUnauthorized) String() string {
	return fmt.Sprintf("[GET /actions/{categoryName}/{actionName}][%d] getActionUnauthorized ", 401)
}

func (o *GetActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionNotFound creates a GetActionNotFound with default headers values
func NewGetActionNotFound() *GetActionNotFound {
	return &GetActionNotFound{}
}

/*
GetActionNotFound describes a response with status code 404, with default header values.

Can not find an action with the specified name
*/
type GetActionNotFound struct {
}

// IsSuccess returns true when this get action not found response has a 2xx status code
func (o *GetActionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action not found response has a 3xx status code
func (o *GetActionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action not found response has a 4xx status code
func (o *GetActionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get action not found response has a 5xx status code
func (o *GetActionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get action not found response a status code equal to that given
func (o *GetActionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get action not found response
func (o *GetActionNotFound) Code() int {
	return 404
}

func (o *GetActionNotFound) Error() string {
	return fmt.Sprintf("[GET /actions/{categoryName}/{actionName}][%d] getActionNotFound ", 404)
}

func (o *GetActionNotFound) String() string {
	return fmt.Sprintf("[GET /actions/{categoryName}/{actionName}][%d] getActionNotFound ", 404)
}

func (o *GetActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
