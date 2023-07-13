package middleware

import (
	"OnlineShop/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 cookie
		cookie, err := c.Cookie("SESSIONID")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// 检查 sessionID 是否存在
		userResponse, err := model.CheckSession(cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// sessionID 存在，继续执行
		c.Set("UserID", userResponse.ID)
		c.Set("IsAdmin", userResponse.IsAdmin)
		log.Printf("[info]util-SessionMiddleware,userResponse:%v\n", userResponse)
		//global.Log.Info("util-SessionMiddleware,userResponse:%v\n", userResponse)
		c.Next()
	}
}
