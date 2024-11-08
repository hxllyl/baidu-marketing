package campaign

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/campaign"
)

// GetBindBusinessPointList 查询计划已使用的推广业务
func GetBindBusinessPointList(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *campaign.GetBindBusinessPointListRequest) (*model.ResponseHeader, []campaign.BusinessPoint, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.GetBindBusinessPointListResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data.Paths, err
}
