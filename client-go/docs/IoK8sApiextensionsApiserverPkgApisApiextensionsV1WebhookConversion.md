# IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientConfig** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig.md) |  | [optional] 
**ConversionReviewVersions** | **[]string** | conversionReviewVersions is an ordered list of preferred &#x60;ConversionReview&#x60; versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail. | 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion(conversionReviewVersions []string, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversionWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversionWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversionWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientConfig

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) GetClientConfig() IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig`

GetClientConfig returns the ClientConfig field if non-nil, zero value otherwise.

### GetClientConfigOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) GetClientConfigOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig, bool)`

GetClientConfigOk returns a tuple with the ClientConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConfig

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) SetClientConfig(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookClientConfig)`

SetClientConfig sets ClientConfig field to given value.

### HasClientConfig

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) HasClientConfig() bool`

HasClientConfig returns a boolean if a field has been set.

### GetConversionReviewVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) GetConversionReviewVersions() []string`

GetConversionReviewVersions returns the ConversionReviewVersions field if non-nil, zero value otherwise.

### GetConversionReviewVersionsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) GetConversionReviewVersionsOk() (*[]string, bool)`

GetConversionReviewVersionsOk returns a tuple with the ConversionReviewVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionReviewVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion) SetConversionReviewVersions(v []string)`

SetConversionReviewVersions sets ConversionReviewVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


