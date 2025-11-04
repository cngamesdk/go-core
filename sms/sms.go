package sms

import "context"

type SmsService struct {
	Phone   string
	content string
	result  string
}

func NewSmsService(phone string) *SmsService {
	return &SmsService{Phone: phone}
}

func (receiver SmsService) GetContent() string {
	return receiver.content
}

func (receiver SmsService) GetResult() string {
	return receiver.result
}

// SendCheckCode 发送验证码
func (receiver SmsService) SendCheckCode(ctx context.Context, req string) (err error) {
	receiver.content = "" // 短信发送内容

	receiver.result = "" // 发送结果
	return
}
