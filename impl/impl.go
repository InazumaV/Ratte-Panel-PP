package impl

import (
	"Ratte-Panel-PP/api"
	"Ratte-Panel-PP/api/models"
	"Ratte-Panel-PP/status"
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

func (i Impl) GetNodeInfo(id int) *panel.GetNodeInfoRsp {
	r, _ := i.remotes.Get(KeyInt(id))
	r.api.GetServerConfig()
	//TODO implement me
	panic("implement me")
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
