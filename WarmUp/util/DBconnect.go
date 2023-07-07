package util

import (
	"WarmUp/config"
	"WarmUp/model"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectDB() {
	connectMysql()
	connectRedis()
}

func connectMysql() {
	//username := "fir"
	//password := "Mysql@123"
	//host := "localhost"
	//port := "3306"
	//database := "WarmUp"
	//charset := "utf8"
	//dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
	//	username,
	//	password,
	//	host,
	//	port,
	//	database,
	//	charset)
	// 使用 gorm 建立 mysql 数据库的连接
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		config.ProjectConfig.MySQL.User,
		config.ProjectConfig.MySQL.Password,
		config.ProjectConfig.MySQL.Host,
		config.ProjectConfig.MySQL.Port,
		config.ProjectConfig.MySQL.Database,
		config.ProjectConfig.MySQL.Charset)
	log.Printf("[info]读取后拼接的dsn，%v\n", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("mysql connect error %v\n", err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Printf("[error]model-connectMysql,迁移出错%v\n", err.Error())
	}
	model.DB = db
}

var (
	//ErrNil = errors.New("no matching record found in redis database")
	Ctx = context.TODO()
)

func connectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.ProjectConfig.Redis.Addr,
		Password: config.ProjectConfig.Redis.Password,
		DB:       config.ProjectConfig.Redis.DB,
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
