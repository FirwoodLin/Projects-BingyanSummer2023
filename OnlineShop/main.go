package main

import (
	"OnlineShop/model"
	"OnlineShop/router"
)

func main() {
	//global.Log = core.Zap() // 初始化zap日志库
	//fmt.Printf("hello")
	model.InitDatabases()
	//model.Test()
	router.StartRouter()
}
