package message_queue

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	MqTaskStatusPending    = "pending"          // 待处理
	MqTaskStatusProcessing = "processing"       // 处理中
	MqTaskStatusSuccess    = sql2.StatusSuccess // 成功
	MqTaskStatusFailed     = sql2.StatusFail    // 失败
)

var (
	MqTaskStatusMap = map[string]string{
		MqTaskStatusPending:    "待处理",
		MqTaskStatusProcessing: "处理中",
		MqTaskStatusSuccess:    "成功",
		MqTaskStatusFailed:     "失败",
	}
)

// OdsMqTaskLogModel 消息队列任务日志表
type OdsMqTaskLogModel struct {
	sql2.SqlBaseModel
	PlatformId   int64                 `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	Topic        string                `json:"topic" gorm:"size:100;column:topic;default:'';comment:消息主题;"`
	Partition    int                   `json:"partition" gorm:"size:32;column:partition;default:0;comment:分区;"`
	Key          string                `json:"key" gorm:"size:100;column:key;default:'';comment:消息键;"`
	Payload      string                `json:"payload" gorm:"type:json;column:payload;comment:消息体;"`
	Status       string                `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	RetryCount   int                   `json:"retry_count" gorm:"size:32;column:retry_count;default:0;comment:重试次数;"`
	MaxRetries   int                   `json:"max_retries" gorm:"size:32;column:max_retries;default:5;comment:最大重试次数;"`
	ErrorMessage string                `json:"error_message" gorm:"type:text;column:error_message;comment:错误信息;"`
	ProcessedAt  sql2.MyCustomDatetime `json:"processed_at" gorm:"type:datetime(0);column:processed_at;comment:处理时间;"`
	Db           func() *gorm.DB       `json:"-" gorm:"-"`
}

func (receiver *OdsMqTaskLogModel) TableName() string {
	return "ods_mq_task_log"
}

func (receiver *OdsMqTaskLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsMqTaskLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsMqTaskLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
