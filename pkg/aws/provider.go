package aws

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sts"
)

// AssumeRoleCredentialsProvider ...
type AssumeRoleCredentialsProvider struct {
	AssumeRoleCredentials *sts.Credentials
}

// NewAssumeRoleCredentialsProvider ...
func NewAssumeRoleCredentialsProvider(credentials *sts.Credentials) *AssumeRoleCredentialsProvider {
	return &AssumeRoleCredentialsProvider{
		AssumeRoleCredentials: credentials,
	}
}

// Retrieve ...
func (c AssumeRoleCredentialsProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     *c.AssumeRoleCredentials.AccessKeyId,
		SecretAccessKey: *c.AssumeRoleCredentials.SecretAccessKey,
		SessionToken:    *c.AssumeRoleCredentials.SessionToken,
		ProviderName:    "AssumeRoleCredentialsProvider",
	}, nil

}

// IsExpired ...
func (c AssumeRoleCredentialsProvider) IsExpired() bool {
	return c.AssumeRoleCredentials.Expiration.After(time.Now())

}
