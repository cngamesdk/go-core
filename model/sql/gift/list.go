package gift

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"gorm.io/gorm"
	"time"
)

const (
	GiftStatusNormal = sql.StatusNormal
	GiftStatusRemove = sql.StatusRemove
)

// OdsGiftListModel 礼包列表
type OdsGiftListModel struct {
	sql.SqlBaseModel
	PlatformId   int64           `json:"platform_id" gorm:"column:platform_id;default:0;comment:平台ID"`
	Icon         string          `json:"icon" gorm:"size:1024;column:icon;default:'';comment:礼包ICON"`
	Title        string          `json:"title" gorm:"size:100;column:title;default:'';comment:标题"`
	Desc         string          `json:"desc" gorm:"size:512;column:desc;default:'';comment:简介"`
	Introduce    string          `json:"introduce" gorm:"type:text;column:introduce;comment:介绍"`
	Status       string          `json:"status" gorm:"size:50;column:status;default:'';comment:状态"`
	StartTime    time.Time       `json:"start_time" gorm:"type:datetime(0);column:start_time;comment:开始时间"`
	EndTime      time.Time       `json:"end_time" gorm:"type:datetime(0);column:end_time;comment:结束时间"`
	TotalNum     int             `json:"total_num" gorm:"column:total_num;default:0;comment:总数"`
	AvailableNum int             `json:"available_num" gorm:"column:available_num;default:0;comment:可用数"`
	Db           func() *gorm.DB `json:"-" gorm:"-"`
}

func (receiver *OdsGiftListModel) TableName() string {
	return "ods_gift_list"
}

func (receiver *OdsGiftListModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

// 是否有效
func (receiver *OdsGiftListModel) Valid() bool {
	if receiver.Status != GiftStatusNormal {
		return false
	}
	if receiver.AvailableNum <= 0 {
		return false
	}
	if receiver.StartTime.Unix() > time.Now().Unix() || time.Now().Unix() > receiver.EndTime.Unix() {
		return false
	}
	return true
}

func (receiver *OdsGiftListModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}
