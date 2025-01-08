package testimpl

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	operationalinsights "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID is not set in the environment variables ")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	}

	resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	workspaceName := terraform.Output(t, ctx.TerratestTerraformOptions(), "workspace_name")
	workspaceId := terraform.Output(t, ctx.TerratestTerraformOptions(), "workspace_id")
	storageInsightsId := terraform.Output(t, ctx.TerratestTerraformOptions(), "storage_insights_id")
	storageInsightsName := terraform.Output(t, ctx.TerratestTerraformOptions(), "storage_insights_name")

	logAnalyticsWorkspaceClient, err := operationalinsights.NewWorkspacesClient(subscriptionID, credential, &options)
	if err != nil {
		t.Fatalf("Error creating log analytics workspace client: %v", err)
	}

	storageInsightsClient, err := operationalinsights.NewStorageInsightConfigsClient(subscriptionID, credential, &options)
	if err != nil {
		t.Fatalf("Error creating storage insights client: %v", err)
	}

	t.Run("doesLogAnalyticsWorkspaceExist", func(t *testing.T) {
		workspace, err := logAnalyticsWorkspaceClient.Get(context.Background(), resourceGroupName, workspaceName, nil)
		if err != nil {
			t.Fatalf("Error getting log analytics workspace: %v", err)
		}
		assert.Equal(t, workspaceId, *workspace.ID, "Expected workspace ID to be %s, but got %s", workspaceId, *workspace.ID)
	})

	t.Run("doesStorageInsightsExist", func(t *testing.T) {
		resp, err := storageInsightsClient.Get(context.Background(), resourceGroupName, workspaceName, storageInsightsName, nil)
		if err != nil {
			t.Fatalf("Error getting storage insights: %v", err)
		}
		// The API response downcases part of the ID, so we will downcase both entire IDs for comparison, but leave them unchanged for the assertion message
		assert.Equal(t, strings.ToLower(storageInsightsId), strings.ToLower(*resp.ID), "Expected storage insights ID to be %s, but got %s", storageInsightsId, *resp.ID)
	})
}
