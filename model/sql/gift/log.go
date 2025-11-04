package gift

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

// OdsGiftClaimLogModel 礼包领取列表
type OdsGiftClaimLogModel struct {
	sql.SqlBaseModel
	UserId     int64           `json:"user_id" gorm:"column:user_id;default:0;comment:用户ID"`
	ActionTime time.Time       `json:"action_time" gorm:"type:datetime(0);column:action_time;comment:领取时间"`
	GiftId     int64           `json:"gift_id" gorm:"column:gift_id;default:0;comment:礼包ID"`
	Code       string          `json:"code" gorm:"size:50;column:code;comment:结束时间"`
	Db         func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsGiftClaimLogModel) TableName() string {
	return "ods_gift_claim_log"
}

func (receiver *OdsGiftClaimLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsGiftClaimLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
