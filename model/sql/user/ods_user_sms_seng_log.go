package user

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

// OdsUserSmsSendLogModel 用户发送短信日志
type OdsUserSmsSendLogModel struct {
	sql.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;index:ix_plat_phone"`
	Phone      string          `json:"phone" gorm:"size:50;column:phone;default:'';comment:手机号;index:ix_plat_phone"`
	ActionTime time.Time       `json:"action_time" gorm:"type:datetime(0);column:action_time;comment:行为时间"`
	Content    string          `json:"content" gorm:"type:text;column:content;comment:发送内容"`
	Result     string          `json:"result" gorm:"size:512;column:result;comment:发送结果"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsUserSmsSendLogModel) TableName() string {
	return "ods_user_sms_send_log"
}

func (receiver *OdsUserSmsSendLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsUserSmsSendLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
