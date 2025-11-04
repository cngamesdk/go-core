package error

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrorInternalSystem  = errors.New("系统异常")
	ErrorRecordIsNotFind = gorm.ErrRecordNotFound
	ErrorRecordIsExists  = errors.New("记录已经存在")
	ErrorDecrypt         = errors.New("解密失败")
	ErrorEncrypt         = errors.New("加密失败")
	ErrorParamEmpty      = errors.New("参数为空")
	ErrorSignVerifyFail  = errors.New("加密校验失败")
)
