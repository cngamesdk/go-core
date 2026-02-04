package material

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	MaterialSourceOriginal = "original" // 原创
	MaterialSourceEditing  = "editing"  // 改动

	MaterialVisibilityPublic  = "public"  // 公开
	MaterialVisibilityPrivate = "private" // 私密

	MaterialTypeImage = "image" // 图片
	MaterialTypeVideo = "video" // 视频
	MaterialTypeAudio = "audio" // 音频
)

var (
	//素材来源
	MaterialSources = map[string]string{
		MaterialSourceOriginal: "原创",
		MaterialSourceEditing:  "原创",
	}
	//素材可见性
	MaterialVisibilities = map[string]string{
		MaterialVisibilityPublic:  "公开",
		MaterialVisibilityPrivate: "私密",
	}
	//素材类型
	MaterialTypes = map[string]string{
		MaterialTypeImage: "图片",
		MaterialTypeVideo: "视频",
		MaterialTypeAudio: "音频",
	}
)

// OdsMaterialLog 素材日志表
type OdsMaterialLogModel struct {
	sql2.SqlBaseModel
	PlatformId   int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name"`
	MaterialName string          `json:"material_name" gorm:"size:100;column:material_name;default:'';comment:素材名称;uniqueIndex:ix_plat_name"`
	ThemeId      int64           `json:"theme_id" gorm:"column:theme_id;default:0;comment:题材ID;"`
	Author       string          `json:"author" gorm:"size:50;column:author;default:'';comment:作者;"`
	Source       string          `json:"source" gorm:"size:50;column:source;default:'';comment:来源;"`
	Status       string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态;"`
	Visibility   string          `json:"visibility" gorm:"size:50;column:visibility;default:'';comment:可见性;"`
	MaterialType string          `json:"material_type" gorm:"size:50;column:material_type;default:'';comment:素材类型;"`
	Db           func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsMaterialLogModel) TableName() string {
	return "ods_material_log"
}

func (receiver *OdsMaterialLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsMaterialLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsMaterialLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
