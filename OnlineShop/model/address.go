package model

type Address struct {
	TimeModel
	AddressID uint   `gorm:"primarykey"`
	Address   string `gorm:"type:varchar(100)"`
	Tel       string
	UserID    uint // 添加 UserID 字段作为外键
	User      User `gorm:"references:UserID"`
}
