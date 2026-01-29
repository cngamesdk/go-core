package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

// DwdGameRegLogModel 按子游戏注册日志
type DwdGameRegLogModel struct {
	sql.SqlBaseModel
	PlatformId      int64                `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_game_user"`
	GameId          int                  `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID;uniqueIndex:ix_plat_game_user"`
	UserId          int64                `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;uniqueIndex:ix_plat_game_user"`
	RegTime         time.Time            `json:"reg_time" gorm:"type:datetime(0);column:reg_time;comment:注册时间"`
	Ad3Id           int64                `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID"`
	LastLoginTime   sql.MyCustomDatetime `json:"last_login_time" gorm:"type:datetime(0);column:last_login_time;comment:最后登录时间"`
	TotalLoginCount int                  `json:"total_login_count" gorm:"size:32;column:total_login_count;default:0;comment:总登录次数"`
	TotalPayCount   int                  `json:"total_pay_count" gorm:"size:32;column:total_pay_count;default:0;comment:总付费次数"`
	TotalPayMoney   int                  `json:"total_pay_money" gorm:"size:32;column:total_pay_money;default:0;comment:总付费金额"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DwdGameRegLogModel) TableName() string {
	return "dwd_game_reg_log"
}

func (receiver *DwdGameRegLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DwdGameRegLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
