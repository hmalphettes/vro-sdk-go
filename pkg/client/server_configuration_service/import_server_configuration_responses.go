// Code generated by go-swagger; DO NOT EDIT.

package server_configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportServerConfigurationReader is a Reader for the ImportServerConfiguration structure.
type ImportServerConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportServerConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportServerConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportServerConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportServerConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportServerConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /server-configuration] importServerConfiguration", response, response.Code())
	}
}

// NewImportServerConfigurationOK creates a ImportServerConfigurationOK with default headers values
func NewImportServerConfigurationOK() *ImportServerConfigurationOK {
	return &ImportServerConfigurationOK{}
}

/*
ImportServerConfigurationOK describes a response with status code 200, with default header values.

The request is successful.
*/
type ImportServerConfigurationOK struct {
}

// IsSuccess returns true when this import server configuration o k response has a 2xx status code
func (o *ImportServerConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import server configuration o k response has a 3xx status code
func (o *ImportServerConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import server configuration o k response has a 4xx status code
func (o *ImportServerConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this import server configuration o k response has a 5xx status code
func (o *ImportServerConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this import server configuration o k response a status code equal to that given
func (o *ImportServerConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the import server configuration o k response
func (o *ImportServerConfigurationOK) Code() int {
	return 200
}

func (o *ImportServerConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationOK ", 200)
}

func (o *ImportServerConfigurationOK) String() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationOK ", 200)
}

func (o *ImportServerConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportServerConfigurationBadRequest creates a ImportServerConfigurationBadRequest with default headers values
func NewImportServerConfigurationBadRequest() *ImportServerConfigurationBadRequest {
	return &ImportServerConfigurationBadRequest{}
}

/*
ImportServerConfigurationBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error).
*/
type ImportServerConfigurationBadRequest struct {
}

// IsSuccess returns true when this import server configuration bad request response has a 2xx status code
func (o *ImportServerConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import server configuration bad request response has a 3xx status code
func (o *ImportServerConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import server configuration bad request response has a 4xx status code
func (o *ImportServerConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this import server configuration bad request response has a 5xx status code
func (o *ImportServerConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this import server configuration bad request response a status code equal to that given
func (o *ImportServerConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the import server configuration bad request response
func (o *ImportServerConfigurationBadRequest) Code() int {
	return 400
}

func (o *ImportServerConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationBadRequest ", 400)
}

func (o *ImportServerConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationBadRequest ", 400)
}

func (o *ImportServerConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportServerConfigurationUnauthorized creates a ImportServerConfigurationUnauthorized with default headers values
func NewImportServerConfigurationUnauthorized() *ImportServerConfigurationUnauthorized {
	return &ImportServerConfigurationUnauthorized{}
}

/*
ImportServerConfigurationUnauthorized describes a response with status code 401, with default header values.

User is not authenticated.
*/
type ImportServerConfigurationUnauthorized struct {
}

// IsSuccess returns true when this import server configuration unauthorized response has a 2xx status code
func (o *ImportServerConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import server configuration unauthorized response has a 3xx status code
func (o *ImportServerConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import server configuration unauthorized response has a 4xx status code
func (o *ImportServerConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this import server configuration unauthorized response has a 5xx status code
func (o *ImportServerConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this import server configuration unauthorized response a status code equal to that given
func (o *ImportServerConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the import server configuration unauthorized response
func (o *ImportServerConfigurationUnauthorized) Code() int {
	return 401
}

func (o *ImportServerConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationUnauthorized ", 401)
}

func (o *ImportServerConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationUnauthorized ", 401)
}

func (o *ImportServerConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportServerConfigurationForbidden creates a ImportServerConfigurationForbidden with default headers values
func NewImportServerConfigurationForbidden() *ImportServerConfigurationForbidden {
	return &ImportServerConfigurationForbidden{}
}

/*
ImportServerConfigurationForbidden describes a response with status code 403, with default header values.

User is not authorized.
*/
type ImportServerConfigurationForbidden struct {
}

// IsSuccess returns true when this import server configuration forbidden response has a 2xx status code
func (o *ImportServerConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import server configuration forbidden response has a 3xx status code
func (o *ImportServerConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import server configuration forbidden response has a 4xx status code
func (o *ImportServerConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this import server configuration forbidden response has a 5xx status code
func (o *ImportServerConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this import server configuration forbidden response a status code equal to that given
func (o *ImportServerConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the import server configuration forbidden response
func (o *ImportServerConfigurationForbidden) Code() int {
	return 403
}

func (o *ImportServerConfigurationForbidden) Error() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationForbidden ", 403)
}

func (o *ImportServerConfigurationForbidden) String() string {
	return fmt.Sprintf("[POST /server-configuration][%d] importServerConfigurationForbidden ", 403)
}

func (o *ImportServerConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
