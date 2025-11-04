package common

import (
	"context"
	sql2 "github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

// DimRootGameModel 主游戏维度
type DimRootGameModel struct {
	sql2.SqlBaseModel
	GameName           string          `json:"game_name" gorm:"size:100;column:game_name;default:'';comment:根游戏名称"`
	ContractName       string          `json:"contract_name" gorm:"size:100;column:contract_name;default:'';comment:合同游戏名称"`
	ProfitSharingRatio int             `json:"profit_sharing_ratio" gorm:"column:profit_sharing_ratio;default:0;comment:与研发分成比例，如30为30%"`
	Db                 func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *DimRootGameModel) TableName() string {
	return "dim_root_game"
}

func (receiver *DimRootGameModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *DimRootGameModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *DimRootGameModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
