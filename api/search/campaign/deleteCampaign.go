package campaign

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/campaign"
)

// DeleteCampaign 删除计划
func DeleteCampaign(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, campaignIds ...uint64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: campaign.DeleteCampaignRequest{
			CampaignIds: campaignIds,
		},
	}

	return clt.Do(ctx, req, nil)
}
