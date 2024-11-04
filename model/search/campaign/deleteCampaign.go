package campaign

import (
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/util"
)

// DeleteCampaignRequest 删除计划 API Request
type DeleteCampaignRequest struct {
	// CampaignIds 计划ID
	CampaignIds []uint64 `json:"campaignIds,omitempty"`
}

func (r DeleteCampaignRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "CampaignService/deleteCampaign")
}
