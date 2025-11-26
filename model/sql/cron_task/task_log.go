package cron_task

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

// OdsCronTaskLogModel 定时任务日志表
type OdsCronTaskLogModel struct {
	sql2.SqlBaseModel
	ConfigId  int64           `json:"config_id" gorm:"column:config_id;default:0;comment:任务Id;index:ix_config_id"`
	StartTime time.Time       `json:"start_time" gorm:"type:datetime(0);column:start_time;comment:开始时间;"`
	EndTime   time.Time       `json:"end_time" gorm:"type:datetime(0);column:end_time;comment:结束时间;"`
	Latency   int             `json:"latency" gorm:"size:11;column:latency;default:0;comment:延迟,单位：秒;"`
	Status    string          `json:"status" gorm:"size:50;column:status;comment:状态;"`
	Result    string          `json:"result" gorm:"type:text;column:result;comment:执行结果;"`
	Db        func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsCronTaskLogModel) TableName() string {
	return "ods_cron_task_log"
}

func (receiver *OdsCronTaskLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsCronTaskLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsCronTaskLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
