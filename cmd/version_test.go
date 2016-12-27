package main

import (
	"bytes"
	"testing"
)

func TestVersion(t *testing.T) {

	tests := []struct {
		name string
		args []string
		fail bool
	}{
		{"default", []string{}, false},
	}

	for _, tt := range tests {
		b := new(bytes.Buffer)

		cmd := newVersionCmd(b)
		cmd.ParseFlags(tt.args)
		if err := cmd.RunE(cmd, tt.args); err != nil {
			if tt.fail {
				continue
			}
			t.Fatal(err)
		}
	}
}
