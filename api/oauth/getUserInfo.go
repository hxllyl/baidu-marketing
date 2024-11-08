package oauth

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model/oauth"
)

// GetUserInfo 查询授权用户信息
func GetUserInfo(ctx context.Context, clt *core.SDKClient, req *oauth.GetUserInfoRequest) (*oauth.UserInfo, error) {
	var resp oauth.UserInfo
	if err := clt.OAuth(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
