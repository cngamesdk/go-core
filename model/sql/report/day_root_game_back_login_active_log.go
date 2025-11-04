package report

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwsDayRootGameBackLoginActiveLogModel 根游戏回流每日活跃报表
type DwsDayRootGameBackLoginActiveLogModel struct {
	sql.SqlBaseModel
	LoginDate   string          `json:"login_date" gorm:"type:date;column:login_date;comment:登录日期;uniqueIndex:ix_unique"`
	RootGameId  int             `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:根游戏ID;uniqueIndex:ix_unique"`
	AgentId     int             `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID;uniqueIndex:ix_unique"`
	SiteId      int             `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID;uniqueIndex:ix_unique"`
	Ad3Id       uint64          `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID;uniqueIndex:ix_unique"`
	RegDate     string          `json:"reg_date" gorm:"type:date;column:reg_date;comment:注册日期;uniqueIndex:ix_unique"`
	ActiveDays  uint            `json:"active_days" gorm:"column:active_days;default:0;comment:活跃天数;"`
	ActiveCount uint            `json:"active_count" gorm:"column:active_count;default:0;comment:活跃人数;"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DwsDayRootGameBackLoginActiveLogModel) TableName() string {
	return "dws_day_root_game_back_login_active_log"
}

func (receiver *DwsDayRootGameBackLoginActiveLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwsDayRootGameBackLoginActiveLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
