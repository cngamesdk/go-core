package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimGameModel 主体维度表
type DimCompanyModel struct {
	sql2.SqlBaseModel
	PlatformId       int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name"`
	CompanyName      string          `json:"company_name" gorm:"size:100;column:company_name;default:'';comment:主体名称;uniqueIndex:ix_plat_name"`
	UserAgreementUrl string          `json:"user_agreement_url" gorm:"size:512;column:user_agreement_url;default:'';comment:用户协议"`
	PrivacyPolicyUrl string          `json:"privacy_policy_url" gorm:"size:512;column:privacy_policy_url;default:'';comment:隐私协议"`
	Db               func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimCompanyModel) TableName() string {
	return "dim_company"
}

func (receiver *DimCompanyModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimCompanyModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimCompanyModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
