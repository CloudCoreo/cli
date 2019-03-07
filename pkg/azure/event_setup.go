package azure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/CloudCoreo/cli/client"
)

type SetupService struct {
	AuthFile string
}

func NewSetupService(input *NewServiceInput) *SetupService {
	return &SetupService{AuthFile: input.AuthFile}
}
func (a *SetupService) SetupEventStream(input *client.EventStreamConfig) error {
	ctx := context.Background()
	err := a.createResourceGroup(ctx, input)
	if err != nil {
		return err
	}

	err = a.deployActionGroup(ctx, input)
	if err != nil {
		return err
	}

	err = a.deployAlert(ctx, input)

	return a.sendSuccessEvent(input)
}

func (a *SetupService) createResourceGroup(ctx context.Context, input *client.EventStreamConfig) error {
	groupsClient, err := a.getGroupsClient(input)
	if err != nil {
		return err
	}
	_, err = groupsClient.CreateOrUpdate(ctx, input.ResourceGroup, resources.Group{Location: to.StringPtr("eastus")})
	return err
}

func (a *SetupService) getGroupsClient(input *client.EventStreamConfig) (*resources.GroupsClient, error) {
	groupsClient := resources.NewGroupsClient(input.SubscriptionID)
	if a.AuthFile != "" {
		au, err := auth.NewAuthorizerFromFile(a.AuthFile)
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

func (a *SetupService) getDeploymentsClient(input *client.EventStreamConfig) (*resources.DeploymentsClient, error) {
	deploymentsClient := resources.NewDeploymentsClient(input.SubscriptionID)
	if a.AuthFile != "" {
		au, err := auth.NewAuthorizerFromFile(a.AuthFile)
		if err != nil {
			return nil, err
		}
		deploymentsClient.Authorizer = au
	} else {
		au, err := auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return nil, err
		}
		deploymentsClient.Authorizer = au
	}
	return &deploymentsClient, nil
}

func (a *SetupService) deployActionGroup(ctx context.Context, input *client.EventStreamConfig) error {
	deploymentsClient, err := a.getDeploymentsClient(input)
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"actionGroupName":      map[string]interface{}{"value": input.ActionGroup},
		"actionGroupShortName": map[string]interface{}{"value": input.ActionGroupShort},
		"webhookReceiverName":  map[string]interface{}{"value": input.WebhookReceiverName},
		"webhookServiceUri":    map[string]interface{}{"value": input.WebhookServiceUri},
	}
	_, err = deploymentsClient.CreateOrUpdate(
		ctx,
		input.ResourceGroup,
		input.ActionDeploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   readJSON(input.ActionDeployFile),
				Parameters: &params,
				Mode:       "Incremental",
			},
		})
	return err
}

func (a *SetupService) deployAlert(ctx context.Context, input *client.EventStreamConfig) error {
	deploymentsClient, err := a.getDeploymentsClient(input)
	if err != nil {
		return err
	}
	actionGroupResourceID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/actionGroups/%s", input.SubscriptionID, input.ResourceGroup, input.ActionGroup)
	params := map[string]interface{}{
		"activityLogAlertName":  map[string]interface{}{"value": input.AlertName},
		"actionGroupResourceId": map[string]interface{}{"value": actionGroupResourceID},
	}
	_, err = deploymentsClient.CreateOrUpdate(
		ctx,
		input.ResourceGroup,
		input.AlertDeploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template:   readJSON(input.AlertDeployFile),
				Parameters: &params,
				Mode:       "Incremental",
			},
		},
	)
	return err
}

func (a *SetupService) sendSuccessEvent(input *client.EventStreamConfig) error {
	//No additional whitespace is allowed in the below string, other with the http request may fail
	//TODO: Discuss to see whether it needs to be a struct
	data := fmt.Sprintf("{\"data\": {\"context\": {\"activityLog\": {\"subscriptionId\": \"%s\", \"operationName\": \"AzureStreamReady\"}}}}", input.SubscriptionID)
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

func readJSON(s string) *map[string]interface{} {
	contents := make(map[string]interface{})
	json.Unmarshal([]byte(s), &contents)
	return &contents
}
