package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const gracefulTimeoutDefault = 10

var (
	Mode                   = "production"
	IsDev                  = false
	Port                   = "8080"
	DSN                    = "file::memory:?cache=shared"
	AllowedOrigins         []string
	GracefulTimeoutSeconds = gracefulTimeoutDefault
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}
	Mode = os.Getenv("MODE")
	IsDev = Mode == "development"
	Port = os.Getenv("PORT")
	DSN = os.Getenv("DSN")
	allowedOriginsString := os.Getenv("COMMA_SEPARATED_ALLOWED_ORIGINS")
	allowedOrigins := strings.Split(allowedOriginsString, ",")
	AllowedOrigins = allowedOrigins
	GracefulTimeoutSeconds, err = strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT_SECONDS"))
	if err != nil {
		fmt.Printf("error parsing GRACEFUL_TIMEOUT_SECONDS, defaulting to %d. Error: %s", gracefulTimeoutDefault, err)
		GracefulTimeoutSeconds = gracefulTimeoutDefault
	}
}
