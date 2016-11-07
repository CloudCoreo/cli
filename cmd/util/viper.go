package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// SaveViperConfig Save viper config to file
func SaveViperConfig() error {
	f, err := os.Create(viper.ConfigFileUsed())
	if err != nil {
		return err
	}
	defer f.Close()

	all := viper.AllSettings()

	b, err := yaml.Marshal(all)
	if err != nil {
		return fmt.Errorf("Panic while encoding into YAML format.")
	}
	if _, err := f.WriteString(string(b)); err != nil {
		return err
	}
	return nil
}
