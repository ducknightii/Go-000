package db

import (
	"fmt"
	"github.com/ducknightii/Go-000/Week04/configs"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// todo 作为一个pkg 全局变量的方式 是否合适？
var DB *gorm.DB

func Init() {
	var err error
	var dsn string
	switch configs.Conf.Database.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s@%s/%s?charset=utf8&parseTime=True&loc=Local&timeout=2s&readTimeout=%ds&writeTimeout=%ds&tx_isolation=%%27READ-COMMITTED%%27", configs.Conf.Database.Mysql.UserPassword, configs.Conf.Database.Mysql.HostPort, configs.Conf.Database.Mysql.DB, configs.Conf.Database.Conn.PerOpTimeout, configs.Conf.Database.Conn.PerOpTimeout)

		DB, err = gorm.Open(configs.Conf.Database.Driver, dsn)
		if err != nil {
			panic(err)
		}
		DB.DB().SetConnMaxLifetime(time.Duration(configs.Conf.Database.Conn.MaxLifeTime) * time.Second)
		DB.DB().SetMaxIdleConns(configs.Conf.Database.Conn.MaxIdle)
		DB.DB().SetMaxOpenConns(configs.Conf.Database.Conn.MaxOpen)
		go func() {
			for {
				DB.DB().Ping()
				time.Sleep(time.Duration(configs.Conf.Database.Conn.PingInterval) * time.Second)
			}
		}()
	case "sqlite3":
		dsn = fmt.Sprintf("%s.database?cache=shared&_synchronous=1&_journal_mode=WAL", configs.Conf.Database.Sqlite3.DB)
		DB, err = gorm.Open(configs.Conf.Database.Driver, dsn)
		if err != nil {
			panic(err)
		}
		DB.DB().SetMaxOpenConns(1)
	default:
		panic(fmt.Sprintf("unknown driver %s", configs.Conf.Database.Driver))
	}

	// 开发环境打印所有sql语句
	if !configs.Conf.IsProduction() {
		DB.LogMode(true)
	}
}
