package common

import (
	"context"
	"fmt"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"github.com/duke-git/lancet/v2/cryptor"
	"gorm.io/gorm"
)

const (
	GameTypeMobileGame             = "mobile-game"
	GameTypeHtml5Game              = "html5-game"
	GameTypeWechatMiniGame         = "wechat-mini-game"
	GameTypeDouyinMiniGame         = "douyin-mini-game"
	OsAndroid                      = "android"
	OsIos                          = "ios"
	GameStatusNormal               = sql2.StatusNormal
	GameStatusRemove               = sql2.StatusRemove
	CooperationModelSelfOperation  = "self-operation"  // 自营
	CooperationModelJointOperation = "joint-operation" // 联运
)

var (
	GameTypes = map[string]string{
		GameTypeMobileGame:     "手游",
		GameTypeHtml5Game:      "H5游戏",
		GameTypeWechatMiniGame: "微信小游戏",
		GameTypeDouyinMiniGame: "抖音小游戏",
	}
	GameOss = map[string]string{
		OsAndroid: "安卓",
		OsIos:     "iOS",
	}
	GameStatuss = map[string]string{
		GameStatusNormal: "正常",
		GameStatusRemove: "下架",
	}
	CooperationModels = map[string]string{
		CooperationModelSelfOperation:  "自营",
		CooperationModelJointOperation: "联运",
	}
)

// GetGameTypeName 获取游戏类型名称
func GetGameTypeName(req string) string {
	resp, ok := GameTypes[req]
	if !ok {
		return ""
	}
	return resp
}

// GetGameOsName 获取游戏系统名称
func GetGameOsName(req string) string {
	resp, ok := GameOss[req]
	if !ok {
		return ""
	}
	return resp
}

// GetStatusName 获取状态名称
func GetStatusName(req string) string {
	resp, ok := GameStatuss[req]
	if !ok {
		return ""
	}
	return resp
}

func GetCooperationModelName(req string) string {
	resp, ok := CooperationModels[req]
	if !ok {
		return ""
	}
	return resp
}

// DimGameModel 游戏维度
type DimGameModel struct {
	sql2.SqlBaseModel
	PlatformId       int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name"`
	GameName         string          `json:"game_name" gorm:"size:100;column:game_name;default:'';comment:游戏名称;uniqueIndex:ix_plat_name"`
	PackageName      string          `json:"package_name" gorm:"size:150;column:package_name;default:'';comment:包名"`
	AppId            string          `json:"app_id" gorm:"size:100;column:app_id;default:'';comment:应用ID"`
	AppName          string          `json:"app_name" gorm:"size:100;column:app_name;default:'';comment:应用名称"`
	GameType         string          `json:"game_type" gorm:"size:50;column:game_type;default:'';comment:游戏类型"`
	Os               string          `json:"os" gorm:"size:50;column:os;default:'';comment:操作系统"`
	MainId           int64           `json:"main_id" gorm:"column:main_id;default:0;comment:主游戏ID"`
	GameCoinName     string          `json:"game_coin_name" gorm:"column:game_coin_name;default:'';comment:游戏币名称"`
	GameRate         int             `json:"game_rate" gorm:"column:game_rate;default:0;comment:游戏币兑换比例"`
	CompanyId        int64           `json:"company_id" gorm:"column:company_id;default:0;comment:主体ID"`
	Status           string          `json:"status" gorm:"size:50;column:status;default:'';comment:游戏状态"`
	CooperationModel string          `json:"cooperation_model" gorm:"size:50;column:cooperation_model;default:'';comment:合作方式"`
	CpGameId         int64           `json:"cp_game_id" gorm:"column:cp_game_id;default:0;comment:研发对接游戏ID"`
	Db               func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimGameModel) TableName() string {
	return "dim_game"
}

func (receiver *DimGameModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimGameModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimGameModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}

// GetGameAppKey 获取与APP前端通信密钥
func GetGameAppKey(gameId int64, key string) string {
	return cryptor.Md5String(fmt.Sprintf("game-app-key-%d-%s", gameId, key))
}

// GetGameLoginKey 获取二次验证通信密钥
func GetGameLoginKey(gameId int64, key string) string {
	return cryptor.Md5String(fmt.Sprintf("game-login-key-%d-%s", gameId, key))
}

// GetGamePayKey 获取支付通信密钥
func GetGamePayKey(gameId int64, key string) string {
	return cryptor.Md5String(fmt.Sprintf("game-pay-key-%d-%s", gameId, key))
}
