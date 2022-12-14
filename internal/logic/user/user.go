package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"go_ticket/internal/dao"
	"go_ticket/internal/errorcode"
	"go_ticket/internal/model"
	"go_ticket/internal/model/entity"
	"go_ticket/internal/service"
	"go_ticket/utility/token"
	"go_ticket/utility/utils"
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
	glog.Info(ctx, "注册输入内容:", in)
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *entity.User
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}
		// 判断用户是否已经存在
		err := s.CheckUsernameUnique(ctx, in.Name)
		if err != nil {
			return err
		}
		user.Password = s.EncryptPassword(user.Password)
		_, err = dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

/**
* CheckUsernameUnique
*  @Description: 检测用户名是否存在
*  @receiver s
*  @param ctx
*  @param name
*  @return error
 */
func (s *sUser) CheckUsernameUnique(ctx context.Context, name string) error {
	n, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Name, name).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`用户"%s"已被占用`, name)
	}
	return nil
}

// EncryptPassword 密码进行加密
func (s *sUser) EncryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, name, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Name:     name,
		dao.User.Columns().Password: s.EncryptPassword(password),
	}).Scan(&user)
	return

}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error) {
	out = &model.UserLoginOutput{}
	userEntity, err := s.GetUserByPassportAndPassword(ctx, in.Name, in.Password)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		glog.Error(ctx, "账号密码错误", in)
		return nil, gerror.New("账号或密码错误")
	}
	// 自动映射同样的字段的值
	out.User = &model.UserInfoOutput{}
	err = utils.MyCopy(ctx, out.User, userEntity)
	if err != nil {
		return nil, err
	}
	return

}

func (s *sUser) Logout(ctx context.Context) error {
	userToken1 := service.Context().Get(ctx)
	if userToken1 == nil {
		glog.Error(ctx, " 获取用户请求token失败")
		return gerror.New("未获取到token信息")
	}
	//myRequestToken, err := token.Instance().DecrypToken(ctx, userToken)
	//if err != nil {
	//	glog.Error(ctx, "解析Token失败", err)
	//	return gerror.New("解析token失败")
	//}
	glog.Info(ctx, "用户注销：", userToken1)
	err := token.Instance().RemoveCache(ctx, token.CacheKeyPrefix+userToken1.Token.UserKey)
	if err != nil {
		glog.Error(ctx, "删除token失败", err)
		return gerror.New("用户注销失败")
	}
	return err
}

func (s *sUser) UserList(ctx context.Context, in model.UserListInput) (out *model.UserListOutput, err error) {
	userList, err := dao.User.Ctx(ctx).All()
	glog.Info(ctx, "logic获取到用户model返回的list信息: ", userList)
	if err != nil {
		glog.Error(ctx, "logic没有获取到用户model list信息", err, userList)
		return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
	}
	out = &model.UserListOutput{}
	err = userList.Structs(&out.List)
	if err != nil {
		glog.Error(ctx, err, userList)
		return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
	}
	glog.Info(ctx, "logic获取到用户UserListOutput返回的list信息", err, out)
	return out, nil
}
