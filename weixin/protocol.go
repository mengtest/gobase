package weixin

// getAccessTokenResponse 获取访问token 的请求信息
type getAccessTokenResponse struct {
	AccessToken string `json:"access_token"` //	获取到的凭证
	ExpiresIn   int64  `json:"expires_in"`   //	凭证有效时间，单位：秒。目前是7200秒之内的值。
}
