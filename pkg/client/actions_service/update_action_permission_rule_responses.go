// Code generated by go-swagger; DO NOT EDIT.

package actions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateActionPermissionRuleReader is a Reader for the UpdateActionPermissionRule structure.
type UpdateActionPermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateActionPermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateActionPermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateActionPermissionRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateActionPermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateActionPermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /actions/{id}/permissions/{ruleId}] updateActionPermissionRule", response, response.Code())
	}
}

// NewUpdateActionPermissionRuleNoContent creates a UpdateActionPermissionRuleNoContent with default headers values
func NewUpdateActionPermissionRuleNoContent() *UpdateActionPermissionRuleNoContent {
	return &UpdateActionPermissionRuleNoContent{}
}

/*
UpdateActionPermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type UpdateActionPermissionRuleNoContent struct {
}

// IsSuccess returns true when this update action permission rule no content response has a 2xx status code
func (o *UpdateActionPermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update action permission rule no content response has a 3xx status code
func (o *UpdateActionPermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action permission rule no content response has a 4xx status code
func (o *UpdateActionPermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update action permission rule no content response has a 5xx status code
func (o *UpdateActionPermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update action permission rule no content response a status code equal to that given
func (o *UpdateActionPermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update action permission rule no content response
func (o *UpdateActionPermissionRuleNoContent) Code() int {
	return 204
}

func (o *UpdateActionPermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleNoContent ", 204)
}

func (o *UpdateActionPermissionRuleNoContent) String() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleNoContent ", 204)
}

func (o *UpdateActionPermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateActionPermissionRuleBadRequest creates a UpdateActionPermissionRuleBadRequest with default headers values
func NewUpdateActionPermissionRuleBadRequest() *UpdateActionPermissionRuleBadRequest {
	return &UpdateActionPermissionRuleBadRequest{}
}

/*
UpdateActionPermissionRuleBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type UpdateActionPermissionRuleBadRequest struct {
}

// IsSuccess returns true when this update action permission rule bad request response has a 2xx status code
func (o *UpdateActionPermissionRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action permission rule bad request response has a 3xx status code
func (o *UpdateActionPermissionRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action permission rule bad request response has a 4xx status code
func (o *UpdateActionPermissionRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action permission rule bad request response has a 5xx status code
func (o *UpdateActionPermissionRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update action permission rule bad request response a status code equal to that given
func (o *UpdateActionPermissionRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update action permission rule bad request response
func (o *UpdateActionPermissionRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateActionPermissionRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleBadRequest ", 400)
}

func (o *UpdateActionPermissionRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleBadRequest ", 400)
}

func (o *UpdateActionPermissionRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateActionPermissionRuleUnauthorized creates a UpdateActionPermissionRuleUnauthorized with default headers values
func NewUpdateActionPermissionRuleUnauthorized() *UpdateActionPermissionRuleUnauthorized {
	return &UpdateActionPermissionRuleUnauthorized{}
}

/*
UpdateActionPermissionRuleUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type UpdateActionPermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this update action permission rule unauthorized response has a 2xx status code
func (o *UpdateActionPermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action permission rule unauthorized response has a 3xx status code
func (o *UpdateActionPermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action permission rule unauthorized response has a 4xx status code
func (o *UpdateActionPermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action permission rule unauthorized response has a 5xx status code
func (o *UpdateActionPermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update action permission rule unauthorized response a status code equal to that given
func (o *UpdateActionPermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update action permission rule unauthorized response
func (o *UpdateActionPermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *UpdateActionPermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleUnauthorized ", 401)
}

func (o *UpdateActionPermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleUnauthorized ", 401)
}

func (o *UpdateActionPermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateActionPermissionRuleNotFound creates a UpdateActionPermissionRuleNotFound with default headers values
func NewUpdateActionPermissionRuleNotFound() *UpdateActionPermissionRuleNotFound {
	return &UpdateActionPermissionRuleNotFound{}
}

/*
UpdateActionPermissionRuleNotFound describes a response with status code 404, with default header values.

Can not find an action with the specified ID, the user does not have 'read' access rights for that action, or the permission rule with the specified ID does not exist
*/
type UpdateActionPermissionRuleNotFound struct {
}

// IsSuccess returns true when this update action permission rule not found response has a 2xx status code
func (o *UpdateActionPermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update action permission rule not found response has a 3xx status code
func (o *UpdateActionPermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update action permission rule not found response has a 4xx status code
func (o *UpdateActionPermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update action permission rule not found response has a 5xx status code
func (o *UpdateActionPermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update action permission rule not found response a status code equal to that given
func (o *UpdateActionPermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update action permission rule not found response
func (o *UpdateActionPermissionRuleNotFound) Code() int {
	return 404
}

func (o *UpdateActionPermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleNotFound ", 404)
}

func (o *UpdateActionPermissionRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /actions/{id}/permissions/{ruleId}][%d] updateActionPermissionRuleNotFound ", 404)
}

func (o *UpdateActionPermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}