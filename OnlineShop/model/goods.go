package model

import (
	"OnlineShop/model/request"
)

type Goods struct {
	TimeModel
	// 设计以下字段：产品编号；商品名称，商品主图，商品缩略图，商品详情，商品价格，销售状态（在售/下架）；商品分类（外键），**商品库存（交易时需要注意）**
	GoodsID      uint   `gorm:"primarykey"`
	Name         string `gorm:"type:varchar(100)"` // 名称
	Detail       string `gorm:"type:varchar(100)"` // 详情(简介)
	PicMain      string `gorm:"type:varchar(100)"` // 主图
	PicThumbnail string `gorm:"type:varchar(100)"` // 缩略图
	Price        int    // 价格
	Stock        int    // 库存
	Sold         int    // 销售量
	/*进行外键关联*/
	CategoryID uint
	Category   Category
	UserID     uint // 持有人/Seller
	User       User
}
type UriBundle struct {
	OriginPicUri     string
	CompressedPicUri string
}

// QueryGoods 按照名称和分类查询商品
func QueryGoods(queryReq *request.GoodsQueryRequest) (goods []Goods, err error) {
	db := DBSql.Model(&Goods{})
	// 双重限制：名字+分类
	if queryReq.Name != "" {
		db.Where("name REGEXP ?", queryReq.Name)
	}
	if len(queryReq.CategoryID) > 0 {
		db.Where("category_id IN ?", queryReq.CategoryID)
	}
	// 使用 Where 指定查询条件后，要进行 Find 操作/First 操作，存储结果
	err = db.Find(&goods).Error
	return
}

// GetGoodsInfo 获取商品详情
func GetGoodsInfo(id uint) (goods Goods, err error) {
	err = DBSql.Find(&goods, id).Error
	return
}

// UpdateGoodsPic 更新商品图片 : 延期上线
func UpdateGoodsPic(id uint, uriBundle *UriBundle) (err error) {
	err = DBSql.Model(&Goods{}).Where("goods_id = ?", id).Updates(map[string]interface{}{"pic_main": uriBundle.OriginPicUri, "pic_thumbnail": uriBundle.CompressedPicUri}).Error
	return
}

// QueryGoodsInfo 根据商品ID查询商品信息
func QueryGoodsInfo(id uint) (goods Goods, err error) {
	err = DBSql.Where("goods_id = ?", id).First(&goods).Error
	return
}
