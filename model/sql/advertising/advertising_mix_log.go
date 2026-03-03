package advertising

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// AdvertisingMixCommonConfig 常规配置
type AdvertisingMixCommonConfig struct {
	GameId int64 `json:"game_id"`
}

// Scan Scanner
func (args *AdvertisingMixCommonConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMixCommonConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

type AdvertisingMixAccountConfig struct {
	List []AdvertisingMixAccountConfigItem `json:"list"`
}

// Scan Scanner
func (args *AdvertisingMixAccountConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMixAccountConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

type AdvertisingMixAccountConfigItem struct {
	AccountId   int64  `json:"account_id"`
	AccountName string `json:"account_name"`
}

// AdvertisingMixAd1Config 广告一级配置
type AdvertisingMixAd1Config struct {
	Name            string                        `json:"name"`
	TargetingConfig AdvertisingMixTargetingConfig `json:"targeting_config"`
}

// Scan Scanner
func (args *AdvertisingMixAd1Config) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMixAd1Config) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingMixAd2Config 广告二级配置
type AdvertisingMixAd2Config struct {
	Name            string                        `json:"name"`
	TargetingConfig AdvertisingMixTargetingConfig `json:"targeting_config"`
}

// Scan Scanner
func (args *AdvertisingMixAd2Config) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMixAd2Config) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingMixAd3Config 广告三级配置
type AdvertisingMixAd3Config struct {
	Name         string                         `json:"name"`
	TextList     []AdvertisingMixTextConfig     `json:"text_list"`
	MaterialList []AdvertisingMixMaterialConfig `json:"material_list"`
}

// Scan Scanner
func (args *AdvertisingMixAd3Config) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMixAd3Config) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingMixConfig 广告组合配置
type AdvertisingMixConfig struct {
	Name string `json:"name"`
}

// Scan Scanner
func (args *AdvertisingMixConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMixConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingMixTargetingConfig 定向配置
type AdvertisingMixTargetingConfig struct {
}

// AdvertisingMixTextConfig 文案配置
type AdvertisingMixTextConfig struct {
}

// AdvertisingMixMaterialConfig 素材配置
type AdvertisingMixMaterialConfig struct {
}

type OdsAdvertisingMixLogModel struct {
	sql2.SqlBaseModel
	Db            func() *gorm.DB             `json:"-" gorm:"-"`
	Name          string                      `json:"name" gorm:"size:50;column:name;default:'';comment:组合名称;uniqueIndex:ix_name;"`
	UserId        int64                       `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;"`
	PlatformId    int64                       `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	Code          string                      `json:"code" gorm:"size:50;column:code;default:'';comment:媒体码;"`
	CommonConfig  AdvertisingMixCommonConfig  `json:"common_config" gorm:"type:json;column:common_config;comment:常规配置;"`
	AccountConfig AdvertisingMixAccountConfig `json:"account_config" gorm:"type:json;column:account_config;comment:帐户配置;"`
	Ad1Config     AdvertisingMixAd1Config     `json:"ad1_config" gorm:"type:json;column:ad1_config;comment:一级配置;"`
	Ad2Config     AdvertisingMixAd2Config     `json:"ad2_config" gorm:"type:json;column:ad2_config;comment:二级配置;"`
	Ad3Config     AdvertisingMixAd3Config     `json:"ad3_config" gorm:"type:json;column:ad3_config;comment:三级配置;"`
	MixConfig     AdvertisingMixConfig        `json:"mix_config" gorm:"type:json;column:mix_config;comment:组合配置;"`
	OtherConfig   sql2.CustomMapType          `json:"other_config" gorm:"type:json;column:other_config;comment:其他配置;"`
	Status        string                      `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
}

func (receiver *OdsAdvertisingMixLogModel) TableName() string {
	return "ods_advertising_mix_log"
}

func (receiver *OdsAdvertisingMixLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsAdvertisingMixLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsAdvertisingMixLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
