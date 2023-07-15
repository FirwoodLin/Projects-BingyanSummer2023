package model

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"log"
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
	Status     int         `gorm:"type:tinyint"`
	OrderID    uint        `gorm:"primarykey"`
	OrderItems []OrderItem // 订单可能对应多个商品
	//OrderItems []OrderItem  `form:"orderItems[]"` // 订单可能对应多个商品
	TotalPrice int          // 前端
	Remark     string       // 备注；前端
	PaySerial  string       // 回调函数
	FinishTime sql.NullTime // 回调函数
	// 进行外键关联
	BuyerID   uint // 后端，根据 SESSION 查到的
	Buyer     User
	SellerID  uint // 每张订单对应一个卖家；后端，根据商品信息查到的
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
}

// SubmitOrder 将订单分项提交到数据库当中
func SubmitOrder(order *Order) (err error) {
	// 新建事务；延迟关闭
	tx := DBSql.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建订单 Order;先查询首件商品的卖家ID
	goods, err := QueryGoodsInfo(order.OrderItems[0].GoodsID)
	if err != nil {
		log.Printf("[error]model-submitOrderItem: 查询首件商品失败%v\n", err)
		return err
	}
	sellerID := goods.UserID
	order.SellerID = sellerID
	if err := createOrder(tx, order); err != nil {
		return err
	}
	// 分项对库存进行修改;创建 OrderItem
	for k, v := range order.OrderItems {
		// 更新库存 Goods
		if err := updateStock(tx, &v); err != nil {
			return err
		}
		// 将订单分项与订单关联;存储 OrderItem
		v.OrderID = order.OrderID
		if err := submitOrderItem(tx, &v); err != nil {
			return err
		}
		// 记录卖家 ID (物品持有人的ID)
		sellerID = v.Goods.UserID
		if k != 0 && sellerID != v.Goods.UserID {
			return errors.New("订单中的商品不属于同一卖家")
		}
	}
	// 更新信息 保存 Order
	//order.SellerID = sellerID        // 将订单与卖家关联; 已经在前面赋值了
	order.Status = OrderStatusUnpaid // 设置订单状态
	err = tx.Save(order).Error
	if err != nil {
		log.Printf("[error]model-submitOrderItem: 保存订单信息失败%v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// createOrder 创建订单
func createOrder(tx *gorm.DB, order *Order) (err error) {
	err = tx.Create(order).Error
	if err != nil {
		log.Printf("[error]model-createOrder: %v", err)
		tx.Rollback()
		return err
	}
	log.Printf("[info]model-createOrder: 创建订单成功,订单ID为%v", order.OrderID)
	return nil
}

// updateStock 更新库存
func updateStock(tx *gorm.DB, orderItem *OrderItem) (err error) {
	var goods Goods
	// 查询商品，加行锁
	if err = tx.Model(&Goods{}).
		Where("goods_id = ?", orderItem.GoodsID).
		Set("gorm:query_option", "FOR UPDATE"). // 关键：悲观锁
		First(&goods).Error; err != nil {
		// 找不到商品，回滚事务
		tx.Rollback()
		return err
	}
	// 减少库存
	if goods.Stock < orderItem.GoodsNum {
		// 库存不足，回滚事务
		tx.Rollback()
		log.Printf("[info]model-submitOrderItem: 库存不足")
		return errors.New("库存不足")
	}
	goods.Stock -= orderItem.GoodsNum
	//goods.Sold += orderItem.GoodsNum // 应该订单结束之后再更新销量
	if err := tx.Save(&goods).Error; err != nil {
		// 更新库存失败，回滚事务
		log.Printf("[error]model-submitOrderItem: %v", err)
		tx.Rollback()
		return err
	}
	return nil
}

// submitOrderItem 保存订单分项
func submitOrderItem(tx *gorm.DB, orderItem *OrderItem) (err error) {
	err = tx.Save(orderItem).Error
	return
}
