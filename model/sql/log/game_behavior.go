package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

const (
	GameBehaviorActionSelectServer = "server-select" // 选择区服
	GameBehaviorActionCreateRole   = "role-create"   // 创建角色
	GameBehaviorActionEnterGame    = "game-enter"    // 进入游戏
	GameBehaviorActionLevelUp      = "level-up"      // 等级提升
)

type OdsGameBehaviorLogModel struct {
	sql.SqlBaseModel
	PlatformId int64     `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID"`
	UserId     int64     `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID"`
	ServerId   int64     `json:"server_id" gorm:"column:server_id;default:0;comment:区服ID"`
	ServerName string    `json:"server_name" gorm:"size:100;column:server_name;default:'';comment:区服名称"`
	RoleId     string    `json:"role_id" gorm:"size:512;column:role_id;default:'';comment:角色ID"`
	RoleName   string    `json:"role_name" gorm:"size:512;column:role_name;default:'';comment:角色名称"`
	Action     string    `json:"action" gorm:"size:50;column:action;default:'';comment:行为"`
	ActionTime time.Time `json:"action_time" gorm:"type:datetime(0);column:action_time;comment:行为时间"`
	RoleLevel  int       `json:"role_level" gorm:"default:0;column:role_level;comment:角色等级"`
	sql.SqlCommonModel
	Db func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsGameBehaviorLogModel) TableName() string {
	return "ods_game_behavior_log"
}

func (receiver *OdsGameBehaviorLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsGameBehaviorLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
