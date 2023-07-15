package request

type SubmitOrderRequest struct {
	// 订单中未必是同一个卖家的商品；卖家信息根据商品信息获取
	Items      []SubmitOrderRequestItem
	TotalPrice int
	//UserID     uint `json:"-"` // BuyerID;买家 ID
	BuyerID uint `json:"-"` // BuyerID;买家 ID
}
type SubmitOrderRequestItem struct {
	GoodsPrice int
	GoodsID    uint
	GoodsNum   int
}
