package api

import "github.com/gin-gonic/gin"

type UserApi struct {
	ID   int
	Name string
}

func NewUserApi() UserApi {
	return UserApi{}
}

func (m UserApi) Login(ctx *gin.Context) {
	// TODO 处理业务逻辑
	Success(ctx, ResponseJson{
		Status: 200,
		Code:   200,
		Msg:    "你好啊",
		Data:   nil,
	})

}
