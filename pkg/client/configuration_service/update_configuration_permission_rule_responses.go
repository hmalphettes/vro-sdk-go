// Code generated by go-swagger; DO NOT EDIT.

package configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateConfigurationPermissionRuleReader is a Reader for the UpdateConfigurationPermissionRule structure.
type UpdateConfigurationPermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigurationPermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateConfigurationPermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConfigurationPermissionRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateConfigurationPermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConfigurationPermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /configurations/{id}/permissions/{ruleId}] updateConfigurationPermissionRule", response, response.Code())
	}
}

// NewUpdateConfigurationPermissionRuleNoContent creates a UpdateConfigurationPermissionRuleNoContent with default headers values
func NewUpdateConfigurationPermissionRuleNoContent() *UpdateConfigurationPermissionRuleNoContent {
	return &UpdateConfigurationPermissionRuleNoContent{}
}

/*
UpdateConfigurationPermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type UpdateConfigurationPermissionRuleNoContent struct {
}

// IsSuccess returns true when this update configuration permission rule no content response has a 2xx status code
func (o *UpdateConfigurationPermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update configuration permission rule no content response has a 3xx status code
func (o *UpdateConfigurationPermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration permission rule no content response has a 4xx status code
func (o *UpdateConfigurationPermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update configuration permission rule no content response has a 5xx status code
func (o *UpdateConfigurationPermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration permission rule no content response a status code equal to that given
func (o *UpdateConfigurationPermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update configuration permission rule no content response
func (o *UpdateConfigurationPermissionRuleNoContent) Code() int {
	return 204
}

func (o *UpdateConfigurationPermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleNoContent ", 204)
}

func (o *UpdateConfigurationPermissionRuleNoContent) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleNoContent ", 204)
}

func (o *UpdateConfigurationPermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigurationPermissionRuleBadRequest creates a UpdateConfigurationPermissionRuleBadRequest with default headers values
func NewUpdateConfigurationPermissionRuleBadRequest() *UpdateConfigurationPermissionRuleBadRequest {
	return &UpdateConfigurationPermissionRuleBadRequest{}
}

/*
UpdateConfigurationPermissionRuleBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type UpdateConfigurationPermissionRuleBadRequest struct {
}

// IsSuccess returns true when this update configuration permission rule bad request response has a 2xx status code
func (o *UpdateConfigurationPermissionRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update configuration permission rule bad request response has a 3xx status code
func (o *UpdateConfigurationPermissionRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration permission rule bad request response has a 4xx status code
func (o *UpdateConfigurationPermissionRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update configuration permission rule bad request response has a 5xx status code
func (o *UpdateConfigurationPermissionRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration permission rule bad request response a status code equal to that given
func (o *UpdateConfigurationPermissionRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update configuration permission rule bad request response
func (o *UpdateConfigurationPermissionRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateConfigurationPermissionRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleBadRequest ", 400)
}

func (o *UpdateConfigurationPermissionRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleBadRequest ", 400)
}

func (o *UpdateConfigurationPermissionRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigurationPermissionRuleUnauthorized creates a UpdateConfigurationPermissionRuleUnauthorized with default headers values
func NewUpdateConfigurationPermissionRuleUnauthorized() *UpdateConfigurationPermissionRuleUnauthorized {
	return &UpdateConfigurationPermissionRuleUnauthorized{}
}

/*
UpdateConfigurationPermissionRuleUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type UpdateConfigurationPermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this update configuration permission rule unauthorized response has a 2xx status code
func (o *UpdateConfigurationPermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update configuration permission rule unauthorized response has a 3xx status code
func (o *UpdateConfigurationPermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration permission rule unauthorized response has a 4xx status code
func (o *UpdateConfigurationPermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update configuration permission rule unauthorized response has a 5xx status code
func (o *UpdateConfigurationPermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration permission rule unauthorized response a status code equal to that given
func (o *UpdateConfigurationPermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update configuration permission rule unauthorized response
func (o *UpdateConfigurationPermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *UpdateConfigurationPermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleUnauthorized ", 401)
}

func (o *UpdateConfigurationPermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleUnauthorized ", 401)
}

func (o *UpdateConfigurationPermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigurationPermissionRuleNotFound creates a UpdateConfigurationPermissionRuleNotFound with default headers values
func NewUpdateConfigurationPermissionRuleNotFound() *UpdateConfigurationPermissionRuleNotFound {
	return &UpdateConfigurationPermissionRuleNotFound{}
}

/*
UpdateConfigurationPermissionRuleNotFound describes a response with status code 404, with default header values.

Cannot find a configuration with the specified ID, the user does not have 'admin' access rights for that configuration, or the permission rule with the specified ID does not exist
*/
type UpdateConfigurationPermissionRuleNotFound struct {
}

// IsSuccess returns true when this update configuration permission rule not found response has a 2xx status code
func (o *UpdateConfigurationPermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update configuration permission rule not found response has a 3xx status code
func (o *UpdateConfigurationPermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update configuration permission rule not found response has a 4xx status code
func (o *UpdateConfigurationPermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update configuration permission rule not found response has a 5xx status code
func (o *UpdateConfigurationPermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update configuration permission rule not found response a status code equal to that given
func (o *UpdateConfigurationPermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update configuration permission rule not found response
func (o *UpdateConfigurationPermissionRuleNotFound) Code() int {
	return 404
}

func (o *UpdateConfigurationPermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleNotFound ", 404)
}

func (o *UpdateConfigurationPermissionRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /configurations/{id}/permissions/{ruleId}][%d] updateConfigurationPermissionRuleNotFound ", 404)
}

func (o *UpdateConfigurationPermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
