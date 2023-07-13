package model

import (
	"OnlineShop/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DBSql *gorm.DB
var (
	RedisClient *redis.Client
	Ctx         = context.TODO()
	RedisCtx    context.Context
)

func connectMysql() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		config.ProjectConfig.MySQL.User,
		config.ProjectConfig.MySQL.Password,
		config.ProjectConfig.MySQL.Host,
		config.ProjectConfig.MySQL.Port,
		config.ProjectConfig.MySQL.Database,
		config.ProjectConfig.MySQL.Charset)
	//log.Printf("[info]读取后拼接的dsn，%v\n", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Printf("mysql connect error %v\n", err)
	}
	DBSql = db
	log.Printf("[info]mysql connect success\n")
}
func initTables() {
	tables := []interface{}{
		&User{}, &Address{},
		&Goods{}, &Category{},
		&CartItem{},
		&Order{}, &OrderItem{},
	}
	err := DBSql.AutoMigrate(tables...)
	if err != nil {
		log.Printf("[error]model:auto migrate tables error %v\n", err)
	}
	log.Printf("[info]model:auto migrate tables success\n")
}
func connectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.ProjectConfig.Redis.Addr,
		Password: config.ProjectConfig.Redis.Password,
		DB:       config.ProjectConfig.Redis.DB,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		log.Printf("[error]redis connect error %v\n", err.Error())
		return
	}
	RedisClient = client
	RedisCtx = context.Background()
	log.Printf("[info]redis connect success\n")
}

func InitDatabases() {
	connectMysql()
	initTables()
	connectRedis()
}
