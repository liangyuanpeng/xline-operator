# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** | strategy specifies how custom resources are converted between versions. Allowed values are: - &#x60;\&quot;None\&quot;&#x60;: The converter only change the apiVersion and would not touch any other field in the custom resource. - &#x60;\&quot;Webhook\&quot;&#x60;: API Server will call to an external webhook to do the conversion. Additional information   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set. | 
**Webhook** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion.md) |  | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion(strategy string, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversionWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversionWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversionWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetWebhook

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) GetWebhook() IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) GetWebhookOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) SetWebhook(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


