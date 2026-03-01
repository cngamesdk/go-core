package advertising

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// AdvertisingTargetingPackageConfig 定向包配置
type AdvertisingTargetingPackageConfig struct {
	Location    []string `json:"location"`
	AgeRange    [2]int   `json:"age_range"`
	Gender      string   `json:"gender"`
	Interests   []string `json:"interests"`
	Platforms   []string `json:"platforms"`
	DeviceTypes []string `json:"device_types"`
	NetworkType string   `json:"network_type"`
}

// Scan Scanner
func (args *AdvertisingTargetingPackageConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingTargetingPackageConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

type OdsAdvertisingTargetingPackageLogModel struct {
	sql2.SqlBaseModel
	Db         func() *gorm.DB                   `json:"-" gorm:"-"`
	PlatformId int64                             `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	UserId     int64                             `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;"`
	Config     AdvertisingTargetingPackageConfig `json:"config" gorm:"type:json;column:config;comment:定向包配置;"`
	Status     string                            `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
}

func (receiver *OdsAdvertisingTargetingPackageLogModel) TableName() string {
	return "ods_advertising_targeting_package_log"
}

func (receiver *OdsAdvertisingTargetingPackageLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsAdvertisingTargetingPackageLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsAdvertisingTargetingPackageLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
