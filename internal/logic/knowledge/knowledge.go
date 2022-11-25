package knowledge

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/dao"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
)

type sKnowledge struct{}

func init() {
	service.RegisterKnowledge(New())
}

func New() *sKnowledge {
	return &sKnowledge{}
}

func (s *sKnowledge) GetKnowledgeData(ctx context.Context, in model.KnowledgeListInput) (out *model.KnowledgeListOutput, err error) {
	m := dao.KnowledgeBase.Ctx(ctx)
	out = &model.KnowledgeListOutput{}
	out.Page = in.Page
	out.Size = in.Size
	glog.Info(ctx, "xxxx", out, in.Keyword)
	out.List = []v1.KnowledgeRes{}
	glog.Info(ctx, "xxxx", out)
	if in.Keyword != "" {
		glog.Info(ctx, "keyword", in.Keyword)
		likePattern := "%" + in.Keyword + "%"
		m = m.WhereOrLike(dao.KnowledgeBase.Columns().FunctionPath, likePattern).
			WhereOrLike(dao.KnowledgeBase.Columns().ProblemDesc, likePattern).
			WhereOrLike(dao.KnowledgeBase.Columns().Summary, likePattern).
			WhereOrLike(dao.KnowledgeBase.Columns().SuggestedSolution, likePattern)
		glog.Info(ctx, m)
	}
	allKnowledge, _ := m.Page(in.Page, in.Size).OrderDesc(dao.KnowledgeBase.Columns().CreatedAt).All()
	out.Total, _ = m.Count()
	err = allKnowledge.Structs(&out.List)
	glog.Info(ctx, allKnowledge)

	return

}
