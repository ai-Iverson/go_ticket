// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"go_ticket/internal/model"
	"go_ticket/internal/model/entity"
)

type IUser interface {
	Register(ctx context.Context, in model.UserRegisterInput) error
	EncryptPassword(password string) string
	GetUserByPassportAndPassword(ctx context.Context, name, password string) (user *entity.User, err error)
	Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
