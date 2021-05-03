package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbConn   *gorm.DB
	user     = "root"
	password = ""
	host     = "localhost"
	port     = "3306"
	dbName   = "orders_by"
)

func newConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetMaxIdleConns(20)
	sqlDb.SetConnMaxLifetime(300 * time.Second)

	return db
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		return newConnection()
	}
	return dbConn
}
