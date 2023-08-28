# CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalhostProfile** | Pointer to **string** | localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet&#39;s configured seccomp profile location. Must only be set if type is \&quot;Localhost\&quot;. | [optional] 
**Type** | **string** | type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.   | 

## Methods

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile(type_ string, ) *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfileWithDefaults

`func NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfileWithDefaults() *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile`

NewCloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfileWithDefaults instantiates a new CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalhostProfile

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) GetLocalhostProfile() string`

GetLocalhostProfile returns the LocalhostProfile field if non-nil, zero value otherwise.

### GetLocalhostProfileOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) GetLocalhostProfileOk() (*string, bool)`

GetLocalhostProfileOk returns a tuple with the LocalhostProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalhostProfile

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) SetLocalhostProfile(v string)`

SetLocalhostProfile sets LocalhostProfile field to given value.

### HasLocalhostProfile

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) HasLocalhostProfile() bool`

HasLocalhostProfile returns a boolean if a field has been set.

### GetType

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudXlineXlineoperatorV1alpha1XlineClusterSpecContainerSecurityContextSeccompProfile) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


