package cron_task

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	TaskTypeSqlCleaning             = "sql-cleaning"          //SQL清洗
	TaskTypeAppApplicationPackaging = "application-packaging" // 应用打包

	ExecutionModeSync  = "sync"  // 同步
	ExecutionModeAsync = "async" // 异步
)

var (
	TaskTypes = map[string]string{
		TaskTypeSqlCleaning:             "SQL清洗",
		TaskTypeAppApplicationPackaging: "应用打包",
	}

	ExecutionModes = map[string]string{
		ExecutionModeSync:  "同步",
		ExecutionModeAsync: "异步",
	}
)

// GetTaskTypeName 获取任务名称
func GetTaskTypeName(req string) string {
	taskType, ok := TaskTypes[req]
	if !ok {
		return ""
	}
	return taskType
}

// GetExecutionModeName 获取执行模式名称
func GetExecutionModeName(req string) string {
	executionMode, ok := ExecutionModes[req]
	if !ok {
		return ""
	}
	return executionMode
}

// DimCronTaskConfigModel 定时任务配置维度表
type DimCronTaskConfigModel struct {
	sql2.SqlBaseModel
	Name          string             `json:"name" gorm:"size:100;column:name;default:'';comment:任务名称;uniqueIndex:ix_name"`
	Spec          string             `json:"spec" gorm:"size:150;column:spec;default:'';comment:任务执行规则;"`
	Remark        string             `json:"remark" gorm:"size:512;column:remark;default:'';comment:备注;"`
	Status        string             `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	Content       string             `json:"content" gorm:"type:longtext;column:content;comment:内容;"`
	Config        sql2.CustomMapType `json:"config" gorm:"type:json;column:config;comment:配置;"`
	Sort          int                `json:"sort" gorm:"size:32;column:sort;default:0;comment:排序,降序;"`
	ParentId      int64              `json:"parent_id" gorm:"column:parent_id;default:0;comment:父节点ID;"`
	ChildrenNum   int                `json:"children_num" gorm:"size:32;column:children_num;default:0;comment:孩子节点数量;"`
	TaskType      string             `json:"task_type" gorm:"size:50;column:task_type;default:'';comment:任务类型;"`
	ExecutionMode string             `json:"execution_mode" gorm:"size:50;column:execution_mode;default:'';comment:执行模式;"`
	Db            func() *gorm.DB    `json:"-" gorm:"-"`
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
