// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PackageDetailsReader is a Reader for the PackageDetails structure.
type PackageDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackageDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackageDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPackageDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPackageDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /packages/{packageName}] packageDetails", response, response.Code())
	}
}

// NewPackageDetailsOK creates a PackageDetailsOK with default headers values
func NewPackageDetailsOK() *PackageDetailsOK {
	return &PackageDetailsOK{}
}

/*
PackageDetailsOK describes a response with status code 200, with default header values.

The request is successful
*/
type PackageDetailsOK struct {
	Payload *models.PackageDetails
}

// IsSuccess returns true when this package details o k response has a 2xx status code
func (o *PackageDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this package details o k response has a 3xx status code
func (o *PackageDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this package details o k response has a 4xx status code
func (o *PackageDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this package details o k response has a 5xx status code
func (o *PackageDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this package details o k response a status code equal to that given
func (o *PackageDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the package details o k response
func (o *PackageDetailsOK) Code() int {
	return 200
}

func (o *PackageDetailsOK) Error() string {
	return fmt.Sprintf("[GET /packages/{packageName}][%d] packageDetailsOK  %+v", 200, o.Payload)
}

func (o *PackageDetailsOK) String() string {
	return fmt.Sprintf("[GET /packages/{packageName}][%d] packageDetailsOK  %+v", 200, o.Payload)
}

func (o *PackageDetailsOK) GetPayload() *models.PackageDetails {
	return o.Payload
}

func (o *PackageDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackageDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackageDetailsUnauthorized creates a PackageDetailsUnauthorized with default headers values
func NewPackageDetailsUnauthorized() *PackageDetailsUnauthorized {
	return &PackageDetailsUnauthorized{}
}

/*
PackageDetailsUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type PackageDetailsUnauthorized struct {
}

// IsSuccess returns true when this package details unauthorized response has a 2xx status code
func (o *PackageDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this package details unauthorized response has a 3xx status code
func (o *PackageDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this package details unauthorized response has a 4xx status code
func (o *PackageDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this package details unauthorized response has a 5xx status code
func (o *PackageDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this package details unauthorized response a status code equal to that given
func (o *PackageDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the package details unauthorized response
func (o *PackageDetailsUnauthorized) Code() int {
	return 401
}

func (o *PackageDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /packages/{packageName}][%d] packageDetailsUnauthorized ", 401)
}

func (o *PackageDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /packages/{packageName}][%d] packageDetailsUnauthorized ", 401)
}

func (o *PackageDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPackageDetailsNotFound creates a PackageDetailsNotFound with default headers values
func NewPackageDetailsNotFound() *PackageDetailsNotFound {
	return &PackageDetailsNotFound{}
}

/*
PackageDetailsNotFound describes a response with status code 404, with default header values.

Cannot find a package with the specified name
*/
type PackageDetailsNotFound struct {
}

// IsSuccess returns true when this package details not found response has a 2xx status code
func (o *PackageDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this package details not found response has a 3xx status code
func (o *PackageDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this package details not found response has a 4xx status code
func (o *PackageDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this package details not found response has a 5xx status code
func (o *PackageDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this package details not found response a status code equal to that given
func (o *PackageDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the package details not found response
func (o *PackageDetailsNotFound) Code() int {
	return 404
}

func (o *PackageDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/{packageName}][%d] packageDetailsNotFound ", 404)
}

func (o *PackageDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /packages/{packageName}][%d] packageDetailsNotFound ", 404)
}

func (o *PackageDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}