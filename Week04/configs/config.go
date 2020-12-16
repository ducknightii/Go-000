package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var Conf *ConfigYaml

type ConfigYaml struct {
	Server struct {
		Listen string `required:"true" flagUsage:"服务监听地址"`
		Env    string `default:"Pro" flagUsage:"服务运行时环境"`
	}
	Database struct {
		Driver string `default:"mysql" flagUsage:"数据库类型：mysql|sqlite3"`
		Mysql  struct {
			HostPort     string `flagUsage:"数据库连接，eg：tcp(127.0.0.1:3306)"`
			UserPassword string `flagUsage:"数据库账号密码"`
			DB           string `flagUsage:"数据库"`
		}
		Sqlite3 struct {
			DB string `flagUsage:"数据库"`
		}
		Conn struct {
			PingInterval int `default:"10" flagUsage:"mysql ping时间间隔，保持连接不被mysql server断开"`
			MaxLifeTime  int `default:"600" flagUsage:"连接最长存活时间，单位s"`
			MaxIdle      int `default:"10" flagUsage:"最多空闲连接数"`
			MaxOpen      int `default:"80" flagUsage:"最多打开连接数"`
			PerOpTimeout int `default:"0" flagUsage:"单次操作超时时间"`
		}
	}
}

func (c *ConfigYaml) IsProduction() bool {
	return strings.ToLower(c.Server.Env) == "pro"
}

// todo 尝试从 yaml 环境变量等多个地方加载配置
func InitConfig() {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Conf = new(ConfigYaml)
	if err := vp.Unmarshal(Conf); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", Conf)
}
