package controller

import (
	"OnlineShop/model"
	"OnlineShop/model/request"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

func SubmitOrder(c *gin.Context) {
	var orderReq request.SubmitOrderRequest
	var order model.Order
	err := c.BindJSON(&orderReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	_ = copier.Copy(&order, &orderReq)
	orderReq.UserID = uint(c.GetInt("userID"))
	//model.
}
