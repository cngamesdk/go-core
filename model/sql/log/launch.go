package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

const (
	LaunchLogActionActive = "active" // 激活
	LaunchLogActionLaunch = "launch" // 启动
)

type OdsLaunchLogModel struct {
	sql.SqlBaseModel
	PlatformId int64     `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID"`
	Action     string    `json:"action" gorm:"size:50;column:action;default:'';comment:行为"`
	ActionTime time.Time `json:"action_time" gorm:"type:datetime(0);column:action_time;comment:行为时间"`
	Ad3Id      int64     `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsLaunchLogModel) TableName() string {
	return "ods_launch_log"
}

func (receiver *OdsLaunchLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsLaunchLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
