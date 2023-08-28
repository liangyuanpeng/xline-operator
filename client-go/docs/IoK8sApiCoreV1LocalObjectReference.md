# IoK8sApiCoreV1LocalObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 

## Methods

### NewIoK8sApiCoreV1LocalObjectReference

`func NewIoK8sApiCoreV1LocalObjectReference() *IoK8sApiCoreV1LocalObjectReference`

NewIoK8sApiCoreV1LocalObjectReference instantiates a new IoK8sApiCoreV1LocalObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1LocalObjectReferenceWithDefaults

`func NewIoK8sApiCoreV1LocalObjectReferenceWithDefaults() *IoK8sApiCoreV1LocalObjectReference`

NewIoK8sApiCoreV1LocalObjectReferenceWithDefaults instantiates a new IoK8sApiCoreV1LocalObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiCoreV1LocalObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1LocalObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1LocalObjectReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1LocalObjectReference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


