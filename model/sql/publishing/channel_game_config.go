package publishing

import (
	"context"
	"encoding/json"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimPublishingChannelGameConfigModel 发行渠道游戏配置表
type DimPublishingChannelGameConfigModel struct {
	sql.SqlBaseModel
	PlatformId   int64                  `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_game_channel_site"`
	GameId       int64                  `json:"game_id" gorm:"column:game_id;default:0;comment:游戏ID;uniqueIndex:ix_plat_game_channel_site"`
	ChannelId    int64                  `json:"channel_id" gorm:"column:channel_id;default:0;comment:发行渠道ID;uniqueIndex:ix_plat_game_channel_site"`
	AgentId      int64                  `json:"gift_id" gorm:"column:gift_id;default:0;comment:渠道ID"`
	SiteId       int64                  `json:"site_id" gorm:"column:site_id;default:0;comment:广告位ID;uniqueIndex:ix_plat_game_channel_site"`
	Config       string                 `json:"config" gorm:"type:text;column:config;comment:json配置"`
	ConfigFormat map[string]interface{} `json:"config_format" gorm:"-"`
	Db           func() *gorm.DB        `json:"-" gorm:"-"`
}

func (receiver *DimPublishingChannelGameConfigModel) TableName() string {
	return "dim_publishing_channel_game_config"
}

func (receiver *DimPublishingChannelGameConfigModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimPublishingChannelGameConfigModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimPublishingChannelGameConfigModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}

func (receiver *DimPublishingChannelGameConfigModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimPublishingChannelGameConfigModel) findHook(tx *gorm.DB) (err error) {
	if receiver.Config != "" && receiver.ConfigFormat == nil {
		receiver.ConfigFormat = make(map[string]interface{})
		if jsonErr := json.Unmarshal([]byte(receiver.Config), &receiver.ConfigFormat); jsonErr != nil {
			err = jsonErr
			return
		}
	}
	return
}
