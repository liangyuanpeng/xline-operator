/*
Kubernetes

Testing FlowcontrolApiserverV1beta2ApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_FlowcontrolApiserverV1beta2ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FlowcontrolApiserverV1beta2ApiService CreateFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.CreateFlowcontrolApiserverV1beta2FlowSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService CreateFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.CreateFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService DeleteFlowcontrolApiserverV1beta2CollectionFlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.DeleteFlowcontrolApiserverV1beta2CollectionFlowSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService DeleteFlowcontrolApiserverV1beta2CollectionPriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.DeleteFlowcontrolApiserverV1beta2CollectionPriorityLevelConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService DeleteFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.DeleteFlowcontrolApiserverV1beta2FlowSchema(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService DeleteFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.DeleteFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService GetFlowcontrolApiserverV1beta2APIResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.GetFlowcontrolApiserverV1beta2APIResources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ListFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ListFlowcontrolApiserverV1beta2FlowSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ListFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ListFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService PatchFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.PatchFlowcontrolApiserverV1beta2FlowSchema(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService PatchFlowcontrolApiserverV1beta2FlowSchemaStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.PatchFlowcontrolApiserverV1beta2FlowSchemaStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService PatchFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.PatchFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService PatchFlowcontrolApiserverV1beta2PriorityLevelConfigurationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.PatchFlowcontrolApiserverV1beta2PriorityLevelConfigurationStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReadFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReadFlowcontrolApiserverV1beta2FlowSchema(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReadFlowcontrolApiserverV1beta2FlowSchemaStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReadFlowcontrolApiserverV1beta2FlowSchemaStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReadFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReadFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReadFlowcontrolApiserverV1beta2PriorityLevelConfigurationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReadFlowcontrolApiserverV1beta2PriorityLevelConfigurationStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReplaceFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReplaceFlowcontrolApiserverV1beta2FlowSchema(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReplaceFlowcontrolApiserverV1beta2FlowSchemaStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReplaceFlowcontrolApiserverV1beta2FlowSchemaStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReplaceFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReplaceFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService ReplaceFlowcontrolApiserverV1beta2PriorityLevelConfigurationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.ReplaceFlowcontrolApiserverV1beta2PriorityLevelConfigurationStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService WatchFlowcontrolApiserverV1beta2FlowSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.WatchFlowcontrolApiserverV1beta2FlowSchema(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService WatchFlowcontrolApiserverV1beta2FlowSchemaList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.WatchFlowcontrolApiserverV1beta2FlowSchemaList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService WatchFlowcontrolApiserverV1beta2PriorityLevelConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.WatchFlowcontrolApiserverV1beta2PriorityLevelConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlowcontrolApiserverV1beta2ApiService WatchFlowcontrolApiserverV1beta2PriorityLevelConfigurationList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlowcontrolApiserverV1beta2Api.WatchFlowcontrolApiserverV1beta2PriorityLevelConfigurationList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
