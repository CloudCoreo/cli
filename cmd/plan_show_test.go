package main

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

func TestPlanShowCmd(t *testing.T) {
	mockPlan := func(PlanName, planID, branch string, refreshInterval int, enabled bool) *client.Plan {
		return &client.Plan{
			ID:              planID,
			Enabled:         enabled,
			Name:            PlanName,
			Branch:          branch,
			RefreshInterval: refreshInterval,
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.Plan
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds: "enable plan show",
			desc: "show a plan",
			flags: []string{
				"--plan-id", "123",
				"--composite-id", "123",
			},
			resp: []*client.Plan{
				mockPlan("Name1", "PlanID1", "Plan1Branch", 1, false),
				mockPlan("Name2", "PlanID2", "Plan2Branch", 1, true)},
			xout: "",
		},
		{
			cmds: "coreo plan show",
			desc: "show a plan",
			args: []string{""},
			err:  true,
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		frc := &fakeReleaseClient{plans: tt.resp}
		if tt.err {
			frc.err = errors.New("Error")
		}

		cmd := newPlanShowCmd(frc, &buf)
		cmd.ParseFlags(tt.flags)
		err := cmd.RunE(cmd, tt.args)

		if (err != nil) != tt.err {
			t.Errorf("%q. expected error, got '%v'", tt.desc, err)
		}

		re := regexp.MustCompile(tt.xout)
		if !re.Match(buf.Bytes()) {
			t.Fatalf("%q\n\t%s:\nexpected\n\t%q\nactual\n\t%q", tt.cmds, tt.desc, tt.xout, buf.String())
		}
		buf.Reset()
	}
}
