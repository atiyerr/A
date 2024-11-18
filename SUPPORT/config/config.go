// 读取配置文件
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		// Host     string
		// Port     string
		// User     string
		// Password string
		// Name     string
		Dsn          string
		MaxidleConns int
		MaxopenConns int
	}
}

var AppConfig *Config //创建配置文件结构体

func InitConfig() {
	viper.SetConfigName("config")   //指定配置文件名字
	viper.SetConfigType("yml")      //指定配置文件格式
	viper.AddConfigPath("./config") //指定配置文件目录
	err := viper.ReadInConfig()     //读取配置文件
	if err != nil {
		log.Fatalf("读取配置文件失败：%v", err)
	}

	AppConfig = &Config{} //解析配置文件
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("解析配置文件失败：%v", err)
	}

	initDB() //初始化数据库连接
}
