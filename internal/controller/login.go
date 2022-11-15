package controller

import (
	"context"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
)

var Login = CLogin{}

type CLogin struct{}

func (c *CLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	user, err := service.User().Login(ctx, model.UserLoginInput{
		Password: req.Name,
		Name:     req.Password,
	},
	)
	res.Token = user.Password
	return res, nil
}
