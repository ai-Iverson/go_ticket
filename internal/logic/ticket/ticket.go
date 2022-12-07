package ticket

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"go_ticket/internal/dao"
	"go_ticket/internal/model"
	"go_ticket/internal/model/entity"
	"go_ticket/internal/service"
)

type sTicket struct {
}

func init() {
	service.RegisterTicket(New())
}

func New() *sTicket {
	return &sTicket{}
}

func (s *sTicket) CreateTicket(ctx context.Context, in model.CreateTicketInput) (err error) {
	code, err := s.generateCode(ctx, in.Type)
	if err != nil {
		return err

	}

	creatUserId := service.Context().Get(ctx).Token.UserKey

	_, err = dao.Ticket.Ctx(ctx).Data(entity.Ticket{
		Code:        code,
		Title:       in.Title,
		Description: in.Description,
		CreatUserId: gconv.Int(creatUserId),
		Type:        in.Type,
		Status:      0,
	}).Insert()

	return err
}

func (s *sTicket) generateCode(ctx context.Context, types int) (code string, err error) {
	// 0 是问题工单
	var lastCode = &model.TickerOutput{}

	if types == 0 {
		sdf := dao.Ticket.Ctx(ctx).Where(dao.Ticket.Columns().Type, 0).Order("created_at desc")
		n, _ := sdf.Count()
		sdf.Limit(1).Scan(lastCode)

		if n == 0 {
			code = "I1"
			return code, nil
		}

		lastcodenum := gconv.Int(lastCode.Code[1:])
		codenum := lastcodenum + 1
		code = "I" + gconv.String(codenum)
		return code, nil

	} else if types == 1 {
		sdf := dao.Ticket.Ctx(ctx).Where(dao.Ticket.Columns().Type, 1).Order("created_at desc")
		n, _ := sdf.Count()
		sdf.Limit(1).Scan(lastCode)

		if n == 0 {
			code = "R1"
			return code, nil
		}

		glog.Info(ctx, lastCode.Code[1:])
		lastcodenum := gconv.Int(lastCode.Code[1:])
		codenum := lastcodenum + 1
		code = "R" + gconv.String(codenum)
		return code, nil

	} else {

		return code, gerror.New("工单类型不合法")
	}
}
