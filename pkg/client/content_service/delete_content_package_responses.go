// Code generated by go-swagger; DO NOT EDIT.

package content_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteContentPackageReader is a Reader for the DeleteContentPackage structure.
type DeleteContentPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteContentPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteContentPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteContentPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /content/packages/{packageName}] deleteContentPackage", response, response.Code())
	}
}

// NewDeleteContentPackageOK creates a DeleteContentPackageOK with default headers values
func NewDeleteContentPackageOK() *DeleteContentPackageOK {
	return &DeleteContentPackageOK{}
}

/*
DeleteContentPackageOK describes a response with status code 200, with default header values.

The request is successful
*/
type DeleteContentPackageOK struct {
}

// IsSuccess returns true when this delete content package o k response has a 2xx status code
func (o *DeleteContentPackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete content package o k response has a 3xx status code
func (o *DeleteContentPackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete content package o k response has a 4xx status code
func (o *DeleteContentPackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete content package o k response has a 5xx status code
func (o *DeleteContentPackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete content package o k response a status code equal to that given
func (o *DeleteContentPackageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete content package o k response
func (o *DeleteContentPackageOK) Code() int {
	return 200
}

func (o *DeleteContentPackageOK) Error() string {
	return fmt.Sprintf("[DELETE /content/packages/{packageName}][%d] deleteContentPackageOK ", 200)
}

func (o *DeleteContentPackageOK) String() string {
	return fmt.Sprintf("[DELETE /content/packages/{packageName}][%d] deleteContentPackageOK ", 200)
}

func (o *DeleteContentPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteContentPackageUnauthorized creates a DeleteContentPackageUnauthorized with default headers values
func NewDeleteContentPackageUnauthorized() *DeleteContentPackageUnauthorized {
	return &DeleteContentPackageUnauthorized{}
}

/*
DeleteContentPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type DeleteContentPackageUnauthorized struct {
}

// IsSuccess returns true when this delete content package unauthorized response has a 2xx status code
func (o *DeleteContentPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete content package unauthorized response has a 3xx status code
func (o *DeleteContentPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete content package unauthorized response has a 4xx status code
func (o *DeleteContentPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete content package unauthorized response has a 5xx status code
func (o *DeleteContentPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete content package unauthorized response a status code equal to that given
func (o *DeleteContentPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete content package unauthorized response
func (o *DeleteContentPackageUnauthorized) Code() int {
	return 401
}

func (o *DeleteContentPackageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /content/packages/{packageName}][%d] deleteContentPackageUnauthorized ", 401)
}

func (o *DeleteContentPackageUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /content/packages/{packageName}][%d] deleteContentPackageUnauthorized ", 401)
}

func (o *DeleteContentPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteContentPackageNotFound creates a DeleteContentPackageNotFound with default headers values
func NewDeleteContentPackageNotFound() *DeleteContentPackageNotFound {
	return &DeleteContentPackageNotFound{}
}

/*
DeleteContentPackageNotFound describes a response with status code 404, with default header values.

Cannot find a package with the specified name.
*/
type DeleteContentPackageNotFound struct {
}

// IsSuccess returns true when this delete content package not found response has a 2xx status code
func (o *DeleteContentPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete content package not found response has a 3xx status code
func (o *DeleteContentPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete content package not found response has a 4xx status code
func (o *DeleteContentPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete content package not found response has a 5xx status code
func (o *DeleteContentPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete content package not found response a status code equal to that given
func (o *DeleteContentPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete content package not found response
func (o *DeleteContentPackageNotFound) Code() int {
	return 404
}

func (o *DeleteContentPackageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /content/packages/{packageName}][%d] deleteContentPackageNotFound ", 404)
}

func (o *DeleteContentPackageNotFound) String() string {
	return fmt.Sprintf("[DELETE /content/packages/{packageName}][%d] deleteContentPackageNotFound ", 404)
}

func (o *DeleteContentPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
