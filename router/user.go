package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*定义user路由参数*/
func InitUserRouter() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Sucess",
			})
		})
		// user路由分组
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.POST("", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "zs"},
					{"id": 2, "name": "ls"},
				},
			})
		})
		rgAuthUser.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "zs",
			})
		})
	})

}
