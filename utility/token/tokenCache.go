package token

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"go_ticket/internal/errorcode"
	"time"
)

/**
* setCache
*  @Description: 设置缓存
*  @param cachekey
*  @param userCache
*  @return error
 */
func (m *MyToken) setCache(ctx context.Context, cachekey string, userCache *MyCacheToken) error {
	switch m.CacheMode {
	case CacheModeCache:
		err := gcache.Set(ctx, cachekey, userCache, gconv.Duration(m.Timeout)*time.Millisecond)
		if err != nil {
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
	case CacheModeRedis:
		cacheValueJson, err := gjson.Encode(userCache)
		if err != nil {
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
		_, err = g.Redis().Do(ctx, "SETEX", cachekey, m.Timeout/1000, cacheValueJson)
		if err != nil {
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
	default:
		return errorcode.NewMyErr(ctx, errorcode.NotSupportedCacheModeError, m.CacheMode)
	}
	return nil
}

/**
* getCache
*  @Description: 获取缓存中的token
*  @param cacheToken
*  @return *MyCacheToken
*  @return error
 */
func (m *MyToken) getCache(ctx context.Context, cacheKey string) (myCacheToken *MyCacheToken, err error) {
	//  根据token存储的模式来获取
	switch m.CacheMode {
	case CacheModeCache:
		//  从缓存中获取
		cacheVaule, err := gcache.Get(ctx, cacheKey)
		if err != nil {
			return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
		if cacheVaule.IsNil() {
			return nil, errorcode.NewMyErr(ctx, errorcode.Unauthorized)
		}
		myCacheToken = &MyCacheToken{}
		err = gconv.Struct(cacheVaule, myCacheToken)
		if err != nil {
			return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
	case CacheModeRedis:
		userCacheJson, err := g.Redis().Do(ctx, "Get", cacheKey)
		if err != nil {
			return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
		if userCacheJson.IsNil() {
			return nil, errorcode.NewMyErr(ctx, errorcode.Unauthorized)
		}
		myCacheToken = &MyCacheToken{}
		err = gjson.DecodeTo(userCacheJson, myCacheToken)
		if err != nil {
			return nil, errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
	default:
		return nil, errorcode.NewMyErr(ctx, errorcode.NotSupportedCacheModeError, m.CacheMode)
	}
	return myCacheToken, nil
}

/**
* RemoveCache
*  @Description: 删除缓存
*  @param ctx
*  @param cacheKey
*  @return error
 */
func (m *MyToken) RemoveCache(ctx context.Context, cacheKey string) error {
	switch m.CacheMode {
	case CacheModeCache:
		_, err := gcache.Remove(ctx, cacheKey)
		if err != nil {
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
	case CacheModeRedis:
		_, err := g.Redis().Do(ctx, "DEL", cacheKey)
		if err != nil {
			return errorcode.NewMyErr(ctx, errorcode.MyInternalError, err)
		}
	default:
		return errorcode.NewMyErr(ctx, errorcode.NotSupportedCacheModeError, m.CacheMode)
	}

	return nil
}
