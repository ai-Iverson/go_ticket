package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" `
	Password string `json:"password"`
}
type LoginDoRes struct {
	Token string `json:"token"`
}
