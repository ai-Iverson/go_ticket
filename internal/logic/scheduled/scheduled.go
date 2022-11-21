package scheduled

//import (
//	"github.com/gogf/gf/v2/os/gctx"
//	"github.com/gogf/gf/v2/os/glog"
//	"github.com/gogf/gf/v2/os/gtimer"
//	"go_ticket/internal/service"
//	"time"
//)
//
//type sScheduled struct{}
//
//func init() {
//	service.RegisterScheduled(New())
//}
//
//func New() *sScheduled {
//	return &sScheduled{}
//}
//
//func (s *sScheduled) singletonTask() {
//	ctx := gctx.New()
//	gtimer.AddSingleton(ctx, time.Second, func() {
//		glog.Info(ctx, "定时任务")
//	})
//}
