package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestCreateAndQueryUserAndAddress(t *testing.T) {
	// 连接数据库
	dsn := "fir:Mysql@123@tcp(127.0.0.1:3306)/OnlineShop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	//db.LogMode(true)
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Address{})
	if err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}
	// 创建用户
	user := User{Name: "Alice", Nickname: "A", Email: "alice@example.com", Tel: "12345678901", HashedPassword: "password", IsAdmin: false}
	db.Create(&user)
	fmt.Printf("@@@User %s created,id:%v\n", user.Name, user.UserID)
	// 创建地址
	address1 := Address{Address: "Address 1", Tel: "12345678", UserID: user.UserID}
	db.Create(&address1)

	address2 := Address{Address: "Address 2", Tel: "87654321", UserID: user.UserID}
	db.Create(&address2)

	//sol1
	// 可以运行
	// 查询用户所有地址
	var addresses []Address
	db.Model(&User{}).Preload("Addresses")
	// preload 查询关联的数据
	//err = db.Model(&user.go).Preload("Addresses").Find(&user.go).Error
	//db.Debug().Where("user_id = ?", user.go.UserID).Preload("Addresses").Find(&addresses)
	err = db.Model(&user).Association("Addresses").Find(&addresses)

	fmt.Printf("User %s's addresses:\n", user.Name)
	for _, address := range addresses {
		fmt.Printf("- %s, %s\n", address.Address, address.Tel)
	}

	//// sol2
	//// 有 bug
	//var addresses []Address
	//db.Model(&user.go).Preload("Users").Find(&addresses)
	//fmt.Printf("User %s's addresses:\n", user.go.Name)
	//for _, address := range addresses {
	//	fmt.Printf("- %s, %s\n", address.Address, address.Tel)
	//}
}
