package platform

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(GetEnv().DSN), &gorm.Config{})
	if err != nil {
		log.Panicf("error connecting to database: %s", err)
	}
	fmt.Println("connected to database")
}
