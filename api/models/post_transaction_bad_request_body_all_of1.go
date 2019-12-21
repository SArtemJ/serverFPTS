// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostTransactionBadRequestBodyAllOf1 post transaction bad request body all of1
// swagger:model postTransactionBadRequestBodyAllOf1
type PostTransactionBadRequestBodyAllOf1 struct {

	// errors
	// Required: true
	Errors *PostTransactionBadRequestBodyAllOf1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this post transaction bad request body all of1
func (m *PostTransactionBadRequestBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTransactionBadRequestBodyAllOf1) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	if m.Errors != nil {
		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errors")
			}
			return err
		}
	}

	return nil
}

func (m *PostTransactionBadRequestBodyAllOf1) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostTransactionBadRequestBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTransactionBadRequestBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res PostTransactionBadRequestBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
