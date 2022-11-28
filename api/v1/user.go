package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserGetInfoReq struct {
	g.Meta `path:"/userinfo" method:"get" summart:"获取用户信息" tags:"获取用户信息"`
	Id     uint `json:"userid" in:"path"`
}

type UserGetInfoRes struct {
	Id        int         `json:"userid"`
	Name      string      `json:"name"`
	IsAdmin   int         `json:"is_amdin"`
	RoleId    int         `json:"role_id"`
	IsDelete  int         `json:"isDelete"  description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}

type UserListReq struct {
	g.Meta `path:"/userlist" method:"post" summart:"获取用户信息" tags:"获取用户信息"`
}

type UserListRes struct {
	List []UserGetInfoRes `json:"users" description:"用户列表"`
}
