// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeVpcResponse openpitrix describe vpc response
// swagger:model openpitrixDescribeVpcResponse
type OpenpitrixDescribeVpcResponse struct {

	// vpc
	Vpc *OpenpitrixVpc `json:"vpc,omitempty"`
}

// Validate validates this openpitrix describe vpc response
func (m *OpenpitrixDescribeVpcResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVpc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDescribeVpcResponse) validateVpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {

		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeVpcResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeVpcResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeVpcResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}