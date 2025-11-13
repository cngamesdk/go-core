package user

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

// OdsUserOperationLogModel 用户操作日志
type OdsUserOperationLogModel struct {
	sql.SqlBaseModel
	PlatformId    int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;index:ix_plat_user"`
	UserId        int64           `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;index:ix_plat_user"`
	OperationTime time.Time       `json:"operation_time" gorm:"column:operation_time;comment:操作时间"`
	Remark        string          `json:"remark" gorm:"type:text;column:remark;comment:备注"`
	Db            func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsUserOperationLogModel) TableName() string {
	return "ods_user_operation_log"
}

func (receiver *OdsUserOperationLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsUserOperationLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
