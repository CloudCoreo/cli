package util

import (
	"github.com/spf13/viper"
	"github.com/CloudCoreo/cli/cmd/content"
)

// GetValueFromConfig to get values from config
func GetValueFromConfig(key string) (value string) {
	if value = viper.GetString(key); value == "" {
		value = content.NONE
	}

	return value
}

// UpdateConfig update value in config
func UpdateConfig(key, value string ) {
	if value != "" {
		viper.Set(key, value)
	}
}