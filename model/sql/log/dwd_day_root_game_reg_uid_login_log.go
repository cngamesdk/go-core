package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwdDayRootGameRegUidLoginLogModel 根注册按天日志表
type DwdDayRootGameRegUidLoginLogModel struct {
	sql.SqlBaseModel
	sql.SqlCommonModel
	PlatformId     int64                `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_rootgame_user_date"`
	RootGameId     int64                `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:根游戏ID;uniqueIndex:ix_plat_rootgame_user_date"`
	UserId         int64                `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;uniqueIndex:ix_plat_rootgame_user_date"`
	LoginDate      string               `json:"login_date" gorm:"type:date;column:login_date;default:'1970-01-01';comment:登录日期;uniqueIndex:ix_plat_rootgame_user_date"`
	FirstLoginTime sql.MyCustomDatetime `json:"first_login_time" gorm:"type:datetime;column:first_login_time;comment:首次登录时间;"`
	LastLoginTime  sql.MyCustomDatetime `json:"last_login_time" gorm:"type:datetime;column:last_login_time;comment:最后登录时间;"`
	LoginCount     int                  `json:"login_count" gorm:"size:32;column:login_count;default:0;comment:登录次数;"`
	RegTime        sql.MyCustomDatetime `json:"reg_time" gorm:"type:datetime;column:reg_time;comment:按根注册时间;"`
	Db             func() *gorm.DB      `json:"-" gorm:"-"`
}

func (receiver *DwdDayRootGameRegUidLoginLogModel) TableName() string {
	return "dwd_day_root_game_reg_uid_login_log"
}

func (receiver *DwdDayRootGameRegUidLoginLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwdDayRootGameRegUidLoginLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
