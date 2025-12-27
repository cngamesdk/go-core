package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimAgentModel 主体维度表
type DimSiteModel struct {
	sql2.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_agent_name"`
	AgentId    int64           `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID;uniqueIndex:ix_plat_agent_name"`
	SiteName   string          `json:"site_name" gorm:"size:100;column:site_name;default:'';comment:广告位名称;uniqueIndex:ix_plat_agent_name"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimSiteModel) TableName() string {
	return "dim_site"
}

func (receiver *DimSiteModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimSiteModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimSiteModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
