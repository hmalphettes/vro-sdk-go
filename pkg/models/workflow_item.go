// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowItem workflow item
//
// swagger:model WorkflowItem
type WorkflowItem struct {

	// alt out name
	AltOutName string `json:"alt-out-name,omitempty" xml:"alt-out-name,attr,omitempty"`

	// business status
	BusinessStatus string `json:"business-status,omitempty" xml:"business-status,attr,omitempty"`

	// catch name
	CatchName string `json:"catch-name,omitempty" xml:"catch-name,attr,omitempty"`

	// color
	Color string `json:"color,omitempty" xml:"color,attr,omitempty"`

	// comparator
	Comparator int32 `json:"comparator,omitempty" xml:"comparator,attr,omitempty"`

	// condition
	Condition []*Condition `json:"condition"`

	// conditions
	Conditions *Conditions `json:"conditions,omitempty"`

	// content mode
	ContentMode string `json:"content-mode,omitempty" xml:"content-mode,attr,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display-name,omitempty"`

	// end mode
	EndMode string `json:"end-mode,omitempty" xml:"end-mode,attr,omitempty"`

	// in binding
	InBinding *InBinding `json:"in-binding,omitempty"`

	// interaction
	Interaction string `json:"interaction,omitempty" xml:"interaction,attr,omitempty"`

	// iteration catch block
	IterationCatchBlock *IterationCatchBlock `json:"iteration-catch-block,omitempty"`

	// launched workflow id
	LaunchedWorkflowID string `json:"launched-workflow-id,omitempty" xml:"launched-workflow-id,attr,omitempty"`

	// linked workflow id
	LinkedWorkflowID string `json:"linked-workflow-id,omitempty" xml:"linked-workflow-id,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// out binding
	OutBinding *OutBinding `json:"out-binding,omitempty"`

	// out name
	OutName string `json:"out-name,omitempty" xml:"out-name,attr,omitempty"`

	// position
	// Required: true
	Position *Position `json:"position"`

	// presentation
	Presentation *Presentation `json:"presentation,omitempty"`

	// prototype id
	PrototypeID string `json:"prototype-id,omitempty" xml:"prototype-id,attr,omitempty"`

	// reference
	Reference *Reference `json:"reference,omitempty"`

	// script
	Script *Script `json:"script,omitempty"`

	// script module
	ScriptModule string `json:"script-module,omitempty" xml:"script-module,attr,omitempty"`

	// throw bind name
	ThrowBindName string `json:"throw-bind-name,omitempty" xml:"throw-bind-name,attr,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`

	// workflow subelements list
	WorkflowSubelementsList *WorkflowSubelementsList `json:"workflow-subelements-list,omitempty"`
}

// Validate validates this workflow item
func (m *WorkflowItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInBinding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIterationCatchBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutBinding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresentation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowSubelementsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowItem) validateCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	for i := 0; i < len(m.Condition); i++ {
		if swag.IsZero(m.Condition[i]) { // not required
			continue
		}

		if m.Condition[i] != nil {
			if err := m.Condition[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("condition" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("condition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowItem) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validateInBinding(formats strfmt.Registry) error {
	if swag.IsZero(m.InBinding) { // not required
		return nil
	}

	if m.InBinding != nil {
		if err := m.InBinding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("in-binding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("in-binding")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validateIterationCatchBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.IterationCatchBlock) { // not required
		return nil
	}

	if m.IterationCatchBlock != nil {
		if err := m.IterationCatchBlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iteration-catch-block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iteration-catch-block")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validateOutBinding(formats strfmt.Registry) error {
	if swag.IsZero(m.OutBinding) { // not required
		return nil
	}

	if m.OutBinding != nil {
		if err := m.OutBinding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("out-binding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("out-binding")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validatePresentation(formats strfmt.Registry) error {
	if swag.IsZero(m.Presentation) { // not required
		return nil
	}

	if m.Presentation != nil {
		if err := m.Presentation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("presentation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("presentation")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validateReference(formats strfmt.Registry) error {
	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if m.Reference != nil {
		if err := m.Reference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reference")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validateScript(formats strfmt.Registry) error {
	if swag.IsZero(m.Script) { // not required
		return nil
	}

	if m.Script != nil {
		if err := m.Script.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) validateWorkflowSubelementsList(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowSubelementsList) { // not required
		return nil
	}

	if m.WorkflowSubelementsList != nil {
		if err := m.WorkflowSubelementsList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow-subelements-list")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflow-subelements-list")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this workflow item based on the context it is used
func (m *WorkflowItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInBinding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIterationCatchBlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutBinding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePresentation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkflowSubelementsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowItem) contextValidateCondition(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Condition); i++ {

		if m.Condition[i] != nil {

			if swag.IsZero(m.Condition[i]) { // not required
				return nil
			}

			if err := m.Condition[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("condition" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("condition" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowItem) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	if m.Conditions != nil {

		if swag.IsZero(m.Conditions) { // not required
			return nil
		}

		if err := m.Conditions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidateInBinding(ctx context.Context, formats strfmt.Registry) error {

	if m.InBinding != nil {

		if swag.IsZero(m.InBinding) { // not required
			return nil
		}

		if err := m.InBinding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("in-binding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("in-binding")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidateIterationCatchBlock(ctx context.Context, formats strfmt.Registry) error {

	if m.IterationCatchBlock != nil {

		if swag.IsZero(m.IterationCatchBlock) { // not required
			return nil
		}

		if err := m.IterationCatchBlock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iteration-catch-block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iteration-catch-block")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidateOutBinding(ctx context.Context, formats strfmt.Registry) error {

	if m.OutBinding != nil {

		if swag.IsZero(m.OutBinding) { // not required
			return nil
		}

		if err := m.OutBinding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("out-binding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("out-binding")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if m.Position != nil {

		if err := m.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidatePresentation(ctx context.Context, formats strfmt.Registry) error {

	if m.Presentation != nil {

		if swag.IsZero(m.Presentation) { // not required
			return nil
		}

		if err := m.Presentation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("presentation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("presentation")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidateReference(ctx context.Context, formats strfmt.Registry) error {

	if m.Reference != nil {

		if swag.IsZero(m.Reference) { // not required
			return nil
		}

		if err := m.Reference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reference")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

	if m.Script != nil {

		if swag.IsZero(m.Script) { // not required
			return nil
		}

		if err := m.Script.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowItem) contextValidateWorkflowSubelementsList(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowSubelementsList != nil {

		if swag.IsZero(m.WorkflowSubelementsList) { // not required
			return nil
		}

		if err := m.WorkflowSubelementsList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow-subelements-list")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflow-subelements-list")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowItem) UnmarshalBinary(b []byte) error {
	var res WorkflowItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}