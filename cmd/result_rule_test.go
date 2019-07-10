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
			"suggestedAction": "fake suggestion",
			"link": "fake link",
			"description": "fake description",
			"displayName": "fake-display-name",
			"level": "Medium",
			"service": "iam",
			"name": "fake-name",
			"include_violations_in_count": true,
			"lastUpdateTime": "2018-10-11T17:21:54.448+00:00"
		},
		"teamAndPlan": [
			{
				"team": {
					"name": "fake-team-name",
					"teamId": "fake-team-id"
				}
			}
		],
		"accounts": [
			"fake-account-id"
		],
		"objects": 1528,
		"regions": [
			"us-east-1"
		]
	}
]
`

const S3AllUserWriteRuleOutput = `[
	{
		"id": "fake-id2",
		"info": {
			"suggestedAction": "fake suggestion",
			"link": "fake link",
			"description": "fake description",
			"displayName": "fake-display-name",
			"level": "High",
			"service": "s3",
			"name": "fake-name",
			"include_violations_in_count": true,
			"lastUpdateTime": "2018-10-11T17:21:55.387+00:00"
		},
		"teamAndPlan": [
			{
				"team": {
					"name": "fake-team-name",
					"teamId": "fake-team-id"
				}
			}
		],
		"accounts": [
			"fake-account-id"
		],
		"objects": 2,
		"regions": [
			"us-west-1"
		]
	}
]
`

func TestResultRuleCmd(t *testing.T) {
	mockRule := func(id string, info client.Info,
		tInfo []client.TeamInfo, cInfo []string, object int, regions []string) *client.ResultRule {
		return &client.ResultRule{
			ID:      id,
			Info:    info,
			TInfo:   []client.TeamInfoWrapper{{TeamInfo: &tInfo[0]}},
			CInfo:   cInfo,
			Object:  object,
			Regions: regions,
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
						SuggestedAction:          "fake suggestion",
						Link:                     "fake link",
						Description:              "fake description",
						DisplayName:              "fake-display-name",
						Level:                    "Medium",
						Service:                  "iam",
						RuleName:                 "fake-name",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:54.448+00:00",
					},
					[]client.TeamInfo{
						{
							Name: "fake-team-name",
							ID:   "fake-team-id",
						},
					},
					[]string{
						"fake-account-id",
					},
					1528, []string{"us-east-1"}),
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
						SuggestedAction:          "fake suggestion",
						Link:                     "fake link",
						Description:              "fake description",
						DisplayName:              "fake-display-name",
						Level:                    "High",
						Service:                  "s3",
						RuleName:                 "fake-name",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:55.387+00:00",
					},
					[]client.TeamInfo{
						{
							Name: "fake-team-name",
							ID:   "fake-team-id",
						},
					},
					[]string{"fake-account-id"},
					2, []string{"us-west-1"}),
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
