package main

import (
	"WarmUp/config"
	"WarmUp/router"
	"WarmUp/util"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ReadIn()
	// 连接数据库
	util.ConnectDB()
	// 开始服务
	r := gin.Default()
	router.NewRouter(r)

}
func InitSettings() {

}
