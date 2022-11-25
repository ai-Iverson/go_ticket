package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
	"go_ticket/utility/utils"
)

var Knowledge = cKnowledge{}

type cKnowledge struct {
}

func (c *cKnowledge) GetAllKnowledge(ctx context.Context, req *v1.KnowledgeListReq) (res *v1.KnowledgeListRes, err error) {
	res = &v1.KnowledgeListRes{}
	glog.Info(ctx, "分页数据定义: ", req.Keyword, req.Page, req.Size)
	allKnowledgeList, err := service.Knowledge().GetKnowledgeData(ctx, model.KnowledgeListInput{
		Keyword: req.Keyword,
		Page:    req.Page,
		Size:    req.Size,
	})
	utils.MyCopy(ctx, res, allKnowledgeList)
	return
}
