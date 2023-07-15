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
		// 商品-更新图片
		rWithSessionMiddleware.POST("/goods/:id/pics", controller.UpdateGoodsPic)
		// 订单
		rWithSessionMiddleware.POST("/orders", controller.SubmitOrder) // 创建订单
		// 地址信息
		rWithSessionMiddleware.POST("/addresses", controller.CreateAddress) // 创建地址
		// 用户 个人信息
		rWithSessionMiddleware.GET("/users/info/:id", controller.QueryOneUser)
	}
	rWithoutSessionMiddleware := r.Group("/v1")
	{
		// 不需要 session 的接口
		rWithoutSessionMiddleware.POST("/users/register", controller.Register)
		rWithoutSessionMiddleware.POST("/users/login", controller.SignIn)
		// 订单
		//rWithoutSessionMiddleware.POST("/orders/notify", controller.Notify) // 接受支付状态信息；支付成功后，支付宝会调用这个接口
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
