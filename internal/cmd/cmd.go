package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"go_ticket/internal/controller"
	"go_ticket/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			SetLoggerDefaultHandler() // 替代默认的log

			s := g.Server()
			s.Use(
				//service.Middleware().MiddlewareCORS,
				service.Middleware().Ctx,
				service.Middleware().I18NMiddleware,
				service.Middleware().ResponseHandler,
			)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.Register,
					controller.Login.Login,
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse,
					service.Middleware().TokenAuth)
				group.Bind(
					controller.Login.Logout,
				)
			})

			s.Run()
			return nil
		},
	}
)

// SetLoggerDefaultHandler 替代默认的日志handler，禁止控制台输出，全部输出到文件
func SetLoggerDefaultHandler() {
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		m := map[string]interface{}{
			"stdout":            g.Config().MustGet(ctx, "logger.globalStdout", true).Bool(), // 使用自定义的全局字段
			"path":              g.Config().MustGet(ctx, "logger.path", "log/").String(),     // 此处必须重新设置，才可以实现db的log写入文件
			"writerColorEnable": true,
		}

		in.Logger.SetConfigWithMap(m)
		in.Next(ctx)
	})
	// 添加日志代码行号和年月日
	glog.SetFlags(glog.F_TIME_STD | glog.F_FILE_SHORT)
}
