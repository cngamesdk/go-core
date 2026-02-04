package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// OdsCampaignLogModel 投放操作日志表
type OdsCampaignLogModel struct {
	sql2.SqlBaseModel
	PlatformId   int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;index:idx_plat_cid_action;"`
	CampaignId   int64           `json:"campaign_id" gorm:"column:campaign_id;default:0;comment:投放ID;index:idx_plat_cid_action;"`
	Action       string          `json:"action" gorm:"size:50;column:action;default:'';comment:操作类型;index:idx_plat_cid_action;"`
	RequestData  string          `json:"request_data" gorm:"type:json;column:request_data;comment:请求数据;"`
	ResponseData string          `json:"response_data" gorm:"type:json;column:response_data;comment:响应数据;"`
	Status       string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	ErrorMessage string          `json:"error_message" gorm:"type:text;column:error_message;comment:错误信息;"`
	Db           func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsCampaignLogModel) TableName() string {
	return "ods_campaign_log"
}

func (receiver *OdsCampaignLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsCampaignLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsCampaignLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
