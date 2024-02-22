package main

import (
	"fmt"
	"log"

	"github.com/sjc5/go-api-template/global"
	"github.com/sjc5/go-api-template/model"
)

func main() {
	err := global.DB.Debug().AutoMigrate(
		// Add new models here whenever you need
		&model.User{},
		&model.Session{},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrated database")
}
