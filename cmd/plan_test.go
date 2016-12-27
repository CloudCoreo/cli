package main

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

func TestPlanListCmd(t *testing.T) {
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
			cmds: "coreo plan list",
			desc: "get list of plans",
			flags: []string{
				"--composite-id", "123",
			},
			resp: []*client.Plan{
				mockPlan("Name1", "PlanID1", "Plan1Branch", 1, false),
				mockPlan("Name2", "PlanID2", "Plan2Branch", 1, true)},
			xout: "------------  --------------  -----------  ----------------  -------------\n" +
				"   Plan ID       Plan Name       Active       Git Branch        Interval  \n" +
				"------------  --------------  -----------  ----------------  -------------\n" +
				"   PlanID1         Name1         false        Plan1Branch          1      \n\n" +
				"   PlanID2         Name2          true        Plan2Branch          1      \n" +
				"------------  --------------  -----------  ----------------  -------------\n\n",
		},
		{
			cmds: "coreo plan list",
			desc: "get list of plan",
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

		cmd := newPlanListCmd(frc, &buf)
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
