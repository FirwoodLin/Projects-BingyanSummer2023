package router

import (
	"OnlineShop/controller"
	"OnlineShop/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func newRouter() *gin.Engine {
	r := gin.Default()
	rWithSessionMiddleware := r.Group("/v1").Use(middleware.SessionMiddleware())
	{
		// 用户登录后才能访问的接口
		// 分类
		rWithSessionMiddleware.GET("/categories", controller.QueryAllCategory)      // 查询所有分类
		rWithSessionMiddleware.GET("/categories/:id", controller.QueryCategoryByID) // 按照 ID 查询某个分类
		// 商品
		rWithSessionMiddleware.GET("/goods", controller.QueryGoods)       //按照名称/分类查询；
		rWithSessionMiddleware.GET("/goods/:id", controller.GetGoodsInfo) //按照 ID 查询
	}
	rWithoutSessionMiddleware := r.Group("/v1")
	{
		// 不需要 session 的接口
		rWithoutSessionMiddleware.POST("/users/register", controller.Register)
		rWithoutSessionMiddleware.POST("/users/login", controller.SignIn)
	}
	return r
}

func StartRouter() {
	r := newRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Printf("[error]router-StartRouter:%v\n", err)
	}
}
