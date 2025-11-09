package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	DimAdvertisingMediaManagementAccountStatusNormal = sql2.StatusNormal
	DimAdvertisingMediaManagementAccountStatusDelete = sql2.StatusDelete
)

var AdvertisingMediaManagementAccountStatuss = map[string]string{
	DimAdvertisingMediaManagementAccountStatusNormal: "正常",
	DimAdvertisingMediaManagementAccountStatusDelete: "删除",
}

func GetAdvertisingMediaManagementAccountStatusName(req string) string {
	resp, ok := AdvertisingMediaManagementAccountStatuss[req]
	if !ok {
		return ""
	}
	return resp
}

// DimAdvertisingMediaManagementAccountModel 广告媒体管理维度表
type DimAdvertisingMediaManagementAccountModel struct {
	sql2.SqlBaseModel
	AccountName string          `json:"account_name" gorm:"size:100;column:account_name;default:'';comment:帐户名称"`
	AccountId   uint64          `json:"account_id" gorm:"size:100;column:account_id;default:0;comment:帐户ID"`
	Status      string          `json:"status" gorm:"size:100;column:status;default:'';comment:状态"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimAdvertisingMediaManagementAccountModel) TableName() string {
	return "dim_advertising_media_management_account"
}

func (receiver *DimAdvertisingMediaManagementAccountModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaManagementAccountModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaManagementAccountModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
