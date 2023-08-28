/*
Kubernetes

Testing NetworkingV1ApiService

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

func Test_openapi_NetworkingV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkingV1ApiService CreateNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.CreateNetworkingV1IngressClass(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService CreateNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.CreateNetworkingV1NamespacedIngress(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService CreateNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.CreateNetworkingV1NamespacedNetworkPolicy(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService DeleteNetworkingV1CollectionIngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.DeleteNetworkingV1CollectionIngressClass(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService DeleteNetworkingV1CollectionNamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.DeleteNetworkingV1CollectionNamespacedIngress(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService DeleteNetworkingV1CollectionNamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.DeleteNetworkingV1CollectionNamespacedNetworkPolicy(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService DeleteNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.NetworkingV1Api.DeleteNetworkingV1IngressClass(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService DeleteNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.DeleteNetworkingV1NamespacedIngress(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService DeleteNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.DeleteNetworkingV1NamespacedNetworkPolicy(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService GetNetworkingV1APIResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.GetNetworkingV1APIResources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ListNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.ListNetworkingV1IngressClass(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ListNetworkingV1IngressForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.ListNetworkingV1IngressForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ListNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ListNetworkingV1NamespacedIngress(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ListNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ListNetworkingV1NamespacedNetworkPolicy(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ListNetworkingV1NetworkPolicyForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.ListNetworkingV1NetworkPolicyForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService PatchNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.NetworkingV1Api.PatchNetworkingV1IngressClass(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService PatchNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.PatchNetworkingV1NamespacedIngress(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService PatchNetworkingV1NamespacedIngressStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.PatchNetworkingV1NamespacedIngressStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService PatchNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.PatchNetworkingV1NamespacedNetworkPolicy(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService PatchNetworkingV1NamespacedNetworkPolicyStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.PatchNetworkingV1NamespacedNetworkPolicyStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReadNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReadNetworkingV1IngressClass(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReadNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReadNetworkingV1NamespacedIngress(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReadNetworkingV1NamespacedIngressStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReadNetworkingV1NamespacedIngressStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReadNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReadNetworkingV1NamespacedNetworkPolicy(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReadNetworkingV1NamespacedNetworkPolicyStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReadNetworkingV1NamespacedNetworkPolicyStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReplaceNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReplaceNetworkingV1IngressClass(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReplaceNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReplaceNetworkingV1NamespacedIngress(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReplaceNetworkingV1NamespacedIngressStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReplaceNetworkingV1NamespacedIngressStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReplaceNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReplaceNetworkingV1NamespacedNetworkPolicy(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService ReplaceNetworkingV1NamespacedNetworkPolicyStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.ReplaceNetworkingV1NamespacedNetworkPolicyStatus(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1IngressClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1IngressClass(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1IngressClassList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1IngressClassList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1IngressListForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1IngressListForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1NamespacedIngress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1NamespacedIngress(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1NamespacedIngressList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1NamespacedIngressList(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1NamespacedNetworkPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1NamespacedNetworkPolicy(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1NamespacedNetworkPolicyList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1NamespacedNetworkPolicyList(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkingV1ApiService WatchNetworkingV1NetworkPolicyListForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkingV1Api.WatchNetworkingV1NetworkPolicyListForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
