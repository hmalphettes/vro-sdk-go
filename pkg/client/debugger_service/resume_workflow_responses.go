// Code generated by go-swagger; DO NOT EDIT.

package debugger_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ResumeWorkflowReader is a Reader for the ResumeWorkflow structure.
type ResumeWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResumeWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResumeWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /debugger/{executionId}/resume] resumeWorkflow", response, response.Code())
	}
}

// NewResumeWorkflowOK creates a ResumeWorkflowOK with default headers values
func NewResumeWorkflowOK() *ResumeWorkflowOK {
	return &ResumeWorkflowOK{}
}

/*
ResumeWorkflowOK describes a response with status code 200, with default header values.

The request is successful
*/
type ResumeWorkflowOK struct {
}

// IsSuccess returns true when this resume workflow o k response has a 2xx status code
func (o *ResumeWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resume workflow o k response has a 3xx status code
func (o *ResumeWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume workflow o k response has a 4xx status code
func (o *ResumeWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resume workflow o k response has a 5xx status code
func (o *ResumeWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resume workflow o k response a status code equal to that given
func (o *ResumeWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resume workflow o k response
func (o *ResumeWorkflowOK) Code() int {
	return 200
}

func (o *ResumeWorkflowOK) Error() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/resume][%d] resumeWorkflowOK ", 200)
}

func (o *ResumeWorkflowOK) String() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/resume][%d] resumeWorkflowOK ", 200)
}

func (o *ResumeWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeWorkflowUnauthorized creates a ResumeWorkflowUnauthorized with default headers values
func NewResumeWorkflowUnauthorized() *ResumeWorkflowUnauthorized {
	return &ResumeWorkflowUnauthorized{}
}

/*
ResumeWorkflowUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type ResumeWorkflowUnauthorized struct {
}

// IsSuccess returns true when this resume workflow unauthorized response has a 2xx status code
func (o *ResumeWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resume workflow unauthorized response has a 3xx status code
func (o *ResumeWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume workflow unauthorized response has a 4xx status code
func (o *ResumeWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this resume workflow unauthorized response has a 5xx status code
func (o *ResumeWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this resume workflow unauthorized response a status code equal to that given
func (o *ResumeWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the resume workflow unauthorized response
func (o *ResumeWorkflowUnauthorized) Code() int {
	return 401
}

func (o *ResumeWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/resume][%d] resumeWorkflowUnauthorized ", 401)
}

func (o *ResumeWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/resume][%d] resumeWorkflowUnauthorized ", 401)
}

func (o *ResumeWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeWorkflowNotFound creates a ResumeWorkflowNotFound with default headers values
func NewResumeWorkflowNotFound() *ResumeWorkflowNotFound {
	return &ResumeWorkflowNotFound{}
}

/*
ResumeWorkflowNotFound describes a response with status code 404, with default header values.

Cannot find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type ResumeWorkflowNotFound struct {
}

// IsSuccess returns true when this resume workflow not found response has a 2xx status code
func (o *ResumeWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resume workflow not found response has a 3xx status code
func (o *ResumeWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume workflow not found response has a 4xx status code
func (o *ResumeWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this resume workflow not found response has a 5xx status code
func (o *ResumeWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this resume workflow not found response a status code equal to that given
func (o *ResumeWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the resume workflow not found response
func (o *ResumeWorkflowNotFound) Code() int {
	return 404
}

func (o *ResumeWorkflowNotFound) Error() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/resume][%d] resumeWorkflowNotFound ", 404)
}

func (o *ResumeWorkflowNotFound) String() string {
	return fmt.Sprintf("[POST /debugger/{executionId}/resume][%d] resumeWorkflowNotFound ", 404)
}

func (o *ResumeWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
