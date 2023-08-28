# IoK8sApiAppsV1DeploymentStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpdate** | Pointer to [**IoK8sApiAppsV1RollingUpdateDeployment**](IoK8sApiAppsV1RollingUpdateDeployment.md) |  | [optional] 
**Type** | Pointer to **string** | Type of deployment. Can be \&quot;Recreate\&quot; or \&quot;RollingUpdate\&quot;. Default is RollingUpdate.  Possible enum values:  - &#x60;\&quot;Recreate\&quot;&#x60; Kill all existing pods before creating new ones.  - &#x60;\&quot;RollingUpdate\&quot;&#x60; Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one. | [optional] 

## Methods

### NewIoK8sApiAppsV1DeploymentStrategy

`func NewIoK8sApiAppsV1DeploymentStrategy() *IoK8sApiAppsV1DeploymentStrategy`

NewIoK8sApiAppsV1DeploymentStrategy instantiates a new IoK8sApiAppsV1DeploymentStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAppsV1DeploymentStrategyWithDefaults

`func NewIoK8sApiAppsV1DeploymentStrategyWithDefaults() *IoK8sApiAppsV1DeploymentStrategy`

NewIoK8sApiAppsV1DeploymentStrategyWithDefaults instantiates a new IoK8sApiAppsV1DeploymentStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpdate

`func (o *IoK8sApiAppsV1DeploymentStrategy) GetRollingUpdate() IoK8sApiAppsV1RollingUpdateDeployment`

GetRollingUpdate returns the RollingUpdate field if non-nil, zero value otherwise.

### GetRollingUpdateOk

`func (o *IoK8sApiAppsV1DeploymentStrategy) GetRollingUpdateOk() (*IoK8sApiAppsV1RollingUpdateDeployment, bool)`

GetRollingUpdateOk returns a tuple with the RollingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpdate

`func (o *IoK8sApiAppsV1DeploymentStrategy) SetRollingUpdate(v IoK8sApiAppsV1RollingUpdateDeployment)`

SetRollingUpdate sets RollingUpdate field to given value.

### HasRollingUpdate

`func (o *IoK8sApiAppsV1DeploymentStrategy) HasRollingUpdate() bool`

HasRollingUpdate returns a boolean if a field has been set.

### GetType

`func (o *IoK8sApiAppsV1DeploymentStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApiAppsV1DeploymentStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApiAppsV1DeploymentStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoK8sApiAppsV1DeploymentStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


