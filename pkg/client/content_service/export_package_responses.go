// Code generated by go-swagger; DO NOT EDIT.

package content_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExportPackageReader is a Reader for the ExportPackage structure.
type ExportPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /content/packages/{packageName}] exportPackage", response, response.Code())
	}
}

// NewExportPackageOK creates a ExportPackageOK with default headers values
func NewExportPackageOK() *ExportPackageOK {
	return &ExportPackageOK{}
}

/*
ExportPackageOK describes a response with status code 200, with default header values.

The request is successful
*/
type ExportPackageOK struct {
}

// IsSuccess returns true when this export package o k response has a 2xx status code
func (o *ExportPackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export package o k response has a 3xx status code
func (o *ExportPackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export package o k response has a 4xx status code
func (o *ExportPackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export package o k response has a 5xx status code
func (o *ExportPackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export package o k response a status code equal to that given
func (o *ExportPackageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export package o k response
func (o *ExportPackageOK) Code() int {
	return 200
}

func (o *ExportPackageOK) Error() string {
	return fmt.Sprintf("[GET /content/packages/{packageName}][%d] exportPackageOK ", 200)
}

func (o *ExportPackageOK) String() string {
	return fmt.Sprintf("[GET /content/packages/{packageName}][%d] exportPackageOK ", 200)
}

func (o *ExportPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportPackageUnauthorized creates a ExportPackageUnauthorized with default headers values
func NewExportPackageUnauthorized() *ExportPackageUnauthorized {
	return &ExportPackageUnauthorized{}
}

/*
ExportPackageUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type ExportPackageUnauthorized struct {
}

// IsSuccess returns true when this export package unauthorized response has a 2xx status code
func (o *ExportPackageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export package unauthorized response has a 3xx status code
func (o *ExportPackageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export package unauthorized response has a 4xx status code
func (o *ExportPackageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export package unauthorized response has a 5xx status code
func (o *ExportPackageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export package unauthorized response a status code equal to that given
func (o *ExportPackageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the export package unauthorized response
func (o *ExportPackageUnauthorized) Code() int {
	return 401
}

func (o *ExportPackageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /content/packages/{packageName}][%d] exportPackageUnauthorized ", 401)
}

func (o *ExportPackageUnauthorized) String() string {
	return fmt.Sprintf("[GET /content/packages/{packageName}][%d] exportPackageUnauthorized ", 401)
}

func (o *ExportPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportPackageNotFound creates a ExportPackageNotFound with default headers values
func NewExportPackageNotFound() *ExportPackageNotFound {
	return &ExportPackageNotFound{}
}

/*
ExportPackageNotFound describes a response with status code 404, with default header values.

Cannot find a package with the specified name.
*/
type ExportPackageNotFound struct {
}

// IsSuccess returns true when this export package not found response has a 2xx status code
func (o *ExportPackageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export package not found response has a 3xx status code
func (o *ExportPackageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export package not found response has a 4xx status code
func (o *ExportPackageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this export package not found response has a 5xx status code
func (o *ExportPackageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this export package not found response a status code equal to that given
func (o *ExportPackageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the export package not found response
func (o *ExportPackageNotFound) Code() int {
	return 404
}

func (o *ExportPackageNotFound) Error() string {
	return fmt.Sprintf("[GET /content/packages/{packageName}][%d] exportPackageNotFound ", 404)
}

func (o *ExportPackageNotFound) String() string {
	return fmt.Sprintf("[GET /content/packages/{packageName}][%d] exportPackageNotFound ", 404)
}

func (o *ExportPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
