# IoK8sApiCoreV1AzureFilePersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnly** | Pointer to **bool** | readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**SecretName** | **string** | secretName is the name of secret that contains Azure Storage Account Name and Key | 
**SecretNamespace** | Pointer to **string** | secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod | [optional] 
**ShareName** | **string** | shareName is the azure Share Name | 

## Methods

### NewIoK8sApiCoreV1AzureFilePersistentVolumeSource

`func NewIoK8sApiCoreV1AzureFilePersistentVolumeSource(secretName string, shareName string, ) *IoK8sApiCoreV1AzureFilePersistentVolumeSource`

NewIoK8sApiCoreV1AzureFilePersistentVolumeSource instantiates a new IoK8sApiCoreV1AzureFilePersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1AzureFilePersistentVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1AzureFilePersistentVolumeSourceWithDefaults() *IoK8sApiCoreV1AzureFilePersistentVolumeSource`

NewIoK8sApiCoreV1AzureFilePersistentVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1AzureFilePersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretName

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetSecretNamespace

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetSecretNamespace() string`

GetSecretNamespace returns the SecretNamespace field if non-nil, zero value otherwise.

### GetSecretNamespaceOk

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetSecretNamespaceOk() (*string, bool)`

GetSecretNamespaceOk returns a tuple with the SecretNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNamespace

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) SetSecretNamespace(v string)`

SetSecretNamespace sets SecretNamespace field to given value.

### HasSecretNamespace

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) HasSecretNamespace() bool`

HasSecretNamespace returns a boolean if a field has been set.

### GetShareName

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *IoK8sApiCoreV1AzureFilePersistentVolumeSource) SetShareName(v string)`

SetShareName sets ShareName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


