package database

import (
	"fmt"

	"github.com/natapapon-flm/go-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(localhost:8889)/go_basics?parseTime=true"
	
  database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success fully connect DATABASE")
	}

	err = database.AutoMigrate(&models.Item{})
	if err != nil {
		return
	} else {
		fmt.Println("Successfully migrate database")
	}

	DB = database
}