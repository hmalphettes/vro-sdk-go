// Code generated by go-swagger; DO NOT EDIT.

package server_configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetScriptingAPIReader is a Reader for the GetScriptingAPI structure.
type GetScriptingAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptingAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptingAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScriptingAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptingAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /server-configuration/api] getScriptingApi", response, response.Code())
	}
}

// NewGetScriptingAPIOK creates a GetScriptingAPIOK with default headers values
func NewGetScriptingAPIOK() *GetScriptingAPIOK {
	return &GetScriptingAPIOK{}
}

/*
GetScriptingAPIOK describes a response with status code 200, with default header values.

The request is successful.
*/
type GetScriptingAPIOK struct {
	Payload *models.APITypes
}

// IsSuccess returns true when this get scripting Api o k response has a 2xx status code
func (o *GetScriptingAPIOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scripting Api o k response has a 3xx status code
func (o *GetScriptingAPIOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripting Api o k response has a 4xx status code
func (o *GetScriptingAPIOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripting Api o k response has a 5xx status code
func (o *GetScriptingAPIOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripting Api o k response a status code equal to that given
func (o *GetScriptingAPIOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get scripting Api o k response
func (o *GetScriptingAPIOK) Code() int {
	return 200
}

func (o *GetScriptingAPIOK) Error() string {
	return fmt.Sprintf("[GET /server-configuration/api][%d] getScriptingApiOK  %+v", 200, o.Payload)
}

func (o *GetScriptingAPIOK) String() string {
	return fmt.Sprintf("[GET /server-configuration/api][%d] getScriptingApiOK  %+v", 200, o.Payload)
}

func (o *GetScriptingAPIOK) GetPayload() *models.APITypes {
	return o.Payload
}

func (o *GetScriptingAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APITypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptingAPIUnauthorized creates a GetScriptingAPIUnauthorized with default headers values
func NewGetScriptingAPIUnauthorized() *GetScriptingAPIUnauthorized {
	return &GetScriptingAPIUnauthorized{}
}

/*
GetScriptingAPIUnauthorized describes a response with status code 401, with default header values.

User is not authenticated.
*/
type GetScriptingAPIUnauthorized struct {
}

// IsSuccess returns true when this get scripting Api unauthorized response has a 2xx status code
func (o *GetScriptingAPIUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripting Api unauthorized response has a 3xx status code
func (o *GetScriptingAPIUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripting Api unauthorized response has a 4xx status code
func (o *GetScriptingAPIUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripting Api unauthorized response has a 5xx status code
func (o *GetScriptingAPIUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripting Api unauthorized response a status code equal to that given
func (o *GetScriptingAPIUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get scripting Api unauthorized response
func (o *GetScriptingAPIUnauthorized) Code() int {
	return 401
}

func (o *GetScriptingAPIUnauthorized) Error() string {
	return fmt.Sprintf("[GET /server-configuration/api][%d] getScriptingApiUnauthorized ", 401)
}

func (o *GetScriptingAPIUnauthorized) String() string {
	return fmt.Sprintf("[GET /server-configuration/api][%d] getScriptingApiUnauthorized ", 401)
}

func (o *GetScriptingAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScriptingAPIForbidden creates a GetScriptingAPIForbidden with default headers values
func NewGetScriptingAPIForbidden() *GetScriptingAPIForbidden {
	return &GetScriptingAPIForbidden{}
}

/*
GetScriptingAPIForbidden describes a response with status code 403, with default header values.

User is not authorized.
*/
type GetScriptingAPIForbidden struct {
}

// IsSuccess returns true when this get scripting Api forbidden response has a 2xx status code
func (o *GetScriptingAPIForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripting Api forbidden response has a 3xx status code
func (o *GetScriptingAPIForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripting Api forbidden response has a 4xx status code
func (o *GetScriptingAPIForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripting Api forbidden response has a 5xx status code
func (o *GetScriptingAPIForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripting Api forbidden response a status code equal to that given
func (o *GetScriptingAPIForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get scripting Api forbidden response
func (o *GetScriptingAPIForbidden) Code() int {
	return 403
}

func (o *GetScriptingAPIForbidden) Error() string {
	return fmt.Sprintf("[GET /server-configuration/api][%d] getScriptingApiForbidden ", 403)
}

func (o *GetScriptingAPIForbidden) String() string {
	return fmt.Sprintf("[GET /server-configuration/api][%d] getScriptingApiForbidden ", 403)
}

func (o *GetScriptingAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
