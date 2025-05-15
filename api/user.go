package api

import (
	"Ratte-Panel-PP/api/client/server"
	"Ratte-Panel-PP/api/models"
	"errors"
)

func (a *Api) GetServerUserList() (*models.GetServerUserListResponse, error) {
	reqV := WithTimeout(&server.GetServerUserListParams{
		Protocol:  a.protocol,
		SecretKey: a.secretKey,
		ServerID:  a.serverID,
	}, a.timeout)
	rsp, err := a.c.GetServerUserList(reqV, WithEtag[server.GetServerUserListOK](a.userEtag))
	if err != nil {
		return nil, err
	}
	if rsp == nil {
		return nil, nil
	}
	if !rsp.IsSuccess() {
		return nil, errors.New(rsp.Error())
	}
	return rsp.GetPayload().Data, nil
}

func (a *Api) PushOnlineUsers(req *models.OnlineUsersRequest) error {
	reqV := WithTimeout(&server.PushOnlineUsersParams{
		Body:      req,
		Protocol:  a.protocol,
		SecretKey: a.secretKey,
		ServerID:  a.serverID,
	}, a.timeout)
	rsp, err := a.c.PushOnlineUsers(reqV)
	if err != nil {
		return err
	}
	if !rsp.IsSuccess() {
		return errors.New(rsp.Error())
	}
	return nil
}

func (a *Api) ServerPushUserTraffic(req *models.ServerPushUserTrafficRequest) error {
	reqV := WithTimeout(&server.ServerPushUserTrafficParams{
		Body:      req,
		Protocol:  a.protocol,
		SecretKey: a.secretKey,
		ServerID:  a.serverID,
	},
		a.timeout)
	rsp, err := a.c.ServerPushUserTraffic(reqV)
	if err != nil {
		return err
	}
	if !rsp.IsSuccess() {
		return errors.New(rsp.Error())
	}
	return nil
}
