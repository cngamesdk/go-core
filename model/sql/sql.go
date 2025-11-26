package sql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

const (
	StatusNormal    = "normal"
	StatusDelete    = "delete"
	StatusForbidden = "forbidden"
	StatusRemove    = "remove"  // 下架
	StatusClaimed   = "claimed" // 已领取
	StatusSuccess   = "success" // 成功
	StatusFail      = "fail"    // 失败
)

type SqlBaseModel struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime(0);column:created_at;autoCreateTime;default:CURRENT_TIMESTAMP;comment:创建时间"`                                            // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime(0);column:updated_at;autoCreateTime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"` // 更新时间
}

type SqlCommonModel struct {
	GameId         int64  `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID"`
	AgentId        int64  `json:"agent_id" gorm:"column:agent_id;default:0;comment:渠道ID"`
	SiteId         int64  `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID"`
	MediaSiteId    int64  `json:"media_site_id" gorm:"column:media_site_id;default:0;comment:媒体广告位Id"`
	Idfv           string `json:"idfv" gorm:"size:100;column:idfv;default:'';comment:iOS idfv"`
	Imei           string `json:"imei" gorm:"size:100;column:imei;default:'';comment:安卓imei/iOS为idfa,获取不到为空"`
	Oaid           string `json:"oaid" gorm:"size:150;column:oaid;default:'';comment:安卓oaid"`
	AndriodId      string `json:"andriod_id" gorm:"size:100;column:andriod_id;default:'';comment:安卓andriod_id"`
	SystemVersion  string `json:"system_version" gorm:"size:50;column:system_version;default:'';comment:系统版本号"`
	AppVersionCode int    `json:"app_version_code" gorm:"column:app_version_code;default:0;comment:app版本号"`
	SdkVersionCode int    `json:"sdk_version_code" gorm:"column:sdk_version_code;default:0;comment:sdk版本号"`
	Network        string `json:"network" gorm:"size:50;column:network;default:'';comment:网络"`
	ClientIp       string `json:"client_ip" gorm:"size:150;column:client_ip;default:'';comment:客户端IP"`
	Ipv4           string `json:"ipv4" gorm:"size:150;column:ipv4;default:'';comment:ipv4地址"`
	Ipv6           string `json:"ipv6" gorm:"size:150;column:ipv6;default:'';comment:ipv6地址"`
	ChannelId      int64  `json:"channel_id" gorm:"column:channel_id;default:0;comment:联运渠道ID"`
	Model          string `json:"model" gorm:"size:50;column:model;default:'';comment:机型"`
	Brand          string `json:"brand" gorm:"size:50;column:brand;default:'';comment:品牌"`
	UserAgent      string `json:"user_agent" gorm:"size:1024;column:user_agent;default:'';comment:用户UA"`
}

type JSON json.RawMessage

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// CustomMapType 自定义MAP类型
type CustomMapType map[string]interface{}

// Scan Scanner
func (args *CustomMapType) Scan(value interface{}) error {
	return scan(args, value)
}

// Value Valuer
func (args *CustomMapType) Value() (driver.Value, error) {
	return value(args)
}

// scan for scanner helper
func scan(data interface{}, value interface{}) error {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), data)
	case string:
		return json.Unmarshal([]byte(value.(string)), data)
	default:
		return fmt.Errorf("val type is valid, is %+v", value)
	}
}

// for valuer helper
func value(data interface{}) (interface{}, error) {
	vi := reflect.ValueOf(data)
	// 判断是否为 0 值
	if vi.IsZero() {
		return nil, nil
	}
	return json.Marshal(data)
}
