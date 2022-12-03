package cron

import "go_ticket/internal/service"

type sCron struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sCron {
	return &sCron{}
}

//func (s *sCron) StartCron() {
//
//}
