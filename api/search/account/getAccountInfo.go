package account

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/account"
)

// GetAccountInfo 查询账户
// 获取username对应的账户信息；如果是来自MCC Token的请求，则返回target对应的账户信息。
func GetAccountInfo(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, accountFields []string) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: &account.GetAccountInfoRequest{
			AccountFields: accountFields,
		},
	}
	var resp account.GetAccountInfoResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
