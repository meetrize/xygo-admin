// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Migration is the golang structure of table xy_migration for DAO operations like Where/Data.
type Migration struct {
	g.Meta     `orm:"table:xy_migration, do:true"`
	Id         any //
	Version    any //
	Name       any //
	ExecutedAt any //
	Checksum   any //
	Success    any //
}
