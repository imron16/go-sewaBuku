package database

import (
	"fmt"
	"log"

	"github.com/imron16/go-sewaBuku/entity"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	fmt.Println("Try Setup Database ...")
	var (
		host     = "localhost"
		username = "root"
		password = "Test1234"
		database = "sewaBuku"
		port     = "3306"
	)

	localDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	DB, err := gorm.Open("mysql", localDsn)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success Connect database")
	}

	DB.LogMode(true)

	return DB

}

// Migrate create/updates database table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Buku{})
	db.AutoMigrate(&entity.Sewa{})
	log.Println("Table migrated")
}
