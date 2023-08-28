# IoK8sApiAppsV1DaemonSetUpdateStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpdate** | Pointer to [**IoK8sApiAppsV1RollingUpdateDaemonSet**](IoK8sApiAppsV1RollingUpdateDaemonSet.md) |  | [optional] 
**Type** | Pointer to **string** | Type of daemon set update. Can be \&quot;RollingUpdate\&quot; or \&quot;OnDelete\&quot;. Default is RollingUpdate.  Possible enum values:  - &#x60;\&quot;OnDelete\&quot;&#x60; Replace the old daemons only when it&#39;s killed  - &#x60;\&quot;RollingUpdate\&quot;&#x60; Replace the old daemons by new ones using rolling update i.e replace them on each node one after the other. | [optional] 

## Methods

### NewIoK8sApiAppsV1DaemonSetUpdateStrategy

`func NewIoK8sApiAppsV1DaemonSetUpdateStrategy() *IoK8sApiAppsV1DaemonSetUpdateStrategy`

NewIoK8sApiAppsV1DaemonSetUpdateStrategy instantiates a new IoK8sApiAppsV1DaemonSetUpdateStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DaemonSetUpdateStrategyWithDefaults

`func NewIoK8sApiAppsV1DaemonSetUpdateStrategyWithDefaults() *IoK8sApiAppsV1DaemonSetUpdateStrategy`

NewIoK8sApiAppsV1DaemonSetUpdateStrategyWithDefaults instantiates a new IoK8sApiAppsV1DaemonSetUpdateStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpdate

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) GetRollingUpdate() IoK8sApiAppsV1RollingUpdateDaemonSet`

GetRollingUpdate returns the RollingUpdate field if non-nil, zero value otherwise.

### GetRollingUpdateOk

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) GetRollingUpdateOk() (*IoK8sApiAppsV1RollingUpdateDaemonSet, bool)`

GetRollingUpdateOk returns a tuple with the RollingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpdate

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) SetRollingUpdate(v IoK8sApiAppsV1RollingUpdateDaemonSet)`

SetRollingUpdate sets RollingUpdate field to given value.

### HasRollingUpdate

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) HasRollingUpdate() bool`

HasRollingUpdate returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiAppsV1DaemonSetUpdateStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


