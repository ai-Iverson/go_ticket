package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateTicketReq struct {
	g.Meta      `path:"/createticket" method:"post" summart:"创建工单" tags:"创建工单"`
	Title       string `json:"title"       description:"工单标题"`
	Type        int    `json:"type"        description:"工单类型，0 问题工单，1 需求工单"`
	Description string `json:"description" description:"问题描述"`
}

type CreateTicketRes struct {
}
