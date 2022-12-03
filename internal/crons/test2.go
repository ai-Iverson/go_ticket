package crons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go_ticket/internal/consts"
	"time"
)

// Test2 测试2任务
var Test2 = &cTest2{name: "test2"}

type cTest2 struct {
	name string
}

func (c *cTest2) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cTest2) Execute(ctx context.Context) {
	args, ok := ctx.Value(consts.CronArgsKey).([]string)
	if !ok {
		g.Log().Warning(ctx, "参数解析失败!")
		return
	}
	if len(args) != 3 {
		g.Log().Warning(ctx, "test2 传入参数不正确!")
		return
	}

	var (
		name = args[0]
		age  = args[1]
		msg  = args[2]
	)

	g.Log().Infof(ctx, "cron test2 Execute:%v, name:%v, age:%v, msg:%v", time.Now(), name, age, msg)
}
