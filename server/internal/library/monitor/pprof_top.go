// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package monitor

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sort"
	"strings"
	"sync"
	"time"

	goprofile "github.com/google/pprof/profile"
)

// PprofTopItem 函数级性能数据项
type PprofTopItem struct {
	Func    string  `json:"func"`    // 函数名（包.函数）
	File    string  `json:"file"`    // 源文件:行号
	Flat    string  `json:"flat"`    // 自身耗时/内存
	FlatPct float64 `json:"flatPct"` // 占比 %
	Cum     string  `json:"cum"`     // 累计耗时/内存
	CumPct  float64 `json:"cumPct"`  // 累计占比 %
}

// PprofTopResult 采样结果
type PprofTopResult struct {
	CpuTop    []PprofTopItem `json:"cpuTop"`
	MemTop    []PprofTopItem `json:"memTop"`
	CpuTime   string         `json:"cpuTime"`   // CPU 采样时长
	Timestamp string         `json:"timestamp"`  // 采样时间
}

var cpuMutex sync.Mutex

// GetPprofTop 获取 CPU + 内存热点 Top N
func GetPprofTop(cpuSeconds, limit int) (*PprofTopResult, error) {
	if cpuSeconds <= 0 {
		cpuSeconds = 3
	}
	if cpuSeconds > 30 {
		cpuSeconds = 30
	}
	if limit <= 0 {
		limit = 15
	}

	result := &PprofTopResult{
		CpuTime:   fmt.Sprintf("%ds", cpuSeconds),
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 1. 内存热点（不阻塞，直接获取当前堆快照）
	runtime.GC() // 触发 GC 获取更准确的堆数据
	memTop, err := getMemoryTop(limit)
	if err == nil {
		result.MemTop = memTop
	}

	// 2. CPU 热点（阻塞 N 秒采样）
	cpuTop, err := getCPUTop(cpuSeconds, limit)
	if err == nil {
		result.CpuTop = cpuTop
	}

	return result, nil
}

// getMemoryTop 获取内存分配热点
func getMemoryTop(limit int) ([]PprofTopItem, error) {
	var buf bytes.Buffer
	p := pprof.Lookup("heap")
	if p == nil {
		return nil, fmt.Errorf("heap profile not found")
	}
	if err := p.WriteTo(&buf, 0); err != nil {
		return nil, fmt.Errorf("write heap profile: %w", err)
	}

	prof, err := goprofile.Parse(&buf)
	if err != nil {
		return nil, fmt.Errorf("parse heap profile: %w", err)
	}

	return extractTop(prof, "alloc_space", limit), nil
}

// getCPUTop 采样 CPU profile
func getCPUTop(seconds, limit int) ([]PprofTopItem, error) {
	if !cpuMutex.TryLock() {
		return nil, fmt.Errorf("CPU profiling already in progress")
	}
	defer cpuMutex.Unlock()

	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		return nil, fmt.Errorf("start CPU profile: %w", err)
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	pprof.StopCPUProfile()

	prof, err := goprofile.Parse(&buf)
	if err != nil {
		return nil, fmt.Errorf("parse CPU profile: %w", err)
	}

	return extractTop(prof, "cpu", limit), nil
}

// extractTop 从 profile 中提取 Top N 函数
func extractTop(prof *goprofile.Profile, sampleType string, limit int) []PprofTopItem {
	// 确定使用哪个 sample type index
	sampleIdx := 0
	for i, st := range prof.SampleType {
		if strings.Contains(st.Type, sampleType) || strings.Contains(st.Type, "samples") || strings.Contains(st.Type, "cpu") {
			sampleIdx = i
			break
		}
		// 对于内存，优先用 alloc_space
		if sampleType == "alloc_space" && (st.Type == "alloc_space" || st.Type == "alloc_objects") {
			sampleIdx = i
			if st.Type == "alloc_space" {
				break
			}
		}
	}

	// 按函数聚合
	type funcStat struct {
		name string
		file string
		flat int64
		cum  int64
	}
	funcMap := make(map[string]*funcStat)
	var total int64

	for _, sample := range prof.Sample {
		value := sample.Value[sampleIdx]
		total += value

		if len(sample.Location) == 0 {
			continue
		}

		// Flat: 只统计栈顶函数
		topLoc := sample.Location[0]
		if len(topLoc.Line) > 0 {
			fn := topLoc.Line[0].Function
			if fn != nil {
				key := fn.Name
				if _, ok := funcMap[key]; !ok {
					file := fmt.Sprintf("%s:%d", filepath.Base(fn.Filename), topLoc.Line[0].Line)
					funcMap[key] = &funcStat{name: shortenFuncName(fn.Name), file: file}
				}
				funcMap[key].flat += value
			}
		}

		// Cum: 统计所有出现过的函数
		seen := make(map[string]bool)
		for _, loc := range sample.Location {
			for _, line := range loc.Line {
				if line.Function != nil && !seen[line.Function.Name] {
					seen[line.Function.Name] = true
					key := line.Function.Name
					if _, ok := funcMap[key]; !ok {
						file := fmt.Sprintf("%s:%d", filepath.Base(line.Function.Filename), line.Line)
						funcMap[key] = &funcStat{name: shortenFuncName(line.Function.Name), file: file}
					}
					funcMap[key].cum += value
				}
			}
		}
	}

	// 排序
	items := make([]*funcStat, 0, len(funcMap))
	for _, fs := range funcMap {
		items = append(items, fs)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].flat > items[j].flat
	})

	// 截取 Top N
	if len(items) > limit {
		items = items[:limit]
	}

	// 判断单位
	isMemory := sampleType == "alloc_space"
	unit := ""
	for _, st := range prof.SampleType {
		if st.Type == sampleType || strings.Contains(st.Type, sampleType) {
			unit = st.Unit
			break
		}
	}

	result := make([]PprofTopItem, 0, len(items))
	for _, fs := range items {
		flatPct := float64(0)
		cumPct := float64(0)
		if total > 0 {
			flatPct = float64(fs.flat) * 100 / float64(total)
			cumPct = float64(fs.cum) * 100 / float64(total)
		}

		result = append(result, PprofTopItem{
			Func:    fs.name,
			File:    fs.file,
			Flat:    formatValue(fs.flat, unit, isMemory),
			FlatPct: pprofRound2(flatPct),
			Cum:     formatValue(fs.cum, unit, isMemory),
			CumPct:  pprofRound2(cumPct),
		})
	}

	return result
}

// shortenFuncName 缩短函数名（去掉长路径前缀）
func shortenFuncName(name string) string {
	// github.com/xxx/yyy/pkg.Func -> pkg.Func
	if idx := strings.LastIndex(name, "/"); idx >= 0 {
		name = name[idx+1:]
	}
	return name
}

// formatValue 格式化值
func formatValue(v int64, unit string, isMemory bool) string {
	if isMemory || unit == "bytes" {
		return pprofFormatBytes(v)
	}
	// CPU: nanoseconds 或 count
	if unit == "nanoseconds" {
		if v >= 1e9 {
			return fmt.Sprintf("%.2fs", float64(v)/1e9)
		}
		if v >= 1e6 {
			return fmt.Sprintf("%.1fms", float64(v)/1e6)
		}
		if v >= 1e3 {
			return fmt.Sprintf("%.0fus", float64(v)/1e3)
		}
		return fmt.Sprintf("%dns", v)
	}
	// count
	return fmt.Sprintf("%d", v)
}

func pprofFormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func pprofRound2(v float64) float64 {
	return float64(int(v*100)) / 100
}
