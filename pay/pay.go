package pay

import "context"

type PreOrderReq struct {
	OrderId     string
	Money       int
	CallbackUrl string
}

type PreOrderResp struct {
	MerchantOrderId string
	Url             string
}

type CallbackReq struct {
	OrderId         string
	MerchantOrderId string
	Sign            string
}

type CallbackResp struct {
}

type PayChannelInterface interface {
	PreOrder(ctx context.Context, req PreOrderReq) (resp PreOrderResp, err error)
	Callback(ctx context.Context, req CallbackReq) (resp CallbackResp, err error)
}
