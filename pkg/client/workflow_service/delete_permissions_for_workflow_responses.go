// Code generated by go-swagger; DO NOT EDIT.

package workflow_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePermissionsForWorkflowReader is a Reader for the DeletePermissionsForWorkflow structure.
type DeletePermissionsForWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePermissionsForWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePermissionsForWorkflowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePermissionsForWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePermissionsForWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /workflows/{id}/permissions] deletePermissionsForWorkflow", response, response.Code())
	}
}

// NewDeletePermissionsForWorkflowNoContent creates a DeletePermissionsForWorkflowNoContent with default headers values
func NewDeletePermissionsForWorkflowNoContent() *DeletePermissionsForWorkflowNoContent {
	return &DeletePermissionsForWorkflowNoContent{}
}

/*
DeletePermissionsForWorkflowNoContent describes a response with status code 204, with default header values.

No content
*/
type DeletePermissionsForWorkflowNoContent struct {
}

// IsSuccess returns true when this delete permissions for workflow no content response has a 2xx status code
func (o *DeletePermissionsForWorkflowNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete permissions for workflow no content response has a 3xx status code
func (o *DeletePermissionsForWorkflowNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete permissions for workflow no content response has a 4xx status code
func (o *DeletePermissionsForWorkflowNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete permissions for workflow no content response has a 5xx status code
func (o *DeletePermissionsForWorkflowNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete permissions for workflow no content response a status code equal to that given
func (o *DeletePermissionsForWorkflowNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete permissions for workflow no content response
func (o *DeletePermissionsForWorkflowNoContent) Code() int {
	return 204
}

func (o *DeletePermissionsForWorkflowNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{id}/permissions][%d] deletePermissionsForWorkflowNoContent ", 204)
}

func (o *DeletePermissionsForWorkflowNoContent) String() string {
	return fmt.Sprintf("[DELETE /workflows/{id}/permissions][%d] deletePermissionsForWorkflowNoContent ", 204)
}

func (o *DeletePermissionsForWorkflowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePermissionsForWorkflowUnauthorized creates a DeletePermissionsForWorkflowUnauthorized with default headers values
func NewDeletePermissionsForWorkflowUnauthorized() *DeletePermissionsForWorkflowUnauthorized {
	return &DeletePermissionsForWorkflowUnauthorized{}
}

/*
DeletePermissionsForWorkflowUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type DeletePermissionsForWorkflowUnauthorized struct {
}

// IsSuccess returns true when this delete permissions for workflow unauthorized response has a 2xx status code
func (o *DeletePermissionsForWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete permissions for workflow unauthorized response has a 3xx status code
func (o *DeletePermissionsForWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete permissions for workflow unauthorized response has a 4xx status code
func (o *DeletePermissionsForWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete permissions for workflow unauthorized response has a 5xx status code
func (o *DeletePermissionsForWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete permissions for workflow unauthorized response a status code equal to that given
func (o *DeletePermissionsForWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete permissions for workflow unauthorized response
func (o *DeletePermissionsForWorkflowUnauthorized) Code() int {
	return 401
}

func (o *DeletePermissionsForWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{id}/permissions][%d] deletePermissionsForWorkflowUnauthorized ", 401)
}

func (o *DeletePermissionsForWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /workflows/{id}/permissions][%d] deletePermissionsForWorkflowUnauthorized ", 401)
}

func (o *DeletePermissionsForWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePermissionsForWorkflowNotFound creates a DeletePermissionsForWorkflowNotFound with default headers values
func NewDeletePermissionsForWorkflowNotFound() *DeletePermissionsForWorkflowNotFound {
	return &DeletePermissionsForWorkflowNotFound{}
}

/*
DeletePermissionsForWorkflowNotFound describes a response with status code 404, with default header values.

Cannot find an action with the specified name
*/
type DeletePermissionsForWorkflowNotFound struct {
}

// IsSuccess returns true when this delete permissions for workflow not found response has a 2xx status code
func (o *DeletePermissionsForWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete permissions for workflow not found response has a 3xx status code
func (o *DeletePermissionsForWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete permissions for workflow not found response has a 4xx status code
func (o *DeletePermissionsForWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete permissions for workflow not found response has a 5xx status code
func (o *DeletePermissionsForWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete permissions for workflow not found response a status code equal to that given
func (o *DeletePermissionsForWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete permissions for workflow not found response
func (o *DeletePermissionsForWorkflowNotFound) Code() int {
	return 404
}

func (o *DeletePermissionsForWorkflowNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{id}/permissions][%d] deletePermissionsForWorkflowNotFound ", 404)
}

func (o *DeletePermissionsForWorkflowNotFound) String() string {
	return fmt.Sprintf("[DELETE /workflows/{id}/permissions][%d] deletePermissionsForWorkflowNotFound ", 404)
}

func (o *DeletePermissionsForWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}