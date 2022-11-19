package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"go_ticket/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			SetLoggerDefaultHandler() // 替代默认的log

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.Register,
					controller.Login,
				)
			})
			s.Run()
			return nil
		},
	}
)

// 替代默认的日志handler，禁止控制台输出，全部输出到文件
func SetLoggerDefaultHandler() {
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		m := map[string]interface{}{
			"stdout": g.Config().MustGet(ctx, "logger.globalStdout", true).Bool(), // 使用自定义的全局字段
			"path":   g.Config().MustGet(ctx, "logger.path", "log/").String(),     // 此处必须重新设置，才可以实现db的log写入文件
		}
		in.Logger.SetConfigWithMap(m)
		in.Next(ctx)
	})
}
