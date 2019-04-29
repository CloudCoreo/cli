package azure

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/CloudCoreo/cli/client"
)

type RemoveService struct {
	authFile string
	region   string
}

// NewRemoveService returns an instance of RemoveService
func NewRemoveService(input *NewServiceInput) *RemoveService {
	return &RemoveService{
		authFile: input.AuthFile,
		region:   input.Region,
	}
}

func (a *RemoveService) RemoveEventStream(input *client.EventRemoveConfig) error {
	ctx := context.Background()
	err := a.removeResourceGroup(ctx, input)
	if err != nil {
		return err
	}
	return a.sendRemoveEvent(input)
}

func (a *RemoveService) removeResourceGroup(ctx context.Context, input *client.EventRemoveConfig) error {
	groupsClient, err := a.getGroupsClient(input)
	if err != nil {
		return err
	}
	_, err = groupsClient.Delete(ctx, input.ResourceGroup)
	return err
}

func (a *RemoveService) getGroupsClient(input *client.EventRemoveConfig) (*resources.GroupsClient, error) {
	groupsClient := resources.NewGroupsClient(input.SubscriptionID)
	if a.authFile != "" {
		au, err := auth.NewAuthorizerFromFile(a.authFile)
		if err != nil {
			return nil, err
		}
		groupsClient.Authorizer = au
	} else {
		au, err := auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return nil, err
		}
		groupsClient.Authorizer = au
	}
	return &groupsClient, nil
}

func (a *RemoveService) sendRemoveEvent(input *client.EventRemoveConfig) error {
	fmt.Println("Sending Event Removal message")
	data := fmt.Sprintf("{\"data\": {\"context\": {\"activityLog\": {\"subscriptionId\": \"%s\", \"operationName\": \"AzureStreamNotReady\"}}}}", input.SubscriptionID)
	req, err := http.NewRequest("POST", input.WebhookServiceUri, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return err
}
