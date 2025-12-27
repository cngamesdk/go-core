package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimChannelGroupModel 渠道组维度表
type DimChannelGroupModel struct {
	sql2.SqlBaseModel
	PlatformId         int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_media_name"`
	AdvertisingMediaId int64           `json:"advertising_media_id" gorm:"column:advertising_media_id;default:0;comment:广告媒体ID;uniqueIndex:ix_plat_media_name"`
	ChannelGroupName   string          `json:"channel_group_name" gorm:"size:100;column:channel_group_name;default:'';comment:渠道组名称;uniqueIndex:ix_plat_media_name"`
	Db                 func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimChannelGroupModel) TableName() string {
	return "dim_channel_group"
}

func (receiver *DimChannelGroupModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimChannelGroupModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimChannelGroupModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
