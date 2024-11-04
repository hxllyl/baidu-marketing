package ocpc

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
)

// ActionCb APP转化数据收集
func ActionCb(ctx context.Context, clt *core.SDKClient, req model.ActionCbRequest) error {
	return clt.ActionCb(ctx, req)
}
