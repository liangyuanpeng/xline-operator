# IoK8sApiCoreV1NodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigSource** | Pointer to [**IoK8sApiCoreV1NodeConfigSource**](IoK8sApiCoreV1NodeConfigSource.md) |  | [optional] 
**ExternalID** | Pointer to **string** | Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966 | [optional] 
**PodCIDR** | Pointer to **string** | PodCIDR represents the pod IP range assigned to the node. | [optional] 
**PodCIDRs** | Pointer to **[]string** | podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6. | [optional] 
**ProviderID** | Pointer to **string** | ID of the node assigned by the cloud provider in the format: &lt;ProviderName&gt;://&lt;ProviderSpecificNodeID&gt; | [optional] 
**Taints** | Pointer to [**[]IoK8sApiCoreV1Taint**](IoK8sApiCoreV1Taint.md) | If specified, the node&#39;s taints. | [optional] 
**Unschedulable** | Pointer to **bool** | Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration | [optional] 

## Methods

### NewIoK8sApiCoreV1NodeSpec

`func NewIoK8sApiCoreV1NodeSpec() *IoK8sApiCoreV1NodeSpec`

NewIoK8sApiCoreV1NodeSpec instantiates a new IoK8sApiCoreV1NodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1NodeSpecWithDefaults

`func NewIoK8sApiCoreV1NodeSpecWithDefaults() *IoK8sApiCoreV1NodeSpec`

NewIoK8sApiCoreV1NodeSpecWithDefaults instantiates a new IoK8sApiCoreV1NodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigSource

`func (o *IoK8sApiCoreV1NodeSpec) GetConfigSource() IoK8sApiCoreV1NodeConfigSource`

GetConfigSource returns the ConfigSource field if non-nil, zero value otherwise.

### GetConfigSourceOk

`func (o *IoK8sApiCoreV1NodeSpec) GetConfigSourceOk() (*IoK8sApiCoreV1NodeConfigSource, bool)`

GetConfigSourceOk returns a tuple with the ConfigSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSource

`func (o *IoK8sApiCoreV1NodeSpec) SetConfigSource(v IoK8sApiCoreV1NodeConfigSource)`

SetConfigSource sets ConfigSource field to given value.

### HasConfigSource

`func (o *IoK8sApiCoreV1NodeSpec) HasConfigSource() bool`

HasConfigSource returns a boolean if a field has been set.

### GetExternalID

`func (o *IoK8sApiCoreV1NodeSpec) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *IoK8sApiCoreV1NodeSpec) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *IoK8sApiCoreV1NodeSpec) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *IoK8sApiCoreV1NodeSpec) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetPodCIDR

`func (o *IoK8sApiCoreV1NodeSpec) GetPodCIDR() string`

GetPodCIDR returns the PodCIDR field if non-nil, zero value otherwise.

### GetPodCIDROk

`func (o *IoK8sApiCoreV1NodeSpec) GetPodCIDROk() (*string, bool)`

GetPodCIDROk returns a tuple with the PodCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCIDR

`func (o *IoK8sApiCoreV1NodeSpec) SetPodCIDR(v string)`

SetPodCIDR sets PodCIDR field to given value.

### HasPodCIDR

`func (o *IoK8sApiCoreV1NodeSpec) HasPodCIDR() bool`

HasPodCIDR returns a boolean if a field has been set.

### GetPodCIDRs

`func (o *IoK8sApiCoreV1NodeSpec) GetPodCIDRs() []string`

GetPodCIDRs returns the PodCIDRs field if non-nil, zero value otherwise.

### GetPodCIDRsOk

`func (o *IoK8sApiCoreV1NodeSpec) GetPodCIDRsOk() (*[]string, bool)`

GetPodCIDRsOk returns a tuple with the PodCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCIDRs

`func (o *IoK8sApiCoreV1NodeSpec) SetPodCIDRs(v []string)`

SetPodCIDRs sets PodCIDRs field to given value.

### HasPodCIDRs

`func (o *IoK8sApiCoreV1NodeSpec) HasPodCIDRs() bool`

HasPodCIDRs returns a boolean if a field has been set.

### GetProviderID

`func (o *IoK8sApiCoreV1NodeSpec) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *IoK8sApiCoreV1NodeSpec) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *IoK8sApiCoreV1NodeSpec) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.

### HasProviderID

`func (o *IoK8sApiCoreV1NodeSpec) HasProviderID() bool`

HasProviderID returns a boolean if a field has been set.

### GetTaints

`func (o *IoK8sApiCoreV1NodeSpec) GetTaints() []IoK8sApiCoreV1Taint`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *IoK8sApiCoreV1NodeSpec) GetTaintsOk() (*[]IoK8sApiCoreV1Taint, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *IoK8sApiCoreV1NodeSpec) SetTaints(v []IoK8sApiCoreV1Taint)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *IoK8sApiCoreV1NodeSpec) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### GetUnschedulable

`func (o *IoK8sApiCoreV1NodeSpec) GetUnschedulable() bool`

GetUnschedulable returns the Unschedulable field if non-nil, zero value otherwise.

### GetUnschedulableOk

`func (o *IoK8sApiCoreV1NodeSpec) GetUnschedulableOk() (*bool, bool)`

GetUnschedulableOk returns a tuple with the Unschedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnschedulable

`func (o *IoK8sApiCoreV1NodeSpec) SetUnschedulable(v bool)`

SetUnschedulable sets Unschedulable field to given value.

### HasUnschedulable

`func (o *IoK8sApiCoreV1NodeSpec) HasUnschedulable() bool`

HasUnschedulable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


