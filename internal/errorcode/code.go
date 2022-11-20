package errorcode

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type MyCode struct {
	code    int
	message string // message 设计为i18n的key
	detail  MyCodeDetail
}
type MyCodeDetail struct {
	HttpCode int
}

func (c MyCode) MyDetail() MyCodeDetail {
	return c.detail
}

func (c MyCode) Code() int {
	return c.code
}

func (c MyCode) Message() string {
	return c.message
}

func (c MyCode) Detail() interface{} {
	return c.detail
}

func New(httpCode int, code int, message string) gcode.Code {
	return MyCode{
		code:    code,
		message: message,
		detail: MyCodeDetail{
			HttpCode: httpCode,
		},
	}
}

func NewMyErr(ctx context.Context, code gcode.Code, params ...interface{}) error {
	tfStr := g.I18n().Tf(ctx, code.Message(), params...)
	return gerror.NewCode(code, tfStr)
}

func MyWrapCode(ctx context.Context, code gcode.Code, err error, params ...interface{}) error {
	tfStr := g.I18n().Tf(ctx, code.Message(), params...)
	return gerror.WrapCode(code, err, tfStr)
}

var (
	// gf框架内置的，参见：github.com\gogf\gf\v2@v2.0.0-rc2\errors\gcode\gcode.go
	CodeNil           = New(200, -1, "")
	CodeNotFound      = New(404, 65, "Not Found")
	CodeInternalError = New(500, 50, "An error occurred internally")

	// 系统起始 10000
	MyInternalError = New(500, 10001, "{#myInternalError}")

	// token 20000起始
	AuthHeaderInvalidError     = New(401, 20001, `{#authHeaderInvalidError}`)
	NotSupportedCacheModeError = New(401, 20002, `{#notSupportedCacheModeError}`)
	TokenEmpty                 = New(401, 20003, `{#tokenEmpty}`)
	TokenKeyEmpty              = New(401, 20004, `{#tokenKeyEmpty}`)
	TokenInvalidError          = New(401, 20005, `{#tokenInvalidError}`)
	Unauthorized               = New(401, 20006, `{#unauthorized}`)
	AuthorizedFailed           = New(401, 20007, `{#authorizedFailed}`)

	//用户30000起始
	UserNotFound        = New(404, 30001, `{#userNotExists}`)
	LoginNameConflicted = New(403, 30002, `{#loginNameConflicted}`)
	PasswordError       = New(401, 30003, `{#passwordError}`)
	LoginFailed         = New(401, 30004, `{#loginFailed}`)
	LogoutFailed        = New(401, 30005, `{#logoutFailed}`)

	// 桌面40000起始
	DesktopNotFound = New(404, 40001, `{#desktopNotExists}`)
)
