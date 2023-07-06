package util

import (
	"WarmUp/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenSessionId(userReq *model.UserRequset) (sessionID string, err error) {
	// user := model.ConvertUserRequestToUser(userReq)
	sessionID = uuid.New().String() // 使用 uuid 作为 sessinID
	if err := model.InsertSession(userReq, sessionID); err != nil {
		return "", err
	}
	return sessionID, nil
}

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
		userId, err := model.CheckSession(cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// sessionID 存在，继续执行

		c.Set("UserId", userId)
		c.Next()
	}
}
