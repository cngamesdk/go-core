package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimProductConfigurationModel 产品通用配置
type DimProductCommonConfigurationModel struct {
	sql2.SqlBaseModel
	ConfigName      string          `json:"config_name" gorm:"size:100;column:config_name;default:'';comment:配置名称;uniqueIndex:ix_config_name"`
	ShippingAddress string          `json:"shipping_address" gorm:"size:512;column:shipping_address;default:'';comment:发货地址"`
	Db              func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimProductCommonConfigurationModel) TableName() string {
	return "dim_product_common_configuration"
}

func (receiver *DimProductCommonConfigurationModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimProductCommonConfigurationModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimProductCommonConfigurationModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
