// Code generated by go-swagger; DO NOT EDIT.

package resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteResourcePermissionRuleReader is a Reader for the DeleteResourcePermissionRule structure.
type DeleteResourcePermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourcePermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteResourcePermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteResourcePermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteResourcePermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /resources/{id}/permissions/{ruleId}] deleteResourcePermissionRule", response, response.Code())
	}
}

// NewDeleteResourcePermissionRuleNoContent creates a DeleteResourcePermissionRuleNoContent with default headers values
func NewDeleteResourcePermissionRuleNoContent() *DeleteResourcePermissionRuleNoContent {
	return &DeleteResourcePermissionRuleNoContent{}
}

/*
DeleteResourcePermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteResourcePermissionRuleNoContent struct {
}

// IsSuccess returns true when this delete resource permission rule no content response has a 2xx status code
func (o *DeleteResourcePermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete resource permission rule no content response has a 3xx status code
func (o *DeleteResourcePermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource permission rule no content response has a 4xx status code
func (o *DeleteResourcePermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete resource permission rule no content response has a 5xx status code
func (o *DeleteResourcePermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource permission rule no content response a status code equal to that given
func (o *DeleteResourcePermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete resource permission rule no content response
func (o *DeleteResourcePermissionRuleNoContent) Code() int {
	return 204
}

func (o *DeleteResourcePermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /resources/{id}/permissions/{ruleId}][%d] deleteResourcePermissionRuleNoContent ", 204)
}

func (o *DeleteResourcePermissionRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /resources/{id}/permissions/{ruleId}][%d] deleteResourcePermissionRuleNoContent ", 204)
}

func (o *DeleteResourcePermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourcePermissionRuleUnauthorized creates a DeleteResourcePermissionRuleUnauthorized with default headers values
func NewDeleteResourcePermissionRuleUnauthorized() *DeleteResourcePermissionRuleUnauthorized {
	return &DeleteResourcePermissionRuleUnauthorized{}
}

/*
DeleteResourcePermissionRuleUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type DeleteResourcePermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this delete resource permission rule unauthorized response has a 2xx status code
func (o *DeleteResourcePermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource permission rule unauthorized response has a 3xx status code
func (o *DeleteResourcePermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource permission rule unauthorized response has a 4xx status code
func (o *DeleteResourcePermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource permission rule unauthorized response has a 5xx status code
func (o *DeleteResourcePermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource permission rule unauthorized response a status code equal to that given
func (o *DeleteResourcePermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete resource permission rule unauthorized response
func (o *DeleteResourcePermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteResourcePermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /resources/{id}/permissions/{ruleId}][%d] deleteResourcePermissionRuleUnauthorized ", 401)
}

func (o *DeleteResourcePermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /resources/{id}/permissions/{ruleId}][%d] deleteResourcePermissionRuleUnauthorized ", 401)
}

func (o *DeleteResourcePermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourcePermissionRuleNotFound creates a DeleteResourcePermissionRuleNotFound with default headers values
func NewDeleteResourcePermissionRuleNotFound() *DeleteResourcePermissionRuleNotFound {
	return &DeleteResourcePermissionRuleNotFound{}
}

/*
DeleteResourcePermissionRuleNotFound describes a response with status code 404, with default header values.

Cannot find a resource with the specified ID, the user does not have 'admin' access rights for that resource, or the permission rule with the specified ID does not exist.
*/
type DeleteResourcePermissionRuleNotFound struct {
}

// IsSuccess returns true when this delete resource permission rule not found response has a 2xx status code
func (o *DeleteResourcePermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource permission rule not found response has a 3xx status code
func (o *DeleteResourcePermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource permission rule not found response has a 4xx status code
func (o *DeleteResourcePermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource permission rule not found response has a 5xx status code
func (o *DeleteResourcePermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource permission rule not found response a status code equal to that given
func (o *DeleteResourcePermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete resource permission rule not found response
func (o *DeleteResourcePermissionRuleNotFound) Code() int {
	return 404
}

func (o *DeleteResourcePermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /resources/{id}/permissions/{ruleId}][%d] deleteResourcePermissionRuleNotFound ", 404)
}

func (o *DeleteResourcePermissionRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /resources/{id}/permissions/{ruleId}][%d] deleteResourcePermissionRuleNotFound ", 404)
}

func (o *DeleteResourcePermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
