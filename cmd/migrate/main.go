package main

import (
	"fmt"
	"log"

	"github.com/sjc5/go-api-template/internal/model"
	"github.com/sjc5/go-api-template/internal/platform"
)

func main() {
	err := platform.DB.Debug().AutoMigrate(
		// Add new models here whenever you need
		&model.User{},
		&model.Session{},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrated database")
}
