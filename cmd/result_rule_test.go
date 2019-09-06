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
			xout: "Findings results are deprecated, please follow the link to swagger API doc `https://api.securestate.vmware.com` \n",
		},
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			xout: "Findings results are deprecated, please follow the link to swagger API doc `https://api.securestate.vmware.com` \n",
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
		cmd.Run(cmd, test.args)

		if buf.String() != test.xout {
			t.Fatalf("%q\n\t%s:\nexpected\n\t%q\nactual\n\t%q", test.cmds, test.desc, test.xout, buf.String())
		}

		buf.Reset()
	}
}
