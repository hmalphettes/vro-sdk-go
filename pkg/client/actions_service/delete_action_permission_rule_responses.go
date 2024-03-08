// Code generated by go-swagger; DO NOT EDIT.

package actions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteActionPermissionRuleReader is a Reader for the DeleteActionPermissionRule structure.
type DeleteActionPermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteActionPermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteActionPermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteActionPermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteActionPermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /actions/{id}/permissions/{ruleId}] deleteActionPermissionRule", response, response.Code())
	}
}

// NewDeleteActionPermissionRuleNoContent creates a DeleteActionPermissionRuleNoContent with default headers values
func NewDeleteActionPermissionRuleNoContent() *DeleteActionPermissionRuleNoContent {
	return &DeleteActionPermissionRuleNoContent{}
}

/*
DeleteActionPermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteActionPermissionRuleNoContent struct {
}

// IsSuccess returns true when this delete action permission rule no content response has a 2xx status code
func (o *DeleteActionPermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete action permission rule no content response has a 3xx status code
func (o *DeleteActionPermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action permission rule no content response has a 4xx status code
func (o *DeleteActionPermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete action permission rule no content response has a 5xx status code
func (o *DeleteActionPermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action permission rule no content response a status code equal to that given
func (o *DeleteActionPermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete action permission rule no content response
func (o *DeleteActionPermissionRuleNoContent) Code() int {
	return 204
}

func (o *DeleteActionPermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/permissions/{ruleId}][%d] deleteActionPermissionRuleNoContent ", 204)
}

func (o *DeleteActionPermissionRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /actions/{id}/permissions/{ruleId}][%d] deleteActionPermissionRuleNoContent ", 204)
}

func (o *DeleteActionPermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteActionPermissionRuleUnauthorized creates a DeleteActionPermissionRuleUnauthorized with default headers values
func NewDeleteActionPermissionRuleUnauthorized() *DeleteActionPermissionRuleUnauthorized {
	return &DeleteActionPermissionRuleUnauthorized{}
}

/*
DeleteActionPermissionRuleUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type DeleteActionPermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this delete action permission rule unauthorized response has a 2xx status code
func (o *DeleteActionPermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action permission rule unauthorized response has a 3xx status code
func (o *DeleteActionPermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action permission rule unauthorized response has a 4xx status code
func (o *DeleteActionPermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete action permission rule unauthorized response has a 5xx status code
func (o *DeleteActionPermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action permission rule unauthorized response a status code equal to that given
func (o *DeleteActionPermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete action permission rule unauthorized response
func (o *DeleteActionPermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteActionPermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/permissions/{ruleId}][%d] deleteActionPermissionRuleUnauthorized ", 401)
}

func (o *DeleteActionPermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /actions/{id}/permissions/{ruleId}][%d] deleteActionPermissionRuleUnauthorized ", 401)
}

func (o *DeleteActionPermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteActionPermissionRuleNotFound creates a DeleteActionPermissionRuleNotFound with default headers values
func NewDeleteActionPermissionRuleNotFound() *DeleteActionPermissionRuleNotFound {
	return &DeleteActionPermissionRuleNotFound{}
}

/*
DeleteActionPermissionRuleNotFound describes a response with status code 404, with default header values.

Can not find an action with the specified ID, the user does not have 'read' access rights for that action, or the permission rule with the specified ID does not exist
*/
type DeleteActionPermissionRuleNotFound struct {
}

// IsSuccess returns true when this delete action permission rule not found response has a 2xx status code
func (o *DeleteActionPermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action permission rule not found response has a 3xx status code
func (o *DeleteActionPermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action permission rule not found response has a 4xx status code
func (o *DeleteActionPermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete action permission rule not found response has a 5xx status code
func (o *DeleteActionPermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action permission rule not found response a status code equal to that given
func (o *DeleteActionPermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete action permission rule not found response
func (o *DeleteActionPermissionRuleNotFound) Code() int {
	return 404
}

func (o *DeleteActionPermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/permissions/{ruleId}][%d] deleteActionPermissionRuleNotFound ", 404)
}

func (o *DeleteActionPermissionRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /actions/{id}/permissions/{ruleId}][%d] deleteActionPermissionRuleNotFound ", 404)
}

func (o *DeleteActionPermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}