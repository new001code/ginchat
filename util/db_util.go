package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ginchat/common"
)

var db *gorm.DB

func init() {
	common.Logger.Info("start database pool init")
	databaseInit()
}

func databaseInit() {
	databaseType := viper.GetString("dataSource.type")

	// mysql
	if databaseType == "mysql" {

		databaseName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8mb4&loc=Local",
			viper.GetString("dataSource.username"),
			viper.GetString("dataSource.password"),
			viper.GetString("dataSource.host"),
			viper.GetInt("dataSource.port"),
			viper.GetString("dataSource.database"),
		)
		conn, err := gorm.Open(mysql.Open(databaseName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("database source type: %s init err: %s", "mysql", err))
		}
		pool, err := conn.DB()
		if err != nil {
			panic(fmt.Errorf("database pool type: %s init err: %s", "mysql", err))
		}
		pool.SetMaxIdleConns(viper.GetInt("databasePool.maxIdleConns"))
		pool.SetMaxOpenConns(viper.GetInt("databasePool.maxOpenConns"))
		pool.SetConnMaxLifetime(time.Second * 600)
		db = conn
	} else if databaseType == "postgres" {
		databaseName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable  TimeZone=Asia/Shanghai",
			viper.GetString("dataSource.host"),
			viper.GetString("dataSource.username"),
			viper.GetString("dataSource.password"),
			viper.GetString("dataSource.database"),
			viper.GetInt("dataSource.port"),
		)
		conn, err := gorm.Open(postgres.Open(databaseName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("database source type: %s init err: %s", "postgres", err))
		}
		pool, err := conn.DB()
		if err != nil {
			panic(fmt.Errorf("database pool type: %s init err: %s", "postgres", err))
		}
		pool.SetMaxIdleConns(viper.GetInt("databasePool.maxIdleConns"))
		pool.SetMaxOpenConns(viper.GetInt("databasePool.maxOpenConns"))
		pool.SetConnMaxLifetime(time.Second * 600)
		db = conn
	} else {
		databaseName := viper.GetString("database.database")
		conn, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("database source type: %s init err: %s", "sqlite", err))
		}
		pool, err := conn.DB()
		if err != nil {
			panic(fmt.Errorf("database pool type: %s init err: %s", "sqlite", err))
		}
		pool.SetMaxIdleConns(viper.GetInt("databasePool.maxIdleConns"))
		pool.SetMaxOpenConns(viper.GetInt("databasePool.maxOpenConns"))
		pool.SetConnMaxLifetime(time.Second * 600)
		db = conn

	}
	common.Logger.Info("success database pool init")

}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		databaseInit()
	}

	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		databaseInit()
	}

	return db
}
