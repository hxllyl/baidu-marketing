package trans

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/adgroup/trans"
)

// UpdateOcpcTransFeed 修改转化追踪
func UpdateOcpcTransFeed(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, list []trans.OcpcTransFeed) (*model.ResponseHeader, []trans.OcpcTransFeed, error) {
	req := &model.Request{
		Header: auth,
		Body:   &trans.UpdateOcpcTransFeedRequest{TransTraceUpdateTypes: list},
	}
	var resp trans.UpdateOcpcTransFeedResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
