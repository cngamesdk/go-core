package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	CampaignStatusDraft   = "draft"
	CampaignStatusPending = "pending"
	CampaignStatusRunning = "running"
	CampaignStatusPaused  = "paused"
	CampaignStatusEnded   = "ended"
	CampaignStatusFailed  = "failed"

	SyncStatusPending = "pending"
	SyncStatusSyncing = "syncing"
	SyncStatusSuccess = "success"
	SyncStatusFailed  = "failed"
)

var (
	CampaignStatus = map[string]string{
		CampaignStatusDraft:   "草稿箱",
		CampaignStatusPending: "待处理",
		CampaignStatusRunning: "执行中",
		CampaignStatusPaused:  "暂停",
		CampaignStatusEnded:   "结束",
		CampaignStatusFailed:  "失败",
	}

	CampaignSyncStatus = map[string]string{
		SyncStatusPending: "待处理",
		SyncStatusSyncing: "同步中",
		SyncStatusSuccess: "成功",
		SyncStatusFailed:  "失败",
	}
)

type DimCampaignModel struct {
	sql2.SqlBaseModel
	PlatformId   int64                 `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;"`
	MediaId      int64                 `json:"media_id" gorm:"column:media_id;default:0;comment:媒体ID;"`
	Name         string                `json:"name" gorm:"size:200;column:name;default:'';comment:投放名称;"`
	GameId       int64                 `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID;"`
	Budget       int                   `json:"budget" gorm:"size:32;column:budget;default:0;comment:总预算,单位：分;"`
	DailyBudget  int                   `json:"daily_budget" gorm:"size:32;column:daily_budget;default:0;comment:日预算,单位：分;"`
	StartTime    sql2.MyCustomDatetime `json:"start_time" gorm:"type:datetime(0);column:start_time;comment:开始时间;index:idx_time_range;"`
	EndTime      sql2.MyCustomDatetime `json:"end_time" gorm:"type:datetime(0);column:end_time;comment:结束时间;index:idx_time_range;"`
	Status       string                `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	Targeting    string                `json:"targeting" gorm:"type:json;column:targeting;comment:定向设置;"`
	Creative     string                `json:"creative" gorm:"type:json;column:creative;comment:创意内容;"`
	Metrics      string                `json:"metrics" gorm:"type:json;column:metrics;comment:投放指标;"`
	SyncStatus   string                `json:"sync_status" gorm:"size:50;column:sync_status;default:'';comment:同步状态;"`
	LastSyncTime sql2.MyCustomDatetime `json:"last_sync_time" gorm:"type:datetime(0);column:last_sync_time;comment:最后同步时间;"`
	CreatedBy    string                `json:"created_by" gorm:"size:50;column:created_by;default:'';comment:创建人;"`
	Db           func() *gorm.DB       `json:"-" gorm:"-"`
}

func (receiver *DimCampaignModel) TableName() string {
	return "dim_campaign"
}

func (receiver *DimCampaignModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimCampaignModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimCampaignModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
