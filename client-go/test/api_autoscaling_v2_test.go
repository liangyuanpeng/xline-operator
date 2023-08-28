/*
Kubernetes

Testing AutoscalingV2ApiService

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

func Test_openapi_AutoscalingV2ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AutoscalingV2ApiService CreateAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.CreateAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService DeleteAutoscalingV2CollectionNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.DeleteAutoscalingV2CollectionNamespacedHorizontalPodAutoscaler(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService DeleteAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.DeleteAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService GetAutoscalingV2APIResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AutoscalingV2Api.GetAutoscalingV2APIResources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService ListAutoscalingV2HorizontalPodAutoscalerForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AutoscalingV2Api.ListAutoscalingV2HorizontalPodAutoscalerForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService ListAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.ListAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService PatchAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.PatchAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService PatchAutoscalingV2NamespacedHorizontalPodAutoscalerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.PatchAutoscalingV2NamespacedHorizontalPodAutoscalerStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService ReadAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.ReadAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService ReadAutoscalingV2NamespacedHorizontalPodAutoscalerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.ReadAutoscalingV2NamespacedHorizontalPodAutoscalerStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService ReplaceAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.ReplaceAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService ReplaceAutoscalingV2NamespacedHorizontalPodAutoscalerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.ReplaceAutoscalingV2NamespacedHorizontalPodAutoscalerStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService WatchAutoscalingV2HorizontalPodAutoscalerListForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AutoscalingV2Api.WatchAutoscalingV2HorizontalPodAutoscalerListForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService WatchAutoscalingV2NamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.WatchAutoscalingV2NamespacedHorizontalPodAutoscaler(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutoscalingV2ApiService WatchAutoscalingV2NamespacedHorizontalPodAutoscalerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.AutoscalingV2Api.WatchAutoscalingV2NamespacedHorizontalPodAutoscalerList(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
