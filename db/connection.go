package db

import (
	"fmt"
	"log"

	"github.com/julianNot/helpet-api-rest/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = ""
var DB *gorm.DB

func DBConnection(host, user, password, nameDb, port string) {
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, nameDb)
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Maria DB Connected")
		log.Println("Listen " + helpers.GetPortServer() + " Port")
	}
}
