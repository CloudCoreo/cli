package util

import (
	"fmt"
	"testing"
)

func TestGetValueFromConfig(t *testing.T) {
	oldViperGetString := viperGetString
	defer func() { viperGetString = oldViperGetString }()
	viperGetString = func(key string) string {
		fmt.Println("printvalue")
		return "someValue"
	}
	GetValueFromConfig("string", true)
}
