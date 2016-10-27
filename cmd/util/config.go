package util

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/CloudCoreo/cli/cmd/content"
)

// GetValueFromConfig to get values from config
func GetValueFromConfig(key string, masked bool) (value string) {
	if value = viper.GetString(key); value == "" {
		value = content.NONE
	} else {
		if masked {
			value = maskConfigKey(value)
		}
	}

	return value
}

// UpdateConfig update value in config
func UpdateConfig(key, value string ) {
	if value != "" {
		viper.Set(key, value)
	}
}

func maskConfigKey(key string) string {
	if len(key) >= content.MASK_LENGTH {
		return fmt.Sprintf("%s%s", content.MASK, key[len(key) - content.MASK_LENGTH:])
	}

	return fmt.Sprintf("%s%s", content.MASK, key)
}
