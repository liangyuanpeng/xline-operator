# IoK8sApiNetworkingV1NetworkPolicyPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpBlock** | Pointer to [**IoK8sApiNetworkingV1IPBlock**](IoK8sApiNetworkingV1IPBlock.md) |  | [optional] 
**NamespaceSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**PodSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 

## Methods

### NewIoK8sApiNetworkingV1NetworkPolicyPeer

`func NewIoK8sApiNetworkingV1NetworkPolicyPeer() *IoK8sApiNetworkingV1NetworkPolicyPeer`

NewIoK8sApiNetworkingV1NetworkPolicyPeer instantiates a new IoK8sApiNetworkingV1NetworkPolicyPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1NetworkPolicyPeerWithDefaults

`func NewIoK8sApiNetworkingV1NetworkPolicyPeerWithDefaults() *IoK8sApiNetworkingV1NetworkPolicyPeer`

NewIoK8sApiNetworkingV1NetworkPolicyPeerWithDefaults instantiates a new IoK8sApiNetworkingV1NetworkPolicyPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpBlock

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) GetIpBlock() IoK8sApiNetworkingV1IPBlock`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) GetIpBlockOk() (*IoK8sApiNetworkingV1IPBlock, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) SetIpBlock(v IoK8sApiNetworkingV1IPBlock)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### GetNamespaceSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) GetNamespaceSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetNamespaceSelector returns the NamespaceSelector field if non-nil, zero value otherwise.

### GetNamespaceSelectorOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) GetNamespaceSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetNamespaceSelectorOk returns a tuple with the NamespaceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) SetNamespaceSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetNamespaceSelector sets NamespaceSelector field to given value.

### HasNamespaceSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) HasNamespaceSelector() bool`

HasNamespaceSelector returns a boolean if a field has been set.

### GetPodSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) GetPodSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetPodSelector returns the PodSelector field if non-nil, zero value otherwise.

### GetPodSelectorOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) GetPodSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetPodSelectorOk returns a tuple with the PodSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) SetPodSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetPodSelector sets PodSelector field to given value.

### HasPodSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicyPeer) HasPodSelector() bool`

HasPodSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


