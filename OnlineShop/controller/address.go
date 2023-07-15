package controller

import (
	"OnlineShop/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	var add model.Address
	if err := c.BindJSON(&add); err != nil {
		log.Printf("[error]controller-CreateAddress:解析错误%v\n", add)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	userID := c.GetUint("UserID")
	add.UserID = userID
	if err := model.CreateAddress(&add); err != nil {
		log.Printf("[error]controller-CreateAddress:创建地址失败%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "创建地址失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":       "创建地址成功",
		"addressID": add.AddressID,
	})
	return
}
