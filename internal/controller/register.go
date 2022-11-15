package controller

import (
	"context"
	"go_ticket/api/v1"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
)

//注册控制器
var Register = cRegister{}

type cRegister struct{}

func (c *cRegister) Register(ctx context.Context, req *v1.RegisterDoReq) (res *v1.RegisterDoRes, err error) {
	if err = service.User().Register(ctx, model.UserRegisterInput{
		Name:     req.Name,
		Password: req.Password,
	}); err != nil {
		return
	}
	v := &v1.RegisterDoRes{
		Referer: "123",
	}
	return v, err
}
