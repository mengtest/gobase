package util

import (
	"github.com/gin-gonic/gin"
)

type comResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func newComResponse(code int, data interface{}) *comResponse {
	return &comResponse{
		Code: code,
		Data: data,
	}
}

// Dispatch 用于下发数据到客户端
func Dispatch(actionCode int, jsonObj interface{}, c *gin.Context) {
	c.JSON(200, newComResponse(actionCode, jsonObj))
}
