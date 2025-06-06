// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetServerUserListParams creates a new GetServerUserListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerUserListParams() *GetServerUserListParams {
	return &GetServerUserListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerUserListParamsWithTimeout creates a new GetServerUserListParams object
// with the ability to set a timeout on a request.
func NewGetServerUserListParamsWithTimeout(timeout time.Duration) *GetServerUserListParams {
	return &GetServerUserListParams{
		timeout: timeout,
	}
}

// NewGetServerUserListParamsWithContext creates a new GetServerUserListParams object
// with the ability to set a context for a request.
func NewGetServerUserListParamsWithContext(ctx context.Context) *GetServerUserListParams {
	return &GetServerUserListParams{
		Context: ctx,
	}
}

// NewGetServerUserListParamsWithHTTPClient creates a new GetServerUserListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerUserListParamsWithHTTPClient(client *http.Client) *GetServerUserListParams {
	return &GetServerUserListParams{
		HTTPClient: client,
	}
}

/*
GetServerUserListParams contains all the parameters to send to the API endpoint

	for the get server user list operation.

	Typically these are written to a http.Request.
*/
type GetServerUserListParams struct {

	// Protocol.
	Protocol string

	// SecretKey.
	SecretKey string

	// ServerID.
	//
	// Format: int64
	ServerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server user list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerUserListParams) WithDefaults() *GetServerUserListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server user list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerUserListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get server user list params
func (o *GetServerUserListParams) WithTimeout(timeout time.Duration) *GetServerUserListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server user list params
func (o *GetServerUserListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server user list params
func (o *GetServerUserListParams) WithContext(ctx context.Context) *GetServerUserListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server user list params
func (o *GetServerUserListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server user list params
func (o *GetServerUserListParams) WithHTTPClient(client *http.Client) *GetServerUserListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server user list params
func (o *GetServerUserListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProtocol adds the protocol to the get server user list params
func (o *GetServerUserListParams) WithProtocol(protocol string) *GetServerUserListParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the get server user list params
func (o *GetServerUserListParams) SetProtocol(protocol string) {
	o.Protocol = protocol
}

// WithSecretKey adds the secretKey to the get server user list params
func (o *GetServerUserListParams) WithSecretKey(secretKey string) *GetServerUserListParams {
	o.SetSecretKey(secretKey)
	return o
}

// SetSecretKey adds the secretKey to the get server user list params
func (o *GetServerUserListParams) SetSecretKey(secretKey string) {
	o.SecretKey = secretKey
}

// WithServerID adds the serverID to the get server user list params
func (o *GetServerUserListParams) WithServerID(serverID int64) *GetServerUserListParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the get server user list params
func (o *GetServerUserListParams) SetServerID(serverID int64) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerUserListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param protocol
	qrProtocol := o.Protocol
	qProtocol := qrProtocol
	if qProtocol != "" {

		if err := r.SetQueryParam("protocol", qProtocol); err != nil {
			return err
		}
	}

	// query param secret_key
	qrSecretKey := o.SecretKey
	qSecretKey := qrSecretKey
	if qSecretKey != "" {

		if err := r.SetQueryParam("secret_key", qSecretKey); err != nil {
			return err
		}
	}

	// query param server_id
	qrServerID := o.ServerID
	qServerID := swag.FormatInt64(qrServerID)
	if qServerID != "" {

		if err := r.SetQueryParam("server_id", qServerID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
