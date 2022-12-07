package model

import "github.com/gogf/gf/v2/os/gtime"

type CreateTicketInput struct {
	Title       string `json:"title"       description:"工单标题"`
	Type        int    `json:"type"        description:"工单类型，0 问题工单，1 需求工单"`
	Description string `json:"description" description:"问题描述"`
}

type TickerOutput struct {
	Id          int         `json:"id"          description:""`
	Code        string      `json:"code"        description:"工单号"`
	Title       string      `json:"title"       description:"工单标题"`
	Description string      `json:"description" description:"问题描述"`
	CreatUserId int         `json:"creatUserId" description:"创建人id"`
	Type        int         `json:"type"        description:"工单类型，0 问题工单，1 需求工单"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"工单创建时间"`
	HandlerId   int         `json:"handlerId"   description:"处理人id"`
	ReasonId    int         `json:"reasonId"    description:"问题原因id"`
	DealReason  string      `json:"dealReason"  description:"解决办法"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
	Status      int         `json:"status"      description:"工单状态"`
}
