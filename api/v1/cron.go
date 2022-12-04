package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go_ticket/internal/model/entity"
)

type AddCronTaskReq struct {
	g.Meta `path:"/addcrontask" tags:"addcrontask" method:"post" summary:"添加定时任务"`
	entity.Cron
}
type AddCronTaskRes struct {
}
