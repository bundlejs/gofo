package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"../component/conf"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
	)

	config := conf.Config()

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.Mysql.User,
		config.Database.Mysql.Password,
		config.Database.Mysql.Host,
		config.Database.Mysql.Name))
	if err != nil {
		panic("failed to connect database")
	}
	defer CloseDB()

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Database.Mysql.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Println(err)
	}
}
