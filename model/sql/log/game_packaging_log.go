package log

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

type OdsGamePackagingLogModel struct {
	sql.SqlBaseModel
	Db              func() *gorm.DB `json:"-" gorm:"-"`
	PlatformId      int             `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID"`
	GameId          int             `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID"`
	AgentId         int             `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID"`
	SiteId          int             `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID"`
	Status          string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态"`
	GamePackagePath string          `json:"game_package_path" gorm:"size:512;column:game_package_path;default:'';comment:游戏包路径"`
	ExecCmd         string          `json:"exec_cmd" gorm:"size:512;column:exec_cmd;default:'';comment:执行的命令"`
	ExecCmdResult   string          `json:"exec_cmd_result" gorm:"type:text;column:exec_cmd_result;comment:执行的命令的结果"`
}

func (receiver *OdsGamePackagingLogModel) TableName() string {
	return "ods_game_packaging_log"
}

func (receiver *OdsGamePackagingLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsGamePackagingLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsGamePackagingLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
