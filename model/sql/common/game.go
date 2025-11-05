package common

import (
	"context"
	"fmt"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"github.com/duke-git/lancet/v2/cryptor"
	"gorm.io/gorm"
)

const (
	GameTypeMobileGame     = "mobile-game"
	GameTypeHtml5Game      = "html5-game"
	GameTypeWechatMiniGame = "wechat-mini-game"
	GameTypeDouyinMiniGame = "douyin-mini-game"
	OsAndroid              = "android"
	OsIos                  = "ios"
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

// DimGameModel 游戏维度
type DimGameModel struct {
	sql2.SqlBaseModel
	GameName    string          `json:"game_name" gorm:"size:100;column:game_name;default:'';comment:游戏名称"`
	PackageName string          `json:"package_name" gorm:"size:100;column:package_name;default:'';comment:包名"`
	GameType    string          `json:"game_type" gorm:"size:50;column:game_type;default:'';comment:游戏类型"`
	Os          string          `json:"os" gorm:"size:50;column:os;default:'';comment:操作系统"`
	CpUrl       string          `json:"cp_url" gorm:"size:1024;column:cp_url;default:'';comment:发货地址"`
	MainId      int64           `json:"main_id" gorm:"column:main_id;default:0;comment:主游戏ID"`
	GameRate    int             `json:"game_rate" gorm:"column:game_rate;default:0;comment:游戏币兑换比例"`
	CpGameId    int64           `json:"cp_game_id" gorm:"column:cp_game_id;default:0;comment:研发对接的游戏ID"`
	CompanyId   int64           `json:"company_id" gorm:"column:company_id;default:0;comment:主体ID"`
	Db          func() *gorm.DB `json:"-" gorm:"-"`
	GetHashKey  func() string   `json:"-" gorm:"-"`
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

func (receiver *DimGameModel) GetAppKey() string {
	return cryptor.Md5String(fmt.Sprintf("game-app-key-%d-%s", receiver.Id, receiver.GetHashKey()))
}

func (receiver *DimGameModel) GetLoginKey() string {
	gameId := receiver.Id
	if receiver.CpGameId > 0 {
		gameId = receiver.CpGameId
	}
	return cryptor.Md5String(fmt.Sprintf("game-login-key-%d-%s", gameId, receiver.GetHashKey()))
}

func (receiver *DimGameModel) GetPayKey() string {
	gameId := receiver.Id
	if receiver.CpGameId > 0 {
		gameId = receiver.CpGameId
	}
	return cryptor.Md5String(fmt.Sprintf("game-pay-key-%d-%s", gameId, receiver.GetHashKey()))
}
