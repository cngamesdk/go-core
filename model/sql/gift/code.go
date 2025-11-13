package gift

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
)

const (
	GiftCodeStatusNormal  = sql.StatusNormal
	GiftCodeStatusClaimed = sql.StatusClaimed
)

// OdsGiftCodeListModel 礼包码列表
type OdsGiftCodeListModel struct {
	sql.SqlBaseModel
	PlatformId int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID;uniqueIndex:ix_plat_gift_code"`
	GiftId     int64           `json:"gift_id" gorm:"column:gift_id;default:0;comment:礼包码ID;uniqueIndex:ix_plat_gift_code"`
	Code       string          `json:"code" gorm:"size:100;column:code;default:'';comment:礼包码;uniqueIndex:ix_plat_gift_code"`
	Status     string          `json:"status" gorm:"size:100;column:status;default:'';comment:领取状态"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsGiftCodeListModel) TableName() string {
	return "ods_gift_code_list"
}

func (receiver *OdsGiftCodeListModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsGiftCodeListModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsGiftCodeListModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}
