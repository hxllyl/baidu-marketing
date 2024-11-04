package oauth

import (
	"context"

	"github.com/hxllyl/baidu-marketing/core"
	"github.com/hxllyl/baidu-marketing/model/oauth"
)

// AccessToken 换取授权令牌接口
func AccessToken(ctx context.Context, clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.AccessToken, error) {
	var resp oauth.AccessToken
	req.AppID = clt.AppID()
	req.SecretKey = clt.Secret()
	if err := clt.OAuth(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
