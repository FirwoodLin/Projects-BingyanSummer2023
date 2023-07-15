package main

import (
	"OnlineShop/initialize"
	"OnlineShop/model"
	"OnlineShop/router"
)

func main() {
	//global.Log = core.Zap() // 初始化zap日志库
	//fmt.Printf("hello")
	model.InitDatabases()
	initialize.ConnectCOS()
	//model.Test()
	router.StartRouter()
}
