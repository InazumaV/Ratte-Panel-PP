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

// AuthMethodConfig AuthMethodConfig
//
// swagger:model AuthMethodConfig
type AuthMethodConfig struct {

	// config
	// Required: true
	Config interface{} `json:"config"`

	// enabled
	// Required: true
	Enabled bool `json:"enabled"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// method
	// Required: true
	Method *string `json:"method"`
}

// Validate validates this auth method config
func (m *AuthMethodConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthMethodConfig) validateConfig(formats strfmt.Registry) error {

	if m.Config == nil {
		return errors.Required("config", "body", nil)
	}

	return nil
}

func (m *AuthMethodConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", bool(m.Enabled)); err != nil {
		return err
	}

	return nil
}

func (m *AuthMethodConfig) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AuthMethodConfig) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this auth method config based on context it is used
func (m *AuthMethodConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthMethodConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthMethodConfig) UnmarshalBinary(b []byte) error {
	var res AuthMethodConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
