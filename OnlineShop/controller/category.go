package controller

import (
	"OnlineShop/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryAllCategory(c *gin.Context) {
	// 查询所有分类，并返回
	//var categories []model.Category
	categories, err := model.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     200,
		"categories": categories,
	})
}

func QueryCategoryByID(c *gin.Context) {
	// 获取分类ID
	categoryIDStr := c.Param("id")
	categoryID, _ := strconv.Atoi(categoryIDStr)
	// 查询分类
	category, err := model.GetCategoryByID(uint(categoryID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"category": category,
	})
}
