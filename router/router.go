package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nomoyu/golearn/global"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// IFnRgRoute /*接收路由预定义函数*/
type IFnRgRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRgRoute
)

// RegisterRoute /*注册路由到切片中暂存*/
func RegisterRoute(fn IFnRgRoute) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)

}

// InitRoute /*初始化路由*/
func InitRoute() {
	// 优雅停机
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()
	// 公共路由，默认公共可访问
	rgPublic := r.Group("/api/v1/public")
	// 鉴权路由，权限可访问
	rgAuth := r.Group("/api/v1")
	// 初始化路由
	InitBasePlatFormRouters()

	for _, router := range gfnRouters {
		// 分配路由
		router(rgPublic, rgAuth)
	}
	// 获取端口
	port := viper.GetString("server.port")
	if port == "" {
		port = "3317"
	}
	// 默认启动方式
	//err := r.Run(fmt.Sprintf(":%s", port))
	//if err != nil {
	//	panic("Start Server Error: " + err.Error())
	//
	//}
	// 原生挂载
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}
	// 异步监听
	go func() {
		global.Log.Info(fmt.Sprintf("Start Listen: %s", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
	}()
	// 等待手动停止或者终止
	<-ctx.Done()
	//cancel()

	ctx, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()
	// 停止后的处理事件
	if err := server.Shutdown(ctx); err != nil {
		//TODO 记录日志
		global.Log.Error("Stop Server Error: %s", err.Error())
		return
	}
	global.Log.Info("Stop Server Success")

}

// InitBasePlatFormRouters /*加载各模块路由配置*/
func InitBasePlatFormRouters() {
	// user 路由配置
	InitUserRouter()
}
