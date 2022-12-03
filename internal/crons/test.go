package crons

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// Test 测试任务
var Test = &cTest{name: "test"}

type cTest struct {
	name string
}

func (c *cTest) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cTest) Execute(ctx context.Context) {
	g.Log().Infof(ctx, "cron test Execute:%v", time.Now())
}
