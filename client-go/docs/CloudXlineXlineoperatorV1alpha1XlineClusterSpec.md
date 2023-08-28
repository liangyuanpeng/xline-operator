# CloudXlineXlineoperatorV1alpha1XlineClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **map[string]interface{}** | Backup specification | [optional] 
**Container** | [**CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer.md) |  | 
**Data** | Pointer to **map[string]interface{}** | The data PVC, if it is not specified, then use emptyDir instead | [optional] 
**Pvcs** | Pointer to **map[string]interface{}** | Some user defined persistent volume claim templates | [optional] 
**Size** | **int32** | Size of the xline cluster, less than 3 is not allowed | 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpec

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpec(container CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer, size int32, ) *CloudXlineXlineoperatorV1alpha1XlineClusterSpec`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpec instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpec`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetBackup() map[string]interface{}`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetBackupOk() (*map[string]interface{}, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetBackup(v map[string]interface{})`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetContainer

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetContainer() CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetContainerOk() (*CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetContainer(v CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainer)`

SetContainer sets Container field to given value.


### GetData

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPvcs

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetPvcs() map[string]interface{}`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetPvcsOk() (*map[string]interface{}, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetPvcs(v map[string]interface{})`

SetPvcs sets Pvcs field to given value.

### HasPvcs

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) HasPvcs() bool`

HasPvcs returns a boolean if a field has been set.

### GetSize

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpec) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


