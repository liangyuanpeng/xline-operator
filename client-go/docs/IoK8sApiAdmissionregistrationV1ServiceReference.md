# IoK8sApiAdmissionregistrationV1ServiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | &#x60;name&#x60; is the name of the service. Required | 
**Namespace** | **string** | &#x60;namespace&#x60; is the namespace of the service. Required | 
**Path** | Pointer to **string** | &#x60;path&#x60; is an optional URL path which will be sent in any request to this service. | [optional] 
**Port** | Pointer to **int32** | If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. &#x60;port&#x60; should be a valid port number (1-65535, inclusive). | [optional] 

## Methods

### NewIoK8sApiAdmissionregistrationV1ServiceReference

`func NewIoK8sApiAdmissionregistrationV1ServiceReference(name string, namespace string, ) *IoK8sApiAdmissionregistrationV1ServiceReference`

NewIoK8sApiAdmissionregistrationV1ServiceReference instantiates a new IoK8sApiAdmissionregistrationV1ServiceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAdmissionregistrationV1ServiceReferenceWithDefaults

`func NewIoK8sApiAdmissionregistrationV1ServiceReferenceWithDefaults() *IoK8sApiAdmissionregistrationV1ServiceReference`

NewIoK8sApiAdmissionregistrationV1ServiceReferenceWithDefaults instantiates a new IoK8sApiAdmissionregistrationV1ServiceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPath

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IoK8sApiAdmissionregistrationV1ServiceReference) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


