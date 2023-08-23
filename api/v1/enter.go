package v1

import (
	"github.com/coder-th/ds-tools/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	BaseApi 	   system.BaseApi
}

var ApiGroupApp = new(ApiGroup)
