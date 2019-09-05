package main

import (
	"bytes"
	"testing"

	"github.com/pkg/errors"
)

func TestResultRuleCmd(t *testing.T) {
	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			xout: "",
		},
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			xout: "",
		},
	}

	var buf bytes.Buffer
	for _, test := range tests {
		frc := &fakeReleaseClient{}
		if test.err {
			frc.err = errors.New("Error")
		}
		cmd := newResultRuleCmd(frc, &buf)
		cmd.ParseFlags(test.flags)

		if buf.String() != test.xout {
			t.Fatalf("%q\n\t%s:\nexpected\n\t%q\nactual\n\t%q", test.cmds, test.desc, test.xout, buf.String())
		}

		buf.Reset()
	}
}
