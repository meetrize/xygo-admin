// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// =================================================================================
// 审计日志 Logic 层
// =================================================================================

package log

import (
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"xygo/internal/dao"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/adminin"
	"xygo/internal/service"
)

func init() {
	service.RegisterAdminLog(NewAdminLog())
}

// sAdminLog 审计日志 logic 实现
type sAdminLog struct{}

func NewAdminLog() *sAdminLog {
	return &sAdminLog{}
}

// ===================== 登录日志 =====================

// RecordLoginLog 记录登录日志
func (s *sAdminLog) RecordLoginLog(ctx context.Context, log *entity.AdminLoginLog) {
	// 自动填充地理位置
	if log.Location == "" {
		log.Location = GetLocationByIP(log.Ip)
	}
	// 排除 id 字段，让数据库自增分配（PG 不会把 id=0 当自增处理）
	_, err := dao.AdminLoginLog.Ctx(ctx).FieldsEx("id").Data(log).Insert()
	if err != nil {
		g.Log().Errorf(ctx, "记录登录日志失败: %v", err)
	}
}

// LoginLogList 登录日志列表
func (s *sAdminLog) LoginLogList(ctx context.Context, in *adminin.LoginLogListInp) (list []adminin.LoginLogItem, total int, err error) {
	m := dao.AdminLoginLog.Ctx(ctx)

	// 筛选条件
	if in.Username != "" {
		m = m.WhereLike(dao.AdminLoginLog.Columns().Username, "%"+in.Username+"%")
	}
	if in.Ip != "" {
		m = m.WhereLike(dao.AdminLoginLog.Columns().Ip, "%"+in.Ip+"%")
	}
	if in.Status != -1 {
		m = m.Where(dao.AdminLoginLog.Columns().Status, in.Status)
	}
	if len(in.DateRange) == 2 && in.DateRange[0] != "" && in.DateRange[1] != "" {
		m = m.WhereBetween(dao.AdminLoginLog.Columns().CreatedAt, in.DateRange[0], in.DateRange[1]+" 23:59:59")
	}

	// 总数
	total, err = m.Count()
	if err != nil {
		return
	}

	// 分页查询
	var entities []*entity.AdminLoginLog
	err = m.OrderDesc(dao.AdminLoginLog.Columns().Id).
		Page(in.Page, in.PageSize).
		Scan(&entities)
	if err != nil {
		return
	}

	// 转换
	list = make([]adminin.LoginLogItem, 0, len(entities))
	for _, e := range entities {
		item := adminin.LoginLogItem{
			Id:       uint(e.Id),
			UserId:   uint(e.UserId),
			Username: e.Username,
			Ip:       e.Ip,
			Location: e.Location,
			Browser:  e.Browser,
			Os:       e.Os,
			Status:   e.Status,
			Message:  e.Message,
		}
		if e.CreatedAt > 0 {
			item.CreatedAt = time.Unix(int64(e.CreatedAt), 0).Format("2006-01-02 15:04:05")
		}
		list = append(list, item)
	}
	return
}

// LoginLogDelete 删除登录日志
func (s *sAdminLog) LoginLogDelete(ctx context.Context, in *adminin.LoginLogDeleteInp) (err error) {
	_, err = dao.AdminLoginLog.Ctx(ctx).WhereIn(dao.AdminLoginLog.Columns().Id, in.Ids).Delete()
	return
}

// LoginLogClear 清空登录日志
func (s *sAdminLog) LoginLogClear(ctx context.Context) (err error) {
	_, err = dao.AdminLoginLog.Ctx(ctx).Where("1=1").Delete()
	return
}

// ===================== 操作日志 =====================

// RecordOperationLog 记录操作日志
func (s *sAdminLog) RecordOperationLog(ctx context.Context, log *entity.AdminOperationLog) {
	// 自动填充地理位置
	if log.Location == "" {
		log.Location = GetLocationByIP(log.Ip)
	}
	// 截断过长的请求体和响应体（避免存储过大）
	if len(log.RequestBody) > 5000 {
		log.RequestBody = log.RequestBody[:5000] + "...(truncated)"
	}
	if len(log.ResponseBody) > 5000 {
		log.ResponseBody = log.ResponseBody[:5000] + "...(truncated)"
	}

	// 排除 id 字段，让数据库自增分配（PG 不会把 id=0 当自增处理）
	_, err := dao.AdminOperationLog.Ctx(ctx).FieldsEx("id").Data(log).Insert()
	if err != nil {
		g.Log().Errorf(ctx, "记录操作日志失败: %v", err)
	}
}

// OperationLogList 操作日志列表
func (s *sAdminLog) OperationLogList(ctx context.Context, in *adminin.OperationLogListInp) (list []adminin.OperationLogItem, total int, err error) {
	m := dao.AdminOperationLog.Ctx(ctx)

	// 筛选条件
	if in.Username != "" {
		m = m.WhereLike(dao.AdminOperationLog.Columns().Username, "%"+in.Username+"%")
	}
	if in.Module != "" {
		m = m.WhereLike(dao.AdminOperationLog.Columns().Module, "%"+in.Module+"%")
	}
	if in.Status != -1 {
		m = m.Where(dao.AdminOperationLog.Columns().Status, in.Status)
	}
	if len(in.DateRange) == 2 && in.DateRange[0] != "" && in.DateRange[1] != "" {
		m = m.WhereBetween(dao.AdminOperationLog.Columns().CreatedAt, in.DateRange[0], in.DateRange[1]+" 23:59:59")
	}

	// 总数
	total, err = m.Count()
	if err != nil {
		return
	}

	// 分页查询（列表不返回 request_body 和 response_body，减少数据量）
	var entities []*entity.AdminOperationLog
	err = m.FieldsEx(dao.AdminOperationLog.Columns().RequestBody, dao.AdminOperationLog.Columns().ResponseBody).
		OrderDesc(dao.AdminOperationLog.Columns().Id).
		Page(in.Page, in.PageSize).
		Scan(&entities)
	if err != nil {
		return
	}

	// 转换
	list = make([]adminin.OperationLogItem, 0, len(entities))
	for _, e := range entities {
		item := adminin.OperationLogItem{
			Id:           uint(e.Id),
			UserId:       uint(e.UserId),
			Username:     e.Username,
			Module:       e.Module,
			Title:        e.Title,
			Method:       e.Method,
			Url:          e.Url,
			Ip:           e.Ip,
			Location:     e.Location,
			ErrorMessage: e.ErrorMessage,
			Status:       e.Status,
			Elapsed:      uint(e.Elapsed),
		}
		if e.CreatedAt > 0 {
			item.CreatedAt = time.Unix(int64(e.CreatedAt), 0).Format("2006-01-02 15:04:05")
		}
		list = append(list, item)
	}
	return
}

// OperationLogDetail 操作日志详情
func (s *sAdminLog) OperationLogDetail(ctx context.Context, in *adminin.OperationLogDetailInp) (out *adminin.OperationLogItem, err error) {
	var e *entity.AdminOperationLog
	err = dao.AdminOperationLog.Ctx(ctx).Where(dao.AdminOperationLog.Columns().Id, in.Id).Scan(&e)
	if err != nil {
		return
	}
	if e == nil {
		return
	}

	out = &adminin.OperationLogItem{
		Id:           uint(e.Id),
		UserId:       uint(e.UserId),
		Username:     e.Username,
		Module:       e.Module,
		Title:        e.Title,
		Method:       e.Method,
		Url:          e.Url,
		Ip:           e.Ip,
		Location:     e.Location,
		RequestBody:  e.RequestBody,
		ResponseBody: e.ResponseBody,
		ErrorMessage: e.ErrorMessage,
		Status:       e.Status,
		Elapsed:      uint(e.Elapsed),
	}
	if e.CreatedAt > 0 {
		out.CreatedAt = time.Unix(int64(e.CreatedAt), 0).Format("2006-01-02 15:04:05")
	}
	return
}

// OperationLogDelete 删除操作日志
func (s *sAdminLog) OperationLogDelete(ctx context.Context, in *adminin.OperationLogDeleteInp) (err error) {
	_, err = dao.AdminOperationLog.Ctx(ctx).WhereIn(dao.AdminOperationLog.Columns().Id, in.Ids).Delete()
	return
}

// OperationLogClear 清空操作日志
func (s *sAdminLog) OperationLogClear(ctx context.Context) (err error) {
	_, err = dao.AdminOperationLog.Ctx(ctx).Where("1=1").Delete()
	return
}

// ParseUserAgent 解析 User-Agent 获取浏览器和操作系统信息（简单实现）
func ParseUserAgent(ua string) (browser, os string) {
	ua = strings.ToLower(ua)

	// 浏览器检测
	switch {
	case strings.Contains(ua, "edg"):
		browser = "Edge"
	case strings.Contains(ua, "chrome") && !strings.Contains(ua, "edg"):
		browser = "Chrome"
	case strings.Contains(ua, "firefox"):
		browser = "Firefox"
	case strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome"):
		browser = "Safari"
	case strings.Contains(ua, "opera") || strings.Contains(ua, "opr"):
		browser = "Opera"
	case strings.Contains(ua, "msie") || strings.Contains(ua, "trident"):
		browser = "IE"
	default:
		browser = "Unknown"
	}

	// 操作系统检测
	switch {
	case strings.Contains(ua, "windows"):
		os = "Windows"
	case strings.Contains(ua, "mac os") || strings.Contains(ua, "macintosh"):
		os = "macOS"
	case strings.Contains(ua, "linux"):
		if strings.Contains(ua, "android") {
			os = "Android"
		} else {
			os = "Linux"
		}
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		os = "iOS"
	default:
		os = "Unknown"
	}

	return
}

// GetClientIP 获取客户端IP（支持代理头）
func GetClientIP(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}

	// 优先从代理头获取
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For 可能包含多个 IP，取第一个
		parts := strings.Split(ip, ",")
		ip = strings.TrimSpace(parts[0])
	} else {
		ip = r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.GetClientIp()
		}
	}

	// 将 IPv6 回环地址转为 IPv4
	if ip == "::1" || ip == "[::1]" {
		ip = "127.0.0.1"
	}

	return ip
}

// GetLocationByIP 根据IP获取地理位置（简单实现）
// 本地/内网IP返回"本机"，外网IP暂时返回空（后续可接入IP定位服务）
func GetLocationByIP(ip string) string {
	if ip == "" {
		return ""
	}
	// 本地/内网IP
	if ip == "127.0.0.1" || ip == "::1" || ip == "localhost" ||
		strings.HasPrefix(ip, "192.168.") ||
		strings.HasPrefix(ip, "10.") ||
		strings.HasPrefix(ip, "172.16.") ||
		strings.HasPrefix(ip, "172.17.") ||
		strings.HasPrefix(ip, "172.18.") ||
		strings.HasPrefix(ip, "172.19.") ||
		strings.HasPrefix(ip, "172.2") ||
		strings.HasPrefix(ip, "172.3") {
		return "本机"
	}
	// TODO: 接入 IP 定位服务（如 ip2region、GeoIP 等）
	return ""
}

// NowTime 获取当前时间戳（秒级 Unix 时间戳）
func NowTime() uint64 {
	return uint64(gtime.Now().Unix())
}
