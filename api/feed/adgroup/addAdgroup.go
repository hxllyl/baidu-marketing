package adgroup

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/adgroup"
)

// AddAdgroup 添加单元
// 新增推广单元
func AddAdgroup(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *adgroup.AddAdgroupRequest) (*model.ResponseHeader, []adgroup.Adgroup, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp adgroup.AddAdgroupResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
