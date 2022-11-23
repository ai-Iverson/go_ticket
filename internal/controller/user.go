package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
	"go_ticket/utility/utils"
)

var User = cUser{}

type cUser struct {
}

func (l *cUser) GetUserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	userList, err := service.User().UserList(ctx, model.UserListInput{})
	glog.Info(ctx, "con_GetUserList用户list信息: ", userList)
	if err != nil {
		glog.Errorf(ctx, "con_GetUserList获取用户信息失败", err)
		return nil, gerror.New("获取用户list信息失败")
	}
	res = &v1.UserListRes{}
	err = utils.MyCopy(ctx, res, userList)
	if err != nil {
		glog.Errorf(ctx, "复制信息出错", err, res, userList)
		return nil, gerror.New("获取用户list信息失败")
	}
	glog.Info(ctx, "con_UserListRes用户list信息: ", res)
	return res, nil
}
