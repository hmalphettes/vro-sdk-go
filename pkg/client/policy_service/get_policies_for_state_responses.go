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

// GetPoliciesForStateReader is a Reader for the GetPoliciesForState structure.
type GetPoliciesForStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPoliciesForStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPoliciesForStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPoliciesForStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPoliciesForStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policies/state/{state}] getPoliciesForState", response, response.Code())
	}
}

// NewGetPoliciesForStateOK creates a GetPoliciesForStateOK with default headers values
func NewGetPoliciesForStateOK() *GetPoliciesForStateOK {
	return &GetPoliciesForStateOK{}
}

/*
GetPoliciesForStateOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetPoliciesForStateOK struct {
	Payload *models.PolicyList
}

// IsSuccess returns true when this get policies for state o k response has a 2xx status code
func (o *GetPoliciesForStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policies for state o k response has a 3xx status code
func (o *GetPoliciesForStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policies for state o k response has a 4xx status code
func (o *GetPoliciesForStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policies for state o k response has a 5xx status code
func (o *GetPoliciesForStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policies for state o k response a status code equal to that given
func (o *GetPoliciesForStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get policies for state o k response
func (o *GetPoliciesForStateOK) Code() int {
	return 200
}

func (o *GetPoliciesForStateOK) Error() string {
	return fmt.Sprintf("[GET /policies/state/{state}][%d] getPoliciesForStateOK  %+v", 200, o.Payload)
}

func (o *GetPoliciesForStateOK) String() string {
	return fmt.Sprintf("[GET /policies/state/{state}][%d] getPoliciesForStateOK  %+v", 200, o.Payload)
}

func (o *GetPoliciesForStateOK) GetPayload() *models.PolicyList {
	return o.Payload
}

func (o *GetPoliciesForStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPoliciesForStateBadRequest creates a GetPoliciesForStateBadRequest with default headers values
func NewGetPoliciesForStateBadRequest() *GetPoliciesForStateBadRequest {
	return &GetPoliciesForStateBadRequest{}
}

/*
GetPoliciesForStateBadRequest describes a response with status code 400, with default header values.

The requested state is invalid
*/
type GetPoliciesForStateBadRequest struct {
}

// IsSuccess returns true when this get policies for state bad request response has a 2xx status code
func (o *GetPoliciesForStateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policies for state bad request response has a 3xx status code
func (o *GetPoliciesForStateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policies for state bad request response has a 4xx status code
func (o *GetPoliciesForStateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policies for state bad request response has a 5xx status code
func (o *GetPoliciesForStateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get policies for state bad request response a status code equal to that given
func (o *GetPoliciesForStateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get policies for state bad request response
func (o *GetPoliciesForStateBadRequest) Code() int {
	return 400
}

func (o *GetPoliciesForStateBadRequest) Error() string {
	return fmt.Sprintf("[GET /policies/state/{state}][%d] getPoliciesForStateBadRequest ", 400)
}

func (o *GetPoliciesForStateBadRequest) String() string {
	return fmt.Sprintf("[GET /policies/state/{state}][%d] getPoliciesForStateBadRequest ", 400)
}

func (o *GetPoliciesForStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPoliciesForStateUnauthorized creates a GetPoliciesForStateUnauthorized with default headers values
func NewGetPoliciesForStateUnauthorized() *GetPoliciesForStateUnauthorized {
	return &GetPoliciesForStateUnauthorized{}
}

/*
GetPoliciesForStateUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetPoliciesForStateUnauthorized struct {
}

// IsSuccess returns true when this get policies for state unauthorized response has a 2xx status code
func (o *GetPoliciesForStateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policies for state unauthorized response has a 3xx status code
func (o *GetPoliciesForStateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policies for state unauthorized response has a 4xx status code
func (o *GetPoliciesForStateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policies for state unauthorized response has a 5xx status code
func (o *GetPoliciesForStateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get policies for state unauthorized response a status code equal to that given
func (o *GetPoliciesForStateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get policies for state unauthorized response
func (o *GetPoliciesForStateUnauthorized) Code() int {
	return 401
}

func (o *GetPoliciesForStateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies/state/{state}][%d] getPoliciesForStateUnauthorized ", 401)
}

func (o *GetPoliciesForStateUnauthorized) String() string {
	return fmt.Sprintf("[GET /policies/state/{state}][%d] getPoliciesForStateUnauthorized ", 401)
}

func (o *GetPoliciesForStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
