// Code generated by go-swagger; DO NOT EDIT.

package authorization_groups_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateAuthorizationGroupReader is a Reader for the CreateAuthorizationGroup structure.
type CreateAuthorizationGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthorizationGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAuthorizationGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAuthorizationGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAuthorizationGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /authorization-groups] createAuthorizationGroup", response, response.Code())
	}
}

// NewCreateAuthorizationGroupOK creates a CreateAuthorizationGroupOK with default headers values
func NewCreateAuthorizationGroupOK() *CreateAuthorizationGroupOK {
	return &CreateAuthorizationGroupOK{}
}

/*
CreateAuthorizationGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateAuthorizationGroupOK struct {
	Payload *models.AuthorizationGroup
}

// IsSuccess returns true when this create authorization group o k response has a 2xx status code
func (o *CreateAuthorizationGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create authorization group o k response has a 3xx status code
func (o *CreateAuthorizationGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authorization group o k response has a 4xx status code
func (o *CreateAuthorizationGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create authorization group o k response has a 5xx status code
func (o *CreateAuthorizationGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create authorization group o k response a status code equal to that given
func (o *CreateAuthorizationGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create authorization group o k response
func (o *CreateAuthorizationGroupOK) Code() int {
	return 200
}

func (o *CreateAuthorizationGroupOK) Error() string {
	return fmt.Sprintf("[POST /authorization-groups][%d] createAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *CreateAuthorizationGroupOK) String() string {
	return fmt.Sprintf("[POST /authorization-groups][%d] createAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *CreateAuthorizationGroupOK) GetPayload() *models.AuthorizationGroup {
	return o.Payload
}

func (o *CreateAuthorizationGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationGroupCreated creates a CreateAuthorizationGroupCreated with default headers values
func NewCreateAuthorizationGroupCreated() *CreateAuthorizationGroupCreated {
	return &CreateAuthorizationGroupCreated{}
}

/*
CreateAuthorizationGroupCreated describes a response with status code 201, with default header values.

The group is created
*/
type CreateAuthorizationGroupCreated struct {
}

// IsSuccess returns true when this create authorization group created response has a 2xx status code
func (o *CreateAuthorizationGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create authorization group created response has a 3xx status code
func (o *CreateAuthorizationGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authorization group created response has a 4xx status code
func (o *CreateAuthorizationGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create authorization group created response has a 5xx status code
func (o *CreateAuthorizationGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create authorization group created response a status code equal to that given
func (o *CreateAuthorizationGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create authorization group created response
func (o *CreateAuthorizationGroupCreated) Code() int {
	return 201
}

func (o *CreateAuthorizationGroupCreated) Error() string {
	return fmt.Sprintf("[POST /authorization-groups][%d] createAuthorizationGroupCreated ", 201)
}

func (o *CreateAuthorizationGroupCreated) String() string {
	return fmt.Sprintf("[POST /authorization-groups][%d] createAuthorizationGroupCreated ", 201)
}

func (o *CreateAuthorizationGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAuthorizationGroupBadRequest creates a CreateAuthorizationGroupBadRequest with default headers values
func NewCreateAuthorizationGroupBadRequest() *CreateAuthorizationGroupBadRequest {
	return &CreateAuthorizationGroupBadRequest{}
}

/*
CreateAuthorizationGroupBadRequest describes a response with status code 400, with default header values.

CreateAuthorizationGroupBadRequest create authorization group bad request
*/
type CreateAuthorizationGroupBadRequest struct {
}

// IsSuccess returns true when this create authorization group bad request response has a 2xx status code
func (o *CreateAuthorizationGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create authorization group bad request response has a 3xx status code
func (o *CreateAuthorizationGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create authorization group bad request response has a 4xx status code
func (o *CreateAuthorizationGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create authorization group bad request response has a 5xx status code
func (o *CreateAuthorizationGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create authorization group bad request response a status code equal to that given
func (o *CreateAuthorizationGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create authorization group bad request response
func (o *CreateAuthorizationGroupBadRequest) Code() int {
	return 400
}

func (o *CreateAuthorizationGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /authorization-groups][%d] createAuthorizationGroupBadRequest ", 400)
}

func (o *CreateAuthorizationGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /authorization-groups][%d] createAuthorizationGroupBadRequest ", 400)
}

func (o *CreateAuthorizationGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}