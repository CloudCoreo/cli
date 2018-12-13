package aws

import (
	"testing"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
)

const rolePolicy = `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Principal": {
				"AWS": "arn:aws:iam::` + "accountID" + `:root"
			},
			"Action": "sts:AssumeRole",
			"Condition": {
				"StringEquals": {
					"sts:ExternalId": "` + "externalID" + `"
				}
			}
		}
	]
}`

func TestCreateAssumeRolePolicyDocument(t *testing.T) {
	ts := httpstub.New()
	defer ts.Close()

	client := NewRoleService(&NewServiceInput{})
	res := client.createAssumeRolePolicyDocument("accountID", "externalID")
	assert.Equal(t, rolePolicy, res)
}
