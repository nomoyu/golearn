package cmd

import (
	"fmt"
	"github.com/nomoyu/golearn/conf"
	"github.com/nomoyu/golearn/global"
	"github.com/nomoyu/golearn/router"
	"github.com/nomoyu/golearn/utils"
)

func Start() {
	var initErr error
	// 初始化系统配置
	conf.InitConfig()
	// 初始化日志组件
	global.Log = conf.InitLog()
	// 初始化数据库
	db, err := conf.InitDb()
	if err != nil {
		initErr = utils.AddError(initErr, err)
	}
	global.DB = db
	if initErr != nil {
		global.Log.Error(initErr.Error())
		panic(initErr.Error())
	}
	// 初始化路由
	router.InitRoute()

}

func Clean() {
	fmt.Println("=====clean=====")
}
