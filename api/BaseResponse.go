package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"status,omitempty"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func (m ResponseJson) isEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func httpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.isEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

func buildStates(resp ResponseJson, defaultStatus int) int {
	if 0 == resp.Status {
		return defaultStatus
	}
	return resp.Status
}

// 请求成功返回结果
func Success(context *gin.Context, resp ResponseJson) {
	httpResponse(context, buildStates(resp, http.StatusOK), resp)
}

// 请求失败返回结果
func Fail(context *gin.Context, resp ResponseJson) {
	httpResponse(context, buildStates(resp, http.StatusBadRequest), resp)
}
