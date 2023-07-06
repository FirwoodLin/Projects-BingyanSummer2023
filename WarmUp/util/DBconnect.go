package util

import (
	"WarmUp/model"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	connectMysql()
	connectRedis()

}

func connectMysql() {
	username := "fir"
	password := "Mysql@123"
	host := "localhost"
	port := "3306"
	database := "WarmUp"
	charset := "utf8"
	// 使用 gorm 建立 mysql 数据库的连接
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("mysql connect error %v\n", err)
	}
	db.AutoMigrate(&model.User{})
	model.DB = db
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func connectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		// return nil, err
		log.Printf("redis connect error %v\n", err.Error())
		return
	}
	// return &Database{
	// 	Client: client,
	// }, nil
	model.RedisClient = client
	model.RedisCtx = context.Background()
}
