// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package logic

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"xygo/internal/service"
)

// UploadConfig 通用上传配置
type UploadConfig struct {
	Driver          string   // 默认驱动
	MaxSizeBytes    int64    // 单文件最大字节数，0 不限
	AllowedSuffixes []string // 允许的文件后缀（小写，无点），空表示不限制
	AllowedMimes    []string // 允许的 MIME（小写），空表示不限制
	LocalBaseDir    string   // 本地存储根目录（绝对或相对）
	// 预留：驱动相关参数，如 OSS/COS/Qiniu 等
}

// LoadUploadConfig 从配置组 oss 读取上传相关配置
func LoadUploadConfig(ctx context.Context) (*UploadConfig, error) {
	cfg, err := service.SysConfig().GetConfigByGroup(ctx, "oss")
	if err != nil {
		return nil, gerror.Wrap(err, "读取上传配置失败")
	}
	c := &UploadConfig{
		Driver: strings.ToLower(firstNonEmpty(cfg["upload_driver"], cfg["oss_driver"])),
	}
	if c.Driver == "" {
		c.Driver = "local"
	}
	if v := firstNonEmpty(cfg["upload_max_size"], cfg["oss.upload_max_size"]); v != "" {
		if bytes := parseSizeToBytes(v); bytes > 0 {
			c.MaxSizeBytes = bytes
		}
	}
	if v := firstNonEmpty(cfg["upload_allowed_suffixes"], cfg["oss.upload_allowed_suffixes"]); v != "" {
		c.AllowedSuffixes = splitAndLower(v)
	}
	if v := firstNonEmpty(cfg["upload_allowed_mime_types"], cfg["oss.upload_allowed_mime_types"]); v != "" {
		c.AllowedMimes = splitAndLower(v)
	}
	// 本地目录固定，不再从数据库读取
	c.LocalBaseDir = "resource/public/attachment/upload"
	return c, nil
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func splitAndLower(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		v := strings.ToLower(strings.TrimSpace(p))
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

// 支持类似 "10mb" "5m" "1024kb" "123456"（字节）
func parseSizeToBytes(v string) int64 {
	s := strings.ToLower(strings.TrimSpace(v))
	if s == "" {
		return 0
	}
	unit := int64(1)
	switch {
	case strings.HasSuffix(s, "kb"):
		unit = 1024
		s = strings.TrimSuffix(s, "kb")
	case strings.HasSuffix(s, "k"):
		unit = 1024
		s = strings.TrimSuffix(s, "k")
	case strings.HasSuffix(s, "mb"):
		unit = 1024 * 1024
		s = strings.TrimSuffix(s, "mb")
	case strings.HasSuffix(s, "m"):
		unit = 1024 * 1024
		s = strings.TrimSuffix(s, "m")
	case strings.HasSuffix(s, "gb"):
		unit = 1024 * 1024 * 1024
		s = strings.TrimSuffix(s, "gb")
	case strings.HasSuffix(s, "g"):
		unit = 1024 * 1024 * 1024
		s = strings.TrimSuffix(s, "g")
	default:
		// 纯数字按字节处理
	}
	n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0
	}
	return int64(n * float64(unit))
}

// NormalizeExt 获取扩展名（不含点，小写）
func NormalizeExt(filename string) string {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	return ext
}

// Sha1Bytes 计算 sha1
func Sha1Bytes(b []byte) string {
	h := sha1.Sum(b)
	return strings.ToLower(hex.EncodeToString(h[:]))
}
