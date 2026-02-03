package material

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	MaterialFileSourceOriginal  = "original"  // 原创
	MaterialFileSourceExtension = "extension" // 扩展
)

var (
	//素材文件类型
	MaterialFileSources = map[string]string{
		MaterialFileSourceOriginal:  "原创",
		MaterialFileSourceExtension: "扩展",
	}
)

// ods_material_file_log 素材文件日志表
type OdsMaterialFileLog struct {
	sql2.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;index:ix_plat_mid"`
	MaterialId int64           `json:"material_id" gorm:"column:material_id;default:0;comment:素材ID;index:ix_plat_mid"`
	Source     string          `json:"source" gorm:"size:50;column:source;default:'';comment:来源;"`
	FileName   string          `json:"file_name" gorm:"size:512;column:file_name;default:'';comment:文件名称;"`
	Url        string          `json:"url" gorm:"size:512;column:url;default:'';comment:文件相对路径;"`
	Status     string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	Visibility string          `json:"visibility" gorm:"size:50;column:visibility;default:'';comment:可见性;"`
	Signature  string          `json:"signature" gorm:"size:50;column:signature;default:'';comment:文件签名;uniqueIndex:ix_sign"`
	Width      int             `json:"width" gorm:"size:32;column:width;default:0;comment:宽度;"`
	Height     int             `json:"height" gorm:"size:32;column:height;default:0;comment:高度;"`
	FileType   string          `json:"file_type" gorm:"size:50;column:file_type;default:'';comment:文件类型;"`
	Duration   int             `json:"duration" gorm:"size:32;column:duration;default:0;comment:时长，单位：毫秒;"`
	Bitrate    int             `json:"bitrate" gorm:"size:32;column:bitrate;default:0;comment:码率kbps;"`
	Size       int             `json:"size" gorm:"size:32;column:size;default:0;comment:文件大小(bit);"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsMaterialFileLog) TableName() string {
	return "ods_material_file_log"
}

func (receiver *OdsMaterialFileLog) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsMaterialFileLog) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsMaterialFileLog) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
