// 多租户控制器预埋（安装租户扩展后此文件会被覆盖）

package tenant

type ControllerV1 struct{}

func NewV1() *ControllerV1 { return &ControllerV1{} }
