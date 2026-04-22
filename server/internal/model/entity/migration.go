// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Migration is the golang structure for table migration.
type Migration struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	Version    string `json:"version"    orm:"version"     description:""` //
	Name       string `json:"name"       orm:"name"        description:""` //
	ExecutedAt int64  `json:"executedAt" orm:"executed_at" description:""` //
	Checksum   string `json:"checksum"   orm:"checksum"    description:""` //
	Success    int    `json:"success"    orm:"success"     description:""` //
}
