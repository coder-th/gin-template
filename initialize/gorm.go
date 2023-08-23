package initialize

import (
	"os"

	"github.com/coder-th/ds-tools/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/coder-th/ds-tools/server/model/system"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.JwtBlacklist{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
