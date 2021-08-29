package config

import (
	"time"

	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
	SetDefault(key string, value interface{})
}

// LoadPrefixedConfigProvider returns a config provider with variables
// loaded from the environment with the prefix. Calls to Get method should
// not include the prefix
func LoadPrefixedConfigProvider(prefix string) Provider {
	return readViperConfig(prefix)
}

// LoadConfigProvider returns a config provider with variables
// loaded from the environment
func LoadConfigProvider() Provider {
	return readViperConfig("")
}

func readViperConfig(prefix string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return v
}
