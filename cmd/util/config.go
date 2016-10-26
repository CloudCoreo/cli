package util

import (
	"github.com/spf13/viper"
	"github.com/CloudCoreo/cli/cmd/content"
)

func GetValueFromConfig(key string) (value string) {
	if value = viper.GetString(key); value == "" {
		value = content.NONE
	}

	return value
}


func UpdateConfig(key, value string ) {
	if value != "" {
		viper.Set(key, value)
	}
}