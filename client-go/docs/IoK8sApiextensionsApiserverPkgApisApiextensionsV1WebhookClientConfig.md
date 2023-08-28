# IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaBundle** | Pointer to **string** | caBundle is a PEM encoded CA bundle which will be used to validate the webhook&#39;s server certificate. If unspecified, system trust roots on the apiserver are used. | [optional] 
**Service** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1ServiceReference**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1ServiceReference.md) |  | [optional] 
**Url** | Pointer to **string** | url gives the location of the webhook, in standard URL form (&#x60;scheme://host:port/path&#x60;). Exactly one of &#x60;url&#x60; or &#x60;service&#x60; must be specified.  The &#x60;host&#x60; should not refer to a service running in the cluster; use the &#x60;service&#x60; field instead. The host might be resolved via external DNS in some apiservers (e.g., &#x60;kube-apiserver&#x60; cannot resolve in-cluster DNS as that would be a layering violation). &#x60;host&#x60; may also be an IP address.  Please note that using &#x60;localhost&#x60; or &#x60;127.0.0.1&#x60; as a &#x60;host&#x60; is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.  The scheme must be \&quot;https\&quot;; the URL must begin with \&quot;https://\&quot;.  A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.  Attempting to use a user or basic auth e.g. \&quot;user:password@\&quot; is not allowed. Fragments (\&quot;#...\&quot;) and query parameters (\&quot;?...\&quot;) are not allowed, either. | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfigWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfigWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfigWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaBundle

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.

### GetService

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) GetService() IoK8sApiextensionsApiserverPkgApisApiextensionsV1ServiceReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) GetServiceOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1ServiceReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) SetService(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1ServiceReference)`

SetService sets Service field to given value.

### HasService

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) HasService() bool`

HasService returns a boolean if a field has been set.

### GetUrl

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


