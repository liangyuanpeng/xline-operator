# IoK8sApiAppsV1RollingUpdateStatefulSetStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUnavailable** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 
**Partition** | Pointer to **int32** | Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0. | [optional] 

## Methods

### NewIoK8sApiAppsV1RollingUpdateStatefulSetStrategy

`func NewIoK8sApiAppsV1RollingUpdateStatefulSetStrategy() *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy`

NewIoK8sApiAppsV1RollingUpdateStatefulSetStrategy instantiates a new IoK8sApiAppsV1RollingUpdateStatefulSetStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1RollingUpdateStatefulSetStrategyWithDefaults

`func NewIoK8sApiAppsV1RollingUpdateStatefulSetStrategyWithDefaults() *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy`

NewIoK8sApiAppsV1RollingUpdateStatefulSetStrategyWithDefaults instantiates a new IoK8sApiAppsV1RollingUpdateStatefulSetStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUnavailable

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetPartition

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *IoK8sApiAppsV1RollingUpdateStatefulSetStrategy) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


