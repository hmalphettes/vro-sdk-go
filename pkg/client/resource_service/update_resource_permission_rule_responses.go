// Code generated by go-swagger; DO NOT EDIT.

package resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateResourcePermissionRuleReader is a Reader for the UpdateResourcePermissionRule structure.
type UpdateResourcePermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateResourcePermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateResourcePermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateResourcePermissionRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateResourcePermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateResourcePermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /resources/{id}/permissions/{ruleId}] updateResourcePermissionRule", response, response.Code())
	}
}

// NewUpdateResourcePermissionRuleNoContent creates a UpdateResourcePermissionRuleNoContent with default headers values
func NewUpdateResourcePermissionRuleNoContent() *UpdateResourcePermissionRuleNoContent {
	return &UpdateResourcePermissionRuleNoContent{}
}

/*
UpdateResourcePermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type UpdateResourcePermissionRuleNoContent struct {
}

// IsSuccess returns true when this update resource permission rule no content response has a 2xx status code
func (o *UpdateResourcePermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update resource permission rule no content response has a 3xx status code
func (o *UpdateResourcePermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource permission rule no content response has a 4xx status code
func (o *UpdateResourcePermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update resource permission rule no content response has a 5xx status code
func (o *UpdateResourcePermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource permission rule no content response a status code equal to that given
func (o *UpdateResourcePermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update resource permission rule no content response
func (o *UpdateResourcePermissionRuleNoContent) Code() int {
	return 204
}

func (o *UpdateResourcePermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleNoContent ", 204)
}

func (o *UpdateResourcePermissionRuleNoContent) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleNoContent ", 204)
}

func (o *UpdateResourcePermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateResourcePermissionRuleBadRequest creates a UpdateResourcePermissionRuleBadRequest with default headers values
func NewUpdateResourcePermissionRuleBadRequest() *UpdateResourcePermissionRuleBadRequest {
	return &UpdateResourcePermissionRuleBadRequest{}
}

/*
UpdateResourcePermissionRuleBadRequest describes a response with status code 400, with default header values.

Request is not valid (validation error)
*/
type UpdateResourcePermissionRuleBadRequest struct {
}

// IsSuccess returns true when this update resource permission rule bad request response has a 2xx status code
func (o *UpdateResourcePermissionRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource permission rule bad request response has a 3xx status code
func (o *UpdateResourcePermissionRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource permission rule bad request response has a 4xx status code
func (o *UpdateResourcePermissionRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource permission rule bad request response has a 5xx status code
func (o *UpdateResourcePermissionRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource permission rule bad request response a status code equal to that given
func (o *UpdateResourcePermissionRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update resource permission rule bad request response
func (o *UpdateResourcePermissionRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateResourcePermissionRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleBadRequest ", 400)
}

func (o *UpdateResourcePermissionRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleBadRequest ", 400)
}

func (o *UpdateResourcePermissionRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateResourcePermissionRuleUnauthorized creates a UpdateResourcePermissionRuleUnauthorized with default headers values
func NewUpdateResourcePermissionRuleUnauthorized() *UpdateResourcePermissionRuleUnauthorized {
	return &UpdateResourcePermissionRuleUnauthorized{}
}

/*
UpdateResourcePermissionRuleUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type UpdateResourcePermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this update resource permission rule unauthorized response has a 2xx status code
func (o *UpdateResourcePermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource permission rule unauthorized response has a 3xx status code
func (o *UpdateResourcePermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource permission rule unauthorized response has a 4xx status code
func (o *UpdateResourcePermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource permission rule unauthorized response has a 5xx status code
func (o *UpdateResourcePermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource permission rule unauthorized response a status code equal to that given
func (o *UpdateResourcePermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update resource permission rule unauthorized response
func (o *UpdateResourcePermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *UpdateResourcePermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleUnauthorized ", 401)
}

func (o *UpdateResourcePermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleUnauthorized ", 401)
}

func (o *UpdateResourcePermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateResourcePermissionRuleNotFound creates a UpdateResourcePermissionRuleNotFound with default headers values
func NewUpdateResourcePermissionRuleNotFound() *UpdateResourcePermissionRuleNotFound {
	return &UpdateResourcePermissionRuleNotFound{}
}

/*
UpdateResourcePermissionRuleNotFound describes a response with status code 404, with default header values.

Cannot find a resource with the specified ID, the user does not have 'admin' access rights for that resource, or the permission rule with the specified ID does not exist.
*/
type UpdateResourcePermissionRuleNotFound struct {
}

// IsSuccess returns true when this update resource permission rule not found response has a 2xx status code
func (o *UpdateResourcePermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource permission rule not found response has a 3xx status code
func (o *UpdateResourcePermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource permission rule not found response has a 4xx status code
func (o *UpdateResourcePermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource permission rule not found response has a 5xx status code
func (o *UpdateResourcePermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource permission rule not found response a status code equal to that given
func (o *UpdateResourcePermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update resource permission rule not found response
func (o *UpdateResourcePermissionRuleNotFound) Code() int {
	return 404
}

func (o *UpdateResourcePermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleNotFound ", 404)
}

func (o *UpdateResourcePermissionRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/permissions/{ruleId}][%d] updateResourcePermissionRuleNotFound ", 404)
}

func (o *UpdateResourcePermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
