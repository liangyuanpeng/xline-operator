# IoK8sApiCoreV1ResourceClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container. | 

## Methods

### NewIoK8sApiCoreV1ResourceClaim

`func NewIoK8sApiCoreV1ResourceClaim(name string, ) *IoK8sApiCoreV1ResourceClaim`

NewIoK8sApiCoreV1ResourceClaim instantiates a new IoK8sApiCoreV1ResourceClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ResourceClaimWithDefaults

`func NewIoK8sApiCoreV1ResourceClaimWithDefaults() *IoK8sApiCoreV1ResourceClaim`

NewIoK8sApiCoreV1ResourceClaimWithDefaults instantiates a new IoK8sApiCoreV1ResourceClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiCoreV1ResourceClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ResourceClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ResourceClaim) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


