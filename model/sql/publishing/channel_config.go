package publishing

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimPublishingChannelConfigModel 发行渠道配置表
type DimPublishingChannelConfigModel struct {
	sql.SqlBaseModel
	ChannelName string          `json:"channel_name" gorm:"size:100;column:channel_name;default:'';comment:发行渠道名称"`
	AgentId     int             `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimPublishingChannelConfigModel) TableName() string {
	return "dim_publishing_channel_config"
}

func (receiver *DimPublishingChannelConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimPublishingChannelConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimPublishingChannelConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
