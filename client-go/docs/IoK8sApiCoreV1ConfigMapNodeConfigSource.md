# IoK8sApiCoreV1ConfigMapNodeConfigSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeletConfigKey** | **string** | KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases. | 
**Name** | **string** | Name is the metadata.name of the referenced ConfigMap. This field is required in all cases. | 
**Namespace** | **string** | Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases. | 
**ResourceVersion** | Pointer to **string** | ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status. | [optional] 
**Uid** | Pointer to **string** | UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status. | [optional] 

## Methods

### NewIoK8sApiCoreV1ConfigMapNodeConfigSource

`func NewIoK8sApiCoreV1ConfigMapNodeConfigSource(kubeletConfigKey string, name string, namespace string, ) *IoK8sApiCoreV1ConfigMapNodeConfigSource`

NewIoK8sApiCoreV1ConfigMapNodeConfigSource instantiates a new IoK8sApiCoreV1ConfigMapNodeConfigSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ConfigMapNodeConfigSourceWithDefaults

`func NewIoK8sApiCoreV1ConfigMapNodeConfigSourceWithDefaults() *IoK8sApiCoreV1ConfigMapNodeConfigSource`

NewIoK8sApiCoreV1ConfigMapNodeConfigSourceWithDefaults instantiates a new IoK8sApiCoreV1ConfigMapNodeConfigSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeletConfigKey

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetKubeletConfigKey() string`

GetKubeletConfigKey returns the KubeletConfigKey field if non-nil, zero value otherwise.

### GetKubeletConfigKeyOk

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetKubeletConfigKeyOk() (*string, bool)`

GetKubeletConfigKeyOk returns a tuple with the KubeletConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletConfigKey

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) SetKubeletConfigKey(v string)`

SetKubeletConfigKey sets KubeletConfigKey field to given value.


### GetName

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetResourceVersion

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetUid

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApiCoreV1ConfigMapNodeConfigSource) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


