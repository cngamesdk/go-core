package report

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DwsDayRootGameBackOverviewLogModel 根游戏回流每日总览表
type DwsDayRootGameBackOverviewLogModel struct {
	sql.SqlBaseModel
	PlatformId       int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_unique"`
	StatDate         string          `json:"stat_date" gorm:"type:date;column:stat_date;comment:统计日期;uniqueIndex:ix_unique"`
	RootGameId       int64           `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:根游戏ID;uniqueIndex:ix_unique"`
	AgentId          int64           `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID;uniqueIndex:ix_unique"`
	SiteId           int64           `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID;uniqueIndex:ix_unique"`
	Ad3Id            int64           `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID;uniqueIndex:ix_unique"`
	Activation       int             `json:"activation" gorm:"column:activation;default:0;comment:激活数;"`
	ActivationDevice int             `json:"activation_device" gorm:"column:activation_device;default:0;comment:激活设备数;"`
	Launch           int             `json:"launch" gorm:"column:launch;default:0;comment:启动数;"`
	LaunchDevice     int             `json:"launch_device" gorm:"column:launch_device;default:0;comment:启动设备数;"`
	Reg              int             `json:"reg" gorm:"column:reg;default:0;comment:注册数;"`
	RegDevice        int             `json:"reg_device" gorm:"column:reg_device;default:0;comment:注册设备数;"`
	Login            int             `json:"login" gorm:"column:login;default:0;comment:登录数;"`
	LoginUser        int             `json:"login_user" gorm:"column:login_user;default:0;comment:登录用户数;"`
	LoginDevice      int             `json:"login_device" gorm:"column:login_device;default:0;comment:登录设备数;"`
	Role             int             `json:"role" gorm:"column:role;default:0;comment:创角数;"`
	RoleUser         int             `json:"role_user" gorm:"column:role_user;default:0;comment:创角用户数;"`
	RoleDevice       int             `json:"role_device" gorm:"column:role_device;default:0;comment:创角设备数;"`
	Pay              int             `json:"pay" gorm:"column:pay;default:0;comment:付费数;"`
	PayUser          int             `json:"pay_user" gorm:"column:pay_user;default:0;comment:付费用户数;"`
	PayDevice        int             `json:"pay_device" gorm:"column:pay_device;default:0;comment:付费设备数;"`
	PayMoney         int             `json:"pay_money" gorm:"column:pay_money;default:0;comment:付费金额;"`
	Db               func() *gorm.DB `json:"-" gorm:"-"`
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
