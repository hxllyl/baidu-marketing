package report

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/report"
)

// CreateReportTask 创建异步任务
func CreateReportTask(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, createRequest *report.CreateReportTaskRequest) (*model.ResponseHeader, string, error) {
	req := &model.Request{
		Header: auth,
		Body:   createRequest,
	}
	var resp report.CreateReportTaskResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data.TaskId, err
}
