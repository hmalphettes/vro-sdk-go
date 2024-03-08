// Code generated by go-swagger; DO NOT EDIT.

package configuration_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetConfigurationPermissionRuleReader is a Reader for the GetConfigurationPermissionRule structure.
type GetConfigurationPermissionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationPermissionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationPermissionRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConfigurationPermissionRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigurationPermissionRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /configurations/{id}/permissions/{ruleId}] getConfigurationPermissionRule", response, response.Code())
	}
}

// NewGetConfigurationPermissionRuleOK creates a GetConfigurationPermissionRuleOK with default headers values
func NewGetConfigurationPermissionRuleOK() *GetConfigurationPermissionRuleOK {
	return &GetConfigurationPermissionRuleOK{}
}

/*
GetConfigurationPermissionRuleOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetConfigurationPermissionRuleOK struct {
	Payload *models.Permission
}

// IsSuccess returns true when this get configuration permission rule o k response has a 2xx status code
func (o *GetConfigurationPermissionRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration permission rule o k response has a 3xx status code
func (o *GetConfigurationPermissionRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration permission rule o k response has a 4xx status code
func (o *GetConfigurationPermissionRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration permission rule o k response has a 5xx status code
func (o *GetConfigurationPermissionRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration permission rule o k response a status code equal to that given
func (o *GetConfigurationPermissionRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get configuration permission rule o k response
func (o *GetConfigurationPermissionRuleOK) Code() int {
	return 200
}

func (o *GetConfigurationPermissionRuleOK) Error() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions/{ruleId}][%d] getConfigurationPermissionRuleOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationPermissionRuleOK) String() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions/{ruleId}][%d] getConfigurationPermissionRuleOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationPermissionRuleOK) GetPayload() *models.Permission {
	return o.Payload
}

func (o *GetConfigurationPermissionRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationPermissionRuleUnauthorized creates a GetConfigurationPermissionRuleUnauthorized with default headers values
func NewGetConfigurationPermissionRuleUnauthorized() *GetConfigurationPermissionRuleUnauthorized {
	return &GetConfigurationPermissionRuleUnauthorized{}
}

/*
GetConfigurationPermissionRuleUnauthorized describes a response with status code 401, with default header values.

User is not authorized
*/
type GetConfigurationPermissionRuleUnauthorized struct {
}

// IsSuccess returns true when this get configuration permission rule unauthorized response has a 2xx status code
func (o *GetConfigurationPermissionRuleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration permission rule unauthorized response has a 3xx status code
func (o *GetConfigurationPermissionRuleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration permission rule unauthorized response has a 4xx status code
func (o *GetConfigurationPermissionRuleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration permission rule unauthorized response has a 5xx status code
func (o *GetConfigurationPermissionRuleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration permission rule unauthorized response a status code equal to that given
func (o *GetConfigurationPermissionRuleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get configuration permission rule unauthorized response
func (o *GetConfigurationPermissionRuleUnauthorized) Code() int {
	return 401
}

func (o *GetConfigurationPermissionRuleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions/{ruleId}][%d] getConfigurationPermissionRuleUnauthorized ", 401)
}

func (o *GetConfigurationPermissionRuleUnauthorized) String() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions/{ruleId}][%d] getConfigurationPermissionRuleUnauthorized ", 401)
}

func (o *GetConfigurationPermissionRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationPermissionRuleNotFound creates a GetConfigurationPermissionRuleNotFound with default headers values
func NewGetConfigurationPermissionRuleNotFound() *GetConfigurationPermissionRuleNotFound {
	return &GetConfigurationPermissionRuleNotFound{}
}

/*
GetConfigurationPermissionRuleNotFound describes a response with status code 404, with default header values.

Cannot find a configuration with the specified ID, the user does not have 'read' access rights for that configuration, or the permission rule with the specified ID does not exist
*/
type GetConfigurationPermissionRuleNotFound struct {
}

// IsSuccess returns true when this get configuration permission rule not found response has a 2xx status code
func (o *GetConfigurationPermissionRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration permission rule not found response has a 3xx status code
func (o *GetConfigurationPermissionRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration permission rule not found response has a 4xx status code
func (o *GetConfigurationPermissionRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration permission rule not found response has a 5xx status code
func (o *GetConfigurationPermissionRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration permission rule not found response a status code equal to that given
func (o *GetConfigurationPermissionRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get configuration permission rule not found response
func (o *GetConfigurationPermissionRuleNotFound) Code() int {
	return 404
}

func (o *GetConfigurationPermissionRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions/{ruleId}][%d] getConfigurationPermissionRuleNotFound ", 404)
}

func (o *GetConfigurationPermissionRuleNotFound) String() string {
	return fmt.Sprintf("[GET /configurations/{id}/permissions/{ruleId}][%d] getConfigurationPermissionRuleNotFound ", 404)
}

func (o *GetConfigurationPermissionRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
