package report

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwsDayRootGameBackOverviewLogModel 根游戏回流每日总览表
type DwsDayRootGameBackOverviewLogModel struct {
	sql.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_unique"`
	StatDate   string          `json:"stat_date" gorm:"type:date;column:stat_date;comment:统计日期;uniqueIndex:ix_unique"`
	RootGameId int64           `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:根游戏ID;uniqueIndex:ix_unique"`
	AgentId    int64           `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID;uniqueIndex:ix_unique"`
	SiteId     int64           `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID;uniqueIndex:ix_unique"`
	Ad3Id      int64           `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID;uniqueIndex:ix_unique"`
	RegCount   int             `json:"reg_count" gorm:"column:reg_count;default:0;comment:注册人数;"`
	LoginCount int             `json:"login_count" gorm:"column:login_count;default:0;comment:登录人数;"`
	LoginSum   int             `json:"login_sum" gorm:"column:login_sum;default:0;comment:登录次数;"`
	PayCount   int             `json:"pay_count" gorm:"column:pay_count;default:0;comment:支付人数;"`
	PaySum     int             `json:"pay_sum" gorm:"column:pay_sum;default:0;comment:支付次数;"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DwsDayRootGameBackOverviewLogModel) TableName() string {
	return "dws_day_root_game_back_overview_log"
}

func (receiver *DwsDayRootGameBackOverviewLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwsDayRootGameBackOverviewLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
