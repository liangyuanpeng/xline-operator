# IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the name of the service | [optional] 
**Namespace** | Pointer to **string** | Namespace is the namespace of the service | [optional] 
**Port** | Pointer to **int32** | If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. &#x60;port&#x60; should be a valid port number (1-65535, inclusive). | [optional] 

## Methods

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference() *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReferenceWithDefaults

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReferenceWithDefaults() *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReferenceWithDefaults instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


