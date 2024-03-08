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

// CreatePolicyReader is a Reader for the CreatePolicy structure.
type CreatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreatePolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreatePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /policies] createPolicy", response, response.Code())
	}
}

// NewCreatePolicyOK creates a CreatePolicyOK with default headers values
func NewCreatePolicyOK() *CreatePolicyOK {
	return &CreatePolicyOK{}
}

/*
CreatePolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type CreatePolicyOK struct {
	Payload *models.Policy
}

// IsSuccess returns true when this create policy o k response has a 2xx status code
func (o *CreatePolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create policy o k response has a 3xx status code
func (o *CreatePolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy o k response has a 4xx status code
func (o *CreatePolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create policy o k response has a 5xx status code
func (o *CreatePolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy o k response a status code equal to that given
func (o *CreatePolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create policy o k response
func (o *CreatePolicyOK) Code() int {
	return 200
}

func (o *CreatePolicyOK) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyOK  %+v", 200, o.Payload)
}

func (o *CreatePolicyOK) String() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyOK  %+v", 200, o.Payload)
}

func (o *CreatePolicyOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *CreatePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyCreated creates a CreatePolicyCreated with default headers values
func NewCreatePolicyCreated() *CreatePolicyCreated {
	return &CreatePolicyCreated{}
}

/*
CreatePolicyCreated describes a response with status code 201, with default header values.

The request is successful
*/
type CreatePolicyCreated struct {
}

// IsSuccess returns true when this create policy created response has a 2xx status code
func (o *CreatePolicyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create policy created response has a 3xx status code
func (o *CreatePolicyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy created response has a 4xx status code
func (o *CreatePolicyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create policy created response has a 5xx status code
func (o *CreatePolicyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy created response a status code equal to that given
func (o *CreatePolicyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create policy created response
func (o *CreatePolicyCreated) Code() int {
	return 201
}

func (o *CreatePolicyCreated) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyCreated ", 201)
}

func (o *CreatePolicyCreated) String() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyCreated ", 201)
}

func (o *CreatePolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePolicyBadRequest creates a CreatePolicyBadRequest with default headers values
func NewCreatePolicyBadRequest() *CreatePolicyBadRequest {
	return &CreatePolicyBadRequest{}
}

/*
CreatePolicyBadRequest describes a response with status code 400, with default header values.

The request is invalid(validation error)
*/
type CreatePolicyBadRequest struct {
}

// IsSuccess returns true when this create policy bad request response has a 2xx status code
func (o *CreatePolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy bad request response has a 3xx status code
func (o *CreatePolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy bad request response has a 4xx status code
func (o *CreatePolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy bad request response has a 5xx status code
func (o *CreatePolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy bad request response a status code equal to that given
func (o *CreatePolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create policy bad request response
func (o *CreatePolicyBadRequest) Code() int {
	return 400
}

func (o *CreatePolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyBadRequest ", 400)
}

func (o *CreatePolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyBadRequest ", 400)
}

func (o *CreatePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePolicyUnauthorized creates a CreatePolicyUnauthorized with default headers values
func NewCreatePolicyUnauthorized() *CreatePolicyUnauthorized {
	return &CreatePolicyUnauthorized{}
}

/*
CreatePolicyUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type CreatePolicyUnauthorized struct {
}

// IsSuccess returns true when this create policy unauthorized response has a 2xx status code
func (o *CreatePolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy unauthorized response has a 3xx status code
func (o *CreatePolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy unauthorized response has a 4xx status code
func (o *CreatePolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy unauthorized response has a 5xx status code
func (o *CreatePolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy unauthorized response a status code equal to that given
func (o *CreatePolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create policy unauthorized response
func (o *CreatePolicyUnauthorized) Code() int {
	return 401
}

func (o *CreatePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyUnauthorized ", 401)
}

func (o *CreatePolicyUnauthorized) String() string {
	return fmt.Sprintf("[POST /policies][%d] createPolicyUnauthorized ", 401)
}

func (o *CreatePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
