package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimPlatformModel 主体维度表
type DimPlatformModel struct {
	sql2.SqlBaseModel
	PlatformName string          `json:"platform_name" gorm:"size:150;column:platform_name;default:'';comment:平台名称;uniqueIndex:ix_name"`
	Db           func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimPlatformModel) TableName() string {
	return "dim_platform"
}

func (receiver *DimPlatformModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimPlatformModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimPlatformModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
