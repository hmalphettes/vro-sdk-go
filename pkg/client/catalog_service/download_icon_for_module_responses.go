// Code generated by go-swagger; DO NOT EDIT.

package catalog_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DownloadIconForModuleReader is a Reader for the DownloadIconForModule structure.
type DownloadIconForModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadIconForModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadIconForModuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadIconForModuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /catalog/{namespace}/metadata/icon] downloadIconForModule", response, response.Code())
	}
}

// NewDownloadIconForModuleOK creates a DownloadIconForModuleOK with default headers values
func NewDownloadIconForModuleOK() *DownloadIconForModuleOK {
	return &DownloadIconForModuleOK{}
}

/*
DownloadIconForModuleOK describes a response with status code 200, with default header values.

The request is successful
*/
type DownloadIconForModuleOK struct {
}

// IsSuccess returns true when this download icon for module o k response has a 2xx status code
func (o *DownloadIconForModuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download icon for module o k response has a 3xx status code
func (o *DownloadIconForModuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download icon for module o k response has a 4xx status code
func (o *DownloadIconForModuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download icon for module o k response has a 5xx status code
func (o *DownloadIconForModuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download icon for module o k response a status code equal to that given
func (o *DownloadIconForModuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download icon for module o k response
func (o *DownloadIconForModuleOK) Code() int {
	return 200
}

func (o *DownloadIconForModuleOK) Error() string {
	return fmt.Sprintf("[GET /catalog/{namespace}/metadata/icon][%d] downloadIconForModuleOK ", 200)
}

func (o *DownloadIconForModuleOK) String() string {
	return fmt.Sprintf("[GET /catalog/{namespace}/metadata/icon][%d] downloadIconForModuleOK ", 200)
}

func (o *DownloadIconForModuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadIconForModuleUnauthorized creates a DownloadIconForModuleUnauthorized with default headers values
func NewDownloadIconForModuleUnauthorized() *DownloadIconForModuleUnauthorized {
	return &DownloadIconForModuleUnauthorized{}
}

/*
DownloadIconForModuleUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type DownloadIconForModuleUnauthorized struct {
}

// IsSuccess returns true when this download icon for module unauthorized response has a 2xx status code
func (o *DownloadIconForModuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download icon for module unauthorized response has a 3xx status code
func (o *DownloadIconForModuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download icon for module unauthorized response has a 4xx status code
func (o *DownloadIconForModuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this download icon for module unauthorized response has a 5xx status code
func (o *DownloadIconForModuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this download icon for module unauthorized response a status code equal to that given
func (o *DownloadIconForModuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the download icon for module unauthorized response
func (o *DownloadIconForModuleUnauthorized) Code() int {
	return 401
}

func (o *DownloadIconForModuleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/{namespace}/metadata/icon][%d] downloadIconForModuleUnauthorized ", 401)
}

func (o *DownloadIconForModuleUnauthorized) String() string {
	return fmt.Sprintf("[GET /catalog/{namespace}/metadata/icon][%d] downloadIconForModuleUnauthorized ", 401)
}

func (o *DownloadIconForModuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}