package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	PayTypeWeiXinPay       = "wei-xin-pay"
	PayTypeAlipay          = "alipay"
	PayTypeDouYinPay       = "dou-yin-pay"
	PayChannelStatusNormal = sql2.StatusNormal
	PayChannelStatusDelete = sql2.StatusDelete
)

var PayTypes = map[string]string{
	PayTypeWeiXinPay: "微信支付",
	PayTypeAlipay:    "支付宝支付",
	PayTypeDouYinPay: "抖音支付",
}

var PayStatuss = map[string]string{
	PayChannelStatusNormal: "正常",
	PayChannelStatusDelete: "删除",
}

// GetPayStatusName 获取状态名称
func GetPayStatusName(req string) string {
	name, ok := PayStatuss[req]
	if !ok {
		return ""
	}
	return name
}

// GetPayTypeName 获取支付名称
func GetPayTypeName(payType string) string {
	name, ok := PayTypes[payType]
	if !ok {
		return ""
	}
	return name
}

// DimRootGameModel 主游戏维度
type DimPayChannelModel struct {
	sql2.SqlBaseModel
	PlatformId  int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_company_name"`
	CompanyId   int64           `json:"company_id" gorm:"column:company_id;default:0;comment:主体ID;uniqueIndex:ix_plat_company_name"`
	ChannelName string          `json:"channel_name" gorm:"size:100;column:channel_name;default:'';comment:渠道名称;uniqueIndex:ix_plat_company_name"`
	PayType     string          `json:"pay_type" gorm:"size:50;column:pay_type;default:'';comment:支付类型"`
	Status      string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态"`
	Rate        int             `json:"rate" gorm:"column:rate;default:0;comment:费率，如5为5%"`
	Config      sql2.JSON       `json:"config" gorm:"column:config;default:'{}';comment:配置"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimPayChannelModel) TableName() string {
	return "dim_pay_channel"
}

func (receiver *DimPayChannelModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimPayChannelModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimPayChannelModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
