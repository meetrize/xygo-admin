// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package cache

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

// cache 缓存驱动
var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

// Init 初始化缓存适配器
func Init(ctx context.Context) {
	var adapter gcache.Adapter

	// 从配置文件读取缓存类型（redis/memory）
	cacheType := g.Cfg().MustGet(ctx, "cache.adapter").String()

	switch cacheType {
	case "redis":
		// 使用 Redis 适配器
		adapter = gcache.NewAdapterRedis(g.Redis())
		g.Log().Info(ctx, "缓存适配器：Redis")
	default:
		// 默认使用内存适配器
		adapter = gcache.NewAdapterMemory()
		g.Log().Info(ctx, "缓存适配器：Memory（内存）")
	}

	// 数据库缓存，默认和通用缓存驱动一致
	g.DB().GetCache().SetAdapter(adapter)

	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(adapter)
}
