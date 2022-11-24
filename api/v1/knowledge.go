package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type KnowledgeRes struct {
	Code              string      `json:"code"              description:""`
	FunctionPath      string      `json:"functionPath"      description:""`
	ModuleName        string      `json:"moduleName"        description:""`
	DescoveryVersion  string      `json:"descoveryVersion"  description:""`
	Summary           string      `json:"summary"           description:""`
	ProblemDesc       string      `json:"problemDesc"       description:""`
	SuggestedSolution string      `json:"suggestedSolution" description:""`
	CreatedAt         *gtime.Time `json:"createdAt"         description:""`
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:""`
}
type KnowledgeListReq struct {
	g.Meta `path:"/knowledge" method:"get" summart:"获取知识库信息" tags:"获取知识库信息"`
	Page   int `json:"page" in:"query" description:"分页码"`
	Size   int `json:"size" in:"query" description:"分页数量"`
}

type KnowledgeListRes struct {
	List  []KnowledgeRes `json:"users" description:"知识库列表"`
	Page  int            `json:"page" description:"分页码"`
	Size  int            `json:"size" description:"分页数量"`
	Total int            `json:"total" description:"数据总数"`
}
