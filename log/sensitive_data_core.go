package log

import (
	"encoding/base64"
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"regexp"
	"strings"
	"sync"
)

// SensitiveDataCore 脱敏核心
type SensitiveDataCore struct {
	zapcore.Core
	policy *MaskingPolicy
}

// MaskingPolicy 脱敏策略
type MaskingPolicy struct {
	patterns map[string]*regexp.Regexp
	mu       sync.RWMutex
}

// NewMaskingPolicy 创建脱敏策略
func NewMaskingPolicy() *MaskingPolicy {
	return &MaskingPolicy{
		patterns: map[string]*regexp.Regexp{
			"phone":     regexp.MustCompile(`(1[3-9]\d)\d{4}(\d{4})`),
			"id_card":   regexp.MustCompile(`(\d{4})\d{10}(\w{4})`),
			"email":     regexp.MustCompile(`(\w{2})[\w.-]*@(\w+\.\w+)`),
			"bank_card": regexp.MustCompile(`(\d{4})\d{8,10}(\d{4})`),
			"password":  regexp.MustCompile(`.+`), // 匹配任何密码内容
			"token":     regexp.MustCompile(`.+`), // 匹配任何TOKEN内容
		},
	}
}

func ReplaceStrSensitiveData(str string, aesKey []byte) string {
	maskingPolicy := NewMaskingPolicy()
	for key, regexpValue := range maskingPolicy.patterns {
		newRegxp := fmt.Sprintf("\"%s\":\"%s\"", key, regexpValue.String())
		findStrs := regexp.MustCompile(newRegxp).FindAllString(str, -1)
		if len(findStrs) > 0 {
			for _, item := range findStrs {
				str = regexp.MustCompile(newRegxp).ReplaceAllString(str, fmt.Sprintf("\"%s\":\"%s\"", key, base64.RawStdEncoding.EncodeToString(cryptor.AesEcbEncrypt([]byte(item), aesKey))))
			}
		}
	}
	return str
}

// AddPattern 添加自定义脱敏模式
func (m *MaskingPolicy) AddPattern(name, pattern string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	m.patterns[name] = re
	return nil
}

// NewSensitiveDataCore 创建脱敏核心
func NewSensitiveDataCore(core zapcore.Core) *SensitiveDataCore {
	return &SensitiveDataCore{
		Core:   core,
		policy: NewMaskingPolicy(),
	}
}

// With 实现 With 方法
func (c *SensitiveDataCore) With(fields []zapcore.Field) zapcore.Core {
	return &SensitiveDataCore{
		Core:   c.Core.With(c.maskSensitiveData(fields)),
		policy: c.policy,
	}
}

// Check 实现 Check 方法
func (c *SensitiveDataCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return c.Core.Check(ent, ce)
}

// Write 实现 Write 方法，在写入前进行脱敏
func (c *SensitiveDataCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	// 对字段进行脱敏处理
	maskedFields := c.maskSensitiveData(fields)
	return c.Core.Write(ent, maskedFields)
}

// Sync 实现 Sync 方法
func (c *SensitiveDataCore) Sync() error {
	return c.Core.Sync()
}

// maskSensitiveData 脱敏敏感数据
func (c *SensitiveDataCore) maskSensitiveData(fields []zapcore.Field) []zapcore.Field {
	if len(fields) == 0 {
		return fields
	}

	maskedFields := make([]zapcore.Field, len(fields))

	for i, field := range fields {
		maskedFields[i] = c.maskField(field)
	}
	return maskedFields
}

// maskField 单个字段脱敏
func (c *SensitiveDataCore) maskField(field zapcore.Field) zapcore.Field {
	fieldName := strings.ToLower(field.Key)

	// 判断字段是否需要脱敏
	if !c.isSensitiveField(fieldName) {
		return field
	}

	// 根据字段类型进行脱敏
	switch field.Type {
	case zapcore.StringType:
		return zap.String(field.Key, c.maskStringValue(fieldName, field.String))
	case zapcore.ByteStringType:
		return zap.ByteString(field.Key, []byte(c.maskStringValue(fieldName, string(field.Interface.([]byte)))))
	default:
		// 对于其他类型，如果字段名敏感，则进行掩码
		return zap.String(field.Key, "***")
	}
}

// isSensitiveField 判断字段是否包含敏感信息
func (c *SensitiveDataCore) isSensitiveField(fieldName string) bool {
	sensitiveKeywords := []string{
		"password", "pwd", "passwd", "pass",
		"id_card", "idcard", "identity", "身份证",
		"phone", "mobile", "tel", "电话", "手机",
		"email", "mail", "邮箱",
		"bank_card", "card_no", "account", "银行卡", "账号",
		"token", "secret", "key", "api_key", "apikey",
		"credit", "credential", "auth",
		"address", "住址",
		"name", "姓名", "realname",
	}

	for _, keyword := range sensitiveKeywords {
		if strings.Contains(fieldName, keyword) {
			return true
		}
	}
	return false
}

// maskStringValue 字符串值脱敏
func (c *SensitiveDataCore) maskStringValue(fieldName, value string) string {
	if value == "" {
		return value
	}

	c.policy.mu.RLock()
	defer c.policy.mu.RUnlock()

	// 根据字段名应用不同的脱敏规则
	switch {
	case strings.Contains(fieldName, "phone") || strings.Contains(fieldName, "mobile") || strings.Contains(fieldName, "tel"):
		if re, ok := c.policy.patterns["phone"]; ok {
			return re.ReplaceAllString(value, "${1}****${2}")
		}

	case strings.Contains(fieldName, "id_card") || strings.Contains(fieldName, "identity"):
		if re, ok := c.policy.patterns["id_card"]; ok {
			return re.ReplaceAllString(value, "${1}**********${2}")
		}

	case strings.Contains(fieldName, "email") || strings.Contains(fieldName, "mail"):
		if re, ok := c.policy.patterns["email"]; ok {
			return re.ReplaceAllString(value, "${1}***@${2}")
		}

	case strings.Contains(fieldName, "bank_card") || strings.Contains(fieldName, "card") || strings.Contains(fieldName, "account"):
		if re, ok := c.policy.patterns["bank_card"]; ok {
			return re.ReplaceAllString(value, "${1}********${2}")
		}

	case strings.Contains(fieldName, "password") || strings.Contains(fieldName, "pwd") ||
		strings.Contains(fieldName, "token") || strings.Contains(fieldName, "secret"):
		return "***"

	default:
		// 对于其他敏感字段，使用通用脱敏规则
		if len(value) <= 2 {
			return "**"
		} else if len(value) <= 6 {
			return value[:1] + "***" + value[len(value)-1:]
		} else {
			return value[:2] + "***" + value[len(value)-2:]
		}
	}

	return value
}
