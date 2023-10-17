package database

import (
	"fmt"
	"os"

	"github.com/natapapon-flm/go-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

  database, err := gorm.Open(mysql.Open(dsn))
	
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success fully connect DATABASE")
	}

	err = database.AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully migrate database")
	}

	DB = database
}