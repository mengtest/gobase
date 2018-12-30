package main

import (
	"fmt"
	"gobase/weixin"
)

func main() {

	appID := "wxf63c1a4f5afe7f7d"
	appSecrect := "6d6c161cc37988df6511fd8a92d53bc7"
	weixin.InitToken(appID, appSecrect)
	fmt.Println(weixin.CheckTextIsLegitimate("123"))
}
