package feed

import (
	"context"

	reportApi "github.com/hxllyl/baidu-marketing/api/report"
	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/report"
)

// Creative 创意报告
func Creative(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 2094816
	return reportApi.GetReportData(ctx, clt, auth, reportRequest)
}
