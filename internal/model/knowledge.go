package model

import v1 "go_ticket/api/v1"

type KnowledgeListInput struct {
	Keyword string `json:"keyword" description:"关键字"`
	Page    int    `json:"page" description:"分页码"`
	Size    int    `json:"size" description:"分页数量"`
}
type KnowledgeListOutput struct {
	v1.KnowledgeListRes
}
