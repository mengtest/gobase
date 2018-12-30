package weixin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type wxCode2Session struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
}

// GetWeiXinOpenID 用于获取WxOpenID
func GetWeiXinOpenID(appid, appSecret, jSCode string) (string, error) {
	var err error
	var resp *http.Response
	var data []byte
	openID := ""
	wxRespObj := &wxCode2Session{}
	if resp, err = http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid,
		appSecret,
		jSCode,
	)); err != nil {
		goto end
	}
	defer resp.Body.Close()
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		goto end
	}
	fmt.Println(string(data))
	if err = json.Unmarshal(data, wxRespObj); err != nil {
		goto end
	}
	openID = wxRespObj.OpenID
end:
	return openID, err
}

// GetWeiXinUserInfo 用于获取用户数据
func GetWeiXinUserInfo() {

}
