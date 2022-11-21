package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterDoReq struct {
	g.Meta   `path:"/register" method:"post" summart:"执行注册请求" tags:"注册"`
	Name     string `json:"name" v:"required|length:4,30#请输入账号|账号长度为:4到:30位"`
	Password string `json:"password" v:"required|password"`
}

type RegisterDoRes struct {
}
