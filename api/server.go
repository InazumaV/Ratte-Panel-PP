package api

import (
	"Ratte-Panel-PP/api/client/server"
	"Ratte-Panel-PP/api/models"
	"encoding/json"
	"errors"
)

type commonConfig struct {
	Port int `json:"port"`
}

type SecurityConfig struct {
	SNI                  string `json:"sni"`
	AllowInsecure        *bool  `json:"allow_insecure"`
	Fingerprint          string `json:"fingerprint"`
	RealityServerAddress string `json:"reality_server_addr"`
	RealityServerPort    int    `json:"reality_server_port"`
	RealityPrivateKey    string `json:"reality_private_key"`
	RealityPublicKey     string `json:"reality_public_key"`
	RealityShortId       string `json:"reality_short_id"`
}

type TransportConfig struct {
	Path        string `json:"path"`
	Host        string `json:"host"`
	ServiceName string `json:"service_name"`
}

type VmessConfig struct {
	commonConfig    `json:",inline"`
	Network         string           `json:"transport"`
	Security        string           `json:"security"`
	SecurityConfig  *SecurityConfig  `json:"security_config"`
	TransportConfig *TransportConfig `json:"transport_config"`
}

type VlessConfig struct {
	VmessConfig `json:",inline"`
	Flow        string `json:"flow"`
}

type TrojanConfig struct {
	commonConfig    `json:",inline"`
	Network         string           `json:"transport"`
	Security        string           `json:"security"`
	SecurityConfig  *SecurityConfig  `json:"security_config"`
	TransportConfig *TransportConfig `json:"transport_config"`
}

type ShadowsocksConfig struct {
	commonConfig `json:",inline"`
	Cipher       string `json:"method"`
	ServerKey    string `json:"server_key"`
}

type TuicNode struct {
	commonConfig   `json:",inline"`
	SecurityConfig *SecurityConfig `json:"security_config"`
}

type Hysteria2Node struct {
	commonConfig   `json:",inline"`
	HopPorts       string          `json:"hop_ports"`
	HopInterval    int             `json:"hop_interval"`
	ObfsPassword   string          `json:"obfs_password"`
	SecurityConfig *SecurityConfig `json:"security_config"`
}

type protocolConfig struct {
	VmessConfig       *VmessConfig       `json:"vmess"`
	VlessConfig       *VlessConfig       `json:"vless"`
	TrojanConfig      *TrojanConfig      `json:"trojan"`
	ShadowsocksConfig *ShadowsocksConfig `json:"shadowsocks"`
	TUICConfig        *TuicNode          `json:"tuic"`
}

func (p *protocolConfig) unmarshalJson(data []byte, protocol string) error {
	var err error
	switch protocol {
	case "vmess":
		err = json.Unmarshal(data, &p.VmessConfig)
	case "vless":
		err = json.Unmarshal(data, &p.VlessConfig)
	case "trojan":
		err = json.Unmarshal(data, &p.TrojanConfig)
	case "shadowsocks":
		err = json.Unmarshal(data, &p.ShadowsocksConfig)
	case "tuic":
		err = json.Unmarshal(data, &p.TUICConfig)
	default:
		err = errors.New("unknown protocol")
	}
	if err != nil {
		return err
	}
	return nil
}

type GetServerConfigRsp struct {
	protocolConfig
	Basic     *models.ServerBasic
	RawConfig json.RawMessage
	Protocol  string
}

func (a *Api) GetServerConfig() (*GetServerConfigRsp, error) {
	req := WithTimeout(&server.GetServerConfigParams{
		Protocol:  a.protocol,
		SecretKey: a.secretKey,
		ServerID:  a.serverID,
	}, a.timeout)
	rsp, err := a.c.GetServerConfig(
		req,
		WithEtag[server.GetServerConfigOK](a.nodeEtag),
	)
	if err != nil {
		return nil, err
	}
	if rsp == nil {
		return nil, nil
	}
	if !rsp.IsSuccess() {
		return nil, errors.New(rsp.Error())
	}
	var config protocolConfig
	err = config.unmarshalJson(rsp.Payload.Data.Config, *rsp.Payload.Data.Protocol)
	return &GetServerConfigRsp{
		Protocol:       *rsp.Payload.Data.Protocol,
		Basic:          rsp.Payload.Data.Basic,
		RawConfig:      rsp.Payload.Data.Config,
		protocolConfig: config,
	}, nil
}

func (a *Api) ServerPushStatus(req *models.ServerPushStatusRequest) error {
	reqV := WithTimeout(&server.ServerPushStatusParams{
		Body:      req,
		Protocol:  a.protocol,
		SecretKey: a.secretKey,
		ServerID:  a.serverID,
	},
		a.timeout)
	rsp, err := a.c.ServerPushStatus(reqV)
	if err != nil {
		return err
	}

	if !rsp.IsSuccess() {
		return errors.New(rsp.Error())
	}
	return nil
}
