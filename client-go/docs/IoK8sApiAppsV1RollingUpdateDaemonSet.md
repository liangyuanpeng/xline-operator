# IoK8sApiAppsV1RollingUpdateDaemonSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSurge** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 
**MaxUnavailable** | Pointer to **string** | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [optional] 

## Methods

### NewIoK8sApiAppsV1RollingUpdateDaemonSet

`func NewIoK8sApiAppsV1RollingUpdateDaemonSet() *IoK8sApiAppsV1RollingUpdateDaemonSet`

NewIoK8sApiAppsV1RollingUpdateDaemonSet instantiates a new IoK8sApiAppsV1RollingUpdateDaemonSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1RollingUpdateDaemonSetWithDefaults

`func NewIoK8sApiAppsV1RollingUpdateDaemonSetWithDefaults() *IoK8sApiAppsV1RollingUpdateDaemonSet`

NewIoK8sApiAppsV1RollingUpdateDaemonSetWithDefaults instantiates a new IoK8sApiAppsV1RollingUpdateDaemonSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSurge

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *IoK8sApiAppsV1RollingUpdateDaemonSet) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


