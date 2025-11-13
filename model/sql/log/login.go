package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

type OdsLoginLogModel struct {
	sql.SqlBaseModel
	PlatformId int64     `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID"`
	LoginTime  time.Time `json:"login_time" gorm:"type:datetime(0);column:login_time;comment:登录时间"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsLoginLogModel) TableName() string {
	return "ods_login_logs"
}

func (receiver *OdsLoginLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsLoginLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
