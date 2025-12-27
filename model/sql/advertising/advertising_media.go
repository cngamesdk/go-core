package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// 常见媒体
const (
	CommonMediaOfficial    = "official"    // 官方
	CommonMediaTencent     = "tencent"     // 腾讯广告
	CommonMediaOceanengine = "oceanengine" // 巨量引擎
	CommonMediaKuaishou    = "kuaishou"    // 快手磁力引擎
	CommonMediaBaidu       = "baidu"       // 百度广告
	CommonMediaUc          = "uc"          // UC广告
	CommonMediaTaptap      = "taptap"      // TAPTAP广告
	CommonMediaZhihu       = "zhihu"       // 知乎广告
	CommonMediaBilibili    = "bilibili"    // 哔哩哔哩广告
	CommonMediaOthers      = "others"      // 其他广告
)

var (
	CommonMediasMap = map[string]string{
		CommonMediaOfficial:    "官方",
		CommonMediaTencent:     "腾讯广告",
		CommonMediaOceanengine: "巨量引擎",
		CommonMediaKuaishou:    "快手磁力引擎",
		CommonMediaBaidu:       "百度",
		CommonMediaUc:          "UC",
		CommonMediaTaptap:      "TAPTAP",
		CommonMediaZhihu:       "知乎",
		CommonMediaBilibili:    "哔哩哔哩",
		CommonMediaOthers:      "其他",
	}
)

// DimAdvertisingMediaModel 广告媒体维度表
type DimAdvertisingMediaModel struct {
	sql2.SqlBaseModel
	PlatformId           int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name"`
	AdvertisingMediaName string          `json:"advertising_media_name" gorm:"size:100;column:advertising_media_name;default:'';comment:广告媒体名称;uniqueIndex:ix_plat_name"`
	BelongCommonMedia    string          `json:"belong_common_media" gorm:"size:100;column:belong_common_media;default:'';comment:归属通用媒体;"`
	Db                   func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimAdvertisingMediaModel) TableName() string {
	return "dim_advertising_media"
}

func (receiver *DimAdvertisingMediaModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimAdvertisingMediaModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
