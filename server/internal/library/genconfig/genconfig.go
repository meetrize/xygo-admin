// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// Package genconfig 代码生成器配置管理
package genconfig

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

// CrudTemplate CRUD 模板配置
type CrudTemplate struct {
	Group          string `json:"group"`
	TemplatePath   string `json:"templatePath"`
	ApiPath        string `json:"apiPath"`
	ControllerPath string `json:"controllerPath"`
	LogicPath      string `json:"logicPath"`
	InputPath      string `json:"inputPath"`
	SqlPath        string `json:"sqlPath"`
	WebApiPath     string `json:"webApiPath"`
	WebViewsPath   string `json:"webViewsPath"`
}

// GenConfig 代码生成器配置
type GenConfig struct {
	AllowedIPs    []string       `json:"allowedIPs"`
	SelectDbs     []string       `json:"selectDbs"`
	DisableTables []string       `json:"disableTables"`
	Application   AppConfig      `json:"application"`
}

// AppConfig 应用配置
type AppConfig struct {
	Crud CrudConfig `json:"crud"`
}

// CrudConfig CRUD 配置
type CrudConfig struct {
	Templates []CrudTemplate `json:"templates"`
}

var (
	config     *GenConfig
	configOnce sync.Once
)

// GetConfig 获取代码生成器配置（单例）
func GetConfig(ctx context.Context) *GenConfig {
	configOnce.Do(func() {
		config = &GenConfig{}
		val, err := g.Cfg().Get(ctx, "xygen")
		if err != nil || val.IsNil() {
			g.Log().Warning(ctx, "xygen config not found, using defaults")
			config = defaultConfig()
			return
		}
		if err := val.Scan(config); err != nil {
			g.Log().Warningf(ctx, "parse xygen config error: %v, using defaults", err)
			config = defaultConfig()
			return
		}
		// 确保至少有一个模板配置
		if len(config.Application.Crud.Templates) == 0 {
			config.Application.Crud.Templates = defaultConfig().Application.Crud.Templates
		}
	})
	return config
}

// GetDefaultTemplate 获取默认模板配置（group=default）
func GetDefaultTemplate(ctx context.Context) *CrudTemplate {
	cfg := GetConfig(ctx)
	for i, t := range cfg.Application.Crud.Templates {
		if t.Group == "default" {
			return &cfg.Application.Crud.Templates[i]
		}
	}
	if len(cfg.Application.Crud.Templates) > 0 {
		return &cfg.Application.Crud.Templates[0]
	}
	tpl := defaultConfig().Application.Crud.Templates[0]
	return &tpl
}

// IsTableDisabled 检查表是否被禁用
func IsTableDisabled(ctx context.Context, tableName string) bool {
	cfg := GetConfig(ctx)
	for _, t := range cfg.DisableTables {
		if t == tableName {
			return true
		}
	}
	return false
}

// ResetConfig 重置配置（用于热更新场景）
func ResetConfig() {
	configOnce = sync.Once{}
	config = nil
}

// defaultConfig 默认配置
func defaultConfig() *GenConfig {
	return &GenConfig{
		AllowedIPs:    []string{"127.0.0.1", "*"},
		SelectDbs:     []string{"default"},
		DisableTables: []string{"xy_sys_gen_codes", "xy_sys_gen_codes_column"},
		Application: AppConfig{
			Crud: CrudConfig{
				Templates: []CrudTemplate{
					{
						Group:          "default",
						TemplatePath:   "./resource/generate/default",
						ApiPath:        "api/admin",
						ControllerPath: "internal/controller/admin",
						LogicPath:      "internal/logic",
						InputPath:      "internal/model/input/adminin",
						SqlPath:        "resource/sql/generate",
						WebApiPath:     "../web/src/api/backend",
						WebViewsPath:   "../web/src/views/backend",
					},
				},
			},
		},
	}
}
