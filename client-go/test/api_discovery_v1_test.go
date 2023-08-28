/*
Kubernetes

Testing DiscoveryV1ApiService

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

func Test_openapi_DiscoveryV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DiscoveryV1ApiService CreateDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.CreateDiscoveryV1NamespacedEndpointSlice(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService DeleteDiscoveryV1CollectionNamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.DeleteDiscoveryV1CollectionNamespacedEndpointSlice(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService DeleteDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.DeleteDiscoveryV1NamespacedEndpointSlice(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService GetDiscoveryV1APIResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DiscoveryV1Api.GetDiscoveryV1APIResources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService ListDiscoveryV1EndpointSliceForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DiscoveryV1Api.ListDiscoveryV1EndpointSliceForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService ListDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.ListDiscoveryV1NamespacedEndpointSlice(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService PatchDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.PatchDiscoveryV1NamespacedEndpointSlice(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService ReadDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.ReadDiscoveryV1NamespacedEndpointSlice(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService ReplaceDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.ReplaceDiscoveryV1NamespacedEndpointSlice(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService WatchDiscoveryV1EndpointSliceListForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DiscoveryV1Api.WatchDiscoveryV1EndpointSliceListForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService WatchDiscoveryV1NamespacedEndpointSlice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.WatchDiscoveryV1NamespacedEndpointSlice(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryV1ApiService WatchDiscoveryV1NamespacedEndpointSliceList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DiscoveryV1Api.WatchDiscoveryV1NamespacedEndpointSliceList(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
