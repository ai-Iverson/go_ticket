package cron

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"go_ticket/internal/consts"
	"go_ticket/internal/crons"
	"go_ticket/internal/dao"
	"go_ticket/internal/model/entity"
	"go_ticket/internal/service"
)

type sCron struct {
}

func init() {
	service.RegisterCron(New())
}

func New() *sCron {
	return &sCron{}
}

func (s *sCron) StartCron(ctx context.Context) {
	var list []*entity.Cron

	err := dao.Cron.Ctx(ctx).Where(dao.Cron.Columns().Status, consts.StatusEnabled).Order("sort asc,id desc").Scan(&list)
	if err != nil {
		glog.Error(ctx, "定时任务获取失败")
		return
	}

	if err := crons.StartALL(list); err != nil {
		g.Log().Fatalf(ctx, "定时任务启动失败, err . %v", err)
		return
	}

}
