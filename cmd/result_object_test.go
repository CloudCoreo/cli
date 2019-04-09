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
				"objectName": "fake-id1",
				"suggested_action": "fake-suggestion",
				"link": "fake-link",
				"description": "fake description",
				"display_name": "fake-display-name",
				"level": "Medium",
				"service": "kms",
				"name": "fake-name",
				"region": "us-east-1",
				"include_violations_in_count": true,
				"timestamp": "2018-10-11T17:21:55.111+00:00",
				"riskScore": 0,
				"team": {
					"name": "fake-team-name",
					"id": "fake-team-id"
				}
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
				"objectName": "fake-id2",
				"suggested_action": "fake-suggestion",
				"link": "fake-link",
				"description": "fake description",
				"display_name": "fake-display-name",
				"level": "Medium",
				"service": "iam",
				"name": "fake-name",
				"region": "global",
				"include_violations_in_count": true,
				"timestamp": "2018-10-11T17:21:54.448+00:00",
				"riskScore": 0,
				"team": {
					"name": "fake-team-name",
					"id": "fake-team-id"
				}
			}
		]
	}
]
`

func TestResultObjectCmd(t *testing.T) {
	mockObject := func(id string, info client.Info,
		tInfo client.TeamInfo) *client.ResultObject {
		return &client.ResultObject{
			ID:    id,
			Info:  info,
			TInfo: tInfo,
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
						SuggestedAction: "fake-suggestion",
						Link:            "fake-link",
						Description:     "fake description",
						DisplayName:     "fake-display-name",
						Level:           "Medium",
						Service:         "kms",
						Name:            "fake-name",
						Region:          "us-east-1",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:55.111+00:00",
					},
					client.TeamInfo{
						Name: "fake-team-name",
						ID:   "fake-team-id",
					}),
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
						SuggestedAction: "fake-suggestion",
						Link:            "fake-link",
						Description:     "fake description",
						DisplayName:     "fake-display-name",
						Level:           "Medium",
						Service:         "iam",
						Name:            "fake-name",
						Region:          "global",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:54.448+00:00",
					},
					client.TeamInfo{
						Name: "fake-team-name",
						ID:   "fake-team-id",
					}),
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
