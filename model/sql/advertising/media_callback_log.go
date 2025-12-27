package advertising

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

const (
	MediaCallbackSourceApi     = "api"
	MediaCallbackSourceSdk     = "sdk"
	MediaCallbackEventActive   = "active" //激活
	MediaCallbackEventReg      = "reg"    //注册
	MediaCallbackEventLogin    = "login"  //登录
	MediaCallbackEventPay      = "pay"    //登录
	MediaCallbackStatusSuccess = sql.StatusSuccess
	MediaCallbackStatusFail    = sql.StatusFail
)

type OdsMediaCallbackLogModel struct {
	sql.SqlBaseModel
	PlatformId   int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台Id"`
	GameId       int64           `json:"game_id" gorm:"column:game_id;default:0;comment:子游戏ID"`
	SiteId       int64           `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID"`
	UserId       int64           `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID"`
	Source       string          `json:"source" gorm:"column:source;default:'';comment:来源"`
	Event        string          `json:"event" gorm:"column:event;default:'';comment:事件"`
	OrderId      string          `json:"order_id" gorm:"size:150;column:order_id;default:'';comment:订单号"`
	ReportMoney  int             `json:"report_money" gorm:"column:report_money;default:0;comment:上报金额"`
	CallbackReq  string          `json:"callback_req" gorm:"type:text;column:callback_req;default:'';comment:上报请求json格式.如:{url:,param:}"`
	CallbackResp string          `json:"callback_resp" gorm:"type:text;column:callback_resp;default:'';comment:上报返回json格式.如:{resp:}"`
	CallbackTime time.Time       `json:"callback_time" gorm:"type:datetime(0);column:callback_time;comment:上报时间"`
	Status       string          `json:"status" gorm:"size:50;column:status;default:'';comment:上报状态"`
	Db           func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsMediaCallbackLogModel) TableName() string {
	return "ods_media_callback_log"
}

func (receiver *OdsMediaCallbackLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsMediaCallbackLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsMediaCallbackLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
