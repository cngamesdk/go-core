package common

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type DimPayChannelSwitchModel struct {
	sql2.SqlBaseModel
	Db          func() *gorm.DB `json:"-" gorm:"-"`
	PlatformId  int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;index:ix_plat_type_status"`
	PayType     string          `json:"pay_type" gorm:"size:50;column:pay_type;default:'';comment:支付类型;index:ix_plat_type_status"`
	Sort        int             `json:"sort" gorm:"size:32;column:sort;default:0;comment:排序，降序"`
	RuleName    string          `json:"rule_name" gorm:"size:50;column:rule_name;default:'';comment:规则名称"`
	RuleKey     string          `json:"rule_key" gorm:"size:150;column:rule_key;default:'';comment:规则键"`
	Rules       Rules           `json:"rules" gorm:"type:json;column:rules;comment:规则"`
	PayChannels PayChannels     `json:"pay_channels" gorm:"type:json;column:pay_channels;comment:支付渠道"`
	Status      string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态;index:ix_plat_type_status"`
}

func (receiver *DimPayChannelSwitchModel) TableName() string {
	return "dim_pay_channel_switch"
}

func (receiver *DimPayChannelSwitchModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimPayChannelSwitchModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimPayChannelSwitchModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}

type Rules []DimPayChannelSwitchRule

// Scan Scanner
func (args *Rules) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args Rules) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

type PayChannels []DimPayChannelSwitchPayChannel

// Scan Scanner
func (args *PayChannels) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args PayChannels) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}

type DimPayChannelSwitchRule struct {
	Name     string        `json:"name"`     // 维度或者指标
	Operator string        `json:"operator"` // 操作符
	Value    []interface{} `json:"value"`    // 值
}

type DimPayChannelSwitchPayChannel struct {
	PayChannelId int64 `json:"pay_channel_id"` // 充值渠道ID
	Weight       int   `json:"weight"`         // 权重
}
