// Code generated by go-swagger; DO NOT EDIT.

package packages_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetImportPackageDetailsReader is a Reader for the GetImportPackageDetails structure.
type GetImportPackageDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImportPackageDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImportPackageDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetImportPackageDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /packages/import-details] getImportPackageDetails", response, response.Code())
	}
}

// NewGetImportPackageDetailsOK creates a GetImportPackageDetailsOK with default headers values
func NewGetImportPackageDetailsOK() *GetImportPackageDetailsOK {
	return &GetImportPackageDetailsOK{}
}

/*
GetImportPackageDetailsOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetImportPackageDetailsOK struct {
	Payload *models.ImportPackageDetails
}

// IsSuccess returns true when this get import package details o k response has a 2xx status code
func (o *GetImportPackageDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get import package details o k response has a 3xx status code
func (o *GetImportPackageDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get import package details o k response has a 4xx status code
func (o *GetImportPackageDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get import package details o k response has a 5xx status code
func (o *GetImportPackageDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get import package details o k response a status code equal to that given
func (o *GetImportPackageDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get import package details o k response
func (o *GetImportPackageDetailsOK) Code() int {
	return 200
}

func (o *GetImportPackageDetailsOK) Error() string {
	return fmt.Sprintf("[POST /packages/import-details][%d] getImportPackageDetailsOK  %+v", 200, o.Payload)
}

func (o *GetImportPackageDetailsOK) String() string {
	return fmt.Sprintf("[POST /packages/import-details][%d] getImportPackageDetailsOK  %+v", 200, o.Payload)
}

func (o *GetImportPackageDetailsOK) GetPayload() *models.ImportPackageDetails {
	return o.Payload
}

func (o *GetImportPackageDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportPackageDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImportPackageDetailsUnauthorized creates a GetImportPackageDetailsUnauthorized with default headers values
func NewGetImportPackageDetailsUnauthorized() *GetImportPackageDetailsUnauthorized {
	return &GetImportPackageDetailsUnauthorized{}
}

/*
GetImportPackageDetailsUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type GetImportPackageDetailsUnauthorized struct {
}

// IsSuccess returns true when this get import package details unauthorized response has a 2xx status code
func (o *GetImportPackageDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get import package details unauthorized response has a 3xx status code
func (o *GetImportPackageDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get import package details unauthorized response has a 4xx status code
func (o *GetImportPackageDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get import package details unauthorized response has a 5xx status code
func (o *GetImportPackageDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get import package details unauthorized response a status code equal to that given
func (o *GetImportPackageDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get import package details unauthorized response
func (o *GetImportPackageDetailsUnauthorized) Code() int {
	return 401
}

func (o *GetImportPackageDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /packages/import-details][%d] getImportPackageDetailsUnauthorized ", 401)
}

func (o *GetImportPackageDetailsUnauthorized) String() string {
	return fmt.Sprintf("[POST /packages/import-details][%d] getImportPackageDetailsUnauthorized ", 401)
}

func (o *GetImportPackageDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
