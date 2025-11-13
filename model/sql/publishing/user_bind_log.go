package publishing

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// OdsPublishingUserBindLogModel 发行渠道用户绑定日志
type OdsPublishingUserBindLogModel struct {
	sql.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_channel_open;index:ix_p_c_u;index:ix_plat_user"`
	ChannelId  int64           `json:"channel_id" gorm:"column:channel_id;default:0;comment:发行渠道ID;uniqueIndex:ix_plat_channel_open;index:ix_p_c_u"`
	OpenId     string          `json:"open_id" gorm:"size:100;column:open_id;default:'';comment:发行渠道用户ID;uniqueIndex:ix_plat_channel_open"`
	UserId     int64           `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;index:ix_plat_user"`
	UnionId    string          `json:"union_id" gorm:"size:100;column:union_id;default:'';comment:union_id;index:ix_p_c_u"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsPublishingUserBindLogModel) TableName() string {
	return "ods_publishing_user_bind_log"
}

func (receiver *OdsPublishingUserBindLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsPublishingUserBindLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsPublishingUserBindLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
