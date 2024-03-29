// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DisablePluginReader is a Reader for the DisablePlugin structure.
type DisablePluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisablePluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisablePluginNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisablePluginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDisablePluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisablePluginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /plugins/{pluginName}/state] disablePlugin", response, response.Code())
	}
}

// NewDisablePluginNoContent creates a DisablePluginNoContent with default headers values
func NewDisablePluginNoContent() *DisablePluginNoContent {
	return &DisablePluginNoContent{}
}

/*
DisablePluginNoContent describes a response with status code 204, with default header values.

No content
*/
type DisablePluginNoContent struct {
}

// IsSuccess returns true when this disable plugin no content response has a 2xx status code
func (o *DisablePluginNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable plugin no content response has a 3xx status code
func (o *DisablePluginNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable plugin no content response has a 4xx status code
func (o *DisablePluginNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable plugin no content response has a 5xx status code
func (o *DisablePluginNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this disable plugin no content response a status code equal to that given
func (o *DisablePluginNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the disable plugin no content response
func (o *DisablePluginNoContent) Code() int {
	return 204
}

func (o *DisablePluginNoContent) Error() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginNoContent ", 204)
}

func (o *DisablePluginNoContent) String() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginNoContent ", 204)
}

func (o *DisablePluginNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisablePluginBadRequest creates a DisablePluginBadRequest with default headers values
func NewDisablePluginBadRequest() *DisablePluginBadRequest {
	return &DisablePluginBadRequest{}
}

/*
DisablePluginBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type DisablePluginBadRequest struct {
}

// IsSuccess returns true when this disable plugin bad request response has a 2xx status code
func (o *DisablePluginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable plugin bad request response has a 3xx status code
func (o *DisablePluginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable plugin bad request response has a 4xx status code
func (o *DisablePluginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable plugin bad request response has a 5xx status code
func (o *DisablePluginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this disable plugin bad request response a status code equal to that given
func (o *DisablePluginBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the disable plugin bad request response
func (o *DisablePluginBadRequest) Code() int {
	return 400
}

func (o *DisablePluginBadRequest) Error() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginBadRequest ", 400)
}

func (o *DisablePluginBadRequest) String() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginBadRequest ", 400)
}

func (o *DisablePluginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisablePluginUnauthorized creates a DisablePluginUnauthorized with default headers values
func NewDisablePluginUnauthorized() *DisablePluginUnauthorized {
	return &DisablePluginUnauthorized{}
}

/*
DisablePluginUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DisablePluginUnauthorized struct {
}

// IsSuccess returns true when this disable plugin unauthorized response has a 2xx status code
func (o *DisablePluginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable plugin unauthorized response has a 3xx status code
func (o *DisablePluginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable plugin unauthorized response has a 4xx status code
func (o *DisablePluginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable plugin unauthorized response has a 5xx status code
func (o *DisablePluginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this disable plugin unauthorized response a status code equal to that given
func (o *DisablePluginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the disable plugin unauthorized response
func (o *DisablePluginUnauthorized) Code() int {
	return 401
}

func (o *DisablePluginUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginUnauthorized ", 401)
}

func (o *DisablePluginUnauthorized) String() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginUnauthorized ", 401)
}

func (o *DisablePluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisablePluginForbidden creates a DisablePluginForbidden with default headers values
func NewDisablePluginForbidden() *DisablePluginForbidden {
	return &DisablePluginForbidden{}
}

/*
DisablePluginForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DisablePluginForbidden struct {
}

// IsSuccess returns true when this disable plugin forbidden response has a 2xx status code
func (o *DisablePluginForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable plugin forbidden response has a 3xx status code
func (o *DisablePluginForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable plugin forbidden response has a 4xx status code
func (o *DisablePluginForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable plugin forbidden response has a 5xx status code
func (o *DisablePluginForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this disable plugin forbidden response a status code equal to that given
func (o *DisablePluginForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the disable plugin forbidden response
func (o *DisablePluginForbidden) Code() int {
	return 403
}

func (o *DisablePluginForbidden) Error() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginForbidden ", 403)
}

func (o *DisablePluginForbidden) String() string {
	return fmt.Sprintf("[PUT /plugins/{pluginName}/state][%d] disablePluginForbidden ", 403)
}

func (o *DisablePluginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
