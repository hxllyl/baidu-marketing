package app

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/app"
)

// AddAdgroupAppBind 添加APP绑定
func AddAdgroupAppBind(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *app.AddAdgroupAppBindRequest) (*model.ResponseHeader, []app.AppBindItem, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.AddAdgroupAppBindResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
