package advertising

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// AdvertisingTaskCommonConfig 广告任务常规配置
type AdvertisingTaskCommonConfig struct {
}

// Scan Scanner
func (args *AdvertisingTaskCommonConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingTaskCommonConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingTaskAd1Config 广告任务一级配置
type AdvertisingTaskAd1Config struct {
}

// Scan Scanner
func (args *AdvertisingTaskAd1Config) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingTaskAd1Config) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingTaskAd2Config 广告任务二级配置
type AdvertisingTaskAd2Config struct {
}

// Scan Scanner
func (args *AdvertisingTaskAd2Config) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingTaskAd2Config) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingTaskAd3Config 广告任务三级配置
type AdvertisingTaskAd3Config struct {
}

// Scan Scanner
func (args *AdvertisingTaskAd3Config) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingTaskAd3Config) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// OdsAdvertisingTaskLogModel 广告任务日志
type OdsAdvertisingTaskLogModel struct {
	sql2.SqlBaseModel
	Db           func() *gorm.DB             `json:"-" gorm:"-"`
	UserId       int64                       `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;"`
	MixId        int64                       `json:"mix_id" gorm:"column:mix_id;default:0;comment:广告组合ID;"`
	PlatformId   int64                       `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	Code         string                      `json:"code" gorm:"size:50;column:code;default:'';comment:媒体码;"`
	AccountId    int64                       `json:"account_id" gorm:"column:account_id;default:0;comment:帐户ID;"`
	Ad1Id        int64                       `json:"ad1_id" gorm:"column:ad1_id;default:0;comment:一级ID;"`
	Ad2Id        int64                       `json:"ad2_id" gorm:"column:ad2_id;default:0;comment:二级ID;"`
	Ad3Id        int64                       `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:三级ID;"`
	CommonConfig AdvertisingTaskCommonConfig `json:"common_config" gorm:"type:json;column:common_config;comment:常规配置;"`
	Ad1Config    AdvertisingTaskAd1Config    `json:"ad1_config" gorm:"type:json;column:ad1_config;comment:一级配置;"`
	Ad2Config    AdvertisingTaskAd2Config    `json:"ad2_config" gorm:"type:json;column:ad2_config;comment:二级配置;"`
	Ad3Config    AdvertisingTaskAd3Config    `json:"ad3_config" gorm:"type:json;column:ad3_config;comment:三级配置;"`
	OtherConfig  sql2.CustomMapType          `json:"other_config" gorm:"type:json;column:other_config;comment:其他配置;"`
}

func (receiver *OdsAdvertisingTaskLogModel) TableName() string {
	return "ods_advertising_task_log"
}

func (receiver *OdsAdvertisingTaskLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsAdvertisingTaskLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsAdvertisingTaskLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
