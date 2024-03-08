// Code generated by go-swagger; DO NOT EDIT.

package category_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCategoryPermissionRuleReader is a Reader for the DeleteCategoryPermissionRule structure.
type DeleteCategoryPermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCategoryPermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCategoryPermissionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCategoryPermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCategoryPermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /categories/{id}/permissions/{ruleId}] deleteCategoryPermissionRule", response, response.Code())
	}
}

// NewDeleteCategoryPermissionRuleNoContent creates a DeleteCategoryPermissionRuleNoContent with default headers values
func NewDeleteCategoryPermissionRuleNoContent() *DeleteCategoryPermissionRuleNoContent {
	return &DeleteCategoryPermissionRuleNoContent{}
}

/*
DeleteCategoryPermissionRuleNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteCategoryPermissionRuleNoContent struct {
}

// IsSuccess returns true when this delete category permission rule no content response has a 2xx status code
func (o *DeleteCategoryPermissionRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete category permission rule no content response has a 3xx status code
func (o *DeleteCategoryPermissionRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category permission rule no content response has a 4xx status code
func (o *DeleteCategoryPermissionRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete category permission rule no content response has a 5xx status code
func (o *DeleteCategoryPermissionRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category permission rule no content response a status code equal to that given
func (o *DeleteCategoryPermissionRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete category permission rule no content response
func (o *DeleteCategoryPermissionRuleNoContent) Code() int {
	return 204
}

func (o *DeleteCategoryPermissionRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /categories/{id}/permissions/{ruleId}][%d] deleteCategoryPermissionRuleNoContent ", 204)
}

func (o *DeleteCategoryPermissionRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /categories/{id}/permissions/{ruleId}][%d] deleteCategoryPermissionRuleNoContent ", 204)
}

func (o *DeleteCategoryPermissionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryPermissionRuleUnauthorized creates a DeleteCategoryPermissionRuleUnauthorized with default headers values
func NewDeleteCategoryPermissionRuleUnauthorized() *DeleteCategoryPermissionRuleUnauthorized {
	return &DeleteCategoryPermissionRuleUnauthorized{}
}

/*
DeleteCategoryPermissionRuleUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type DeleteCategoryPermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this delete category permission rule unauthorized response has a 2xx status code
func (o *DeleteCategoryPermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete category permission rule unauthorized response has a 3xx status code
func (o *DeleteCategoryPermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category permission rule unauthorized response has a 4xx status code
func (o *DeleteCategoryPermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete category permission rule unauthorized response has a 5xx status code
func (o *DeleteCategoryPermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category permission rule unauthorized response a status code equal to that given
func (o *DeleteCategoryPermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete category permission rule unauthorized response
func (o *DeleteCategoryPermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *DeleteCategoryPermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /categories/{id}/permissions/{ruleId}][%d] deleteCategoryPermissionRuleUnauthorized ", 401)
}

func (o *DeleteCategoryPermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /categories/{id}/permissions/{ruleId}][%d] deleteCategoryPermissionRuleUnauthorized ", 401)
}

func (o *DeleteCategoryPermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoryPermissionRuleNotFound creates a DeleteCategoryPermissionRuleNotFound with default headers values
func NewDeleteCategoryPermissionRuleNotFound() *DeleteCategoryPermissionRuleNotFound {
	return &DeleteCategoryPermissionRuleNotFound{}
}

/*
DeleteCategoryPermissionRuleNotFound describes a response with status code 404, with default header values.

Cannot find a category with the specified ID, the user does not have 'admin' access rights for that category, or the permission rule with the specified ID does not exist
*/
type DeleteCategoryPermissionRuleNotFound struct {
}

// IsSuccess returns true when this delete category permission rule not found response has a 2xx status code
func (o *DeleteCategoryPermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete category permission rule not found response has a 3xx status code
func (o *DeleteCategoryPermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category permission rule not found response has a 4xx status code
func (o *DeleteCategoryPermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete category permission rule not found response has a 5xx status code
func (o *DeleteCategoryPermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category permission rule not found response a status code equal to that given
func (o *DeleteCategoryPermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete category permission rule not found response
func (o *DeleteCategoryPermissionRuleNotFound) Code() int {
	return 404
}

func (o *DeleteCategoryPermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /categories/{id}/permissions/{ruleId}][%d] deleteCategoryPermissionRuleNotFound ", 404)
}

func (o *DeleteCategoryPermissionRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /categories/{id}/permissions/{ruleId}][%d] deleteCategoryPermissionRuleNotFound ", 404)
}

func (o *DeleteCategoryPermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
