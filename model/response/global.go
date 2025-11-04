package response

type GlobalResp struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

func NewGlobalResp() GlobalResp {
	return GlobalResp{}
}

func (g GlobalResp) SetCode(code int) GlobalResp {
	g.Code = code
	return g
}

func (g GlobalResp) SetMsg(msg string) GlobalResp {
	g.Msg = msg
	return g
}

func (g GlobalResp) SetData(data interface{}) GlobalResp {
	g.Data = data
	return g
}

func (g GlobalResp) SetRequestId(requestId string) GlobalResp {
	g.RequestId = requestId
	return g
}
