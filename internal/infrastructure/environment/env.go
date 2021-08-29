package environment

import (
	"errors"
	"fmt"
	"strings"

	"github.com/adolsalamanca/go-rest-boilerplate/internal/infrastructure/config"
)

const (
	DBHost        = "DB_HOST"
	DBName        = "DB_NAME"
	DBPass        = "DB_PASS"
	DBPort        = "DB_PORT"
	DBUser        = "DB_USER"
	ServerPort    = "SERVER_PORT"
	StatsdAddress = "STATSD_HOST"
	StatsdPort    = "STATSD_PORT"
)

var (
	RequiredVars = []string{
		DBHost,
		DBName,
		DBPort,
		DBUser,
	}

	OptionalVars = []string{
		DBPass,
		ServerPort,
		StatsdAddress,
		StatsdPort,
	}
)

func Verify(cfg config.Provider) error {
	for _, key := range OptionalVars {
		val := cfg.Get(key)
		if val == nil {
			fmt.Printf("optional environment vars '%s' missing \n", key)
		}
	}

	keys := make([]string, 0)
	for _, key := range RequiredVars {
		val := cfg.Get(key)
		if val == nil {
			keys = append(keys, key)
		}
	}

	if len(keys) > 0 {
		return errors.New(fmt.Sprintf("required environment vars %s missing \n", strings.Join(keys, ", ")))
	}

	return nil
}
