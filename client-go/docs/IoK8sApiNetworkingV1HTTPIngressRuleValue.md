# IoK8sApiNetworkingV1HTTPIngressRuleValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]IoK8sApiNetworkingV1HTTPIngressPath**](IoK8sApiNetworkingV1HTTPIngressPath.md) | paths is a collection of paths that map requests to backends. | 

## Methods

### NewIoK8sApiNetworkingV1HTTPIngressRuleValue

`func NewIoK8sApiNetworkingV1HTTPIngressRuleValue(paths []IoK8sApiNetworkingV1HTTPIngressPath, ) *IoK8sApiNetworkingV1HTTPIngressRuleValue`

NewIoK8sApiNetworkingV1HTTPIngressRuleValue instantiates a new IoK8sApiNetworkingV1HTTPIngressRuleValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1HTTPIngressRuleValueWithDefaults

`func NewIoK8sApiNetworkingV1HTTPIngressRuleValueWithDefaults() *IoK8sApiNetworkingV1HTTPIngressRuleValue`

NewIoK8sApiNetworkingV1HTTPIngressRuleValueWithDefaults instantiates a new IoK8sApiNetworkingV1HTTPIngressRuleValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *IoK8sApiNetworkingV1HTTPIngressRuleValue) GetPaths() []IoK8sApiNetworkingV1HTTPIngressPath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *IoK8sApiNetworkingV1HTTPIngressRuleValue) GetPathsOk() (*[]IoK8sApiNetworkingV1HTTPIngressPath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *IoK8sApiNetworkingV1HTTPIngressRuleValue) SetPaths(v []IoK8sApiNetworkingV1HTTPIngressPath)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


