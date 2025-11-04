package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimMainGameModel 主游戏维度
type DimMainGameModel struct {
	sql2.SqlBaseModel
	RootGameId int64           `json:"root_game_id" gorm:"column:root_game_id;default:0;comment:根游戏ID"`
	GameName   string          `json:"game_name" gorm:"size:100;column:game_name;default:'';comment:主游戏名称"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimMainGameModel) TableName() string {
	return "dim_main_game"
}

func (receiver *DimMainGameModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimMainGameModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimMainGameModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
