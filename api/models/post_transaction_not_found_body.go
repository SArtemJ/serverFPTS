// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostTransactionNotFoundBody post transaction not found body
// swagger:model postTransactionNotFoundBody
type PostTransactionNotFoundBody struct {
	ErrorData
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostTransactionNotFoundBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ErrorData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ErrorData = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostTransactionNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ErrorData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post transaction not found body
func (m *PostTransactionNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ErrorData
	if err := m.ErrorData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostTransactionNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTransactionNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostTransactionNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
