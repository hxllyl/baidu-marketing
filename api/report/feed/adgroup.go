package feed

import (
	"context"

	reportApi "github.com/hxllyl/baidu-marketing/api/report"
	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/report"
)

// Adgroup 单元报告
func Adgroup(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, *report.ReportData, error) {
	reportRequest.ReportType = 2330652
	return reportApi.GetReportData(ctx, clt, auth, reportRequest)
}
