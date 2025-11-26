package cron_task

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimCronTaskConfigModel 定时任务配置维度表
type DimCronTaskConfigModel struct {
	sql2.SqlBaseModel
	TaskId string          `json:"task_id" gorm:"size:100;column:task_id;default:'';comment:任务ID;uniqueIndex:ix_task"`
	Name   string          `json:"name" gorm:"size:100;column:name;default:'';comment:任务名称;"`
	Spec   string          `json:"spec" gorm:"size:150;column:spec;default:'';comment:任务执行规则;"`
	Remark string          `json:"remark" gorm:"size:512;column:remark;default:'';comment:备注;"`
	Status string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	Config sql2.JSON       `json:"config" gorm:"type:json;column:config;comment:配置;"`
	Db     func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimCronTaskConfigModel) TableName() string {
	return "dim_cron_task_config"
}

func (receiver *DimCronTaskConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimCronTaskConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimCronTaskConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}

func (receiver *DimCronTaskConfigModel) Save(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Save(receiver).Error
	return
}
