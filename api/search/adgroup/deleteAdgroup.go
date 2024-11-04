package adgroup

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/adgroup"
)

// DeleteAdgroup 删除单元
func DeleteAdgroup(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, adgroupIds ...uint64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: adgroup.DeleteAdgroupRequest{
			AdgroupIds: adgroupIds,
		},
	}

	return clt.Do(ctx, req, nil)
}
