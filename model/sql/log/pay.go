package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

const (
	PayStatusPreorder = "preorder"
	PayStatusSuccess  = "success"
	PayStatusCancel   = "cancel"
	CpStatusSuccess   = "success"
	CpStatusFail      = "fail"
)

type OdsPayLogModel struct {
	sql.SqlBaseModel
	PlatformId        int64     `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_order"`
	OrderId           string    `json:"order_id" gorm:"size:100;column:order_id;default:'';comment:订单ID;uniqueIndex:ix_plat_order"`
	MerchantOrderId   string    `json:"merchant_order_id" gorm:"size:512;column:merchant_order_id;default:'';comment:商户订单ID"`
	UserId            int64     `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID"`
	ServerId          int64     `json:"server_id" gorm:"column:server_id;default:0;comment:区服ID"`
	ServerName        string    `json:"server_name" gorm:"size:512;column:server_name;default:'';comment:区服名称"`
	RoleId            string    `json:"role_id" gorm:"size:512;column:role_id;default:'';comment:角色ID"`
	RoleName          string    `json:"role_name" gorm:"size:512;column:role_name;default:'';comment:角色名称"`
	ProductId         string    `json:"product_id" gorm:"size:512;column:product_id;default:'';comment:产品ID"`
	ProductName       string    `json:"product_name" gorm:"size:512;column:product_name;default:'';comment:产品名称"`
	Money             int       `json:"money" gorm:"column:money;default:0;comment:金额，单位分"`
	PayStatus         string    `json:"pay_status" gorm:"size:50;column:pay_status;default:'';comment:支付状态"`
	PayTime           time.Time `json:"pay_time" gorm:"type:datetime(0);column:pay_time;comment:下单时间"`
	CallbackTime      time.Time `json:"callback_time" gorm:"type:datetime(0);column:callback_time;comment:支付回调时间"`
	Ext               string    `json:"ext" gorm:"size:1024;column:ext;default:'';comment:透传参数"`
	CpStatus          string    `json:"cp_status" gorm:"size:50;column:cp_status;default:'';comment:cp发货状态"`
	CpMoney           int       `json:"cp_money" gorm:"column:cp_money;default:0;comment:cp发货金额"`
	CpUrl             string    `json:"cp_url" gorm:"size:1024;column:cp_url;default:'';comment:cp发货请求"`
	CpResult          string    `json:"cp_result" gorm:"type:text;column:cp_result;comment:cp发货结果"`
	CpSendTime        time.Time `json:"cp_send_time" gorm:"type:datetime(0);column:cp_send_time;comment:cp发货时间"`
	CpSendRepeatTimes int       `json:"cp_send_repeat_times" gorm:"column:cp_send_repeat_times;default:0;comment:cp发货重试次数"`
	PayChannelId      int       `json:"pay_channel_id" gorm:"column:pay_channel_id;default:0;comment:支付渠道ID"`
	TestOrder         int       `json:"test_order" gorm:"size:6;column:test_order;default:0;comment:测试订单"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsPayLogModel) TableName() string {
	return "ods_pay_log"
}

func (receiver *OdsPayLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsPayLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsPayLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
