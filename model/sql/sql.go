package sql

import "time"

const (
	StatusNormal    = "normal"
	StatusDelete    = "delete"
	StatusForbidden = "forbidden"
	StatusRemove    = "remove"  // 下架
	StatusClaimed   = "claimed" // 已领取
)

type SqlBaseModel struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime(0);column:created_at;autoCreateTime;default:CURRENT_TIMESTAMP;comment:创建时间"`                                            // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime(0);column:updated_at;autoCreateTime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
}

type SqlCommonModel struct {
	GameId         int    `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID"`
	AgentId        int    `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID"`
	SiteId         int    `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID"`
	MediaSiteId    int    `json:"media_site_id" gorm:"column:media_site_id;default:0;comment:媒体广告位Id"`
	Idfa           string `json:"idfa" gorm:"size:100;column:idfa;default:'';comment:iOS idfa"`
	Idfv           string `json:"idfv" gorm:"size:100;column:idfv;default:'';comment:iOS idfv"`
	Imei           string `json:"imei" gorm:"size:100;column:imei;default:'';comment:安卓imei"`
	Oaid           string `json:"oaid" gorm:"size:150;column:oaid;default:'';comment:安卓oaid"`
	AndriodId      string `json:"andriod_id" gorm:"size:100;column:andriod_id;default:'';comment:安卓andriod_id"`
	SystemVersion  string `json:"system_version" gorm:"size:50;column:system_version;default:'';comment:系统版本号"`
	AppVersionCode int    `json:"app_version_code" gorm:"column:app_version_code;default:0;comment:app版本号"`
	SdkVersionCode int    `json:"sdk_version_code" gorm:"column:sdk_version_code;default:0;comment:sdk版本号"`
	Network        string `json:"network" gorm:"size:50;column:network;default:'';comment:网络"`
	ClientIp       string `json:"client_ip" gorm:"size:150;column:client_ip;default:'';comment:客户端IP"`
	Ipv4           string `json:"ipv4" gorm:"size:150;column:ipv4;default:'';comment:ipv4地址"`
	Ipv6           string `json:"ipv6" gorm:"size:150;column:ipv6;default:'';comment:ipv6地址"`
	ChannelId      int    `json:"channel_id" gorm:"column:channel_id;default:0;comment:联运渠道ID"`
	Model          string `json:"model" gorm:"size:50;column:model;default:'';comment:机型"`
	Brand          string `json:"brand" gorm:"size:50;column:brand;default:'';comment:品牌"`
	UserAgent      string `json:"user_agent" gorm:"size:1024;column:user_agent;default:'';comment:用户UA"`
}
