package request

type GoodsQueryRequest struct {
	Name       string `json:"name" form:"name"`
	CategoryID []uint `json:"category_id" form:"category_id"`
}
