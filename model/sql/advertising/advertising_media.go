package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimAdvertisingMediaModel 广告媒体维度表
type DimAdvertisingMediaModel struct {
	sql2.SqlBaseModel
	PlatformId           int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name"`
	AdvertisingMediaName string          `json:"advertising_media_name" gorm:"size:100;column:advertising_media_name;default:'';comment:广告媒体名称;uniqueIndex:ix_plat_name"`
	Db                   func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimAdvertisingMediaModel) TableName() string {
	return "dim_advertising_media"
}

func (receiver *DimAdvertisingMediaModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
