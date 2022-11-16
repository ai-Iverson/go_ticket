package token

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
