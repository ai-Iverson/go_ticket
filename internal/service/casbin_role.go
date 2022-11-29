// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
)

type ICasbinRole interface {
	Verify(ctx context.Context, path, method string) bool
	VerifySuperId(ctx context.Context) bool
}

var localCasbinRole ICasbinRole

func CasbinRole() ICasbinRole {
	if localCasbinRole == nil {
		panic("implement not found for interface ICasbinRole, forgot register?")
	}
	return localCasbinRole
}

func RegisterCasbinRole(i ICasbinRole) {
	localCasbinRole = i
}
