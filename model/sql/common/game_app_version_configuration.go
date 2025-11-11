package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type DimGameAppVersionConfiguration struct {
	sql2.SqlBaseModel
	GameId          int64           `json:"game_id" gorm:"column:game_id;default:0;comment:子游戏ID;uniqueIndex:ix_game_code"`
	AppVersionCode  int64           `json:"app_version_code" gorm:"column:app_version_code;default:0;comment:整形AppCode;uniqueIndex:ix_game_code"`
	AppVersionName  string          `json:"app_version_name" gorm:"column:app_version_name;default:'';comment:字符串AppName"`
	Remark          string          `json:"remark" gorm:"size:512;column:remark;default:'';comment:备注"`
	ProductConfigId int64           `json:"product_config_id" gorm:"column:product_config_id;default:0;comment:产品配置ID"`
	Db              func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimGameAppVersionConfiguration) TableName() string {
	return "dim_game_app_version_configuration"
}

func (receiver *DimGameAppVersionConfiguration) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimGameAppVersionConfiguration) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimGameAppVersionConfiguration) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
