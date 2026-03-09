// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package field

import (
	"sync"
)

// ResourceField 资源字段定义
type ResourceField struct {
	Name        string `json:"name"`        // 字段名称
	Label       string `json:"label"`       // 字段显示名称
	Type        string `json:"type"`        // 字段类型（string, number, date, array...）
	Source      string `json:"source"`      // 来源表（主表或关联表）
	IsVirtual   bool   `json:"isVirtual"`   // 是否虚拟字段（来自JOIN或计算）
	IsSensitive bool   `json:"isSensitive"` // 是否敏感字段
}

// Resource 资源定义
type Resource struct {
	Code   string          `json:"code"`   // 资源编码（表名）
	Label  string          `json:"label"`  // 资源显示名（中文）
	Module string          `json:"module"` // 所属模块（admin, system, business...）
	Fields []ResourceField `json:"fields"` // 字段列表
}

var (
	resourceMap = make(map[string]*Resource)
	mu          sync.RWMutex
)

// Register 注册资源（在各个field文件的init中调用）
func Register(code, label, module string, fields []ResourceField) {
	mu.Lock()
	defer mu.Unlock()

	resourceMap[code] = &Resource{
		Code:   code,
		Label:  label,
		Module: module,
		Fields: fields,
	}
}

// Get 获取指定资源定义
func Get(code string) *Resource {
	mu.RLock()
	defer mu.RUnlock()
	return resourceMap[code]
}

// GetFields 获取指定资源的字段列表
func GetFields(code string) []ResourceField {
	resource := Get(code)
	if resource == nil {
		return []ResourceField{}
	}
	return resource.Fields
}

// GetAll 获取所有已注册的资源
func GetAll() []*Resource {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]*Resource, 0, len(resourceMap))
	for _, res := range resourceMap {
		result = append(result, res)
	}
	return result
}

// GetAllCodes 获取所有资源编码列表
func GetAllCodes() []string {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]string, 0, len(resourceMap))
	for code := range resourceMap {
		result = append(result, code)
	}
	return result
}
