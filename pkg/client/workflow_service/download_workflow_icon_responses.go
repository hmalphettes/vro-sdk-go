// Code generated by go-swagger; DO NOT EDIT.

package workflow_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DownloadWorkflowIconReader is a Reader for the DownloadWorkflowIcon structure.
type DownloadWorkflowIconReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadWorkflowIconReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadWorkflowIconOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadWorkflowIconUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadWorkflowIconNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{id}/icon] downloadWorkflowIcon", response, response.Code())
	}
}

// NewDownloadWorkflowIconOK creates a DownloadWorkflowIconOK with default headers values
func NewDownloadWorkflowIconOK() *DownloadWorkflowIconOK {
	return &DownloadWorkflowIconOK{}
}

/*
DownloadWorkflowIconOK describes a response with status code 200, with default header values.

The request is successful
*/
type DownloadWorkflowIconOK struct {
}

// IsSuccess returns true when this download workflow icon o k response has a 2xx status code
func (o *DownloadWorkflowIconOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download workflow icon o k response has a 3xx status code
func (o *DownloadWorkflowIconOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download workflow icon o k response has a 4xx status code
func (o *DownloadWorkflowIconOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download workflow icon o k response has a 5xx status code
func (o *DownloadWorkflowIconOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download workflow icon o k response a status code equal to that given
func (o *DownloadWorkflowIconOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download workflow icon o k response
func (o *DownloadWorkflowIconOK) Code() int {
	return 200
}

func (o *DownloadWorkflowIconOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{id}/icon][%d] downloadWorkflowIconOK ", 200)
}

func (o *DownloadWorkflowIconOK) String() string {
	return fmt.Sprintf("[GET /workflows/{id}/icon][%d] downloadWorkflowIconOK ", 200)
}

func (o *DownloadWorkflowIconOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadWorkflowIconUnauthorized creates a DownloadWorkflowIconUnauthorized with default headers values
func NewDownloadWorkflowIconUnauthorized() *DownloadWorkflowIconUnauthorized {
	return &DownloadWorkflowIconUnauthorized{}
}

/*
DownloadWorkflowIconUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type DownloadWorkflowIconUnauthorized struct {
}

// IsSuccess returns true when this download workflow icon unauthorized response has a 2xx status code
func (o *DownloadWorkflowIconUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download workflow icon unauthorized response has a 3xx status code
func (o *DownloadWorkflowIconUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download workflow icon unauthorized response has a 4xx status code
func (o *DownloadWorkflowIconUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this download workflow icon unauthorized response has a 5xx status code
func (o *DownloadWorkflowIconUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this download workflow icon unauthorized response a status code equal to that given
func (o *DownloadWorkflowIconUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the download workflow icon unauthorized response
func (o *DownloadWorkflowIconUnauthorized) Code() int {
	return 401
}

func (o *DownloadWorkflowIconUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{id}/icon][%d] downloadWorkflowIconUnauthorized ", 401)
}

func (o *DownloadWorkflowIconUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{id}/icon][%d] downloadWorkflowIconUnauthorized ", 401)
}

func (o *DownloadWorkflowIconUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadWorkflowIconNotFound creates a DownloadWorkflowIconNotFound with default headers values
func NewDownloadWorkflowIconNotFound() *DownloadWorkflowIconNotFound {
	return &DownloadWorkflowIconNotFound{}
}

/*
DownloadWorkflowIconNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type DownloadWorkflowIconNotFound struct {
}

// IsSuccess returns true when this download workflow icon not found response has a 2xx status code
func (o *DownloadWorkflowIconNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download workflow icon not found response has a 3xx status code
func (o *DownloadWorkflowIconNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download workflow icon not found response has a 4xx status code
func (o *DownloadWorkflowIconNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this download workflow icon not found response has a 5xx status code
func (o *DownloadWorkflowIconNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this download workflow icon not found response a status code equal to that given
func (o *DownloadWorkflowIconNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the download workflow icon not found response
func (o *DownloadWorkflowIconNotFound) Code() int {
	return 404
}

func (o *DownloadWorkflowIconNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{id}/icon][%d] downloadWorkflowIconNotFound ", 404)
}

func (o *DownloadWorkflowIconNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{id}/icon][%d] downloadWorkflowIconNotFound ", 404)
}

func (o *DownloadWorkflowIconNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}