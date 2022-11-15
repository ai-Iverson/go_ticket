package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterDoReq struct {
	g.Meta   `path:"/register" method:"post" summart:"执行注册请求" tags:"注册"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterDoRes struct {
}
