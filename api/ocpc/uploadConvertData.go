package ocpc

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/ocpc"
)

// UploadConvertData 广告主回传转化数据接口
// 广告主通过调用该接口，将匹配到的转化数据发送给百度服务器。
func UploadConvertData(ctx context.Context, clt *core.SDKClient, req *ocpc.UploadConvertDataRequest) (*model.ResponseHeader, error) {
	return clt.Conversion(ctx, req, nil)
}
