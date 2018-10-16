package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

const iamInactiveKeyNoRotationRuleOutput = `[
	{
		"id": "iam-inactive-key-no-rotation",
		"info": {
			"suggested_action": "If you regularly use the AWS access keys, we recommend that you also regularly rotate or delete them.",
			"link": "http://kb.cloudcoreo.com/mydoc_iam-inactive-key-no-rotation.html",
			"description": "User has inactive keys that have not been rotated in the last 90 days.",
			"display_name": "User Has Access Keys Inactive and Un-rotated",
			"level": "Medium",
			"service": "iam",
			"name": "iam-inactive-key-no-rotation",
			"region": "global",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:54.448+00:00"
		},
		"teams": [
			{
				"name": "zechen2",
				"id": "5bb6a4956365930011a41a0b"
			}
		],
		"accounts": [
			{
				"name": "new-test",
				"id": "530342348278"
			}
		],
		"objects": 1528
	}
]
`

const S3AllUserWriteRuleOutput = `[
	{
		"id": "s3-allusers-write",
		"info": {
			"suggested_action": "Remove the entry from the bucket permissions that allows everyone to write.",
			"link": "http://kb.cloudcoreo.com/mydoc_s3-allusers-write.html",
			"description": "Bucket has permissions (ACL) which let all users write to the bucket.",
			"display_name": "All users can write to the affected bucket",
			"level": "High",
			"service": "s3",
			"name": "s3-allusers-write",
			"region": "us-east-1",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:55.387+00:00"
		},
		"teams": [
			{
				"name": "zechen2",
				"id": "5bb6a4956365930011a41a0b"
			}
		],
		"accounts": [
			{
				"name": "new-test",
				"id": "530342348278"
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
					"iam-inactive-key-no-rotation",
					client.Info{
						SuggestedAction: "If you regularly use the AWS access keys, we recommend that you also regularly rotate or delete them.",
						Link:            "http://kb.cloudcoreo.com/mydoc_iam-inactive-key-no-rotation.html",
						Description:     "User has inactive keys that have not been rotated in the last 90 days.",
						DisplayName:     "User Has Access Keys Inactive and Un-rotated",
						Level:           "Medium",
						Service:         "iam",
						Name:            "iam-inactive-key-no-rotation",
						Region:          "global",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:54.448+00:00",
					},
					[]client.TeamInfo{
						{
							Name: "zechen2",
							ID:   "5bb6a4956365930011a41a0b",
						},
					},
					[]client.CloudAccountInfo{
						{
							Name: "new-test",
							ID:   "530342348278",
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
					"s3-allusers-write",
					client.Info{
						SuggestedAction: "Remove the entry from the bucket permissions that allows everyone to write.",
						Link:            "http://kb.cloudcoreo.com/mydoc_s3-allusers-write.html",
						Description:     "Bucket has permissions (ACL) which let all users write to the bucket.",
						DisplayName:     "All users can write to the affected bucket",
						Level:           "High",
						Service:         "s3",
						Name:            "s3-allusers-write",
						Region:          "us-east-1",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:55.387+00:00",
					},
					[]client.TeamInfo{
						{
							Name: "zechen2",
							ID:   "5bb6a4956365930011a41a0b",
						},
					},
					[]client.CloudAccountInfo{
						{
							Name: "new-test",
							ID:   "530342348278",
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
