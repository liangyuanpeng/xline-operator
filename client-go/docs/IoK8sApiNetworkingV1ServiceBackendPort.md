# IoK8sApiNetworkingV1ServiceBackendPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name is the name of the port on the Service. This is a mutually exclusive setting with \&quot;Number\&quot;. | [optional] 
**Number** | Pointer to **int32** | number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with \&quot;Name\&quot;. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1ServiceBackendPort

`func NewIoK8sApiNetworkingV1ServiceBackendPort() *IoK8sApiNetworkingV1ServiceBackendPort`

NewIoK8sApiNetworkingV1ServiceBackendPort instantiates a new IoK8sApiNetworkingV1ServiceBackendPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1ServiceBackendPortWithDefaults

`func NewIoK8sApiNetworkingV1ServiceBackendPortWithDefaults() *IoK8sApiNetworkingV1ServiceBackendPort`

NewIoK8sApiNetworkingV1ServiceBackendPortWithDefaults instantiates a new IoK8sApiNetworkingV1ServiceBackendPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *IoK8sApiNetworkingV1ServiceBackendPort) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


