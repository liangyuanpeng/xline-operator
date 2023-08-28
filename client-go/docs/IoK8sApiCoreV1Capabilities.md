# IoK8sApiCoreV1Capabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | Added capabilities | [optional] 
**Drop** | Pointer to **[]string** | Removed capabilities | [optional] 

## Methods

### NewIoK8sApiCoreV1Capabilities

`func NewIoK8sApiCoreV1Capabilities() *IoK8sApiCoreV1Capabilities`

NewIoK8sApiCoreV1Capabilities instantiates a new IoK8sApiCoreV1Capabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1CapabilitiesWithDefaults

`func NewIoK8sApiCoreV1CapabilitiesWithDefaults() *IoK8sApiCoreV1Capabilities`

NewIoK8sApiCoreV1CapabilitiesWithDefaults instantiates a new IoK8sApiCoreV1Capabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *IoK8sApiCoreV1Capabilities) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *IoK8sApiCoreV1Capabilities) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *IoK8sApiCoreV1Capabilities) SetAdd(v []string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *IoK8sApiCoreV1Capabilities) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetDrop

`func (o *IoK8sApiCoreV1Capabilities) GetDrop() []string`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *IoK8sApiCoreV1Capabilities) GetDropOk() (*[]string, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *IoK8sApiCoreV1Capabilities) SetDrop(v []string)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *IoK8sApiCoreV1Capabilities) HasDrop() bool`

HasDrop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


