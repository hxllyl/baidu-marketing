package campaign

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/campaign"
)

// UpdateCampaign 更新计划
// 根据指定的计划ID更新推广计划的属性
func UpdateCampaign(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, campaigns []campaign.Campaign) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body: campaign.UpdateCampaignRequest{
			CampaignTypes: campaigns,
		},
	}
	var resp campaign.UpdateCampaignResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
