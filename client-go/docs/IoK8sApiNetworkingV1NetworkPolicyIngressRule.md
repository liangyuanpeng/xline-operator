# IoK8sApiNetworkingV1NetworkPolicyIngressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**[]IoK8sApiNetworkingV1NetworkPolicyPeer**](IoK8sApiNetworkingV1NetworkPolicyPeer.md) | from is a list of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list. | [optional] 
**Ports** | Pointer to [**[]IoK8sApiNetworkingV1NetworkPolicyPort**](IoK8sApiNetworkingV1NetworkPolicyPort.md) | ports is a list of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1NetworkPolicyIngressRule

`func NewIoK8sApiNetworkingV1NetworkPolicyIngressRule() *IoK8sApiNetworkingV1NetworkPolicyIngressRule`

NewIoK8sApiNetworkingV1NetworkPolicyIngressRule instantiates a new IoK8sApiNetworkingV1NetworkPolicyIngressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1NetworkPolicyIngressRuleWithDefaults

`func NewIoK8sApiNetworkingV1NetworkPolicyIngressRuleWithDefaults() *IoK8sApiNetworkingV1NetworkPolicyIngressRule`

NewIoK8sApiNetworkingV1NetworkPolicyIngressRuleWithDefaults instantiates a new IoK8sApiNetworkingV1NetworkPolicyIngressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) GetFrom() []IoK8sApiNetworkingV1NetworkPolicyPeer`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) GetFromOk() (*[]IoK8sApiNetworkingV1NetworkPolicyPeer, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) SetFrom(v []IoK8sApiNetworkingV1NetworkPolicyPeer)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetPorts

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) GetPorts() []IoK8sApiNetworkingV1NetworkPolicyPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) GetPortsOk() (*[]IoK8sApiNetworkingV1NetworkPolicyPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) SetPorts(v []IoK8sApiNetworkingV1NetworkPolicyPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiNetworkingV1NetworkPolicyIngressRule) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


