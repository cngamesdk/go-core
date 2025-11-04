package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	DimPayChannelModelPayTypeWeiXin    = "wei-xin"
	DimPayChannelModelPayTypeAlipay    = "alipay"
	DimPayChannelModelPayTypeDouYinPay = "dou-yin-pay"
)

var PayTypes = map[string]string{
	DimPayChannelModelPayTypeWeiXin:    "微信支付",
	DimPayChannelModelPayTypeAlipay:    "支付宝支付",
	DimPayChannelModelPayTypeDouYinPay: "抖音支付",
}

// ValidPayType 是否有效支付方式
func ValidPayType(payType string) bool {
	_, ok := PayTypes[payType]
	return ok
}

// DimRootGameModel 主游戏维度
type DimPayChannelModel struct {
	sql2.SqlBaseModel
	ChannelName        string          `json:"channel_name" gorm:"size:100;column:channel_name;default:'';comment:渠道名称"`
	CompanyId          int64           `json:"company_id" gorm:"column:company_id;default:0;comment:主体ID"`
	PayType            string          `json:"pay_type" gorm:"size:50;column:pay_type;default:'';comment:支付类型"`
	Status             string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态"`
	ProfitSharingRatio int             `json:"profit_sharing_ratio" gorm:"column:profit_sharing_ratio;default:0;comment:分成比例，如30为30%"`
	Db                 func() *gorm.DB `json:"-" gorm:"-"`
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
