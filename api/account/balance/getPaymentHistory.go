package balance

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/account/balance"
)

// GetPaymentHistory 查询待加款信息
// 支持KA账户查询付款记录，此接口仅供KA账户使用
func GetPaymentHistory(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *balance.GetPaymentHistoryRequest) (*model.ResponseHeader, []balance.PaymentHistory, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp balance.GetPaymentHistoryResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
