package controller

import (
	"OnlineShop/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SubmitOrder(c *gin.Context) {
	// 创建订单
	var order model.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	order.BuyerID = c.GetUint("UserID")
	if err := model.SubmitOrder(&order); err != nil {
		log.Printf("[error]controller-SubmitOrder:%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	// 开始支付
	// TODO: 支付接口
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"orderID": order.OrderID,
	})

}
