package user

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	UserStatusNormal    = sql.StatusNormal    //正常
	UserStatusDelete    = sql.StatusDelete    //注销
	UserStatusForbidden = sql.StatusForbidden //封禁
)

var (
	UserStatusMap = map[string]string{
		UserStatusNormal:    "正常",
		UserStatusDelete:    "注销",
		UserStatusForbidden: "封禁",
	}
)

type OdsUserInfoLogModel struct {
	sql.SqlBaseModel
	UserName      string                        `json:"user_name" gorm:"size:50;column:user_name;default:'';comment:用户名;uniqueIndex:ix_user_name"`
	Password      string                        `json:"-" gorm:"-"`
	PasswordCrypt string                        `json:"-" gorm:"size:50;column:password_crypt;default:'';comment:密码加密"`
	Phone         string                        `json:"phone" gorm:"-"`
	PhoneCrypt    string                        `json:"phone_crypt" gorm:"size:512;column:phone_crypt;default:'';comment:手机号加密"`
	TrueName      string                        `json:"true_name" gorm:"size:50;column:true_name;default:'';comment:真实姓名"`
	IdCard        string                        `json:"id_card" gorm:"-"`
	IdCardCrypt   string                        `json:"id_card_crypt" gorm:"size:512;column:id_card_crypt;default:'';comment:身份证加密"`
	Version       int64                         `json:"version" gorm:"column:version;default:0;comment:用户版本"`
	Status        string                        `json:"status" gorm:"size:50;column:status;default:'normal';comment:用户状态"`
	Salt          string                        `json:"salt" gorm:"size:50;column:salt;default:'';comment:密码加盐"`
	Db            func() *gorm.DB               `json:"-" gorm:"-"`
	GetAesKey     func() string                 `json:"-" gorm:"-"`
	GetHashKey    func() string                 `json:"-" gorm:"-"`
	UpdateHook    func(tx *gorm.DB) (err error) `json:"-" gorm:"-"`
}

func (receiver *OdsUserInfoLogModel) TableName() string {
	return "ods_user_info_log"
}

func (receiver *OdsUserInfoLogModel) Take(ctx context.Context, fields string, query string, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Select(fields).Where(query, args...).Take(receiver).Error
	return
}

func (receiver *OdsUserInfoLogModel) Create(ctx context.Context) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Create(receiver).Error
	return
}

func (receiver *OdsUserInfoLogModel) Updates(ctx context.Context, query interface{}, args ...interface{}) (err error) {
	err = receiver.Db().WithContext(ctx).Table(receiver.TableName()).Where(query, args).Updates(receiver).Error
	return
}

func (receiver *OdsUserInfoLogModel) Valid() (resp bool) {
	return receiver.Status == UserStatusNormal
}

func (receiver *OdsUserInfoLogModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *OdsUserInfoLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *OdsUserInfoLogModel) BeforeSave(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *OdsUserInfoLogModel) BeforeUpdate(tx *gorm.DB) (err error) {
	return receiver.beforeCreateOrUpdateHook(tx)
}

func (receiver *OdsUserInfoLogModel) beforeCreateOrUpdateHook(tx *gorm.DB) (err error) {
	if receiver.Phone != "" && receiver.PhoneCrypt == "" {
		receiver.PhoneCrypt = cryptor.Base64StdEncode(string(cryptor.AesEcbEncrypt([]byte(receiver.Phone), []byte(receiver.GetAesKey()))))
	}
	if receiver.IdCard != "" && receiver.IdCardCrypt == "" {
		receiver.IdCardCrypt = cryptor.Base64StdEncode(string(cryptor.AesEcbEncrypt([]byte(receiver.IdCard), []byte(receiver.GetAesKey()))))
	}

	//密码加密
	if receiver.Password != "" && receiver.PasswordCrypt == "" {
		if receiver.Salt == "" {
			receiver.Salt = random.RandString(5)
		}
		receiver.PasswordCrypt = cryptor.Md5String(receiver.GetHashKey() + receiver.Salt + receiver.Password)
	}

	//修改密码|创建用户
	if receiver.Password != "" {
		receiver.Version = time.Now().Unix()
	}
	return
}

// ValidatePassword 验证密码
func (receiver *OdsUserInfoLogModel) ValidatePassword(password string) bool {
	return receiver.PasswordCrypt == cryptor.Md5String(receiver.GetHashKey()+receiver.Salt+password)
}

func (receiver *OdsUserInfoLogModel) GetEcbEncrypt(str string) string {
	return cryptor.Base64StdEncode(string(cryptor.AesEcbEncrypt([]byte(str), []byte(receiver.GetAesKey()))))
}

func (receiver *OdsUserInfoLogModel) findHook(tx *gorm.DB) (err error) {
	if receiver.PhoneCrypt != "" && receiver.Phone == "" {
		receiver.Phone = string(cryptor.AesEcbDecrypt([]byte(cryptor.Base64StdDecode(receiver.PhoneCrypt)), []byte(receiver.GetAesKey())))
	}
	if receiver.IdCardCrypt != "" && receiver.IdCard == "" {
		receiver.IdCard = string(cryptor.AesEcbDecrypt([]byte(cryptor.Base64StdDecode(receiver.IdCardCrypt)), []byte(receiver.GetAesKey())))
	}
	return
}

func (receiver *OdsUserInfoLogModel) AfterUpdate(tx *gorm.DB) (err error) {
	if receiver.UpdateHook != nil {
		err = receiver.UpdateHook(tx)
		return
	}
	return
}

func (receiver *OdsUserInfoLogModel) AfterSave(tx *gorm.DB) (err error) {
	return
}

func (receiver *OdsUserInfoLogModel) SaveUserOperationLog(ctx context.Context) (err error) {
	var remarks []string
	if receiver.UserName != "" {
		remarks = append(remarks, fmt.Sprintf("[用户名]变更为[%s]", receiver.UserName))
	}
	if receiver.Password != "" {
		remarks = append(remarks, "[密码]修改")
	}
	if receiver.Phone != "" {
		remarks = append(remarks, fmt.Sprintf("[手机号]变更为[%s]", receiver.Phone))
	}
	if receiver.TrueName != "" {
		remarks = append(remarks, fmt.Sprintf("[真实姓名]变更为[%s]", receiver.TrueName))
	}
	if receiver.IdCard != "" {
		remarks = append(remarks, fmt.Sprintf("[身份证]变更为[%s]", receiver.IdCard))
	}
	if receiver.Status != "" {
		remarks = append(remarks, fmt.Sprintf("[用户状态]变更为[%s]", receiver.Status))
	}
	if len(remarks) > 0 {
		userLogModel := &OdsUserOperationLogModel{}
		userLogModel.Db = func() *gorm.DB {
			return receiver.Db()
		}
		userLogModel.UserId = receiver.Id
		userLogModel.OperationTime = time.Now()
		userLogModel.Remark = strings.Join(remarks, ",")
		if createErr := userLogModel.Create(ctx); createErr != nil {
			err = createErr
			return
		}
	}
	return
}
