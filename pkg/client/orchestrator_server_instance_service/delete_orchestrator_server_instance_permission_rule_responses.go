// Code generated by go-swagger; DO NOT EDIT.

package orchestrator_server_instance_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrchestratorServerInstancePermissionRuleReader is a Reader for the DeleteOrchestratorServerInstancePermissionRule structure.
type DeleteOrchestratorServerInstancePermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrchestratorServerInstancePermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrchestratorServerInstancePermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOrchestratorServerInstancePermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrchestratorServerInstancePermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /server/permissions/{ruleId}] deleteOrchestratorServerInstancePermissionRule", response, response.Code())
	}
}

// NewDeleteOrchestratorServerInstancePermissionRuleNoContent creates a DeleteOrchestratorServerInstancePermissionRuleNoContent with default headers values
func NewDeleteOrchestratorServerInstancePermissionRuleNoContent() *DeleteOrchestratorServerInstancePermissionRuleNoContent {
	return &DeleteOrchestratorServerInstancePermissionRuleNoContent{}
}

/*
DeleteOrchestratorServerInstancePermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteOrchestratorServerInstancePermissionRuleNoContent struct {
}

// IsSuccess returns true when this delete orchestrator server instance permission rule no content response has a 2xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete orchestrator server instance permission rule no content response has a 3xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete orchestrator server instance permission rule no content response has a 4xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete orchestrator server instance permission rule no content response has a 5xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete orchestrator server instance permission rule no content response a status code equal to that given
func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete orchestrator server instance permission rule no content response
func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) Code() int {
	return 204
}

func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /server/permissions/{ruleId}][%d] deleteOrchestratorServerInstancePermissionRuleNoContent ", 204)
}

func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /server/permissions/{ruleId}][%d] deleteOrchestratorServerInstancePermissionRuleNoContent ", 204)
}

func (o *DeleteOrchestratorServerInstancePermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrchestratorServerInstancePermissionRuleUnauthorized creates a DeleteOrchestratorServerInstancePermissionRuleUnauthorized with default headers values
func NewDeleteOrchestratorServerInstancePermissionRuleUnauthorized() *DeleteOrchestratorServerInstancePermissionRuleUnauthorized {
	return &DeleteOrchestratorServerInstancePermissionRuleUnauthorized{}
}

/*
DeleteOrchestratorServerInstancePermissionRuleUnauthorized describes a response with status code 401, with default header values.

The user is not authorized.
*/
type DeleteOrchestratorServerInstancePermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this delete orchestrator server instance permission rule unauthorized response has a 2xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete orchestrator server instance permission rule unauthorized response has a 3xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete orchestrator server instance permission rule unauthorized response has a 4xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete orchestrator server instance permission rule unauthorized response has a 5xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete orchestrator server instance permission rule unauthorized response a status code equal to that given
func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete orchestrator server instance permission rule unauthorized response
func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /server/permissions/{ruleId}][%d] deleteOrchestratorServerInstancePermissionRuleUnauthorized ", 401)
}

func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /server/permissions/{ruleId}][%d] deleteOrchestratorServerInstancePermissionRuleUnauthorized ", 401)
}

func (o *DeleteOrchestratorServerInstancePermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrchestratorServerInstancePermissionRuleNotFound creates a DeleteOrchestratorServerInstancePermissionRuleNotFound with default headers values
func NewDeleteOrchestratorServerInstancePermissionRuleNotFound() *DeleteOrchestratorServerInstancePermissionRuleNotFound {
	return &DeleteOrchestratorServerInstancePermissionRuleNotFound{}
}

/*
DeleteOrchestratorServerInstancePermissionRuleNotFound describes a response with status code 404, with default header values.

The user does not have 'read' access rights for the server object, or the permission rule with the specified ID does not exist.
*/
type DeleteOrchestratorServerInstancePermissionRuleNotFound struct {
}

// IsSuccess returns true when this delete orchestrator server instance permission rule not found response has a 2xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete orchestrator server instance permission rule not found response has a 3xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete orchestrator server instance permission rule not found response has a 4xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete orchestrator server instance permission rule not found response has a 5xx status code
func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete orchestrator server instance permission rule not found response a status code equal to that given
func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete orchestrator server instance permission rule not found response
func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) Code() int {
	return 404
}

func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /server/permissions/{ruleId}][%d] deleteOrchestratorServerInstancePermissionRuleNotFound ", 404)
}

func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /server/permissions/{ruleId}][%d] deleteOrchestratorServerInstancePermissionRuleNotFound ", 404)
}

func (o *DeleteOrchestratorServerInstancePermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
