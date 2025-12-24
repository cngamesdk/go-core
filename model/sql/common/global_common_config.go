package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type DimGlobalCommonConfigModel struct {
	sql2.SqlBaseModel
	Db                    func() *gorm.DB `json:"-" gorm:"-"`
	PlatformId            int             `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat"`
	GamePackagingToolPath string          `json:"game_packaging_tool_path" gorm:"size:512;column:game_packaging_tool_path;default:'';comment:游戏打包工具路径"`
}

func (receiver *DimGlobalCommonConfigModel) TableName() string {
	return "dim_global_common_config"
}

func (receiver *DimGlobalCommonConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimGlobalCommonConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimGlobalCommonConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
