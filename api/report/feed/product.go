package feed

import (
	"context"

	reportApi "github.com/hxllyl/baidu-marketing/api/report"
	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/report"
)

// Product 信息流商品报告
func Product(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 2094829
	return reportApi.GetReportData(ctx, clt, auth, reportRequest)
}
