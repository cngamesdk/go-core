package advertising

import (
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"gorm.io/gorm"
)

// DimAdvertisingDeveloperConfigModel 广告开发者配置
type DimAdvertisingDeveloperConfigModel struct {
	Name        string          `json:"name" gorm:"size:50;column:name;default:'';comment:配置名称;"`
	PlatformId  int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	Code        string          `json:"code" gorm:"size:50;column:code;default:'';comment:媒体码;"`
	CompanyId   int64           `json:"company_id" gorm:"column:company_id;default:0;comment:主体ID;"`
	AppId       string          `json:"app_id" gorm:"size:512;column:app_id;default:'';comment:开发者ID;uniqueIndex:ix_app_id"`
	Secret      string          `json:"secret" gorm:"-"`
	SecretCrypt string          `json:"secret_crypt" gorm:"size:1024;column:secret_crypt;default:'';comment:加密密钥;"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
	AesKey      func() string   `json:"-" gorm:"-"`
}

func (receiver *DimAdvertisingDeveloperConfigModel) TableName() string {
	return "dim_advertising_developer_config"
}

func (receiver *DimAdvertisingDeveloperConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimAdvertisingDeveloperConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimAdvertisingDeveloperConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}

func (receiver *DimAdvertisingDeveloperConfigModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimAdvertisingDeveloperConfigModel) BeforeCreate(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *DimAdvertisingDeveloperConfigModel) BeforeSave(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *DimAdvertisingDeveloperConfigModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *DimAdvertisingDeveloperConfigModel) beforeCreateOrUpdateHook(tx *gorm.DB) (err error) {
	if receiver.Secret != "" && receiver.SecretCrypt == "" {
		receiver.SecretCrypt = cryptor.Base64StdEncode(string(cryptor.AesEcbEncrypt([]byte(receiver.Secret), []byte(receiver.AesKey()))))
	}
	return
}

func (receiver *DimAdvertisingDeveloperConfigModel) findHook(tx *gorm.DB) (err error) {
	if receiver.SecretCrypt != "" && receiver.Secret == "" {
		receiver.Secret = string(cryptor.AesEcbDecrypt([]byte(cryptor.Base64StdDecode(receiver.SecretCrypt)), []byte(receiver.AesKey())))
	}
	return
}
