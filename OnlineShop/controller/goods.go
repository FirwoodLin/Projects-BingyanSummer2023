package controller

import (
	"OnlineShop/model"
	"OnlineShop/model/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// QueryGoods 按照名称 & 分类 ID 查询
func QueryGoods(c *gin.Context) {
	// 解析请求中的参数
	var queryGoodsRequest request.GoodsQueryRequest
	err := c.BindQuery(&queryGoodsRequest)
	if err != nil {
		log.Printf("[error]controller:QueryGoods-解析请求参数错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	log.Printf("[info]controller:QueryGoods-解析请求参数成功:%v\n", queryGoodsRequest)

	// 查询商品
	goods, err := model.QueryGoods(&queryGoodsRequest)
	// 返回结果
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"goods":  goods,
	})
}

// GetGoodsInfo 获得商品详情页
func GetGoodsInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("[error]controller:GetFoodsInfo-解析参数错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	goods, err := model.GetGoodsInfo(uint(id))
	if err != nil {
		log.Printf("[error]controller:GetFoodsInfo-查询商品错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"goods":  goods,
	})
}

/*
// UpdateGoodsPic 更新商品图片
func UpdateGoodsPic(c *gin.Context) {
	// 获取商品 ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("[error]controller:UpdateGoodsPic-解析参数错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	// 获取图片
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("[error]controller:UpdateGoodsPic-获取图片错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	// 上传图片;返回原图和缩略图的地址 bundle
	uri,err := util.UploadFile(file)
	if err != nil {
		log.Printf("[error]controller:UpdateGoodsPic-保存图片错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	// 更新数据库中的图片地址
	err = model.UpdateGoodsPic(uint(id), uri)
	if err != nil {
		log.Printf("[error]controller:UpdateGoodsPic-更新数据库错误:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
}*/
