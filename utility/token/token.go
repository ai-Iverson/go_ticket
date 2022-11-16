package token

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"go_ticket/internal/errorcode"
	"strings"
)

var (
	// 创建一个空的stranymap对象
	// 参数并发安装默认位false
	instances = gmap.NewStrAnyMap(true)
)

type MyToken struct {
	Timeout   int // 超时时间(毫秒)
	CacheMode int // 缓存类型: 内存1,redis2
}
type MyCacheToken struct {
	Token         string
	UserKey       string
	Uuid          string
	UserData      interface{}
	CreatedAt     int64 // Token 生成的时间
	NextFreshTime int64 // 下次token刷新时间, =0，一次性token
}

type MyRequestToken struct {
	UserKey string
	Uuid    string
	Token   string
}

// Instance 获取配置文件Token的配置信息
func Instance(name ...string) *MyToken {
	key := "default"
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	//GetOrSetFunc按键返回值，
	//如果回调函数'f'不存在，则使用返回值设置值
	//然后返回这个值。
	return instances.GetOrSetFuncLock(key, func() interface{} {
		// MustGet获取不到会报错
		//context.Background() 空值初始值
		timeout := g.Cfg().MustGet(context.Background(), "token.timeout", CacheTimeout).Int()
		cacheMode := g.Cfg().MustGet(context.Background(), "token.cacheMode", CacheModeCache).Int()
		token := &MyToken{
			Timeout:   timeout,
			CacheMode: cacheMode,
		}
		return token
	}).(*MyToken)
}

/**
*  GetRequestToken
*   @Description:获取请求header中的token,只支持http header中，Authorization: Bearer 类型
*   @param r
*   @return token
*   @return err
 */
func GetRequestToken(r *ghttp.Request) (token string, err error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if parts[0] != "Bearer" {
			return "", errorcode.NewMyErr(r.Context(), errorcode.AuthHeaderInvalidError, authHeader)
		}
		return parts[1], nil

	}
	return "", nil
}

/**
* ValidToken
*  @Description: 验证token
*  @param ctx
*  @param token
*  @return MyCacheToken
 */
//func (m *MyToken) ValidToken(ctx context.Context, token string) *MyCacheToken {
//
//}

/**
* EncryptToken
*  @Description:加密生成token.
               token的生成规则是base64(gaes.Encrypt(base64(userKey)+TokenDelimiter+uuid)): 其中TokenDelimiter默认为_;
               为什么要base64(userKey)，因为可能userKey包含_; 标准base64是使用 `数字`+`大小写字母`+`/`+`+`以及`=`组成
               解释：为什么还要对token进行加解密？答：加密因为token携带了userKey信息，且便于过滤掉不合法token；
*  @param ctx
*  @param userKey 用户唯一标识
*  @return *MyRequestToken
*  @return error
*/
func (m *MyToken) EncryptToken(ctx context.Context, userKey string) (*MyRequestToken, error) {
	if userKey == "" {
		return nil, errorcode.NewMyErr(ctx, errorcode.TokenEmpty)
	}
	uuid := guid.S()
	tokenstr := gbase64.EncodeToString([]byte(userKey)) + TokenDelimiter + uuid

	token, err := gaes.Encrypt([]byte(tokenstr), []byte(EncryptKey))
	if err != nil {
		return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
	}
	return &MyRequestToken{
		UserKey: userKey,
		Uuid:    uuid,
		Token:   gbase64.EncodeToString(token),
	}, nil
}
