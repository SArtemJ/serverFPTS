// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostTransactionMethodNotAllowedBody post transaction method not allowed body
// swagger:model postTransactionMethodNotAllowedBody
type PostTransactionMethodNotAllowedBody struct {
	Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostTransactionMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Error405Data
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Error405Data = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostTransactionMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post transaction method not allowed body
func (m *PostTransactionMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Error405Data
	if err := m.Error405Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostTransactionMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTransactionMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostTransactionMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
