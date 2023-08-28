# IoK8sApiFlowcontrolV1beta3ServiceAccountSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | &#x60;name&#x60; is the name of matching ServiceAccount objects, or \&quot;*\&quot; to match regardless of name. Required. | 
**Namespace** | **string** | &#x60;namespace&#x60; is the namespace of matching ServiceAccount objects. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta3ServiceAccountSubject

`func NewIoK8sApiFlowcontrolV1beta3ServiceAccountSubject(name string, namespace string, ) *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject`

NewIoK8sApiFlowcontrolV1beta3ServiceAccountSubject instantiates a new IoK8sApiFlowcontrolV1beta3ServiceAccountSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3ServiceAccountSubjectWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3ServiceAccountSubjectWithDefaults() *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject`

NewIoK8sApiFlowcontrolV1beta3ServiceAccountSubjectWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3ServiceAccountSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoK8sApiFlowcontrolV1beta3ServiceAccountSubject) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


