package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/client"

	"github.com/pkg/errors"
)

const kmsKeyRotatesObjectOutput = `[
	{
		"totalItems": 200,
		"violations": [
			{
				"ObjectName": "fake-id1",
				"SuggestedAction": "fake-suggestion",
				"KnowledgeBase": "fake-link",
				"RuleDescription": "fake description",
				"Name": "fake-display-name",
				"Severity": "um",
				"RuleService": "kms",
				"RuleName": "fake-name",
				"RiskScore": 0,
				"RiskScoreSum": 0,
				"TeamName": "fake-team-name",
				"FindingRegion": "us-east-1"
			}
		]
	}
]
`

const iamInactiveKeyNoRotationObjectOutput = `[
	{
		"totalItems": 100,
		"violations": [
			{
				"ObjectName": "fake-id2",
				"SuggestedAction": "fake-suggestion",
				"KnowledgeBase": "fake-link",
				"RuleDescription": "fake description",
				"Name": "fake-display-name",
				"Severity": "Medium",
				"RuleService": "iam",
				"RuleName": "fake-name",
				"RiskScore": 0,
				"RiskScoreSum": 0,
				"TeamName": "fake-team-name",
				"FindingRegion": "us-west-2"
			}
		]
	}
]
`

func TestResultObjectCmd(t *testing.T) {
	mockObject := func(id string, info client.Info,
		tInfo client.TeamInfo, region string) *client.ResultObject {
		return &client.ResultObject{
			ObjectID: id,
			Info:     info,
			TInfo:    tInfo,
			Region:   region,
		}
	}

	mockInt := func(num int) *int {
		integer := num
		return &integer
	}
	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.ResultObject
		num   *int
		json  bool
		err   bool
		xout  string
	}{
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			resp: []*client.ResultObject{
				mockObject(
					"fake-id1",
					client.Info{
						SuggestedAction:          "fake-suggestion",
						Link:                     "fake-link",
						Description:              "fake description",
						DisplayName:              "fake-display-name",
						Level:                    "um",
						Service:                  "kms",
						RuleName:                 "fake-name",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:55.111+00:00",
					},
					client.TeamInfo{
						Name: "fake-team-name",
						ID:   "fake-team-id",
					}, "us-east-1"),
			},
			num:  mockInt(200),
			xout: kmsKeyRotatesObjectOutput,
		},
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			resp: []*client.ResultObject{
				mockObject(
					"fake-id2",
					client.Info{
						SuggestedAction:          "fake-suggestion",
						Link:                     "fake-link",
						Description:              "fake description",
						DisplayName:              "fake-display-name",
						Level:                    "Medium",
						Service:                  "iam",
						RuleName:                 "fake-name",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:54.448+00:00",
					},
					client.TeamInfo{
						Name: "fake-team-name",
						ID:   "fake-team-id",
					}, "us-west-2"),
			},
			num:  mockInt(100),
			xout: iamInactiveKeyNoRotationObjectOutput,
		},
	}

	var buf bytes.Buffer
	for _, test := range tests {
		objectWrapper := &client.ResultObjectWrapper{
			Objects:    test.resp,
			TotalItems: *test.num,
		}
		frc := &fakeReleaseClient{objects: objectWrapper}
		if test.err {
			frc.err = errors.New("Error")
		}

		cmd := newResultObjectCmd(frc, &buf)
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
