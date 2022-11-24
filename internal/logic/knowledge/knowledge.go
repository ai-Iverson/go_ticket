package knowledge

import (
	"context"
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
	out.List = []v1.KnowledgeRes{}
	allKnowledge, err := m.Page(in.Page, in.Size).All()
	out.Total, err = m.Count()
	err = allKnowledge.Structs(&out.List)
	return

}
