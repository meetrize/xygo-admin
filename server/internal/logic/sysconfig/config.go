// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package sysconfig

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/model/entity"
	"xygo/internal/service"
)

type sSysConfig struct{}

func init() {
	service.RegisterSysConfig(New())
}

// New 构造
func New() *sSysConfig {
	return &sSysConfig{}
}

// InitConfig 启动初始化
func (s *sSysConfig) InitConfig(ctx context.Context) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Fatalf(ctx, "InitConfig fail: %+v", err)
	}
}

// LoadConfig 预加载配置（预留：可在此将部分分组注入业务组件）
func (s *sSysConfig) LoadConfig(ctx context.Context) error {
	// 目前仅做占位，后续需要把基础配置写入组件时可在此扩展
	return nil
}

// GetConfigByGroup 获取指定分组配置
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, group string) (map[string]string, error) {
	if group == "" {
		return nil, gerror.New("分组不能为空")
	}

	var items []entity.SysConfig
	if err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Group, group).
		OrderAsc(dao.SysConfig.Columns().Group).OrderAsc("sort").OrderAsc("id").
		Scan(&items); err != nil {
		return nil, gerror.Wrapf(err, "获取配置分组[%v]失败，请稍后重试", group)
	}
	if len(items) == 0 {
		return nil, gerror.NewCode(consts.CodeDataNotFound, "分组不存在或无配置项")
	}

	m := make(map[string]string, len(items))
	for _, it := range items {
		m[it.Key] = it.Value
	}
	return m, nil
}

