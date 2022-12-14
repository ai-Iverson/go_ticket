package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" summart:"用户登录" tags:"登录"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginRes struct {
	User  *UserGetInfoRes
	Token string `json:"token"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" summart:"用户注销" tags:"注销"`
}
type LogoutRes struct {
}
