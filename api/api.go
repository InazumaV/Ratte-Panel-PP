package api

import (
	"Ratte-Panel-PP/api/client"
	"Ratte-Panel-PP/api/client/server"
	"net/url"
	"time"
)

type Api struct {
	c         server.ClientService
	protocol  string
	secretKey string
	serverID  int64
	nodeEtag  string
	userEtag  string
	timeout   time.Duration
}

func New(
	baseUrl,
	key string,
	protocol string,
	serverID int64,
	timeout time.Duration,
) (*Api, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	cfg := client.DefaultTransportConfig()
	cfg.Host = baseUrl
	cfg.BasePath = u.Path
	cfg.Schemes = []string{u.Scheme}
	c := client.NewHTTPClientWithConfig(nil, cfg)
	return &Api{
		c:         c.Server,
		protocol:  protocol,
		secretKey: key,
		serverID:  serverID,
		timeout:   timeout,
	}, nil
}
