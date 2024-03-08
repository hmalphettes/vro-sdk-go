// Code generated by go-swagger; DO NOT EDIT.

package policy_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetPolicyLogsReader is a Reader for the GetPolicyLogs structure.
type GetPolicyLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPolicyLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicyLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policies/{id}/syslogs] getPolicyLogs", response, response.Code())
	}
}

// NewGetPolicyLogsOK creates a GetPolicyLogsOK with default headers values
func NewGetPolicyLogsOK() *GetPolicyLogsOK {
	return &GetPolicyLogsOK{}
}

/*
GetPolicyLogsOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPolicyLogsOK struct {
	Payload *models.SysLogs
}

// IsSuccess returns true when this get policy logs o k response has a 2xx status code
func (o *GetPolicyLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policy logs o k response has a 3xx status code
func (o *GetPolicyLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy logs o k response has a 4xx status code
func (o *GetPolicyLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policy logs o k response has a 5xx status code
func (o *GetPolicyLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy logs o k response a status code equal to that given
func (o *GetPolicyLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get policy logs o k response
func (o *GetPolicyLogsOK) Code() int {
	return 200
}

func (o *GetPolicyLogsOK) Error() string {
	return fmt.Sprintf("[GET /policies/{id}/syslogs][%d] getPolicyLogsOK  %+v", 200, o.Payload)
}

func (o *GetPolicyLogsOK) String() string {
	return fmt.Sprintf("[GET /policies/{id}/syslogs][%d] getPolicyLogsOK  %+v", 200, o.Payload)
}

func (o *GetPolicyLogsOK) GetPayload() *models.SysLogs {
	return o.Payload
}

func (o *GetPolicyLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyLogsUnauthorized creates a GetPolicyLogsUnauthorized with default headers values
func NewGetPolicyLogsUnauthorized() *GetPolicyLogsUnauthorized {
	return &GetPolicyLogsUnauthorized{}
}

/*
GetPolicyLogsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetPolicyLogsUnauthorized struct {
}

// IsSuccess returns true when this get policy logs unauthorized response has a 2xx status code
func (o *GetPolicyLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy logs unauthorized response has a 3xx status code
func (o *GetPolicyLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy logs unauthorized response has a 4xx status code
func (o *GetPolicyLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy logs unauthorized response has a 5xx status code
func (o *GetPolicyLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy logs unauthorized response a status code equal to that given
func (o *GetPolicyLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get policy logs unauthorized response
func (o *GetPolicyLogsUnauthorized) Code() int {
	return 401
}

func (o *GetPolicyLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies/{id}/syslogs][%d] getPolicyLogsUnauthorized ", 401)
}

func (o *GetPolicyLogsUnauthorized) String() string {
	return fmt.Sprintf("[GET /policies/{id}/syslogs][%d] getPolicyLogsUnauthorized ", 401)
}

func (o *GetPolicyLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyLogsNotFound creates a GetPolicyLogsNotFound with default headers values
func NewGetPolicyLogsNotFound() *GetPolicyLogsNotFound {
	return &GetPolicyLogsNotFound{}
}

/*
GetPolicyLogsNotFound describes a response with status code 404, with default header values.

Can not find a policy with the specified ID or the user does not have 'read' access rights for that policy
*/
type GetPolicyLogsNotFound struct {
}

// IsSuccess returns true when this get policy logs not found response has a 2xx status code
func (o *GetPolicyLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy logs not found response has a 3xx status code
func (o *GetPolicyLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy logs not found response has a 4xx status code
func (o *GetPolicyLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy logs not found response has a 5xx status code
func (o *GetPolicyLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy logs not found response a status code equal to that given
func (o *GetPolicyLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get policy logs not found response
func (o *GetPolicyLogsNotFound) Code() int {
	return 404
}

func (o *GetPolicyLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /policies/{id}/syslogs][%d] getPolicyLogsNotFound ", 404)
}

func (o *GetPolicyLogsNotFound) String() string {
	return fmt.Sprintf("[GET /policies/{id}/syslogs][%d] getPolicyLogsNotFound ", 404)
}

func (o *GetPolicyLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
