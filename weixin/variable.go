package weixin

import (
	"encoding/json"
	"fmt"
	"gobase/util"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// Token 用于描述微信token的信息
type Token struct {
	AppID          string
	AppSecrect     string
	AccessToken    string
	NextUpdateTime int64
	Lock           sync.Mutex
}

var tokenInstance *Token

const (
	tokenUpdateInterval = 6 // s单位
)

// InitToken 用于初始化Token
func InitToken(appid, appSecrect string) {
	tokenInstance = &Token{
		AppID:       appid,
		AppSecrect:  appSecrect,
		AccessToken: "",
	}

}

// GetToken 用于更新微信小程序的Token
func getToken() string {
	tokenInstance.Lock.Lock()
	defer tokenInstance.Lock.Unlock()
	currentTime := util.GetCurrentTimestamp()
	if currentTime > tokenInstance.NextUpdateTime || tokenInstance.AccessToken == "" {
		tokenInstance.NextUpdateTime = currentTime + tokenUpdateInterval
		getWeixinAccessToken()
	}
	return tokenInstance.AccessToken
}

// getWeixinAccessToken 用于获取微信的访问TOKEN https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func getWeixinAccessToken() {
	resp, err := http.Get(fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		tokenInstance.AppID,
		tokenInstance.AppSecrect,
	))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	dataResp := &getAccessTokenResponse{}
	json.Unmarshal(data, dataResp)
	tokenInstance.AccessToken = dataResp.AccessToken
}

type textResp struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// CheckTextIsLegitimate 用于检测文本是不是合格的
func CheckTextIsLegitimate(content string) bool {
	resp, err := http.Post(
		fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", getToken()),
		"application/json",
		strings.NewReader(fmt.Sprintf("{\"content\":\"%s\"}", content)),
	)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	tR := &textResp{}
	err = json.Unmarshal(data, tR)
	if err != nil {
		return false
	}
	if tR.ErrCode != 0 {
		return false
	}
	return true
}
