package databasex

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-config"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func NewDatabase() *gorm.DB {
	// new gorm.DB struct
	db, err = gorm.Open(config.Get("SqlSettings", "DriverName").String(""), config.Get("SqlSettings", "DataSource").String(""))
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func GetDatabase() *gorm.DB {
	// Get
	return db
}
