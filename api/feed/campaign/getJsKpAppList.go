package campaign

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/feed/campaign"
)

// GetJskAppList 查询计划
// 根据指定的计划ID获取推广计划(ID可批量)
func GetJsKpAppList(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *campaign.GetJsKpAppListRequest) (*model.ResponseHeader, []campaign.AppInfo, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.GetJsKpAppListResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
