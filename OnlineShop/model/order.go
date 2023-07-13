package model

import (
	"database/sql"
	"errors"
)

// 定义枚举量，用于表示订单状态

const (
	OrderStatusUnpaid = iota
	OrderStatusPaid
	OrderStatusDelivered
	OrderStatusReceived
	OrderStatusClosed
)

type Order struct {
	TimeModel
	OrderID    uint `gorm:"primarykey"`
	TotalPrice int
	PaySerial  string
	Status     int    `gorm:"type:tinyint"`
	Remark     string // 备注
	FinishTime sql.NullTime
	OrderItems []OrderItem
	// 进行外键关联
	BuyerID   uint
	Buyer     User
	SellerID  uint
	Seller    User
	AddressID uint
	Address   Address
}
type OrderItem struct {
	TimeModel
	OrderItemID uint `gorm:"primarykey"`
	// 要填写的字段
	GoodsNum   int
	GoodsPrice int
	// 进行外键关联
	OrderID uint
	Order   Order
	GoodsID uint
	Goods   Goods
	// define these fields: 商品图片，商品单价
}

// SubmitOrder 将订单整体提交到数据库当中
func SubmitOrder(order *Order) (err error) {
	for _, v := range order.OrderItems {
		if err := submitOrderItem(&v); err != nil {
			return err
		}
	}
	return nil
}
func submitOrderItem(orderItem *OrderItem) (err error) {
	// 开启数据库事务
	tx := DBSql.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 查询商品以及库存
	var goods Goods
	if err := tx.Model(&Goods{}).
		Where("goods_id = ?", orderItem.GoodsID).
		Set("gorm:query_option", "FOR UPDATE").
		First(&goods).Error; err != nil {
		// 找不到商品，回滚事务
		tx.Rollback()
		return err
	}
	// 减少库存
	if goods.Stock < orderItem.GoodsNum {
		tx.Rollback()
		return errors.New("库存不足")
	}
	goods.Stock -= orderItem.GoodsNum
	goods.Sold += orderItem.GoodsNum
	if err := tx.Save(&goods).Error; err != nil {
		tx.Rollback()
		return err
	}
	return
}
