package impl

import (
	"github.com/InazumaV/Ratte-Interface/panel"
	cmap "github.com/orcaman/concurrent-map/v2"
	"resty.dev/v3"
	"strconv"
)

type KeyInt int

func (k KeyInt) String() string {
	return strconv.Itoa(int(k))
}

type Remote struct {
	nodeEtag string
	userEtag string
	*panel.AddRemoteParams
}

type Impl struct {
	client  *resty.Client
	remotes cmap.ConcurrentMap[KeyInt, *Remote]
}

func (i Impl) CustomMethod(method string, args any, reply *any) error {
	//TODO implement me
	panic("implement me")
}

func (i Impl) AddRemote(params *panel.AddRemoteParams) *panel.AddRemoteRsp {
	//TODO implement me
	panic("implement me")
}

func (i Impl) DelRemote(id int) error {
	//TODO implement me
	panic("implement me")
}

func (i Impl) GetNodeInfo(id int) *panel.GetNodeInfoRsp {
	//TODO implement me
	panic("implement me")
}

func (i Impl) GetUserList(id int) *panel.GetUserListRsp {
	//TODO implement me
	panic("implement me")
}

func (i Impl) ReportUserTraffic(p *panel.ReportUserTrafficParams) error {
	//TODO implement me
	panic("implement me")
}

func New() *Impl {
	return &Impl{
		client:  resty.New(),
		remotes: cmap.NewStringer[KeyInt, *Remote](),
	}
}
