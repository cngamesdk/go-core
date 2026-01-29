package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwdRootGameBackRegLogModel 30天按根游戏回流注册日志
type DwdRootGameBackRegLogModel struct {
	sql.SqlBaseModel
	PlatformId       int64                `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_game_user_time"`
	RootGameId       int                  `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:游戏ID;uniqueIndex:ix_plat_game_user_time"`
	UserId           int64                `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;uniqueIndex:ix_plat_game_user_time"`
	RegTime          sql.MyCustomDatetime `json:"reg_time" gorm:"type:datetime(0);column:reg_time;comment:注册时间;uniqueIndex:ix_plat_game_user_time"`
	LastTime         sql.MyCustomDatetime `json:"last_time" gorm:"type:datetime(0);column:last_time;comment:最后时间:最后登录时间+30天;"`
	Ad3Id            int64                `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID"`
	FirstDayPayTime  sql.MyCustomDatetime `json:"first_day_pay_time" gorm:"type:datetime(0);column:first_day_pay_time;comment:首日付费时间;"`
	FirstDayPayCount int                  `json:"first_day_pay_count" gorm:"size:32;column:first_day_pay_count;default:0;comment:首日付费次数;"`
	FirstDayPayMoney int                  `json:"first_day_pay_money" gorm:"size:32;column:first_day_pay_money;default:0;comment:首日付费金额,单位:分;"`
	FirstPayTime     sql.MyCustomDatetime `json:"first_pay_time" gorm:"type:datetime(0);column:first_pay_time;comment:首次付费时间;"`
	FirstPayMoney    int                  `json:"first_pay_money" gorm:"size:32;column:first_pay_money;default:0;comment:首次付费金额,单位:分;"`
	TotalPayCount    int                  `json:"total_pay_count" gorm:"size:32;column:total_pay_count;default:0;comment:累计付费次数;"`
	TotalPayMoney    int                  `json:"total_pay_money" gorm:"size:32;column:total_pay_money;default:0;comment:累计付费金额,单位:分;"`
	LastLoginTime    sql.MyCustomDatetime `json:"last_login_time" gorm:"type:datetime(0);column:last_login_time;comment:最后登录时间;"`
	TotalLoginCount  int                  `json:"total_login_count" gorm:"size:32;column:total_login_count;default:0;comment:累计登录次数;"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DwdRootGameBackRegLogModel) TableName() string {
	return "dwd_root_game_back_reg_log"
}

func (receiver *DwdRootGameBackRegLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwdRootGameBackRegLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
