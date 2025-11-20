package advertising

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

type OdsMediaAdClickLogModel struct {
	sql.SqlBaseModel
	PlatformId    int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台Id"`
	MediaId       int64           `json:"media_id" gorm:"column:media_id;default:0;comment:媒体Id"`
	GameId        int64           `json:"game_id" gorm:"column:game_id;default:0;comment:子游戏ID"`
	AgentId       int64           `json:"agent_id" gorm:"column:agent;default:0;comment:渠道ID"`
	SiteId        int64           `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID"`
	AccountId     int64           `json:"account_id" gorm:"column:account_id;default:0;comment:帐户ID"`
	Ad1Id         int64           `json:"ad1_id" gorm:"column:ad1_id;default:0;comment:广告一级ID"`
	Ad2Id         int64           `json:"ad2_id" gorm:"column:ad2_id;default:0;comment:广告二级ID"`
	Ad3Id         int64           `json:"ad3_id" gorm:"column:ad3_id;default:0;comment:广告三级ID"`
	ClickTime     time.Time       `json:"click_time" gorm:"type:datetime(0);column:click_time;comment:点击时间"`
	ClickId       string          `json:"click_id" gorm:"size:100;column:click_id;comment:点击ID"`
	RequestId     string          `json:"request_id" gorm:"size:100;column:request_id;comment:点击请求ID"`
	Csite         string          `json:"csite" gorm:"size:50;column:csite;comment:广告资源位"`
	Ipv4          string          `json:"ipv4" gorm:"size:20;column:ipv4;comment:ipv4地址"`
	Ipv6          string          `json:"ipv6" gorm:"size:50;column:ipv6;comment:ipv6地址"`
	Muid          string          `json:"muid" gorm:"size:50;column:muid;comment:设备号MD5值"`
	Oaid          string          `json:"oaid" gorm:"size:100;column:oaid;comment:oaid"`
	Moaid         string          `json:"moaid" gorm:"size:50;column:moaid;comment:oaid MD5值"`
	DeviceModel   string          `json:"device_model" gorm:"size:50;column:device_model;comment:设备机型"`
	UserAgent     string          `json:"user_agent" gorm:"size:512;column:user_agent;comment:用户UA"`
	Callback      string          `json:"callback" gorm:"size:512;column:callback;comment:回调地址"`
	CallbackParam string          `json:"callback_param" gorm:"size:512;column:callback_param;comment:回调参数"`
	CaidData      string          `json:"caid_data" gorm:"type:text;column:callback;comment:caid数据"`
	Mcaid1        string          `json:"mcaid1" gorm:"size:50;column:mcaid1;comment:caid1 MD5值"`
	Mcaid2        string          `json:"mcaid2" gorm:"size:50;column:mcaid2;comment:caid2 MD5值"`
	LiveAccount   string          `json:"live_account" gorm:"size:50;column:live_account;comment:直播账号"`
	OpenId        string          `json:"open_id" gorm:"size:50;column:open_id;comment:open_id"`
	Db            func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsMediaAdClickLogModel) TableName() string {
	return "ods_media_ad_click_log"
}

func (receiver *OdsMediaAdClickLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsMediaAdClickLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsMediaAdClickLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
