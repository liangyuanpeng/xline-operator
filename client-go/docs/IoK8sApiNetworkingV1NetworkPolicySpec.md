# IoK8sApiNetworkingV1NetworkPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Egress** | Pointer to [**[]IoK8sApiNetworkingV1NetworkPolicyEgressRule**](IoK8sApiNetworkingV1NetworkPolicyEgressRule.md) | egress is a list of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8 | [optional] 
**Ingress** | Pointer to [**[]IoK8sApiNetworkingV1NetworkPolicyIngressRule**](IoK8sApiNetworkingV1NetworkPolicyIngressRule.md) | ingress is a list of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod&#39;s local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default) | [optional] 
**PodSelector** | [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | 
**PolicyTypes** | Pointer to **[]string** | policyTypes is a list of rule types that the NetworkPolicy relates to. Valid options are [\&quot;Ingress\&quot;], [\&quot;Egress\&quot;], or [\&quot;Ingress\&quot;, \&quot;Egress\&quot;]. If this field is not specified, it will default based on the existence of ingress or egress rules; policies that contain an egress section are assumed to affect egress, and all policies (whether or not they contain an ingress section) are assumed to affect ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ \&quot;Egress\&quot; ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include \&quot;Egress\&quot; (since such a policy would not include an egress section and would otherwise default to just [ \&quot;Ingress\&quot; ]). This field is beta-level in 1.8 | [optional] 

## Methods

### NewIoK8sApiNetworkingV1NetworkPolicySpec

`func NewIoK8sApiNetworkingV1NetworkPolicySpec(podSelector IoK8sApimachineryPkgApisMetaV1LabelSelector, ) *IoK8sApiNetworkingV1NetworkPolicySpec`

NewIoK8sApiNetworkingV1NetworkPolicySpec instantiates a new IoK8sApiNetworkingV1NetworkPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1NetworkPolicySpecWithDefaults

`func NewIoK8sApiNetworkingV1NetworkPolicySpecWithDefaults() *IoK8sApiNetworkingV1NetworkPolicySpec`

NewIoK8sApiNetworkingV1NetworkPolicySpecWithDefaults instantiates a new IoK8sApiNetworkingV1NetworkPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgress

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetEgress() []IoK8sApiNetworkingV1NetworkPolicyEgressRule`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetEgressOk() (*[]IoK8sApiNetworkingV1NetworkPolicyEgressRule, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) SetEgress(v []IoK8sApiNetworkingV1NetworkPolicyEgressRule)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) HasEgress() bool`

HasEgress returns a boolean if a field has been set.

### GetIngress

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetIngress() []IoK8sApiNetworkingV1NetworkPolicyIngressRule`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetIngressOk() (*[]IoK8sApiNetworkingV1NetworkPolicyIngressRule, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) SetIngress(v []IoK8sApiNetworkingV1NetworkPolicyIngressRule)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) HasIngress() bool`

HasIngress returns a boolean if a field has been set.

### GetPodSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetPodSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetPodSelector returns the PodSelector field if non-nil, zero value otherwise.

### GetPodSelectorOk

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetPodSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetPodSelectorOk returns a tuple with the PodSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSelector

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) SetPodSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetPodSelector sets PodSelector field to given value.


### GetPolicyTypes

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *IoK8sApiNetworkingV1NetworkPolicySpec) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


