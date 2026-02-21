package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	DimAdvertisingMediaAccountStatusNormal = sql2.StatusNormal
	DimAdvertisingMediaAccountStatusDelete = sql2.StatusDelete
)

var AdvertisingMediaAccountStatuss = map[string]string{
	DimAdvertisingMediaAccountStatusNormal: "正常",
	DimAdvertisingMediaAccountStatusDelete: "删除",
}

func GetAdvertisingMediaAccountStatusName(req string) string {
	resp, ok := AdvertisingMediaAccountStatuss[req]
	if !ok {
		return ""
	}
	return resp
}

// DimAdvertisingMediaAccountModel 广告媒体维度表
type DimAdvertisingMediaAccountModel struct {
	sql2.SqlBaseModel
	PlatformId            int64                 `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_account"`
	AccountName           string                `json:"account_name" gorm:"size:100;column:account_name;default:'';comment:帐户名称"`
	AccountId             uint64                `json:"account_id" gorm:"size:100;column:account_id;default:0;comment:帐户ID;uniqueIndex:ix_plat_account"`
	Status                string                `json:"status" gorm:"size:100;column:status;default:'';comment:状态"`
	ManageId              int64                 `json:"manage_id" gorm:"column:manage_id;default:0;comment:超管主键ID"`
	Role                  string                `json:"role" gorm:"column:role;default:'';comment:角色"`
	Extension             string                `json:"extension" gorm:"type:json;column:extension;comment:扩展字段"`
	AccessToken           string                `json:"access_token" gorm:"size:100;column:access_token;default:'';comment:授权令牌"`
	RefreshToken          string                `json:"refresh_token" gorm:"size:512;column:refresh_token;default:'';comment:刷新令牌"`
	ExpiresAt             sql2.MyCustomDatetime `json:"expires_at" gorm:"type:datetime(0);column:expires_at;comment:Access_Token过期时间"`
	RefreshTokenExpiresAt sql2.MyCustomDatetime `json:"refresh_token_expires_at" gorm:"type:datetime(0);column:refresh_token_expires_at;comment:refresh_token过期时间"`
	DeveloperId           int64                 `json:"developer_id" gorm:"column:developer_id;default:0;comment:开发者主键ID"`
	Db                    func() *gorm.DB       `json:"-" gorm:"-"`
}

func (receiver *DimAdvertisingMediaAccountModel) TableName() string {
	return "dim_advertising_media_account"
}

func (receiver *DimAdvertisingMediaAccountModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaAccountModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaAccountModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
