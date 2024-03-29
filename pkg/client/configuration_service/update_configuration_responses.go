// Code generated by go-swagger; DO NOT EDIT.

package configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateConfigurationReader is a Reader for the UpdateConfiguration structure.
type UpdateConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /configurations/{id}] updateConfiguration", response, response.Code())
	}
}

// NewUpdateConfigurationNoContent creates a UpdateConfigurationNoContent with default headers values
func NewUpdateConfigurationNoContent() *UpdateConfigurationNoContent {
	return &UpdateConfigurationNoContent{}
}

/*
UpdateConfigurationNoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateConfigurationNoContent struct {
}

// IsSuccess returns true when this update configuration no content response has a 2xx status code
func (o *UpdateConfigurationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update configuration no content response has a 3xx status code
func (o *UpdateConfigurationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration no content response has a 4xx status code
func (o *UpdateConfigurationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update configuration no content response has a 5xx status code
func (o *UpdateConfigurationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration no content response a status code equal to that given
func (o *UpdateConfigurationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update configuration no content response
func (o *UpdateConfigurationNoContent) Code() int {
	return 204
}

func (o *UpdateConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationNoContent ", 204)
}

func (o *UpdateConfigurationNoContent) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationNoContent ", 204)
}

func (o *UpdateConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigurationBadRequest creates a UpdateConfigurationBadRequest with default headers values
func NewUpdateConfigurationBadRequest() *UpdateConfigurationBadRequest {
	return &UpdateConfigurationBadRequest{}
}

/*
UpdateConfigurationBadRequest describes a response with status code 400, with default header values.

The request is invalid
*/
type UpdateConfigurationBadRequest struct {
}

// IsSuccess returns true when this update configuration bad request response has a 2xx status code
func (o *UpdateConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update configuration bad request response has a 3xx status code
func (o *UpdateConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration bad request response has a 4xx status code
func (o *UpdateConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update configuration bad request response has a 5xx status code
func (o *UpdateConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration bad request response a status code equal to that given
func (o *UpdateConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update configuration bad request response
func (o *UpdateConfigurationBadRequest) Code() int {
	return 400
}

func (o *UpdateConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationBadRequest ", 400)
}

func (o *UpdateConfigurationBadRequest) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationBadRequest ", 400)
}

func (o *UpdateConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigurationUnauthorized creates a UpdateConfigurationUnauthorized with default headers values
func NewUpdateConfigurationUnauthorized() *UpdateConfigurationUnauthorized {
	return &UpdateConfigurationUnauthorized{}
}

/*
UpdateConfigurationUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type UpdateConfigurationUnauthorized struct {
}

// IsSuccess returns true when this update configuration unauthorized response has a 2xx status code
func (o *UpdateConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update configuration unauthorized response has a 3xx status code
func (o *UpdateConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration unauthorized response has a 4xx status code
func (o *UpdateConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update configuration unauthorized response has a 5xx status code
func (o *UpdateConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration unauthorized response a status code equal to that given
func (o *UpdateConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update configuration unauthorized response
func (o *UpdateConfigurationUnauthorized) Code() int {
	return 401
}

func (o *UpdateConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationUnauthorized ", 401)
}

func (o *UpdateConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationUnauthorized ", 401)
}

func (o *UpdateConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigurationNotFound creates a UpdateConfigurationNotFound with default headers values
func NewUpdateConfigurationNotFound() *UpdateConfigurationNotFound {
	return &UpdateConfigurationNotFound{}
}

/*
UpdateConfigurationNotFound describes a response with status code 404, with default header values.

The resource is not found
*/
type UpdateConfigurationNotFound struct {
}

// IsSuccess returns true when this update configuration not found response has a 2xx status code
func (o *UpdateConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update configuration not found response has a 3xx status code
func (o *UpdateConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration not found response has a 4xx status code
func (o *UpdateConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update configuration not found response has a 5xx status code
func (o *UpdateConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration not found response a status code equal to that given
func (o *UpdateConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update configuration not found response
func (o *UpdateConfigurationNotFound) Code() int {
	return 404
}

func (o *UpdateConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationNotFound ", 404)
}

func (o *UpdateConfigurationNotFound) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}][%d] updateConfigurationNotFound ", 404)
}

func (o *UpdateConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
