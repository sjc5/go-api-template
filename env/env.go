package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	Mode                   string
	IsDev                  bool
	Port                   string
	DSN                    string
	AllowedOrigins         []string
	GracefulTimeoutSeconds int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %s", err)
	}
	Mode = GetEnvAsString("MODE", "production")
	IsDev = Mode == "development"
	Port = GetEnvAsString("PORT", "8080")
	DSN = GetEnvAsString("DSN", "file::memory:?cache=shared")
	AllowedOrigins = strings.Split(GetEnvAsString("COMMA_SEPARATED_ALLOWED_ORIGINS", "*"), ",")
	GracefulTimeoutSeconds = GetEnvAsInt("GRACEFUL_TIMEOUT_SECONDS", 10)
}

func GetEnvAsString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	strValue := GetEnvAsString(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(strValue)
	if err == nil {
		return value
	}
	fmt.Printf("error parsing %s, defaulting to %d. error: %s", key, defaultValue, err)
	return defaultValue
}
