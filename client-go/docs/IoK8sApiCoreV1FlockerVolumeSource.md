# IoK8sApiCoreV1FlockerVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | Pointer to **string** | datasetName is Name of the dataset stored as metadata -&gt; name on the dataset for Flocker should be considered as deprecated | [optional] 
**DatasetUUID** | Pointer to **string** | datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset | [optional] 

## Methods

### NewIoK8sApiCoreV1FlockerVolumeSource

`func NewIoK8sApiCoreV1FlockerVolumeSource() *IoK8sApiCoreV1FlockerVolumeSource`

NewIoK8sApiCoreV1FlockerVolumeSource instantiates a new IoK8sApiCoreV1FlockerVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1FlockerVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1FlockerVolumeSourceWithDefaults() *IoK8sApiCoreV1FlockerVolumeSource`

NewIoK8sApiCoreV1FlockerVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1FlockerVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *IoK8sApiCoreV1FlockerVolumeSource) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *IoK8sApiCoreV1FlockerVolumeSource) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *IoK8sApiCoreV1FlockerVolumeSource) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *IoK8sApiCoreV1FlockerVolumeSource) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetDatasetUUID

`func (o *IoK8sApiCoreV1FlockerVolumeSource) GetDatasetUUID() string`

GetDatasetUUID returns the DatasetUUID field if non-nil, zero value otherwise.

### GetDatasetUUIDOk

`func (o *IoK8sApiCoreV1FlockerVolumeSource) GetDatasetUUIDOk() (*string, bool)`

GetDatasetUUIDOk returns a tuple with the DatasetUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetUUID

`func (o *IoK8sApiCoreV1FlockerVolumeSource) SetDatasetUUID(v string)`

SetDatasetUUID sets DatasetUUID field to given value.

### HasDatasetUUID

`func (o *IoK8sApiCoreV1FlockerVolumeSource) HasDatasetUUID() bool`

HasDatasetUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


