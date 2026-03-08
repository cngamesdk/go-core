package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// OdsAdvertisingOceanengineAppLogModel 巨量引擎安卓应用列表
// https://open.oceanengine.com/labels/7/docs/1846773030696265?origin=left_nav
type OdsAdvertisingOceanengineAppLogModel struct {
	sql2.SqlBaseModel
	Db                 func() *gorm.DB       `json:"-" gorm:"-"`
	PlatformId         int64                 `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	AccountId          int64                 `json:"account_id" gorm:"column:account_id;default:0;comment:帐户ID;"`
	BasicPackageId     string                `json:"basic_package_id" gorm:"size:100;column:basic_package_id;default:'';comment:应用包ID（创建分包需要使用此id入参）;uniqueIndex:ix_package_id;"`
	AppName            string                `json:"app_name" gorm:"size:100;column:app_name;default:'';comment:应用名称;"`
	AppNameEn          string                `json:"app_name_en" gorm:"size:100;column:app_name_en;default:'';comment:英文应用名称;"`
	PackageName        string                `json:"package_name" gorm:"size:100;column:package_name;default:'';comment:包名;uniqueIndex:ix_package_name;"`
	VersionCode        string                `json:"version_code" gorm:"size:100;column:version_code;default:'';comment:版本号;"`
	VersionName        string                `json:"version_name" gorm:"size:100;column:version_name;default:'';comment:版本名称;"`
	AppLogo            string                `json:"app_logo" gorm:"size:512;column:app_logo;default:'';comment:应用logo;"`
	PublishTime        sql2.MyCustomDatetime `json:"publish_time" gorm:"type:type:datetime(0);column:publish_time;comment:发布时间;"`
	Reason             string                `json:"reason" gorm:"type:text;column:reason;comment:拒审原因;"`
	SuccessReason      string                `json:"success_reason" gorm:"type:text;column:success_reason;comment:审核成功透传信息;"`
	HistoryAccountId   int64                 `json:"history_account_id" gorm:"column:history_account_id;default:0;comment:历史来源账户id;"`
	HistoryAccountType int64                 `json:"history_account_type" gorm:"column:history_account_type;default:0;comment:历史来源账户类型;"`
	HistoryAccountName string                `json:"history_account_name" gorm:"size:100;column:history_account_name;default:'';comment:历史来源账户名称;"`
	IsEbpAsset         bool                  `json:"is_ebp_asset" gorm:"column:is_ebp_asset;comment:是否为EBP资产;"`
	HasExtendPackage   bool                  `json:"has_extend_package" gorm:"column:has_extend_package;comment:是否有分包;"`
	DownloadUrl        string                `json:"download_url" gorm:"size:512;column:download_url;default:'';comment:下载链接;"`
	CreateTime         sql2.MyCustomDatetime `json:"create_time" gorm:"type:datetime(0);column:create_time;comment:创建时间;"`
	Extension          string                `json:"extension" gorm:"type:json;column:extension;comment:扩展字段;"`
}

func (receiver *OdsAdvertisingOceanengineAppLogModel) TableName() string {
	return "ods_advertising_oceanengine_app_log"
}

func (receiver *OdsAdvertisingOceanengineAppLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsAdvertisingOceanengineAppLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsAdvertisingOceanengineAppLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
