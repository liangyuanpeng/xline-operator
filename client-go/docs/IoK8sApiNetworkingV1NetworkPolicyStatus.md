# IoK8sApiNetworkingV1NetworkPolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1Condition**](IoK8sApimachineryPkgApisMetaV1Condition.md) | conditions holds an array of metav1.Condition that describe the state of the NetworkPolicy. Current service state | [optional] 

## Methods

### NewIoK8sApiNetworkingV1NetworkPolicyStatus

`func NewIoK8sApiNetworkingV1NetworkPolicyStatus() *IoK8sApiNetworkingV1NetworkPolicyStatus`

NewIoK8sApiNetworkingV1NetworkPolicyStatus instantiates a new IoK8sApiNetworkingV1NetworkPolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1NetworkPolicyStatusWithDefaults

`func NewIoK8sApiNetworkingV1NetworkPolicyStatusWithDefaults() *IoK8sApiNetworkingV1NetworkPolicyStatus`

NewIoK8sApiNetworkingV1NetworkPolicyStatusWithDefaults instantiates a new IoK8sApiNetworkingV1NetworkPolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sApiNetworkingV1NetworkPolicyStatus) GetConditions() []IoK8sApimachineryPkgApisMetaV1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyStatus) GetConditionsOk() (*[]IoK8sApimachineryPkgApisMetaV1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiNetworkingV1NetworkPolicyStatus) SetConditions(v []IoK8sApimachineryPkgApisMetaV1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiNetworkingV1NetworkPolicyStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


