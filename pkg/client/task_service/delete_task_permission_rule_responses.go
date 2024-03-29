// Code generated by go-swagger; DO NOT EDIT.

package task_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTaskPermissionRuleReader is a Reader for the DeleteTaskPermissionRule structure.
type DeleteTaskPermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTaskPermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteTaskPermissionRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTaskPermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTaskPermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /tasks/{id}/permissions/{ruleId}] deleteTaskPermissionRule", response, response.Code())
	}
}

// NewDeleteTaskPermissionRuleAccepted creates a DeleteTaskPermissionRuleAccepted with default headers values
func NewDeleteTaskPermissionRuleAccepted() *DeleteTaskPermissionRuleAccepted {
	return &DeleteTaskPermissionRuleAccepted{}
}

/*
DeleteTaskPermissionRuleAccepted describes a response with status code 202, with default header values.

No content
*/
type DeleteTaskPermissionRuleAccepted struct {
}

// IsSuccess returns true when this delete task permission rule accepted response has a 2xx status code
func (o *DeleteTaskPermissionRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete task permission rule accepted response has a 3xx status code
func (o *DeleteTaskPermissionRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task permission rule accepted response has a 4xx status code
func (o *DeleteTaskPermissionRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete task permission rule accepted response has a 5xx status code
func (o *DeleteTaskPermissionRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task permission rule accepted response a status code equal to that given
func (o *DeleteTaskPermissionRuleAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete task permission rule accepted response
func (o *DeleteTaskPermissionRuleAccepted) Code() int {
	return 202
}

func (o *DeleteTaskPermissionRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}/permissions/{ruleId}][%d] deleteTaskPermissionRuleAccepted ", 202)
}

func (o *DeleteTaskPermissionRuleAccepted) String() string {
	return fmt.Sprintf("[DELETE /tasks/{id}/permissions/{ruleId}][%d] deleteTaskPermissionRuleAccepted ", 202)
}

func (o *DeleteTaskPermissionRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaskPermissionRuleUnauthorized creates a DeleteTaskPermissionRuleUnauthorized with default headers values
func NewDeleteTaskPermissionRuleUnauthorized() *DeleteTaskPermissionRuleUnauthorized {
	return &DeleteTaskPermissionRuleUnauthorized{}
}

/*
DeleteTaskPermissionRuleUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type DeleteTaskPermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this delete task permission rule unauthorized response has a 2xx status code
func (o *DeleteTaskPermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task permission rule unauthorized response has a 3xx status code
func (o *DeleteTaskPermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task permission rule unauthorized response has a 4xx status code
func (o *DeleteTaskPermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task permission rule unauthorized response has a 5xx status code
func (o *DeleteTaskPermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task permission rule unauthorized response a status code equal to that given
func (o *DeleteTaskPermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete task permission rule unauthorized response
func (o *DeleteTaskPermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteTaskPermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}/permissions/{ruleId}][%d] deleteTaskPermissionRuleUnauthorized ", 401)
}

func (o *DeleteTaskPermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /tasks/{id}/permissions/{ruleId}][%d] deleteTaskPermissionRuleUnauthorized ", 401)
}

func (o *DeleteTaskPermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaskPermissionRuleNotFound creates a DeleteTaskPermissionRuleNotFound with default headers values
func NewDeleteTaskPermissionRuleNotFound() *DeleteTaskPermissionRuleNotFound {
	return &DeleteTaskPermissionRuleNotFound{}
}

/*
DeleteTaskPermissionRuleNotFound describes a response with status code 404, with default header values.

Can not find a task with the specified ID, the user does not have 'admin' access rights for that task, or the permission rule with the specified ID does not exist.
*/
type DeleteTaskPermissionRuleNotFound struct {
}

// IsSuccess returns true when this delete task permission rule not found response has a 2xx status code
func (o *DeleteTaskPermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task permission rule not found response has a 3xx status code
func (o *DeleteTaskPermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task permission rule not found response has a 4xx status code
func (o *DeleteTaskPermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task permission rule not found response has a 5xx status code
func (o *DeleteTaskPermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task permission rule not found response a status code equal to that given
func (o *DeleteTaskPermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete task permission rule not found response
func (o *DeleteTaskPermissionRuleNotFound) Code() int {
	return 404
}

func (o *DeleteTaskPermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}/permissions/{ruleId}][%d] deleteTaskPermissionRuleNotFound ", 404)
}

func (o *DeleteTaskPermissionRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /tasks/{id}/permissions/{ruleId}][%d] deleteTaskPermissionRuleNotFound ", 404)
}

func (o *DeleteTaskPermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
