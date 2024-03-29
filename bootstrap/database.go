package bootstrap

import (
	"errors"
	"fmt"
	"gin-skeleton/pkg/config"
	"gin-skeleton/pkg/database"
	"gin-skeleton/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func SetupDB()  {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("database connection not supported"))
	}

	// connection db && set gorm log mode
	database.Connect(dbConfig, logger.NewGormLogger())
	// set max connection
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// set max free conncetion
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// set connection expire time
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}
