package controller

import (
	"context"
	v1 "go_ticket/api/v1"
	"go_ticket/internal/model"
	"go_ticket/internal/service"
)

var Ticket = cTicket{}

type cTicket struct {
}

func (c *cTicket) CreateTicket(ctx context.Context, req *v1.CreateTicketReq) (res *v1.CreateTicketRes, err error) {

	err = service.Ticket().CreateTicket(ctx, model.CreateTicketInput{
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return
}
