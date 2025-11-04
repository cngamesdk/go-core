package idcard

import "context"

type IdCardService struct {
	IdCard   string
	TrueName string
}

func NewIdCardService(idCard string, trueName string) IdCardService {
	return IdCardService{IdCard: idCard, TrueName: trueName}
}

func (receiver IdCardService) Check(ctx context.Context) (err error) {
	return
}
