package keyword

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/keyword"
)

// AddWord 添加关键词
// 新增关键词及商品目录模板词(绑定在商品目录营销目标计划单元下的关键词)（可批量)，一次请求最多新建10000个关键词
func AddWord(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *keyword.AddWordRequest) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp keyword.AddWordResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
