package common

import (
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type GlobalCommonConfigModel struct {
	sql2.SqlBaseModel
	Db                    func() *gorm.DB `json:"-" gorm:"-"`
	PlatformId            int             `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID"`
	GamePackagingToolPath string          `json:"game_packaging_tool_path" gorm:"column:game_packaging_tool_path;default:'';comment:游戏打包工具路径"`
}
