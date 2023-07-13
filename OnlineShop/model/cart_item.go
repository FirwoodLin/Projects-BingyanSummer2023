package model

type CartItem struct {
	// 不紧急

	// 只需要存储购物车中的每一项，不需要存储购物车本身 使用时，通过用户ID来获取购物车中的所有项
	TimeModel
	CartID   uint `gorm:"primarykey"`
	GoodsNum int  // 商品数量
	/*进行外键关联*/
	GoodsID uint
	Goods   Goods
	UserID  uint
	User    User
}
