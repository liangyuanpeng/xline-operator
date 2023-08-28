# IoK8sApiCoreV1ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1Condition**](IoK8sApimachineryPkgApisMetaV1Condition.md) | Current service state | [optional] 
**LoadBalancer** | Pointer to [**IoK8sApiCoreV1LoadBalancerStatus**](IoK8sApiCoreV1LoadBalancerStatus.md) |  | [optional] 

## Methods

### NewIoK8sApiCoreV1ServiceStatus

`func NewIoK8sApiCoreV1ServiceStatus() *IoK8sApiCoreV1ServiceStatus`

NewIoK8sApiCoreV1ServiceStatus instantiates a new IoK8sApiCoreV1ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ServiceStatusWithDefaults

`func NewIoK8sApiCoreV1ServiceStatusWithDefaults() *IoK8sApiCoreV1ServiceStatus`

NewIoK8sApiCoreV1ServiceStatusWithDefaults instantiates a new IoK8sApiCoreV1ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiCoreV1ServiceStatus) GetConditions() []IoK8sApimachineryPkgApisMetaV1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCoreV1ServiceStatus) GetConditionsOk() (*[]IoK8sApimachineryPkgApisMetaV1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCoreV1ServiceStatus) SetConditions(v []IoK8sApimachineryPkgApisMetaV1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCoreV1ServiceStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *IoK8sApiCoreV1ServiceStatus) GetLoadBalancer() IoK8sApiCoreV1LoadBalancerStatus`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *IoK8sApiCoreV1ServiceStatus) GetLoadBalancerOk() (*IoK8sApiCoreV1LoadBalancerStatus, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *IoK8sApiCoreV1ServiceStatus) SetLoadBalancer(v IoK8sApiCoreV1LoadBalancerStatus)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *IoK8sApiCoreV1ServiceStatus) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


