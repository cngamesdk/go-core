package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwdDayGameRegUidLoginLogModel 子注册
type DwdDayGameRegUidLoginLogModel struct {
	sql.SqlBaseModel
	sql.SqlCommonModel
	GameId         int64                `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID;uniqueIndex:ix_plat_game_user_date"`
	UserId         int64                `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;uniqueIndex:ix_plat_game_user_date"`
	LoginDate      string               `json:"login_date" gorm:"type:date;column:login_date;default:'1970-01-01';comment:登录日期;uniqueIndex:ix_plat_game_user_date"`
	FirstLoginTime sql.MyCustomDatetime `json:"first_login_time" gorm:"type:datetime;column:first_login_time;default:'';comment:首次登录时间;"`
	LastLoginTime  sql.MyCustomDatetime `json:"last_login_time" gorm:"type:datetime;column:last_login_time;default:'';comment:最后登录时间;"`
	LoginCount     int                  `json:"login_count" gorm:"size:32;column:login_count;default:0;comment:登录次数;"`
	RegTime        sql.MyCustomDatetime `json:"reg_time" gorm:"type:datetime;column:login_count;default:0;comment:登录次数;"`
	Db             func() *gorm.DB      `json:"-" gorm:"-"`
}

func (receiver *DwdDayGameRegUidLoginLogModel) TableName() string {
	return "dwd_day_game_reg_uid_login_log"
}

func (receiver *DwdDayGameRegUidLoginLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwdDayGameRegUidLoginLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
