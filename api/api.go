package api

import (
	"Ratte-Panel-PP/api/client"
	"Ratte-Panel-PP/api/client/server"
	"Ratte-Panel-PP/api/models"
	"errors"
	"net/url"
	"time"
)

type Api struct {
	c         server.ClientService
	protocol  string
	secretKey string
	serverID  int64
	timeout   time.Duration
}

func New(
	baseUrl,
	key string,
	protocol string,
	serverID int64,
	timeout time.Duration,
) (Api, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return Api{}, err
	}
	cfg := client.DefaultTransportConfig()
	cfg.Host = baseUrl
	cfg.BasePath = u.Path
	cfg.Schemes = []string{u.Scheme}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	return Api{
		c:         c.Server,
		protocol:  protocol,
		secretKey: key,
		serverID:  serverID,
		timeout:   timeout,
	}, nil
}

func (a *Api) GetServerConfig() (*models.GetServerConfigResponse, error) {
	rsp, err := a.c.GetServerConfig(
		SetTimeout(&server.GetServerConfigParams{
			Protocol:  a.protocol,
			SecretKey: a.secretKey,
			ServerID:  a.serverID,
		},
			a.timeout))
	if err != nil {
		return nil, err
	}
	if !rsp.IsSuccess() {
		return nil, errors.New(rsp.Error())
	}
	return rsp.GetPayload().Data, nil
}

func (a *Api) GetServerUserList() (*models.GetServerUserListResponse, error) {
	rsp, err := a.c.GetServerUserList(
		SetTimeout(&server.GetServerUserListParams{
			Protocol:  a.protocol,
			SecretKey: a.secretKey,
			ServerID:  a.serverID,
		},
			a.timeout))
	if err != nil {
		return nil, err
	}
	if !rsp.IsSuccess() {
		return nil, errors.New(rsp.Error())
	}
	return rsp.GetPayload().Data, nil
}

func (a *Api) PushOnlineUsers(req *models.OnlineUsersRequest) error {
	rsp, err := a.c.PushOnlineUsers(
		SetTimeout(&server.PushOnlineUsersParams{
			Body:      req,
			Protocol:  a.protocol,
			SecretKey: a.secretKey,
			ServerID:  a.serverID,
		},
			a.timeout))
	if err != nil {
		return err
	}
	if !rsp.IsSuccess() {
		return errors.New(rsp.Error())
	}
	return nil
}

func (a *Api) ServerPushStatus(req *models.ServerPushStatusRequest) error {
	rsp, err := a.c.ServerPushStatus(
		SetTimeout(&server.ServerPushStatusParams{
			Body:      req,
			Protocol:  a.protocol,
			SecretKey: a.secretKey,
			ServerID:  a.serverID,
		},
			a.timeout))
	if err != nil {
		return err
	}
	if !rsp.IsSuccess() {
		return errors.New(rsp.Error())
	}
	return nil
}

func (a *Api) ServerPushUserTraffic(req *models.ServerPushUserTrafficRequest) error {
	rsp, err := a.c.ServerPushUserTraffic(
		SetTimeout(&server.ServerPushUserTrafficParams{
			Body:      req,
			Protocol:  a.protocol,
			SecretKey: a.secretKey,
			ServerID:  a.serverID,
		},
			a.timeout))
	if err != nil {
		return err
	}
	if !rsp.IsSuccess() {
		return errors.New(rsp.Error())
	}
	return nil
}
