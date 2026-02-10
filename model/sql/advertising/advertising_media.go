package advertising

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// 常见媒体码
const (
	MediaCodeOfficial    = "official"    // 官方
	MediaCodeTencent     = "tencent"     // 腾讯广告
	MediaCodeOceanengine = "oceanengine" // 巨量引擎
	MediaCodeKuaishou    = "kuaishou"    // 快手磁力引擎
	MediaCodeBaidu       = "baidu"       // 百度广告
	MediaCodeUc          = "uc"          // UC广告
	MediaCodeTaptap      = "taptap"      // TAPTAP广告
	MediaCodeZhihu       = "zhihu"       // 知乎广告
	MediaCodeBilibili    = "bilibili"    // 哔哩哔哩广告
	MediaCodeOthers      = "others"      // 其他广告
)

var (
	MediaCodesMap = map[string]string{
		MediaCodeOfficial:    "官方",
		MediaCodeTencent:     "腾讯广告",
		MediaCodeOceanengine: "巨量引擎",
		MediaCodeKuaishou:    "快手磁力引擎",
		MediaCodeBaidu:       "百度",
		MediaCodeUc:          "UC",
		MediaCodeTaptap:      "TAPTAP",
		MediaCodeZhihu:       "知乎",
		MediaCodeBilibili:    "哔哩哔哩",
		MediaCodeOthers:      "其他",
	}
)

// DimAdvertisingMediaModel 广告媒体维度表
type DimAdvertisingMediaModel struct {
	sql2.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name;uniqueIndex:ix_plat_code;"`
	MediaName  string          `json:"media_name" gorm:"size:100;column:media_name;default:'';comment:媒体名称;uniqueIndex:ix_plat_name"`
	Code       string          `json:"code" gorm:"size:100;column:code;default:'';comment:媒体码;uniqueIndex:ix_plat_code;"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
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
