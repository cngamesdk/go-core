package advertising

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type AdvertisingMaterialPackageConfigItem struct {
	MaterialId     int64 `json:"material_id"`
	MaterialFileId int64 `json:"material_file_id"`
}

type AdvertisingMaterialPackageConfig struct {
	List []AdvertisingMaterialPackageConfigItem `json:"list"`
}

// Scan Scanner
func (args *AdvertisingMaterialPackageConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args AdvertisingMaterialPackageConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

// OdsAdvertisingMaterialPackageLogModel 广告素材包模型
type OdsAdvertisingMaterialPackageLogModel struct {
	sql2.SqlBaseModel
	Db         func() *gorm.DB                  `json:"-" gorm:"-"`
	PlatformId int64                            `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	UserId     int64                            `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;"`
	Config     AdvertisingMaterialPackageConfig `json:"config" gorm:"type:json;column:config;comment:素材包配置;"`
	Status     string                           `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
}

func (receiver *OdsAdvertisingMaterialPackageLogModel) TableName() string {
	return "ods_advertising_material_package_log"
}

func (receiver *OdsAdvertisingMaterialPackageLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsAdvertisingMaterialPackageLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsAdvertisingMaterialPackageLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
