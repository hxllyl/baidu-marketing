package creative

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/creative"
)

// AddCreative 新增推广创意
func AddCreative(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *creative.AddCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.AddCreativeResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
