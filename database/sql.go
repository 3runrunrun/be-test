package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql dialect - core
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect for gorm
)

// ConnectSQL to MySQL database
func ConnectSQL() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Parkour3run@/laundry?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("Connection to MySQL is failed: ", err)
	}
	log.Println("Connecting to MySQL ... success!")
	return db
}
