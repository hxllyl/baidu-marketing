package keyword

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model"
	"github.com/hxllyl/baidu-marketing/model/search/keyword"
)

// UpdateWord 更新关键词
func UpdateWord(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, keywords []keyword.Keyword) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body: &keyword.UpdateWordRequest{
			KeywordTypes: keywords,
		},
	}
	var resp keyword.UpdateWordResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
