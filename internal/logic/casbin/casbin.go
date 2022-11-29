package casbin

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"go_ticket/internal/dao"
	"go_ticket/internal/model/entity"
	"go_ticket/internal/service"
)

type sCasbinRole struct {
}

func init() {
	service.RegisterCasbinRole(New())
}

func New() *sCasbinRole {
	return &sCasbinRole{}
}

// Verify 验证权限
func (s sCasbinRole) Verify(ctx context.Context, path, method string) bool {
	// TODO 验证白名单

	// 判断如果是超级用户就不需要验证
	isSuperadmin := s.VerifySuperId(ctx)
	if isSuperadmin {
		return true
	}
	return false
}

// 验证是否是超级管理员
func (s sCasbinRole) VerifySuperId(ctx context.Context) bool {
	var userRole *entity.UserRole
	var role *entity.Role
	user := service.Context().Get(ctx)
	userId := user.Token.UserKey
	dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, userId).Scan(&userRole)
	glog.Info(ctx, userRole)
	dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, userRole.RoleId).Scan(&role)
	glog.Info(ctx, role.Name)
	if role.Name == "superadmin" {
		return true
	}
	return false
}
