package upload

import (
	"mime/multipart"

	"github.com/coder-th/ds-tools/server/global"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "aliyun-oss":
		return &AliyunOSS{}
	default:
		return &Local{}
	}
}