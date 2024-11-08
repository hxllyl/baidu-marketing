package keyword

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/keyword"
)

// GetWord 查询关键词
func GetWord(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *keyword.GetWordRequest) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp keyword.GetWordResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
