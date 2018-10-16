package main

import (
	"bytes"
	"testing"

	"github.com/CloudCoreo/cli/client"
	"github.com/pkg/errors"
)

const kmsKeyRotatesObjectOutput = `[
	{
		"id": "a7288f05-157a-4043-ab1a-f55709457807",
		"rule_report": {
			"suggested_action": "It is recommended that CMK key rotation be enabled.",
			"link": "http://kb.cloudcoreo.com/mydoc_kms-key-rotates.html",
			"description": "AWS Key Management Service (KMS) allows customers to rotate the backing key which is key material stored within the KMS which is tied to the key ID of the Customer Created customer master key (CMK). It is the backing key that is used to perform cryptographic operations such as encryption and decryption. Automated key rotation currently retains all prior backing keys so that decryption of encrypted data can take place transparently.",
			"display_name": "Verify rotation for customer created CMKs is enabled",
			"level": "Medium",
			"service": "kms",
			"name": "kms-key-rotates",
			"region": "us-east-1",
			"include_violations_in_count": true,
			"timestamp": "2018-10-11T17:21:55.111+00:00"
		},
		"team": {
			"name": "zechen2",
			"id": "5bb6a4956365930011a41a0b"
		},
		"cloud_account": {
			"name": "new-test",
			"id": "530342348278"
		},
		"run_id": "1050436168129818625"
	}
]
`

const iamInactiveKeyNoRotationObjectOutput = `[
	{
		"id": "coreo-team-5b6c76cc2bc8452fe4586bce",
		"rule_report": {
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
		"team": {
			"name": "zechen2",
			"id": "5bb6a4956365930011a41a0b"
		},
		"cloud_account": {
			"name": "new-test",
			"id": "530342348278"
		},
		"run_id": "1050436168129818625"
	}
]
`

func TestResultObjectCmd(t *testing.T) {
	mockObject := func(id string, info client.Info,
		tInfo client.TeamInfo, cInfo client.CloudAccountInfo, runId string) *client.ResultObject {
		return &client.ResultObject{
			ID:    id,
			Info:  info,
			TInfo: tInfo,
			CInfo: cInfo,
			RunID: runId,
		}
	}

	tests := []struct {
		cmds  string
		desc  string
		flags []string
		args  []string
		resp  []*client.ResultObject
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
					"a7288f05-157a-4043-ab1a-f55709457807",
					client.Info{
						SuggestedAction: "It is recommended that CMK key rotation be enabled.",
						Link:            "http://kb.cloudcoreo.com/mydoc_kms-key-rotates.html",
						Description:     "AWS Key Management Service (KMS) allows customers to rotate the backing key which is key material stored within the KMS which is tied to the key ID of the Customer Created customer master key (CMK). It is the backing key that is used to perform cryptographic operations such as encryption and decryption. Automated key rotation currently retains all prior backing keys so that decryption of encrypted data can take place transparently.",
						DisplayName:     "Verify rotation for customer created CMKs is enabled",
						Level:           "Medium",
						Service:         "kms",
						Name:            "kms-key-rotates",
						Region:          "us-east-1",
						IncludeViolationsInCount: true,
						TimeStamp:                "2018-10-11T17:21:55.111+00:00",
					},
					client.TeamInfo{
						Name: "zechen2",
						ID:   "5bb6a4956365930011a41a0b",
					},
					client.CloudAccountInfo{
						Name: "new-test",
						ID:   "530342348278",
					},
					"1050436168129818625"),
			},
			xout: kmsKeyRotatesObjectOutput,
		},
		{
			cmds:  "coreo result object",
			desc:  "Show violating objects",
			flags: []string{},
			resp: []*client.ResultObject{
				mockObject(
					"coreo-team-5b6c76cc2bc8452fe4586bce",
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
					client.TeamInfo{
						Name: "zechen2",
						ID:   "5bb6a4956365930011a41a0b",
					},
					client.CloudAccountInfo{
						Name: "new-test",
						ID:   "530342348278",
					},
					"1050436168129818625"),
			},
			xout: iamInactiveKeyNoRotationObjectOutput,
		},
	}

	var buf bytes.Buffer
	for _, test := range tests {
		frc := &fakeReleaseClient{objects: test.resp}
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
