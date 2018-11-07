package client

import (
	"net/http"
	"testing"

	"github.com/CloudCoreo/cli/cmd/content"

	"github.com/jharlap/httpstub"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
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
]`

const kmsKeyRotatesObjectOutput = `[{
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
	}]`

func TestGetAllResultRuleSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/rule").WithMethod("GET").WithBody(iamInactiveKeyNoRotationRuleOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultRule(context.Background())
	assert.Nil(t, err, "GetAllResultRule shouldn't return error")
}

func TestGetAllResultObjectSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/object").WithMethod("GET").WithBody(kmsKeyRotatesObjectOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultObject(context.Background())
	assert.Nil(t, err, "GetAllResultObject shouldn't return error")
}

func TestGetAllResultRuleFailureNoViolatedRule(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/rule").WithMethod("GET").WithBody("[]").WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultRule(context.Background())
	assert.NotNil(t, err, "GetAllResultRule should return error.")
	assert.Equal(t, "No violated rule", err.Error())
}

func TestGetAllResultRuleFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/rule").WithMethod("GET").WithBody(iamInactiveKeyNoRotationRuleOutput).WithStatus(http.StatusBadRequest)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultRule(context.Background())
	assert.NotNil(t, err, "GetAllResultRule should return error.")
}

func TestGetAllResultObjectFailureBadRequest(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/object").WithMethod("GET").WithBody(kmsKeyRotatesObjectOutput).WithStatus(http.StatusBadRequest)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.getAllResultObject(context.Background())
	assert.NotNil(t, err, "GetAllResultObject should return error.")
}

func TestGetResultRuleSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/rule").WithMethod("GET").WithBody(iamInactiveKeyNoRotationRuleOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultRule(context.Background(), "5bb6a4956365930011a41a0b", "530342348278", "Medium")
	assert.Nil(t, err, "GetResultRule shouldn't return error")
}

func TestShowResultObjectSuccess(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/object").WithMethod("GET").WithBody(kmsKeyRotatesObjectOutput).WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultObject(context.Background(), content.None, content.None, content.None)
	assert.Nil(t, err, "GetResultObject shouldn't return error")
}

func TestGetResultRuleFailureNoViolatedRule(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/rule").WithMethod("GET").WithBody("[]").WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultRule(context.Background(), "5bb6a4956365930011a41a0b", "530342348278", "Medium")
	assert.NotNil(t, err, "GetResultRule should return error.")
	assert.Equal(t, "No violated rule", err.Error())
}

func TestShowResultObjectFailureNoViolatedObject(t *testing.T) {
	ts := httpstub.New()
	ts.Path("/result/object").WithMethod("GET").WithBody("[]").WithStatus(http.StatusOK)
	defer ts.Close()

	client, _ := MakeClient("ApiKey", "SecretKey", ts.URL)
	_, err := client.ShowResultObject(context.Background(), content.None, content.None, content.None)
	assert.NotNil(t, err, "GetResultObject should return error.")
	assert.Equal(t, "No violated object", err.Error())
}
