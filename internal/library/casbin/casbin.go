package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"strings"
)

var Enforcer *casbin.Enforcer

func InitEnforcer(ctx context.Context) {
	mqLink, err := g.Config().Get(ctx, "database.default.link")
	glog.Info(ctx, mqLink.String())
	if err != nil {
		glog.Error(ctx, "未获取到carbin数据库连接配置信息")
		return
	}
	mqConfig := strings.SplitN(mqLink.String(), ":", 2)
	if len(mqConfig) != 2 {
		gerror.New("invalid database link")
		return
	}
	a, _ := gormadapter.NewAdapter(mqConfig[0], mqConfig[1], true)
	e, _ := casbin.NewEnforcer("./manifest/config/casbin.conf", a, a)
	Enforcer = e

}
