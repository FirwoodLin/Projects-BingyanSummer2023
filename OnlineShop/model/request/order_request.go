package request

type SubmitOrderRequest struct {
	//GoodsID uint
	//Num     uint
	Items      []SubmitOrderRequestItem
	TotalPrice int

	UserID uint `json:"-"`
}
type SubmitOrderRequestItem struct {
	GoodsPrice int
	GoodsID    uint

	GoodsNum int
}
