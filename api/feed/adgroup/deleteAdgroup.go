package adgroup

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/adgroup"
)

// DeleteAdgroup 删除单元
func DeleteAdgroup(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, adgroupFeedIds ...uint64) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body: &adgroup.DeleteAdgroupRequest{
			AdgroupFeedIds: adgroupFeedIds,
		},
	}

	var resp adgroup.DeleteAdgroupResponse
	headers, err := clt.Do(ctx, req, &resp)
	return headers, resp.Data, err
}
