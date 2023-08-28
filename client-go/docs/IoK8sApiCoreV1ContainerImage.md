# IoK8sApiCoreV1ContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **[]string** | Names by which this image is known. e.g. [\&quot;kubernetes.example/hyperkube:v1.0.7\&quot;, \&quot;cloud-vendor.registry.example/cloud-vendor/hyperkube:v1.0.7\&quot;] | [optional] 
**SizeBytes** | Pointer to **int64** | The size of the image in bytes. | [optional] 

## Methods

### NewIoK8sApiCoreV1ContainerImage

`func NewIoK8sApiCoreV1ContainerImage() *IoK8sApiCoreV1ContainerImage`

NewIoK8sApiCoreV1ContainerImage instantiates a new IoK8sApiCoreV1ContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ContainerImageWithDefaults

`func NewIoK8sApiCoreV1ContainerImageWithDefaults() *IoK8sApiCoreV1ContainerImage`

NewIoK8sApiCoreV1ContainerImageWithDefaults instantiates a new IoK8sApiCoreV1ContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *IoK8sApiCoreV1ContainerImage) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *IoK8sApiCoreV1ContainerImage) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *IoK8sApiCoreV1ContainerImage) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *IoK8sApiCoreV1ContainerImage) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSizeBytes

`func (o *IoK8sApiCoreV1ContainerImage) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *IoK8sApiCoreV1ContainerImage) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *IoK8sApiCoreV1ContainerImage) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *IoK8sApiCoreV1ContainerImage) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


