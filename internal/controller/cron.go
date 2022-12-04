package controller

import (
	"context"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
	"go_ticket/utility/utils"
)

var (
	Cron = cCron{}
)

type cCron struct{}

func (c *cCron) AddCronTaskReq(ctx context.Context, req *v1.AddCronTaskReq) (res *v1.AddCronTaskRes, err error) {
	var in = model.AddCronTaskInput{}
	utils.MyCopy(ctx, &in, req)
	err = service.Cron().AddCronTask(ctx, in)
	if err != nil {
		return nil, err
	}
	return nil, err
}
