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

// RechargeOrderResponse RechargeOrderResponse
//
// swagger:model RechargeOrderResponse
type RechargeOrderResponse struct {

	// order no
	// Required: true
	OrderNo *string `json:"order_no"`
}

// Validate validates this recharge order response
func (m *RechargeOrderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RechargeOrderResponse) validateOrderNo(formats strfmt.Registry) error {

	if err := validate.Required("order_no", "body", m.OrderNo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recharge order response based on context it is used
func (m *RechargeOrderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RechargeOrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RechargeOrderResponse) UnmarshalBinary(b []byte) error {
	var res RechargeOrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
