// Copyright © 2016 Paul Allen <paul@cloudcoreo.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"

	"github.com/CloudCoreo/cli/cmd/content"
	"github.com/spf13/viper"
)

var viperGetString = viper.GetString

// GetValueFromConfig to get values from config
func GetValueFromConfig(key string, masked bool) (value string) {
	if value = viper.GetString(key); value == "" {
		value = content.None
	} else {
		if masked {
			value = maskConfigKey(value)
		}
	}

	return value
}

// UpdateConfig update value in config
func UpdateConfig(key, value string) {
	if value != "" {
		viper.Set(key, value)
	}
}

func maskConfigKey(key string) string {
	if len(key) >= content.MaskLength {
		return fmt.Sprintf("%s%s", content.Mask, key[len(key)-content.MaskLength:])
	}

	return fmt.Sprintf("%s%s", content.Mask, key)
}
