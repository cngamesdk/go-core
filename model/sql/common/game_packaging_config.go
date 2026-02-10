package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	UseStatusNormal = sql2.StatusNormal
	UseStatusRemove = sql2.StatusRemove
)

type DimGamePackagingConfigModel struct {
	sql2.SqlBaseModel
	Db              func() *gorm.DB `json:"-" gorm:"-"`
	PlatformId      int             `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;index:ix_plat_game_media"`
	GameId          int             `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID;index:ix_plat_game_media"`
	MediaCode       string          `json:"media_code" gorm:"size:50;column:media_code;default:'';comment:媒体码;index:ix_plat_game_media"`
	GamePackagePath string          `json:"game_package_path" gorm:"size:512;column:game_package_path;default:'';comment:游戏包路径"`
	GamePackageHash string          `json:"game_package_hash" gorm:"size:100;column:game_package_hash;default:'';comment:游戏包hash摘要"`
	Status          string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态"`
	UseStatus       string          `json:"use_status" gorm:"size:50;column:use_status;default:'';comment:使用状态"`
}

func (receiver *DimGamePackagingConfigModel) TableName() string {
	return "dim_game_packaging_config"
}

func (receiver *DimGamePackagingConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimGamePackagingConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimGamePackagingConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
