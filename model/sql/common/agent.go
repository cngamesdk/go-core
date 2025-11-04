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

// DimAgentModel 主体维度表
type DimAgentModel struct {
	sql2.SqlBaseModel
	AgentName      string          `json:"agent_name" gorm:"size:100;column:agent_name;default:'';comment:渠道名称"`
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
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
