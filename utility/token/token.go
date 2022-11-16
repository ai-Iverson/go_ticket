package token

//
//import (
//	"context"
//	"fmt"
//	jwt "github.com/gogf/gf-jwt/v2"
//	"github.com/gogf/gf/v2/crypto/gaes"
//	"github.com/gogf/gf/v2/encoding/gbase64"
//	"github.com/gogf/gf/v2/errors/gcode"
//	"github.com/gogf/gf/v2/errors/gerror"
//	"github.com/gogf/gf/v2/frame/g"
//	"github.com/gogf/gf/v2/net/ghttp"
//	"github.com/gogf/gf/v2/os/gtime"
//	"github.com/gogf/gf/v2/text/gstr"
//	"github.com/gogf/gf/v2/util/gconv"
//	v1 "go_ticket/api/v1"
//	"go_ticket/internal/service"
//	"net/http"
//	"time"
//)
//
//var authService *jwt.GfJWTMiddleware
//
//// 权限包管理
//func Auth() *jwt.GfJWTMiddleware {
//	return authService
//}
//
//// 初始化
//func init() {
//	auth := jwt.New(&jwt.GfJWTMiddleware{
//		//用户的领域名称，必传
//		Realm: "go_ticket",
//		// 签名算法
//		SigningAlgorithm: "HS256",
//		// 签名密钥
//		Key: []byte("go_ticket"),
//		// 时效
//		Timeout: time.Minute * 5,
//		// 	token过期后，可凭借旧token获取新token的刷新时间
//		MaxRefresh: time.Minute * 5,
//		// 身份验证的key值
//		IdentityKey: "id",
//		//token检索模式，用于提取token-> Authorization
//		// TokenLookup: "header: Authorization, query: token, cookie: jwt",
//		TokenLookup: "header: Authorization, query: token, cookie: jwt",
//		// token在请求头时的名称，默认值为Bearer.客户端在header中传入"Authorization":"token xxxxxx"
//		TokenHeadName: "Bearer",
//		TimeFunc:      time.Now,
//		// 用户标识 map  私有属性
//		// 根据登录信息对用户进行身份验证的回调函数
//		Authenticator: Authenticator,
//		// 处理不进行授权的逻辑
//		Unauthorized: Unauthorized,
//		//登录期间的设置私有载荷的函数，默认设置Authenticator函数回调的所有内容
//		PayloadFunc: PayloadFunc,
//		// 解析并设置用户身份信息，并设置身份信息至每次请求中
//		IdentityHandler: IdentityHandler,
//	})
//	authService = auth
//}
//func PayloadFunc(data interface{}) jwt.MapClaims {
//	claims := jwt.MapClaims{}
//	params := data.(map[string]interface{})
//	if len(params) > 0 {
//		for k, v := range params {
//			claims[k] = v
//		}
//	}
//	return claims
//}
//func IdentityHandler(ctx context.Context) interface{} {
//	claims := jwt.ExtractClaims(ctx)
//	fmt.Println(claims[authService.IdentityKey])
//	return claims[authService.IdentityKey]
//}
//
//func Unauthorized(ctx context.Context, code int, message string) {
//	r := g.RequestFromCtx(ctx)
//	r.Response.WriteJson(g.Map{
//		"code": code,
//		"msg":  message,
//	})
//	r.ExitAll()
//}
//func Authenticator(ctx context.Context) (interface{}, error) {
//	var (
//		apiReq     v1.UserRegisterReq
//		serviceReq = g.RequestFromCtx(ctx)
//	)
//	if err := serviceReq.Parse(&apiReq); err != nil {
//		return "", err
//	}
//	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
//		return "", err
//	}
//
//	// user {"id": 1, "username": "admin"}
//	user := service.User().CheckUserPassword(ctx, apiReq.UserName, apiReq.PassWord)
//	if user != nil {
//		return user, nil
//	}
//	return nil, jwt.ErrFailedAuthentication
//}
//
//// 权限中间件
//type middlewareService struct{}
//
//var middleware = middlewareService{}
//
//func Middleware() *middlewareService {
//	return &middleware
//}
//
//func (s *middlewareService) Auth(r *ghttp.Request) {
//	// GfJWTMiddleware gf jwt集成的中间件
//	// Auth是权限service中配置的gf jwt
//	Auth().MiddlewareFunc()(r)
//	r.Middleware.Next()
//}
//
//type DefaultHandlerRes struct {
//	ResultCode int         `json:"resultCode"    dc:"Error code"`
//	Message    string      `json:"message" dc:"Error message" d:""`
//	Data       interface{} `json:"data"    dc:"Result data for certain request according API definition"`
//}
//
//func (s *middlewareService) CustomResponse(r *ghttp.Request) {
//	r.Middleware.Next()
//
//	// There's custom buffer content, it then exits current handler.
//	if r.Response.BufferLength() > 0 {
//		return
//	}
//
//	var (
//		msg  string
//		err  = r.GetError()
//		res  = r.GetHandlerResponse()
//		code = gerror.Code(err)
//	)
//	if err != nil {
//		if code == gcode.CodeNil {
//			code = gcode.CodeInternalError
//		}
//		msg = err.Error()
//	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
//		msg = http.StatusText(r.Response.Status)
//		switch r.Response.Status {
//		case http.StatusNotFound:
//			code = gcode.CodeNotFound
//		case http.StatusForbidden:
//			code = gcode.CodeNotAuthorized
//		default:
//			code = gcode.CodeUnknown
//		}
//	} else {
//		code = gcode.New(200, "success", "")
//	}
//	r.Response.WriteJson(DefaultHandlerRes{
//		ResultCode: code.Code(),
//		Message:    msg,
//		Data:       res,
//	})
//}
//
///**
// * @description 解密token。token的生成规则是base64(gaes.Encrypt(base64(userKey)+TokenDelimiter+uuid))
//
// * @param userKey 用户的标识，一般使用用户名称或者用户的uuid
// * @param uuid 可以使用外部提供的uuid，如果为空，会重新生成
// **/
//func (m *MyToken) decryptToken(ctx context.Context, token string) (tokenDecrypted *MyRequestToken, err error) {
//	if token == "" {
//		return nil, errorCode.NewMyErr(ctx, errorCode.TokenEmpty)
//	}
//	token64, err := gbase64.Decode([]byte(token))
//	if err != nil {
//		return nil, errorCode.NewMyErr(ctx, errorCode.TokenInvalidError, token, err)
//	}
//	decryptTokenStr, err := gaes.Decrypt(token64, []byte(EncryptKey))
//	if err != nil {
//		return nil, errorCode.NewMyErr(ctx, errorCode.TokenInvalidError, token, err)
//	}
//	tokenArray := gstr.Split(string(decryptTokenStr), TokenDelimiter)
//	if len(tokenArray) < 2 {
//		return nil, errorCode.NewMyErr(ctx, errorCode.TokenInvalidError, token, err)
//	}
//	userKey, err := gbase64.Decode([]byte(tokenArray[0]))
//	if err != nil {
//		return nil, errorCode.NewMyErr(ctx, errorCode.TokenInvalidError, token, err)
//	}
//
//	return &MyRequestToken{string(userKey), tokenArray[1], token}, nil
//}
//
//func (m *MyToken) getAndFreshCacheToken(ctx context.Context, userKey string) (*MyCacheToken, error) {
//	cacheKey := CacheKeyPrefix + userKey
//
//	cacheToken, err := m.getCache(ctx, cacheKey)
//	if err != nil {
//		return nil, err
//	}
//
//	nowTime := gtime.Now().TimestampMilli()
//
//	// 需要进行缓存超时时间刷新
//	// cacheToken.NextFreshTime == 0, 表明是一个一次性的token
//	if gconv.Int64(cacheToken.NextFreshTime) == 0 || nowTime > gconv.Int64(cacheToken.NextFreshTime) {
//		cacheToken.CreatedAt = gtime.Now().TimestampMilli()
//		cacheToken.NextFreshTime = gtime.Now().TimestampMilli() + gconv.Int64(CacheMaxRefresh)
//		m.setCache(ctx, cacheKey, cacheToken)
//	}
//	return cacheToken, nil
//}
//
//func (m *MyToken) GenerateToken(ctx context.Context, userKey string, data interface{}) (*MyCacheToken, error) {
//	myRequestToke, err := m.EncryptToken(ctx, userKey, "")
//	if err != nil {
//		return nil, err
//	}
//
//	cacheKey := CacheKeyPrefix + userKey
//	nowTime := gtime.Now().TimestampMilli()
//	myCacheToken := &MyCacheToken{
//		Token:         myRequestToke.Token,
//		Uuid:          myRequestToke.Uuid,
//		UserKey:       userKey,
//		UserData:      data,
//		CreatedAt:     nowTime,
//		NextFreshTime: nowTime + gconv.Int64(CacheMaxRefresh),
//	}
//
//	err = m.setCache(ctx, cacheKey, myCacheToken)
//	if err != nil {
//		return nil, errorCode.NewMyErr(ctx, errorCode.MyInternalError, err)
//	}
//
//	return myCacheToken, nil
//}
