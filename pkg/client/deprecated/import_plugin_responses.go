// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportPluginReader is a Reader for the ImportPlugin structure.
type ImportPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewImportPluginNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImportPluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportPluginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImportPluginConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /plugins] importPlugin", response, response.Code())
	}
}

// NewImportPluginNoContent creates a ImportPluginNoContent with default headers values
func NewImportPluginNoContent() *ImportPluginNoContent {
	return &ImportPluginNoContent{}
}

/*
ImportPluginNoContent describes a response with status code 204, with default header values.

No content
*/
type ImportPluginNoContent struct {
}

// IsSuccess returns true when this import plugin no content response has a 2xx status code
func (o *ImportPluginNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import plugin no content response has a 3xx status code
func (o *ImportPluginNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import plugin no content response has a 4xx status code
func (o *ImportPluginNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this import plugin no content response has a 5xx status code
func (o *ImportPluginNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this import plugin no content response a status code equal to that given
func (o *ImportPluginNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the import plugin no content response
func (o *ImportPluginNoContent) Code() int {
	return 204
}

func (o *ImportPluginNoContent) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginNoContent ", 204)
}

func (o *ImportPluginNoContent) String() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginNoContent ", 204)
}

func (o *ImportPluginNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportPluginUnauthorized creates a ImportPluginUnauthorized with default headers values
func NewImportPluginUnauthorized() *ImportPluginUnauthorized {
	return &ImportPluginUnauthorized{}
}

/*
ImportPluginUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ImportPluginUnauthorized struct {
}

// IsSuccess returns true when this import plugin unauthorized response has a 2xx status code
func (o *ImportPluginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import plugin unauthorized response has a 3xx status code
func (o *ImportPluginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import plugin unauthorized response has a 4xx status code
func (o *ImportPluginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this import plugin unauthorized response has a 5xx status code
func (o *ImportPluginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this import plugin unauthorized response a status code equal to that given
func (o *ImportPluginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the import plugin unauthorized response
func (o *ImportPluginUnauthorized) Code() int {
	return 401
}

func (o *ImportPluginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginUnauthorized ", 401)
}

func (o *ImportPluginUnauthorized) String() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginUnauthorized ", 401)
}

func (o *ImportPluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportPluginForbidden creates a ImportPluginForbidden with default headers values
func NewImportPluginForbidden() *ImportPluginForbidden {
	return &ImportPluginForbidden{}
}

/*
ImportPluginForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ImportPluginForbidden struct {
}

// IsSuccess returns true when this import plugin forbidden response has a 2xx status code
func (o *ImportPluginForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import plugin forbidden response has a 3xx status code
func (o *ImportPluginForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import plugin forbidden response has a 4xx status code
func (o *ImportPluginForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this import plugin forbidden response has a 5xx status code
func (o *ImportPluginForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this import plugin forbidden response a status code equal to that given
func (o *ImportPluginForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the import plugin forbidden response
func (o *ImportPluginForbidden) Code() int {
	return 403
}

func (o *ImportPluginForbidden) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginForbidden ", 403)
}

func (o *ImportPluginForbidden) String() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginForbidden ", 403)
}

func (o *ImportPluginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportPluginConflict creates a ImportPluginConflict with default headers values
func NewImportPluginConflict() *ImportPluginConflict {
	return &ImportPluginConflict{}
}

/*
ImportPluginConflict describes a response with status code 409, with default header values.

Plug-in already exists
*/
type ImportPluginConflict struct {
}

// IsSuccess returns true when this import plugin conflict response has a 2xx status code
func (o *ImportPluginConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import plugin conflict response has a 3xx status code
func (o *ImportPluginConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import plugin conflict response has a 4xx status code
func (o *ImportPluginConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this import plugin conflict response has a 5xx status code
func (o *ImportPluginConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this import plugin conflict response a status code equal to that given
func (o *ImportPluginConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the import plugin conflict response
func (o *ImportPluginConflict) Code() int {
	return 409
}

func (o *ImportPluginConflict) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginConflict ", 409)
}

func (o *ImportPluginConflict) String() string {
	return fmt.Sprintf("[POST /plugins][%d] importPluginConflict ", 409)
}

func (o *ImportPluginConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
