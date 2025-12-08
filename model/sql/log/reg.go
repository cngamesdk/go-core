package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

type OdsRegLogModel struct {
	sql.SqlBaseModel
	PlatformId int64     `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_user"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;uniqueIndex:ix_plat_user"`
	RegTime    time.Time `json:"reg_time" gorm:"type:datetime(0);column:reg_time;comment:注册时间"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsRegLogModel) TableName() string {
	return "ods_reg_log"
}

func (receiver *OdsRegLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsRegLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsRegLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *OdsRegLogModel) BeforeSave(tx *gorm.DB) (err error) {
	return
}

func (receiver *OdsRegLogModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (receiver *OdsRegLogModel) beforeCreateOrUpdateHook(tx *gorm.DB) (err error) {
	if receiver.UniqueDevice == "" {
		receiver.SqlCommonModel.FormatUniqueDevice()
	}
	return
}
