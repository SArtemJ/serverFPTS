// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SuccessData success data
// swagger:model Success_data
type SuccessData struct {
	SuccessDataAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuccessData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SuccessDataAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SuccessDataAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuccessData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SuccessDataAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this success data
func (m *SuccessData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuccessDataAllOf0
	if err := m.SuccessDataAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessData) UnmarshalBinary(b []byte) error {
	var res SuccessData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
