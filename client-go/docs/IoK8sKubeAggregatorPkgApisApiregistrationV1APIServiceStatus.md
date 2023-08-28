# IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition**](IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition.md) | Current service state of apiService. | [optional] 

## Methods

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatusWithDefaults

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatusWithDefaults() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatusWithDefaults instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus) GetConditions() []IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus) GetConditionsOk() (*[]IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus) SetConditions(v []IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


