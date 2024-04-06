package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// 总配置文件
type config struct {
	Server        server        `yaml:"server"`
	Db            db            `yaml:"db"`
	Redis         redis         `yaml:"redis"`
	Log           log           `yaml:"log"`
	ImageSettings imageSettings `yaml:"imageSettings"`
}

type server struct {
	Address string `yaml:"address"`
	Mode    string `yaml:"mode"`
}

// DB 数据库配置
type db struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis 配置
type redis struct {
	Address  string `yaml:"address"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

// log 配置
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model.go"`
}

// 图片配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
}

var Config *config

// Init 初始化配置
func Init() {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")    // 还可以在工作目录中查找配置
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config) // 绑定结构体
	if err != nil {
		fmt.Printf("bind viper err:%v", err)
		return
	}
	//fmt.Println(Config)

}
