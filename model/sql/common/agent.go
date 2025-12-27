package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	SettlementTypeFree = "free" //结算类型-免费
	SettlementTypeCps  = "cps"  //结算类型-CPS
	SettlementTypeCpa  = "cpa"  //结算类型-CPA
	SettlementTypeCpt  = "cpt"  //结算类型-CPT
	SettlementTypeCpm  = "cpm"  //结算类型-CPM
	SettlementTypeCpc  = "cpc"  //结算类型-CPC
)

var SettlementTypes = map[string]string{
	SettlementTypeFree: "免费",
	SettlementTypeCps:  "CPS",
	SettlementTypeCpa:  "CPA",
	SettlementTypeCpt:  "CPT",
	SettlementTypeCpm:  "CPM",
	SettlementTypeCpc:  "CPC",
}

// GetSettlementTypeName 获取结算类型名称
func GetSettlementTypeName(req string) string {
	resp, ok := SettlementTypes[req]
	if !ok {
		return ""
	}
	return resp
}

// DimAgentModel 主体维度表
type DimAgentModel struct {
	sql2.SqlBaseModel
	PlatformId     int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_group_name"`
	ChannelGroupId int64           `json:"channel_group_id" gorm:"column:channel_group_id;default:0;comment:渠道组ID;uniqueIndex:ix_plat_group_name"`
	AgentName      string          `json:"agent_name" gorm:"size:100;column:agent_name;default:'';comment:渠道名称;uniqueIndex:ix_plat_group_name"`
	SettlementType string          `json:"settlement_type" gorm:"size:50;column:settlement_type;default:'';comment:结算类型"`
	Db             func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimAgentModel) TableName() string {
	return "dim_agent"
}

func (receiver *DimAgentModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimAgentModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimAgentModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
