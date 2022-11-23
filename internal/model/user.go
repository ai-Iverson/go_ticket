package model

import "github.com/gogf/gf/v2/os/gtime"

// UserRegisterInput 创建用户
type UserRegisterInput struct {
	Name     string
	Password string
}

// UserLoginInput 登录输入
type UserLoginInput struct {
	Name     string
	Password string
}

type UserLoginOutput struct {
	User *UserInfoOutput
}

type UserInfoOutput struct {
	Id        int
	Name      string
	IsDelete  int
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
}

type UserListInput struct{}

type UserListOutput struct {
	List []UserInfoOutput
}
