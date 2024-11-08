package creative

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/creative"
)

// DeleteCreative 删除基础创意
func DeleteCreative(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, creativeIds ...uint64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: &creative.DeleteCreativeRequest{
			CreativeIds: creativeIds,
		},
	}

	return clt.Do(ctx, req, nil)
}
