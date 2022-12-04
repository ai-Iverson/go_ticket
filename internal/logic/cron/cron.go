package cron

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"go_ticket/internal/consts"
	"go_ticket/internal/crons"
	"go_ticket/internal/dao"
	"go_ticket/internal/model"
	"go_ticket/internal/model/entity"
	"go_ticket/internal/service"
	"strings"
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

func (s *sCron) AddCronTask(ctx context.Context, in model.AddCronTaskInput) (err error) {
	// TODO 状态数据状态校验
	glog.Info(ctx, "状态：", in.Status)

	count, _ := dao.Cron.Ctx(ctx).Where(dao.Cron.Columns().Name, in.Name).Count()
	if count > 0 {
		return gerror.Newf("此任务已经存在:%v", in.Name)
	}

	// 判断添加的任务状态
	if in.Status == consts.StatusEnabled {
		// 	在gcron队列中添加任务
		f := crons.Inst.Get(in.Func)
		if f == nil {
			return gerror.Newf("该任务没有加入任务列表:%v", in.Func)
		}
		if gcron.Search(in.Name) == nil {
			var (
				t  *gcron.Entry
				ct = context.WithValue(gctx.New(), consts.CronArgsKey, strings.Split(in.Params, consts.CronSplitStr))
			)
			switch in.Policy {
			case consts.CronPolicySame:
				t, err = gcron.Add(ctx, in.Pattern, f.Fun, in.Name)

			case consts.CronPolicySingle:
				t, err = gcron.AddSingleton(ct, in.Pattern, f.Fun, in.Name)

			case consts.CronPolicyOnce:
				t, err = gcron.AddOnce(ct, in.Pattern, f.Fun, in.Name)

			case consts.CronPolicyTimes:
				if f.Count <= 0 {
					f.Count = 1
				}
				t, err = gcron.AddTimes(ct, in.Pattern, int(in.Count), f.Fun, in.Name)

			default:
				return gerror.Newf("使用无效的策略, cron.Policy=%v", in.Policy)
			}
			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("添加任务失败")
			}
		} else {
			return gerror.Newf("此任务已经存在:%v", in.Name)
		}

		// 在数据库中记录

		_, err = dao.Cron.Ctx(ctx).Data(in).Insert()

		if err != nil {
			return gerror.Newf("内部Orm错误:%v", err)
		}
	} else {

		_, err = dao.Cron.Ctx(ctx).Data(in).Insert()
		if err != nil {
			return gerror.Newf("内部Orm错误:%v", err)
		}
	}
	return nil
}
