// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package queue

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DiskDriver 基于文件的简易队列驱动（开发/单机兜底用）
type DiskDriver struct {
	basePath string
	mu       sync.Mutex
}

// NewDiskDriver 创建磁盘驱动
func NewDiskDriver(basePath string) *DiskDriver {
	_ = os.MkdirAll(basePath, 0755)
	return &DiskDriver{basePath: basePath}
}

func (d *DiskDriver) filePath(topic string) string {
	return filepath.Join(d.basePath, topic+".jsonl")
}

func (d *DiskDriver) Push(_ context.Context, topic string, msg *Message) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(d.filePath(topic), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s\n", data)
	return err
}

func (d *DiskDriver) Pop(_ context.Context, topic string, timeout time.Duration) (*Message, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	path := d.filePath(topic)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// 没有文件，等一会
			d.mu.Unlock()
			time.Sleep(timeout)
			d.mu.Lock()
			return nil, nil
		}
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	f.Close()

	if len(lines) == 0 {
		d.mu.Unlock()
		time.Sleep(timeout)
		d.mu.Lock()
		return nil, nil
	}

	// 取第一行
	first := lines[0]
	rest := lines[1:]

	// 重写文件（去掉第一行）
	wf, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	w := bufio.NewWriter(wf)
	for _, line := range rest {
		fmt.Fprintf(w, "%s\n", line)
	}
	w.Flush()
	wf.Close()

	var msg Message
	if err := json.Unmarshal([]byte(first), &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func (d *DiskDriver) Len(_ context.Context, topic string) (int64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	f, err := os.Open(d.filePath(topic))
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	defer f.Close()

	var count int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			count++
		}
	}
	return count, nil
}

func (d *DiskDriver) Close() error {
	return nil
}
