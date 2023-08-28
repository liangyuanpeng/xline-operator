# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelectorPath** | Pointer to **string** | labelSelectorPath defines the JSON path inside of a custom resource that corresponds to Scale &#x60;status.selector&#x60;. Only JSON paths without the array notation are allowed. Must be a JSON Path under &#x60;.status&#x60; or &#x60;.spec&#x60;. Must be set to work with HorizontalPodAutoscaler. The field pointed by this JSON path must be a string field (not a complex selector struct) which contains a serialized label selector in string form. More info: https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions#scale-subresource If there is no value under the given path in the custom resource, the &#x60;status.selector&#x60; value in the &#x60;/scale&#x60; subresource will default to the empty string. | [optional] 
**SpecReplicasPath** | **string** | specReplicasPath defines the JSON path inside of a custom resource that corresponds to Scale &#x60;spec.replicas&#x60;. Only JSON paths without the array notation are allowed. Must be a JSON Path under &#x60;.spec&#x60;. If there is no value under the given path in the custom resource, the &#x60;/scale&#x60; subresource will return an error on GET. | 
**StatusReplicasPath** | **string** | statusReplicasPath defines the JSON path inside of a custom resource that corresponds to Scale &#x60;status.replicas&#x60;. Only JSON paths without the array notation are allowed. Must be a JSON Path under &#x60;.status&#x60;. If there is no value under the given path in the custom resource, the &#x60;status.replicas&#x60; value in the &#x60;/scale&#x60; subresource will default to 0. | 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale(specReplicasPath string, statusReplicasPath string, ) *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScaleWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScaleWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScaleWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSelectorPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) GetLabelSelectorPath() string`

GetLabelSelectorPath returns the LabelSelectorPath field if non-nil, zero value otherwise.

### GetLabelSelectorPathOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) GetLabelSelectorPathOk() (*string, bool)`

GetLabelSelectorPathOk returns a tuple with the LabelSelectorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelectorPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) SetLabelSelectorPath(v string)`

SetLabelSelectorPath sets LabelSelectorPath field to given value.

### HasLabelSelectorPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) HasLabelSelectorPath() bool`

HasLabelSelectorPath returns a boolean if a field has been set.

### GetSpecReplicasPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) GetSpecReplicasPath() string`

GetSpecReplicasPath returns the SpecReplicasPath field if non-nil, zero value otherwise.

### GetSpecReplicasPathOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) GetSpecReplicasPathOk() (*string, bool)`

GetSpecReplicasPathOk returns a tuple with the SpecReplicasPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecReplicasPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) SetSpecReplicasPath(v string)`

SetSpecReplicasPath sets SpecReplicasPath field to given value.


### GetStatusReplicasPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) GetStatusReplicasPath() string`

GetStatusReplicasPath returns the StatusReplicasPath field if non-nil, zero value otherwise.

### GetStatusReplicasPathOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) GetStatusReplicasPathOk() (*string, bool)`

GetStatusReplicasPathOk returns a tuple with the StatusReplicasPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReplicasPath

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceSubresourceScale) SetStatusReplicasPath(v string)`

SetStatusReplicasPath sets StatusReplicasPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


