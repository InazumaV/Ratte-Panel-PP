package impl

import (
	"Ratte-Panel-PP/api"
	"Ratte-Panel-PP/api/models"
	"Ratte-Panel-PP/status"
	"errors"
	"github.com/InazumaV/Ratte-Interface/panel"
	"github.com/InazumaV/Ratte-Interface/params"
	cmap "github.com/orcaman/concurrent-map/v2"
	"strconv"
	"time"
)

type KeyInt int

func (k KeyInt) String() string {
	return strconv.Itoa(int(k))
}

type Remote struct {
	api *api.Api
	*panel.AddRemoteParams
}

type Impl struct {
	remotes cmap.ConcurrentMap[KeyInt, *Remote]
}

func (i Impl) CustomMethod(method string, args any, reply *any) error {
	//TODO implement me
	panic("implement me")
}

func (i Impl) AddRemote(params *panel.AddRemoteParams) *panel.AddRemoteRsp {
	rid := i.remotes.Count()
	a, err := api.New(
		params.Baseurl,
		params.Key,
		params.NodeType,
		int64(params.NodeId),
		time.Duration(params.Timeout)*time.Second,
	)
	if err != nil {
		return &panel.AddRemoteRsp{
			Err: err,
		}
	}
	i.remotes.Set(KeyInt(rid), &Remote{
		AddRemoteParams: params,
		api:             a,
	})
	return &panel.AddRemoteRsp{
		RemoteId: rid,
	}
}

func (i Impl) DelRemote(id int) error {
	i.remotes.Remove(KeyInt(id))
	return nil
}

func parseSecurity(t string, as *api.SecurityConfig, ps *params.SecurityConfig) {
	switch t {
	case "tls":
		ps.TlsSettings = params.TlsSettings{
			ServerName:    as.SNI,
			AllowInsecure: *as.AllowInsecure,
			Fingerprint:   as.Fingerprint,
		}
	case "reality":
		ps.RealityConfig = params.RealityConfig{
			ServerName: as.RealityServerAddress,
			ServerPort: as.RealityServerPort,
			ShortId:    as.RealityShortId,
			PrivateKey: as.RealityPrivateKey,
		}
	}
}

func (i Impl) GetNodeInfo(id int) *panel.GetNodeInfoRsp {
	r, _ := i.remotes.Get(KeyInt(id))
	c, err := r.api.GetServerConfig()
	if err != nil {
		return &panel.GetNodeInfoRsp{
			Err: err,
		}
	}
	p := panel.NodeInfo{
		Type: c.Protocol,
		Name: r.Name,
		ExpandParams: params.ExpandParams{
			CustomData: c.RawConfig,
		},
	}
	switch c.Protocol {
	case "vmess":
		p.Port = c.VmessConfig.Port
		p.Security = c.VmessConfig.Security
		p.SecurityConfig = new(params.SecurityConfig)
		parseSecurity(c.VmessConfig.Security, c.VmessConfig.SecurityConfig, p.SecurityConfig)
		p.VMess = &params.VMess{
			ExpandParams: params.ExpandParams{
				Options: map[string]any{
					"ServiceName": c.VmessConfig.TransportConfig.ServiceName,
					"Path":        c.VmessConfig.TransportConfig.Path,
				},
			},
			Network: c.VmessConfig.Network,
		}
		switch p.VLess.Network {
		case "grpc":
			p.VLess.NetworkSettings.Grpc = params.GrpcSettings{
				Authority:   c.VlessConfig.TransportConfig.Host,
				ServiceName: c.VlessConfig.TransportConfig.ServiceName,
			}
		case "ws":
			p.VLess.NetworkSettings.Ws = params.WsSettings{
				Host: c.VlessConfig.TransportConfig.Host,
				Path: c.VlessConfig.TransportConfig.Path,
			}
		}
	case "vless":
		p.Port = c.VlessConfig.Port
		p.Security = c.VlessConfig.Security
		p.SecurityConfig = new(params.SecurityConfig)
		parseSecurity(c.VlessConfig.Security, c.VlessConfig.SecurityConfig, p.SecurityConfig)
		p.VLess = &params.VLess{
			Flow: c.VlessConfig.Flow,
			VMess: params.VMess{
				ExpandParams: params.ExpandParams{
					Options: map[string]any{
						"ServiceName": c.VlessConfig.TransportConfig.ServiceName,
						"Path":        c.VlessConfig.TransportConfig.Path,
					},
				},
				Network: c.VmessConfig.Network,
			},
		}
		switch p.VLess.Network {
		case "grpc":
			p.VLess.NetworkSettings.Grpc = params.GrpcSettings{
				Authority:   c.VlessConfig.TransportConfig.Host,
				ServiceName: c.VlessConfig.TransportConfig.ServiceName,
			}
		case "ws":
			p.VLess.NetworkSettings.Ws = params.WsSettings{
				Host: c.VlessConfig.TransportConfig.Host,
				Path: c.VlessConfig.TransportConfig.Path,
			}
		}
	case "shadowsocks":
		p.Port = c.ShadowsocksConfig.Port
		p.Shadowsocks = &params.Shadowsocks{
			Cipher:    c.ShadowsocksConfig.Cipher,
			ServerKey: c.ShadowsocksConfig.ServerKey,
		}
	case "trojan":
		p.Port = c.TrojanConfig.Port
		p.Security = c.TrojanConfig.Security
		p.SecurityConfig = new(params.SecurityConfig)
		parseSecurity(c.TrojanConfig.Security, c.TrojanConfig.SecurityConfig, p.SecurityConfig)
		p.Trojan = &params.Trojan{
			Host: c.TrojanConfig.TransportConfig.Host,
		}
	default:
		return &panel.GetNodeInfoRsp{
			Err: errors.New("not supported"),
		}
	}
	return &panel.GetNodeInfoRsp{
		NodeInfo: p,
	}
}

func (i Impl) GetUserList(id int) *panel.GetUserListRsp {
	r, _ := i.remotes.Get(KeyInt(id))
	rsp, err := r.api.GetServerUserList()
	if err != nil {
		return &panel.GetUserListRsp{
			Err: err,
		}
	}
	us := make([]panel.UserInfo, 0, len(rsp.Users))
	for _, v := range rsp.Users {
		us = append(us, panel.UserInfo{
			HashOrKey: *v.UUID +
				strconv.FormatInt(*v.SpeedLimit, 10) +
				strconv.FormatInt(*v.DeviceLimit, 10),
			UserInfo: params.UserInfo{
				Id:  int(*v.ID),
				Key: []string{*v.UUID},
			},
		})
	}
	return &panel.GetUserListRsp{
		Users: us,
	}
}

func (i Impl) ReportUserTraffic(p *panel.ReportUserTrafficParams) error {
	r, _ := i.remotes.Get(KeyInt(p.Id))
	if r == nil {
		return nil
	}

	{
		ous := make([]*models.OnlineUser, 0, len(p.Users))
		for _, v := range p.Users {
			id := int64(v.Id)
			ous = append(ous, &models.OnlineUser{
				UID: &id,
			})
		}
		err := r.api.PushOnlineUsers(&models.OnlineUsersRequest{
			Users: ous,
		})
		if err != nil {
			return err
		}
	}

	{
		uts := make([]*models.UserTraffic, 0, len(p.Users))
		for _, v := range p.Users {
			id := int64(v.Id)
			uts = append(uts, &models.UserTraffic{
				UID:      &id,
				Upload:   &v.Upload,
				Download: &v.Download,
			})
		}
		err := r.api.ServerPushUserTraffic(&models.ServerPushUserTrafficRequest{
			Traffic: uts,
		})
		if err != nil {
			return err
		}
	}

	ss, err := status.GetSystemStatus()
	if err != nil {
		return err
	}
	t := time.Now().Unix()
	err = r.api.ServerPushStatus(&models.ServerPushStatusRequest{
		CPU:       &ss.CPUUsage,
		Mem:       &ss.MemoryUsage,
		Disk:      &ss.DiskUsage,
		UpdatedAt: &t,
	})
	if err != nil {
		return err
	}
	return nil
}

func New() *Impl {
	return &Impl{
		remotes: cmap.NewStringer[KeyInt, *Remote](),
	}
}
