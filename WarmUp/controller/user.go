package controller

import (
	"WarmUp/model"
	"WarmUp/util"
	"log"
	"runtime"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var userReq model.UserRequest
	// 解析请求体；使用 BindJson 简化流程
	if err := c.BindJSON(&userReq); err != nil {
		// err.Error 返回的是 字符串形式的错误信息
		_, file, line, _ := runtime.Caller(0)
		log.Printf("[srror]file: %s, line: %d\n%v\n", file, line, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 进行注册
	if err := model.Register(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 注册成功；返回 SessionId
	log.Printf("[info]准备生成 session 并返回 user:%v\n", userReq)
	sessionID, err := util.GenSessionId(&model.UserSignInResponse{ID: userReq.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("SESSIONID", sessionID, 60*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"sessionID": sessionID, "userID": userReq.ID})
}

// SignIn 用户登录
func SignIn(c *gin.Context) {
	var userSignInRequest model.UserSignInRequest
	if err := c.BindJSON(&userSignInRequest); err != nil {
		log.Printf("[error]SignIn:解析请求体错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 检查用户邮箱和密码是否匹配
	userResponse, err := model.CheckUser(&userSignInRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 颁发 sessionID
	sessionID, err := util.GenSessionId(userResponse)
	if err != nil {
		log.Printf("[error]SignIn:生成 SessionID 错误:%v\n", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("SESSIONID", sessionID, 60*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"sessionID": sessionID})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	var userReq model.UserUpdateRequest
	if err := c.BindJSON(&userReq); err != nil {
		log.Printf("[error]controller:更新用户信息-BindJson出错%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userReq.ID = c.GetUint("UserID")
	log.Printf("[info]controller-UpdateUser,userReq:%v\n", userReq)
	model.UpdateUser(&userReq)
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}

// DeleteUser 删除用户 - 管理员权限
func DeleteUser(c *gin.Context) {
	if !c.GetBool("IsAdmin") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "没有删除权限"})
		return
	}
	deteleIdStr := c.Param("id")
	deleteId, _ := strconv.ParseUint(deteleIdStr, 10, 64)
	if err := model.DeleteUser(uint(deleteId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
}

// QueryOneUser 查询个人信息
func QueryOneUser(c *gin.Context) {
	queryIdStr := c.Param("id")
	queryId, _ := strconv.ParseUint(queryIdStr, 10, 64)
	if queryId != c.GetUint64("UserID") && !c.GetBool("IsAdmin") {
		// 不是本人，也不是管理员
		c.JSON(http.StatusForbidden, gin.H{"error": "无查询权限"})
		return
	}
	userResponse, err := model.QueryOneUser(uint(queryId))

	if err != nil {
		// 数据库查询出错
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userResponse)
}

// QueryUsers 查询所有人的信息 - 管理员权限
func QueryUsers(c *gin.Context) {
	if !c.GetBool("IsAdmin") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "没有查询权限"})
		return
	}
	users, err := model.QueryUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器查询出错"})
	}
	c.JSON(http.StatusOK, users)

}
