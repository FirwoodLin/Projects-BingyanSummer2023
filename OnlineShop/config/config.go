package config

import (
	"github.com/spf13/viper"
	"log"
)

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"DB"`
}

type Config struct {
	MySQL MySQLConfig `mapstructure:"Mysql"`
	Redis RedisConfig `mapstructure:"Redis"`
}

var ProjectConfig Config

func ReadIn() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config-配置信息读取出错，%v\n", err.Error())
	}
	if err := viper.Unmarshal(&ProjectConfig); err != nil {
		log.Printf("config-配置信息解码出错，%v\n", err.Error())
	}
	log.Printf("[info]读取到的配置信息：%v\n", ProjectConfig)

}
func init() {
	ReadIn()
}
