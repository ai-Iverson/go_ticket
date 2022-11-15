package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"go_ticket/internal/dao"
	"go_ticket/internal/model"
	"go_ticket/internal/model/entity"
	"go_ticket/internal/service"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.UserRegisterInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *entity.User
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}
		user.Password = s.EncryptPassword(user.Password)
		_, err := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

// EncryptPassword 密码进行加密
func (s *sUser) EncryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (user *entity.User, err error) {
	user, err = s.GetUserByNameAndPassword(ctx, in.Name, s.EncryptPassword(in.Password))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New(`用户名或密码错误`)
	}
	return

}

func (s *sUser) GetUserByNameAndPassword(ctx context.Context, name, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(gdb.Map{
		dao.User.Columns().Name:     name,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}
