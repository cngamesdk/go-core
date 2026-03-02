package advertising

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type AdvertisingCopyPackageConfigItem struct {
	Text string `json:"text"`
}

type AdvertisingCopyPackageConfig struct {
	List []AdvertisingCopyPackageConfigItem `json:"list"`
}

// Scan Scanner
func (args *AdvertisingCopyPackageConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingCopyPackageConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// AdvertisingCopyPackageLogModel 广告文案包
type AdvertisingCopyPackageLogModel struct {
	sql2.SqlBaseModel
	Db         func() *gorm.DB              `json:"-" gorm:"-"`
	PlatformId int64                        `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	UserId     int64                        `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;"`
	Config     AdvertisingCopyPackageConfig `json:"config" gorm:"type:json;column:config;comment:文案包配置;"`
	Status     string                       `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
}

func (receiver *AdvertisingCopyPackageLogModel) TableName() string {
	return "ods_advertising_copy_package_log"
}

func (receiver *AdvertisingCopyPackageLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *AdvertisingCopyPackageLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *AdvertisingCopyPackageLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
