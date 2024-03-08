// Code generated by go-swagger; DO NOT EDIT.

package task_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tasks/{id}] getTask", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*
GetTaskOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this get task o k response has a 2xx status code
func (o *GetTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task o k response has a 3xx status code
func (o *GetTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task o k response has a 4xx status code
func (o *GetTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task o k response has a 5xx status code
func (o *GetTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task o k response a status code equal to that given
func (o *GetTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task o k response
func (o *GetTaskOK) Code() int {
	return 200
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) String() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskUnauthorized creates a GetTaskUnauthorized with default headers values
func NewGetTaskUnauthorized() *GetTaskUnauthorized {
	return &GetTaskUnauthorized{}
}

/*
GetTaskUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetTaskUnauthorized struct {
}

// IsSuccess returns true when this get task unauthorized response has a 2xx status code
func (o *GetTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task unauthorized response has a 3xx status code
func (o *GetTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task unauthorized response has a 4xx status code
func (o *GetTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task unauthorized response has a 5xx status code
func (o *GetTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get task unauthorized response a status code equal to that given
func (o *GetTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get task unauthorized response
func (o *GetTaskUnauthorized) Code() int {
	return 401
}

func (o *GetTaskUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskUnauthorized ", 401)
}

func (o *GetTaskUnauthorized) String() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskUnauthorized ", 401)
}

func (o *GetTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*
GetTaskNotFound describes a response with status code 404, with default header values.

Can not find a task with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetTaskNotFound struct {
}

// IsSuccess returns true when this get task not found response has a 2xx status code
func (o *GetTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task not found response has a 3xx status code
func (o *GetTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task not found response has a 4xx status code
func (o *GetTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task not found response has a 5xx status code
func (o *GetTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task not found response a status code equal to that given
func (o *GetTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get task not found response
func (o *GetTaskNotFound) Code() int {
	return 404
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskNotFound ", 404)
}

func (o *GetTaskNotFound) String() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] getTaskNotFound ", 404)
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}