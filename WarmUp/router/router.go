package router

import (
	"WarmUp/controller"
	"WarmUp/util"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	// 注册登陆
	r.POST("/register", controller.Register)
	r.POST("/signin", controller.SignIn)
	// 修改个人信息
	r.PATCH("/user", util.SessionMiddleware(), controller.UpdateUser)
	r.GET("/user/:id", util.SessionMiddleware(), controller.QueryOneUser)
	// ===管理员===
	r.DELETE("/user/:id", util.SessionMiddleware(), controller.DeleteUser)
	r.GET("/user", util.SessionMiddleware(), controller.QueryUsers)
	_ = r.Run(":8080")
}
