package report

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwsDayRootGameBackPayActiveLogModel 根游戏回流每日活跃报表
type DwsDayRootGameBackPayActiveLogModel struct {
	sql.SqlBaseModel
	PlatformId  int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_unique"`
	PayDate     string          `json:"pay_date" gorm:"type:date;column:pay_date;comment:付费日期;uniqueIndex:ix_unique"`
	RootGameId  int64           `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:根游戏ID;uniqueIndex:ix_unique"`
	AgentId     int64           `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID;uniqueIndex:ix_unique"`
	SiteId      int64           `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID;uniqueIndex:ix_unique"`
	Ad3Id       int64           `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID;uniqueIndex:ix_unique"`
	RegDate     string          `json:"reg_date" gorm:"type:date;column:reg_date;comment:注册日期;uniqueIndex:ix_unique"`
	ActiveDays  int             `json:"active_days" gorm:"column:active_days;default:0;comment:注册到付费的活跃天数;"`
	ActiveCount int             `json:"active_count" gorm:"column:active_count;default:0;comment:注册到付费的活跃人数;"`
	PayMoneySum int             `json:"pay_money_sum" gorm:"column:pay_money_sum;default:0;comment:付费金额;"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DwsDayRootGameBackPayActiveLogModel) TableName() string {
	return "dws_day_root_game_back_pay_active_log"
}

func (receiver *DwsDayRootGameBackPayActiveLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwsDayRootGameBackPayActiveLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
