// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new server API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new server API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetServerConfig(params *GetServerConfigParams, opts ...ClientOption) (*GetServerConfigOK, error)

	GetServerUserList(params *GetServerUserListParams, opts ...ClientOption) (*GetServerUserListOK, error)

	PushOnlineUsers(params *PushOnlineUsersParams, opts ...ClientOption) (*PushOnlineUsersOK, error)

	ServerPushStatus(params *ServerPushStatusParams, opts ...ClientOption) (*ServerPushStatusOK, error)

	ServerPushUserTraffic(params *ServerPushUserTrafficParams, opts ...ClientOption) (*ServerPushUserTrafficOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetServerConfig gets server config
*/
func (a *Client) GetServerConfig(params *GetServerConfigParams, opts ...ClientOption) (*GetServerConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetServerConfig",
		Method:             "GET",
		PathPattern:        "/v1/server/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServerConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetServerConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServerUserList gets user list
*/
func (a *Client) GetServerUserList(params *GetServerUserListParams, opts ...ClientOption) (*GetServerUserListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerUserListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetServerUserList",
		Method:             "GET",
		PathPattern:        "/v1/server/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServerUserListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerUserListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetServerUserList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PushOnlineUsers pushes online users
*/
func (a *Client) PushOnlineUsers(params *PushOnlineUsersParams, opts ...ClientOption) (*PushOnlineUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPushOnlineUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PushOnlineUsers",
		Method:             "POST",
		PathPattern:        "/v1/server/online",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PushOnlineUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushOnlineUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PushOnlineUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServerPushStatus pushes server status
*/
func (a *Client) ServerPushStatus(params *ServerPushStatusParams, opts ...ClientOption) (*ServerPushStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerPushStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerPushStatus",
		Method:             "POST",
		PathPattern:        "/v1/server/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerPushStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerPushStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ServerPushStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServerPushUserTraffic pushes user traffic
*/
func (a *Client) ServerPushUserTraffic(params *ServerPushUserTrafficParams, opts ...ClientOption) (*ServerPushUserTrafficOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerPushUserTrafficParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerPushUserTraffic",
		Method:             "POST",
		PathPattern:        "/v1/server/push",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerPushUserTrafficReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerPushUserTrafficOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ServerPushUserTraffic: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
