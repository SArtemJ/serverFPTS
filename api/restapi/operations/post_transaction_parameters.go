// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SArtemJ/serverFPTS/api/models"
)

// NewPostTransactionParams creates a new PostTransactionParams object
// no default values defined in spec.
func NewPostTransactionParams() PostTransactionParams {

	return PostTransactionParams{}
}

// PostTransactionParams contains all the bound params for the post transaction operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostTransaction
type PostTransactionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Source-Type
	  Required: true
	  In: header
	*/
	SourceType string
	/*
	  Required: true
	  In: body
	*/
	Body *models.PostTransactionParamsBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostTransactionParams() beforehand.
func (o *PostTransactionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindSourceType(r.Header[http.CanonicalHeaderKey("Source-Type")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PostTransactionParamsBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSourceType binds and validates parameter SourceType from header.
func (o *PostTransactionParams) bindSourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Source-Type", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Source-Type", "header", raw); err != nil {
		return err
	}

	o.SourceType = raw

	return nil
}
