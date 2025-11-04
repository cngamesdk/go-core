package pay

import "context"

type Alipay struct {
}

func (receiver Alipay) PreOrder(ctx context.Context, req PreOrderReq) (resp PreOrderResp, err error) {
	return
}
