package account

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/account"
)

// UpdateAccountFeed 更新账户信息
func UpdateAccountFeed(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, budget float64) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: &account.UpdateAccountFeedRequest{
			AccountFeedType: &account.Account{
				Budget: budget,
			},
		},
	}
	var resp account.UpdateAccountFeedResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
