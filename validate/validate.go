package validate

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/validator"
	"strings"
)

// UserName 验证用户名格式
func UserName(str string) (err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("用户名为空")
		return
	}
	if !validator.ContainLetter(str) {
		err = errors.New("用户名必须包含字母")
		return
	}
	if !validator.IsAlphaNumeric(str) {
		err = errors.New("用户名只能是字母和数字")
		return
	}
	minLen := 6
	maxLen := 20
	strLen := len(str)
	if strLen < 6 || strLen > 20 {
		err = errors.New(fmt.Sprintf("用户名长度必须在%d-%d之间", minLen, maxLen))
		return
	}
	return
}

// Password 验证密码格式
func Password(str string) (err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("密码为空")
		return
	}
	minLen := 6
	maxLen := 20
	strLen := len(str)
	if strLen < 6 || strLen > 20 {
		err = errors.New(fmt.Sprintf("密码长度必须在%d-%d之间", minLen, maxLen))
		return
	}
	return
}

// ChineseMobile 验证中国手机号
func ChineseMobile(str string) (err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("手机号为空")
		return
	}
	if !validator.IsChineseMobile(str) {
		err = errors.New("手机号格式异常")
		return
	}
	return
}

// ChineseIdCard 验证中国身份证
func ChineseIdCard(str string) (err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("身份证为空")
		return
	}
	if !validator.IsChineseIdNum(str) {
		err = errors.New("身份证格式异常")
		return
	}
	return
}

// EmptyString 是否空字符串
func EmptyString(str string) (err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("字符串为空")
		return
	}
	return
}
