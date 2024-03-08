// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskData task data
//
// swagger:model task-data
type TaskData struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// recurrence cycle
	// Enum: [ONE_TIME EVERY_MINUTES EVERY_HOURS EVERY_DAYS EVERY_WEEKS EVERY_MONTHS]
	RecurrenceCycle string `json:"recurrence-cycle,omitempty"`

	// recurrence end date
	// Format: date-time
	RecurrenceEndDate strfmt.DateTime `json:"recurrence-end-date,omitempty"`

	// recurrence pattern
	RecurrencePattern string `json:"recurrence-pattern,omitempty"`

	// recurrence start date
	// Format: date-time
	RecurrenceStartDate strfmt.DateTime `json:"recurrence-start-date,omitempty"`

	// start mode
	// Enum: [NORMAL START_IN_THE_PAST]
	StartMode string `json:"start-mode,omitempty"`

	// state
	// Enum: [FINISHED CANCELED ERROR PENDING RUNNING SUSPENDED UNKNOWN]
	State string `json:"state,omitempty"`
}

// Validate validates this task data
func (m *TaskData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecurrenceCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurrenceEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurrenceStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var taskDataTypeRecurrenceCyclePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONE_TIME","EVERY_MINUTES","EVERY_HOURS","EVERY_DAYS","EVERY_WEEKS","EVERY_MONTHS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskDataTypeRecurrenceCyclePropEnum = append(taskDataTypeRecurrenceCyclePropEnum, v)
	}
}

const (

	// TaskDataRecurrenceCycleONETIME captures enum value "ONE_TIME"
	TaskDataRecurrenceCycleONETIME string = "ONE_TIME"

	// TaskDataRecurrenceCycleEVERYMINUTES captures enum value "EVERY_MINUTES"
	TaskDataRecurrenceCycleEVERYMINUTES string = "EVERY_MINUTES"

	// TaskDataRecurrenceCycleEVERYHOURS captures enum value "EVERY_HOURS"
	TaskDataRecurrenceCycleEVERYHOURS string = "EVERY_HOURS"

	// TaskDataRecurrenceCycleEVERYDAYS captures enum value "EVERY_DAYS"
	TaskDataRecurrenceCycleEVERYDAYS string = "EVERY_DAYS"

	// TaskDataRecurrenceCycleEVERYWEEKS captures enum value "EVERY_WEEKS"
	TaskDataRecurrenceCycleEVERYWEEKS string = "EVERY_WEEKS"

	// TaskDataRecurrenceCycleEVERYMONTHS captures enum value "EVERY_MONTHS"
	TaskDataRecurrenceCycleEVERYMONTHS string = "EVERY_MONTHS"
)

// prop value enum
func (m *TaskData) validateRecurrenceCycleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taskDataTypeRecurrenceCyclePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaskData) validateRecurrenceCycle(formats strfmt.Registry) error {
	if swag.IsZero(m.RecurrenceCycle) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecurrenceCycleEnum("recurrence-cycle", "body", m.RecurrenceCycle); err != nil {
		return err
	}

	return nil
}

func (m *TaskData) validateRecurrenceEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RecurrenceEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("recurrence-end-date", "body", "date-time", m.RecurrenceEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskData) validateRecurrenceStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RecurrenceStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("recurrence-start-date", "body", "date-time", m.RecurrenceStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var taskDataTypeStartModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NORMAL","START_IN_THE_PAST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskDataTypeStartModePropEnum = append(taskDataTypeStartModePropEnum, v)
	}
}

const (

	// TaskDataStartModeNORMAL captures enum value "NORMAL"
	TaskDataStartModeNORMAL string = "NORMAL"

	// TaskDataStartModeSTARTINTHEPAST captures enum value "START_IN_THE_PAST"
	TaskDataStartModeSTARTINTHEPAST string = "START_IN_THE_PAST"
)

// prop value enum
func (m *TaskData) validateStartModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taskDataTypeStartModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaskData) validateStartMode(formats strfmt.Registry) error {
	if swag.IsZero(m.StartMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStartModeEnum("start-mode", "body", m.StartMode); err != nil {
		return err
	}

	return nil
}

var taskDataTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FINISHED","CANCELED","ERROR","PENDING","RUNNING","SUSPENDED","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskDataTypeStatePropEnum = append(taskDataTypeStatePropEnum, v)
	}
}

const (

	// TaskDataStateFINISHED captures enum value "FINISHED"
	TaskDataStateFINISHED string = "FINISHED"

	// TaskDataStateCANCELED captures enum value "CANCELED"
	TaskDataStateCANCELED string = "CANCELED"

	// TaskDataStateERROR captures enum value "ERROR"
	TaskDataStateERROR string = "ERROR"

	// TaskDataStatePENDING captures enum value "PENDING"
	TaskDataStatePENDING string = "PENDING"

	// TaskDataStateRUNNING captures enum value "RUNNING"
	TaskDataStateRUNNING string = "RUNNING"

	// TaskDataStateSUSPENDED captures enum value "SUSPENDED"
	TaskDataStateSUSPENDED string = "SUSPENDED"

	// TaskDataStateUNKNOWN captures enum value "UNKNOWN"
	TaskDataStateUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *TaskData) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taskDataTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaskData) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task data based on context it is used
func (m *TaskData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskData) UnmarshalBinary(b []byte) error {
	var res TaskData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}