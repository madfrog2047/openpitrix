// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRetryTasksRequest openpitrix retry tasks request
// swagger:model openpitrixRetryTasksRequest
type OpenpitrixRetryTasksRequest struct {

	// ids of task to retry
	TaskID []string `json:"task_id"`
}

// Validate validates this openpitrix retry tasks request
func (m *OpenpitrixRetryTasksRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaskID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixRetryTasksRequest) validateTaskID(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRetryTasksRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRetryTasksRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRetryTasksRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
