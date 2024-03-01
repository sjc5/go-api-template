package platform

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sjc5/kit/pkg/envutil"
)

type envType struct {
	Mode                   string
	IsDev                  bool
	Port                   string
	DSN                    string
	AllowedOrigins         []string
	GracefulTimeoutSeconds int
}

var env *envType = nil

func GetEnv() *envType {
	if env != nil {
		return env
	}
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %s", err)
	}
	env = &envType{}
	env.Mode = envutil.GetStr("MODE", "production")
	env.IsDev = env.Mode == "development"
	env.Port = envutil.GetStr("PORT", "8080")
	env.DSN = envutil.GetStr("DSN", "file::memory:?cache=shared")
	env.AllowedOrigins = strings.Split(envutil.GetStr("COMMA_SEPARATED_ALLOWED_ORIGINS", "*"), ",")
	env.GracefulTimeoutSeconds = envutil.GetInt("GRACEFUL_TIMEOUT_SECONDS", 10)
	return env
}
