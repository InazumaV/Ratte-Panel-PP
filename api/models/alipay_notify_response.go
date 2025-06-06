// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlipayNotifyResponse AlipayNotifyResponse
//
// swagger:model AlipayNotifyResponse
type AlipayNotifyResponse struct {

	// return code
	// Required: true
	ReturnCode *string `json:"return_code"`
}

// Validate validates this alipay notify response
func (m *AlipayNotifyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReturnCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlipayNotifyResponse) validateReturnCode(formats strfmt.Registry) error {

	if err := validate.Required("return_code", "body", m.ReturnCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alipay notify response based on context it is used
func (m *AlipayNotifyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlipayNotifyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlipayNotifyResponse) UnmarshalBinary(b []byte) error {
	var res AlipayNotifyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
