package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const gracefulTimeoutDefault = 10

var (
	Mode                   = GetEnvAsString("MODE", "production")
	IsDev                  = Mode == "development"
	Port                   = GetEnvAsString("PORT", "8080")
	DSN                    = GetEnvAsString("DSN", "file::memory:?cache=shared")
	AllowedOrigins         = GetEnvAsStringSlice("COMMA_SEPARATED_ALLOWED_ORIGINS", []string{})
	GracefulTimeoutSeconds = GetEnvAsInt("GRACEFUL_TIMEOUT_SECONDS", gracefulTimeoutDefault)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %s", err)
	}
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

func GetEnvAsStringSlice(key string, defaultValue []string) []string {
	strValue := GetEnvAsString(key, strings.Join(defaultValue, ","))
	return strings.Split(strValue, ",")
}
