package controller

import (
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"go_ticket/internal/service"
)

var (
	Schedules cSchedules
)

type cSchedules struct {
}

//var ctx context.Context

func (s *cSchedules) Initialize() error {
	var xxx = gctx.New()
	ctx := xxx
	gcron.New()
	_, err := gcron.AddSingleton(ctx, "* * * * * *", service.Scheduled().GetTicketData)
	if err != nil {
		return err
	}
	//_, err = gcron.AddSingleton(ctx, "* * * * * *", service.Scheduled().Task1)
	//if err != nil {
	//	return err
	//}
	return nil
}
