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

// GetAllPolicyTemplatesReader is a Reader for the GetAllPolicyTemplates structure.
type GetAllPolicyTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllPolicyTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllPolicyTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetAllPolicyTemplatesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllPolicyTemplatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policies/templates] getAllPolicyTemplates", response, response.Code())
	}
}

// NewGetAllPolicyTemplatesOK creates a GetAllPolicyTemplatesOK with default headers values
func NewGetAllPolicyTemplatesOK() *GetAllPolicyTemplatesOK {
	return &GetAllPolicyTemplatesOK{}
}

/*
GetAllPolicyTemplatesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllPolicyTemplatesOK struct {
	Payload *models.PolicyTemplates
}

// IsSuccess returns true when this get all policy templates o k response has a 2xx status code
func (o *GetAllPolicyTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all policy templates o k response has a 3xx status code
func (o *GetAllPolicyTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all policy templates o k response has a 4xx status code
func (o *GetAllPolicyTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all policy templates o k response has a 5xx status code
func (o *GetAllPolicyTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all policy templates o k response a status code equal to that given
func (o *GetAllPolicyTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all policy templates o k response
func (o *GetAllPolicyTemplatesOK) Code() int {
	return 200
}

func (o *GetAllPolicyTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /policies/templates][%d] getAllPolicyTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetAllPolicyTemplatesOK) String() string {
	return fmt.Sprintf("[GET /policies/templates][%d] getAllPolicyTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetAllPolicyTemplatesOK) GetPayload() *models.PolicyTemplates {
	return o.Payload
}

func (o *GetAllPolicyTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTemplates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPolicyTemplatesNoContent creates a GetAllPolicyTemplatesNoContent with default headers values
func NewGetAllPolicyTemplatesNoContent() *GetAllPolicyTemplatesNoContent {
	return &GetAllPolicyTemplatesNoContent{}
}

/*
GetAllPolicyTemplatesNoContent describes a response with status code 204, with default header values.

No content
*/
type GetAllPolicyTemplatesNoContent struct {
}

// IsSuccess returns true when this get all policy templates no content response has a 2xx status code
func (o *GetAllPolicyTemplatesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all policy templates no content response has a 3xx status code
func (o *GetAllPolicyTemplatesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all policy templates no content response has a 4xx status code
func (o *GetAllPolicyTemplatesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all policy templates no content response has a 5xx status code
func (o *GetAllPolicyTemplatesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get all policy templates no content response a status code equal to that given
func (o *GetAllPolicyTemplatesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get all policy templates no content response
func (o *GetAllPolicyTemplatesNoContent) Code() int {
	return 204
}

func (o *GetAllPolicyTemplatesNoContent) Error() string {
	return fmt.Sprintf("[GET /policies/templates][%d] getAllPolicyTemplatesNoContent ", 204)
}

func (o *GetAllPolicyTemplatesNoContent) String() string {
	return fmt.Sprintf("[GET /policies/templates][%d] getAllPolicyTemplatesNoContent ", 204)
}

func (o *GetAllPolicyTemplatesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllPolicyTemplatesUnauthorized creates a GetAllPolicyTemplatesUnauthorized with default headers values
func NewGetAllPolicyTemplatesUnauthorized() *GetAllPolicyTemplatesUnauthorized {
	return &GetAllPolicyTemplatesUnauthorized{}
}

/*
GetAllPolicyTemplatesUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetAllPolicyTemplatesUnauthorized struct {
}

// IsSuccess returns true when this get all policy templates unauthorized response has a 2xx status code
func (o *GetAllPolicyTemplatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all policy templates unauthorized response has a 3xx status code
func (o *GetAllPolicyTemplatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all policy templates unauthorized response has a 4xx status code
func (o *GetAllPolicyTemplatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all policy templates unauthorized response has a 5xx status code
func (o *GetAllPolicyTemplatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all policy templates unauthorized response a status code equal to that given
func (o *GetAllPolicyTemplatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all policy templates unauthorized response
func (o *GetAllPolicyTemplatesUnauthorized) Code() int {
	return 401
}

func (o *GetAllPolicyTemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies/templates][%d] getAllPolicyTemplatesUnauthorized ", 401)
}

func (o *GetAllPolicyTemplatesUnauthorized) String() string {
	return fmt.Sprintf("[GET /policies/templates][%d] getAllPolicyTemplatesUnauthorized ", 401)
}

func (o *GetAllPolicyTemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
