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
	PlatformId int64     `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_game_user"`
	GameId     int       `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID;uniqueIndex:ix_plat_game_user"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID;uniqueIndex:ix_plat_game_user"`
	RegTime    time.Time `json:"reg_time" gorm:"type:datetime(0);column:reg_time;comment:注册时间"`
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
