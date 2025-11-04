package pay

import "context"

type WeiXinPay struct {
}

func (receiver *WeiXinPay) PreOrder(ctx context.Context, req PreOrderReq) (resp PreOrderResp, err error) {
	return
}
