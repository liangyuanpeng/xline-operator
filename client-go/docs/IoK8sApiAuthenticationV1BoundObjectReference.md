# IoK8sApiAuthenticationV1BoundObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API version of the referent. | [optional] 
**Kind** | Pointer to **string** | Kind of the referent. Valid kinds are &#39;Pod&#39; and &#39;Secret&#39;. | [optional] 
**Name** | Pointer to **string** | Name of the referent. | [optional] 
**Uid** | Pointer to **string** | UID of the referent. | [optional] 

## Methods

### NewIoK8sApiAuthenticationV1BoundObjectReference

`func NewIoK8sApiAuthenticationV1BoundObjectReference() *IoK8sApiAuthenticationV1BoundObjectReference`

NewIoK8sApiAuthenticationV1BoundObjectReference instantiates a new IoK8sApiAuthenticationV1BoundObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAuthenticationV1BoundObjectReferenceWithDefaults

`func NewIoK8sApiAuthenticationV1BoundObjectReferenceWithDefaults() *IoK8sApiAuthenticationV1BoundObjectReference`

NewIoK8sApiAuthenticationV1BoundObjectReferenceWithDefaults instantiates a new IoK8sApiAuthenticationV1BoundObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoK8sApiAuthenticationV1BoundObjectReference) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


