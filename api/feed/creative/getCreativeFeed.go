package creative

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/creative"
)

// GetCreativeFeed 查询创意
func GetCreativeFeed(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *creative.GetCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.GetCreativeResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
