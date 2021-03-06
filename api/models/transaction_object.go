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

// TransactionObject Transaction_object
// swagger:model Transaction_object
type TransactionObject struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// state
	// Required: true
	State *string `json:"state"`

	// transactionID
	// Required: true
	// Format: uuid
	TransactionID *strfmt.UUID `json:"transactionID"`

	// userGUID
	// Required: true
	// Format: uuid
	UserGUID *strfmt.UUID `json:"userGUID"`
}

// Validate validates this transaction object
func (m *TransactionObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionObject) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TransactionObject) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *TransactionObject) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transactionID", "body", m.TransactionID); err != nil {
		return err
	}

	if err := validate.FormatOf("transactionID", "body", "uuid", m.TransactionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransactionObject) validateUserGUID(formats strfmt.Registry) error {

	if err := validate.Required("userGUID", "body", m.UserGUID); err != nil {
		return err
	}

	if err := validate.FormatOf("userGUID", "body", "uuid", m.UserGUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionObject) UnmarshalBinary(b []byte) error {
	var res TransactionObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
