package material

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimMaterialTheme 素材题材维度表
type DimMaterialTheme struct {
	sql2.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_name_pid"`
	ThemeName  string          `json:"theme_name" gorm:"size:100;column:theme_name;default:'';comment:题材名称;uniqueIndex:ix_plat_name_pid"`
	ParentId   int64           `json:"parent_id" gorm:"column:parent_id;default:0;comment:父题材ID;uniqueIndex:ix_plat_name_pid"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimMaterialTheme) TableName() string {
	return "dim_material_theme"
}

func (receiver *DimMaterialTheme) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimMaterialTheme) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimMaterialTheme) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args...).Updates(receiver).Error
	return
}
