package request

import (
	"github.com/richardgong1987/server/model/common/request"
	"github.com/richardgong1987/server/model/system"
)

type SysApiTokenSearch struct {
	system.SysApiToken
	request.PageInfo
	Status *bool `json:"status" form:"status"`
}
