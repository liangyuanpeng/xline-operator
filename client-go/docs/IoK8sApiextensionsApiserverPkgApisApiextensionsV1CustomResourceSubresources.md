# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scale** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** | CustomResourceSubresourceStatus defines how to serve the status subresource for CustomResources. Status is represented by the &#x60;.status&#x60; JSON path inside of a CustomResource. When set, * exposes a /status subresource for the custom resource * PUT requests to the /status subresource take a custom resource object, and ignore changes to anything except the status stanza * PUT/POST/PATCH requests to the custom resource ignore changes to the status stanza | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourcesWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourcesWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourcesWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScale

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) GetScale() IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) GetScaleOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) SetScale(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale)`

SetScale sets Scale field to given value.

### HasScale

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetStatus

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresources) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


