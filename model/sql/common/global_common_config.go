package common

import (
	"context"
	"database/sql/driver"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type DimGlobalCommonConfigModel struct {
	sql2.SqlBaseModel
	Db                         func() *gorm.DB            `json:"-" gorm:"-"`
	PlatformId                 int                        `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat"`
	GamePackagingToolPath      string                     `json:"game_packaging_tool_path" gorm:"size:512;column:game_packaging_tool_path;default:'';comment:游戏打包工具路径"`
	JavaExecutionPath          string                     `json:"java_execution_path" gorm:"size:512;column:java_execution_path;default:'';comment:java执行路径"`
	PaymentChannelSwitchConfig PaymentChannelSwitchConfig `json:"payment_channel_switch_config" gorm:"type:json;column:payment_channel_switch_config;comment:充值渠道切换配置"`
}

func (receiver *DimGlobalCommonConfigModel) TableName() string {
	return "dim_global_common_config"
}

func (receiver *DimGlobalCommonConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimGlobalCommonConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimGlobalCommonConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}

type PaymentChannelSwitchConfig struct {
	PayType string                           `json:"pay_type"` // 充值方式
	Config  []PaymentChannelSwitchItemConfig `json:"config"`   // 配置
}

type PaymentChannelSwitchItemConfig struct {
	RuleName    string                                     `json:"rule_name"`    // 规则名称
	RuleKey     string                                     `json:"rule_key"`     // 规则键值
	Value       []PaymentChannelSwitchItemConfigValue      `json:"value"`        // 规则项
	PayChannels []PaymentChannelSwitchItemConfigPayChannel `json:"pay_channels"` // 具体充值渠道
}

type PaymentChannelSwitchItemConfigValue struct {
	Name     string        `json:"name"`     // 维度或者指标
	Operator string        `json:"operator"` // 操作符
	Value    []interface{} `json:"value"`    // 值
}

type PaymentChannelSwitchItemConfigPayChannel struct {
	PayChannelId int64 `json:"pay_channel_id"` // 充值渠道ID
	Weight       int   `json:"weight"`         // 权重
}

// Scan Scanner
func (args *PaymentChannelSwitchConfig) Scan(value interface{}) error {
	return sql2.JsonScan(args, value)
}

// Value Valuer
func (args PaymentChannelSwitchConfig) Value() (driver.Value, error) {
	return sql2.JsonValue(args)
}
