# IoK8sApiCoreV1ClaimSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceClaimName** | Pointer to **string** | ResourceClaimName is the name of a ResourceClaim object in the same namespace as this pod. | [optional] 
**ResourceClaimTemplateName** | Pointer to **string** | ResourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace as this pod.  The template will be used to create a new ResourceClaim, which will be bound to this pod. When this pod is deleted, the ResourceClaim will also be deleted. The name of the ResourceClaim will be &lt;pod name&gt;-&lt;resource name&gt;, where &lt;resource name&gt; is the PodResourceClaim.Name. Pod validation will reject the pod if the concatenated name is not valid for a ResourceClaim (e.g. too long).  An existing ResourceClaim with that name that is not owned by the pod will not be used for the pod to avoid using an unrelated resource by mistake. Scheduling and pod startup are then blocked until the unrelated ResourceClaim is removed.  This field is immutable and no changes will be made to the corresponding ResourceClaim by the control plane after creating the ResourceClaim. | [optional] 

## Methods

### NewIoK8sApiCoreV1ClaimSource

`func NewIoK8sApiCoreV1ClaimSource() *IoK8sApiCoreV1ClaimSource`

NewIoK8sApiCoreV1ClaimSource instantiates a new IoK8sApiCoreV1ClaimSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ClaimSourceWithDefaults

`func NewIoK8sApiCoreV1ClaimSourceWithDefaults() *IoK8sApiCoreV1ClaimSource`

NewIoK8sApiCoreV1ClaimSourceWithDefaults instantiates a new IoK8sApiCoreV1ClaimSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceClaimName

`func (o *IoK8sApiCoreV1ClaimSource) GetResourceClaimName() string`

GetResourceClaimName returns the ResourceClaimName field if non-nil, zero value otherwise.

### GetResourceClaimNameOk

`func (o *IoK8sApiCoreV1ClaimSource) GetResourceClaimNameOk() (*string, bool)`

GetResourceClaimNameOk returns a tuple with the ResourceClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceClaimName

`func (o *IoK8sApiCoreV1ClaimSource) SetResourceClaimName(v string)`

SetResourceClaimName sets ResourceClaimName field to given value.

### HasResourceClaimName

`func (o *IoK8sApiCoreV1ClaimSource) HasResourceClaimName() bool`

HasResourceClaimName returns a boolean if a field has been set.

### GetResourceClaimTemplateName

`func (o *IoK8sApiCoreV1ClaimSource) GetResourceClaimTemplateName() string`

GetResourceClaimTemplateName returns the ResourceClaimTemplateName field if non-nil, zero value otherwise.

### GetResourceClaimTemplateNameOk

`func (o *IoK8sApiCoreV1ClaimSource) GetResourceClaimTemplateNameOk() (*string, bool)`

GetResourceClaimTemplateNameOk returns a tuple with the ResourceClaimTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceClaimTemplateName

`func (o *IoK8sApiCoreV1ClaimSource) SetResourceClaimTemplateName(v string)`

SetResourceClaimTemplateName sets ResourceClaimTemplateName field to given value.

### HasResourceClaimTemplateName

`func (o *IoK8sApiCoreV1ClaimSource) HasResourceClaimTemplateName() bool`

HasResourceClaimTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


