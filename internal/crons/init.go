package crons

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"go_ticket/internal/consts"
	"go_ticket/internal/dao"
	"go_ticket/internal/model/entity"
	"strings"
	"sync"
)

var (
	cronList = []cronStrategy{
		Test,  // 测试无参数任务
		Test2, // 测试有参数任务
	}
	Inst = new(tasks)
)

type cronStrategy interface {
	GetName() string
	Execute(ctx context.Context)
}

type tasks struct {
	list []*TaskItem
	// 读写锁
	sync.RWMutex
}

type TaskItem struct {
	Pattern string        // 表达式，参考：https://goframe.org/pages/viewpage.action?pageId=30736411
	Name    string        // 唯一的任务名称
	Params  string        // 函数参数，多个用,隔开
	Fun     gcron.JobFunc // 执行的函数接口
	Policy  int64         // 策略 1：并行 2：单例 3：单次 4：多次
	Count   int           // 执行次数，仅Policy=4时有效
}

func init() {
	for _, cron := range cronList {
		Inst.Add(&TaskItem{
			Name: cron.GetName(),
			Fun:  cron.Execute,
		})
	}

}

// StartALL 启动任务
func StartALL(sysCron []*entity.Cron) error {
	var (
		err error
		ct  = gctx.New()
	)

	if len(sysCron) == 0 {
		g.Log().Info(ct, "没有可用的定时任务")
		return nil
	}

	for _, cron := range sysCron {
		f := Inst.Get(cron.Name)
		if f == nil {
			return gerror.Newf("该任务没有加入任务列表:%v", cron.Name)
		}

		// 没有则添加
		// 根据数据库中任务记录的策略依次注册到gcron的队列中
		if gcron.Search(cron.Name) == nil {
			var (
				t   *gcron.Entry
				ctx = context.WithValue(gctx.New(), consts.CronArgsKey, strings.Split(cron.Params, consts.CronSplitStr))
			)
			switch cron.Policy {
			case consts.CronPolicySame:
				t, err = gcron.Add(ctx, cron.Pattern, f.Fun, cron.Name)

			case consts.CronPolicySingle:
				t, err = gcron.AddSingleton(ctx, cron.Pattern, f.Fun, cron.Name)

			case consts.CronPolicyOnce:
				t, err = gcron.AddOnce(ctx, cron.Pattern, f.Fun, cron.Name)

			case consts.CronPolicyTimes:
				if f.Count <= 0 {
					f.Count = 1
				}
				t, err = gcron.AddTimes(ctx, cron.Pattern, int(cron.Count), f.Fun, cron.Name)

			default:
				return gerror.Newf("使用无效的策略, cron.Policy=%v", cron.Policy)
			}

			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("启动任务失败")
			}
		}

		gcron.Start(cron.Name)

		// 执行完毕，单次和多次执行的任务更新状态
		if cron.Policy == consts.CronPolicyOnce || cron.Policy == consts.CronPolicyTimes {
			_, err = dao.Cron.Ctx(ct).Where("id", cron.Id).
				Data(g.Map{"status": consts.StatusDisable, "updated_at": gtime.Now()}).
				Update()
			if err != nil {
				err = gerror.Wrap(err, "sql执行异常")
				return err
			}
		}
	}

	g.Log().Info(ct, "定时任务启动完毕...")
	return nil
}

/// Stop 停止单个任务
func Stop(sysCron *entity.Cron) error {
	return nil
}

// Once 立即执行一次某个任务
func Once(sysCron *entity.Cron) error {
	return nil
}

// Delete 删除任务
func Delete(sysCron *entity.Cron) error {
	// ...

	return Stop(sysCron)
}

// Start 启动单个任务
func Start(sysCron *entity.Cron) error {
	return nil
}

// Add 添加任务
// 只是把所有task加入到tasks的list中供后面好查找
func (t *tasks) Add(task *TaskItem) *tasks {
	if task.Name == "" || task.Fun == nil {
		return t
	}
	t.Lock()
	defer t.Unlock()
	t.list = append(t.list, task)
	return t
}

// Get 找到任务
// 根据数据库中存入的任务名称到task.list中找对应的名称的task
func (t *tasks) Get(name string) *TaskItem {
	if len(t.list) == 0 {
		return nil
	}

	for _, item := range t.list {
		if item.Name == name {
			return item
		}
	}
	return nil
}

// Del 删除任务
func (t *tasks) Del(name string) (newList []*TaskItem) {
	if len(t.list) == 0 {
		return nil
	}
	t.Lock()
	defer t.Unlock()

	for _, item := range t.list {
		if item.Name == name {
			continue
		}
		newList = append(newList, item)
	}
	return newList
}
