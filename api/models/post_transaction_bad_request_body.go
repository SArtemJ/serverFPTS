// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostTransactionBadRequestBody post transaction bad request body
// swagger:model postTransactionBadRequestBody
type PostTransactionBadRequestBody struct {
	Error400Data

	PostTransactionBadRequestBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostTransactionBadRequestBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Error400Data
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Error400Data = aO0

	// AO1
	var aO1 PostTransactionBadRequestBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PostTransactionBadRequestBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostTransactionBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PostTransactionBadRequestBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post transaction bad request body
func (m *PostTransactionBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Error400Data
	if err := m.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PostTransactionBadRequestBodyAllOf1
	if err := m.PostTransactionBadRequestBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostTransactionBadRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTransactionBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostTransactionBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}