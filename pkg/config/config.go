package config

import (
	"github.com/spf13/viper"
)

// Provider defines read-only methods for accessing application's configuration
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetString(key string) string
}

// LoadConfigProvider returns a config provider with vars loaded from the environment
func LoadConfigProvider() Provider {
	return readViperConfig("")
}

func readViperConfig(prefix string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return v
}
