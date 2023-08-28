# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesClaimsInner**](CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesClaimsInner.md) | Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers. | [optional] 
**Limits** | Pointer to **map[string]string** | Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 
**Requests** | Pointer to **map[string]string** | Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) GetClaims() []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesClaimsInner`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) GetClaimsOk() (*[]CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesClaimsInner, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) SetClaims(v []CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResourcesClaimsInner)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetLimits

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) GetLimits() map[string]string`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) GetLimitsOk() (*map[string]string, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) SetLimits(v map[string]string)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRequests

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) GetRequests() map[string]string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) GetRequestsOk() (*map[string]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) SetRequests(v map[string]string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerResources) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


