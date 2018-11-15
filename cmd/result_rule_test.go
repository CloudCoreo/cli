package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/client"

	"github.com/pkg/errors"
)

const iamInactiveKeyNoRotationRuleOutput = `[
	{
		"id": "fake-id1",
		"info": {
			"suggested_action": "fake suggestion",
			"link": "fake link",
			"description": "fake description",
			"display_name": "fake-display-name",
			"level": "Medium",
			"service": "iam",
			"name": "fake-name",
			"region": "global",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:54.448+00:00"
		},
		"teams": [
			{
				"name": "fake-team-name",
				"id": "fake-team-id"
			}
		],
		"accounts": [
			{
				"name": "fake-account-name",
				"id": "fake-account-id"
			}
		],
		"objects": 1528
	}
]
`

const S3AllUserWriteRuleOutput = `[
	{
		"id": "fake-id2",
		"info": {
			"suggested_action": "fake suggestion",
			"link": "fake link",
			"description": "fake description",
			"display_name": "fake-display-name",
			"level": "High",
			"service": "s3",
			"name": "fake-name",
			"region": "us-east-1",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:55.387+00:00"
		},
		"teams": [
			{
				"name": "fake-team-name",
				"id": "fake-team-id"
			}
		],
		"accounts": [
			{
				"name": "fake-account-name",
				"id": "fake-account-id"
			}
		],
		"objects": 2
	}
]
`

func TestResultRuleCmd(t *testing.T) {
	mockRule := func(id string, info client.Info,
		tInfo []client.TeamInfo, cInfo []client.CloudAccountInfo, object int) *client.ResultRule {
		return &client.ResultRule{
			ID:     id,
			Info:   info,
			TInfo:  tInfo,
			CInfo:  cInfo,
			Object: object,
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.ResultRule
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			resp: []*client.ResultRule{
				mockRule(
					"fake-id1",
					client.Info{
						SuggestedAction: "fake suggestion",
						Link:            "fake link",
						Description:     "fake description",
						DisplayName:     "fake-display-name",
						Level:           "Medium",
						Service:         "iam",
						Name:            "fake-name",
						Region:          "global",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:54.448+00:00",
					},
					[]client.TeamInfo{
						{
							Name: "fake-team-name",
							ID:   "fake-team-id",
						},
					},
					[]client.CloudAccountInfo{
						{
							Name: "fake-account-name",
							ID:   "fake-account-id",
						},
					},
					1528),
			},
			xout: iamInactiveKeyNoRotationRuleOutput,
		},
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			resp: []*client.ResultRule{
				mockRule(
					"fake-id2",
					client.Info{
						SuggestedAction: "fake suggestion",
						Link:            "fake link",
						Description:     "fake description",
						DisplayName:     "fake-display-name",
						Level:           "High",
						Service:         "s3",
						Name:            "fake-name",
						Region:          "us-east-1",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:55.387+00:00",
					},
					[]client.TeamInfo{
						{
							Name: "fake-team-name",
							ID:   "fake-team-id",
						},
					},
					[]client.CloudAccountInfo{
						{
							Name: "fake-account-name",
							ID:   "fake-account-id",
						},
					},
					2),
			},
			xout: S3AllUserWriteRuleOutput,
		},
	}

	var buf bytes.Buffer
	for _, test := range tests {
		frc := &fakeReleaseClient{rules: test.resp}
		if test.err {
			frc.err = errors.New("Error")
		}
		cmd := newResultRuleCmd(frc, &buf)
		cmd.ParseFlags(test.flags)
		err := cmd.RunE(cmd, test.args)

		if (err != nil) != test.err {
			t.Errorf("%q. expected error, got '%v'", test.desc, err)
		}

		if buf.String() != test.xout {
			t.Fatalf("%q\n\t%s:\nexpected\n\t%q\nactual\n\t%q", test.cmds, test.desc, test.xout, buf.String())
		}

		buf.Reset()
	}
}
