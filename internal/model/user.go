package model

// UserRegisterInput 创建用户
type UserRegisterInput struct {
	Name     string
	Password string
}

type UserLoginInput struct {
	Password string // 密码
	Name     string
}
