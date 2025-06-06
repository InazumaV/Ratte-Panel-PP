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

// UserLoginLog UserLoginLog
//
// swagger:model UserLoginLog
type UserLoginLog struct {

	// created at
	// Required: true
	CreatedAt *int64 `json:"created_at"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// login ip
	// Required: true
	LoginIP *string `json:"login_ip"`

	// success
	// Required: true
	Success bool `json:"success"`

	// user agent
	// Required: true
	UserAgent *string `json:"user_agent"`

	// user id
	// Required: true
	UserID *int64 `json:"user_id"`
}

// Validate validates this user login log
func (m *UserLoginLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoginIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserLoginLog) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *UserLoginLog) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserLoginLog) validateLoginIP(formats strfmt.Registry) error {

	if err := validate.Required("login_ip", "body", m.LoginIP); err != nil {
		return err
	}

	return nil
}

func (m *UserLoginLog) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", bool(m.Success)); err != nil {
		return err
	}

	return nil
}

func (m *UserLoginLog) validateUserAgent(formats strfmt.Registry) error {

	if err := validate.Required("user_agent", "body", m.UserAgent); err != nil {
		return err
	}

	return nil
}

func (m *UserLoginLog) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user login log based on context it is used
func (m *UserLoginLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserLoginLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLoginLog) UnmarshalBinary(b []byte) error {
	var res UserLoginLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
