package ocpc

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/ocpc"
)

// UploadInvalidConvertData 广告主回传无效转化数据接口
// 广告主通过调用该接口，将认为是无效的转化数据发送给百度服务器。
func UploadInvalidConvertData(ctx context.Context, clt *core.SDKClient, req *ocpc.UploadInvalidConvertDataRequest) (*model.ResponseHeader, error) {
	return clt.Conversion(ctx, req, nil)
}
