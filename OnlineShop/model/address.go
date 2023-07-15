package model

type Address struct {
	TimeModel
	AddressID uint   `gorm:"primarykey"`
	Address   string `gorm:"type:varchar(100)"`
	Tel       string `gorm:"type:varchar(20)"`
	// 添加外键关联
	UserID uint `json:"-"` // 添加 UserID 字段作为外键
	User   User `json:"-" gorm:"references:UserID"`
}

func CreateAddress(add *Address) error {
	if err := DBSql.Create(add).Error; err != nil {
		return err
	}
	return nil
}
