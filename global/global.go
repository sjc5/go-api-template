package global

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/sjc5/go-api-template/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Validate *validator.Validate

func init() {
	// DB
	var err error
	DB, err = gorm.Open(sqlite.Open(env.DSN), &gorm.Config{})
	if err != nil {
		log.Panicf("error connecting to database: %s", err)
	}
	fmt.Println("connected to database")

	// Validator
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
