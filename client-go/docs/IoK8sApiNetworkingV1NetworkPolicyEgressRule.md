# IoK8sApiNetworkingV1NetworkPolicyEgressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]IoK8sApiNetworkingV1NetworkPolicyPort**](IoK8sApiNetworkingV1NetworkPolicyPort.md) | ports is a list of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list. | [optional] 
**To** | Pointer to [**[]IoK8sApiNetworkingV1NetworkPolicyPeer**](IoK8sApiNetworkingV1NetworkPolicyPeer.md) | to is a list of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list. | [optional] 

## Methods

### NewIoK8sApiNetworkingV1NetworkPolicyEgressRule

`func NewIoK8sApiNetworkingV1NetworkPolicyEgressRule() *IoK8sApiNetworkingV1NetworkPolicyEgressRule`

NewIoK8sApiNetworkingV1NetworkPolicyEgressRule instantiates a new IoK8sApiNetworkingV1NetworkPolicyEgressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1NetworkPolicyEgressRuleWithDefaults

`func NewIoK8sApiNetworkingV1NetworkPolicyEgressRuleWithDefaults() *IoK8sApiNetworkingV1NetworkPolicyEgressRule`

NewIoK8sApiNetworkingV1NetworkPolicyEgressRuleWithDefaults instantiates a new IoK8sApiNetworkingV1NetworkPolicyEgressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) GetPorts() []IoK8sApiNetworkingV1NetworkPolicyPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) GetPortsOk() (*[]IoK8sApiNetworkingV1NetworkPolicyPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) SetPorts(v []IoK8sApiNetworkingV1NetworkPolicyPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTo

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) GetTo() []IoK8sApiNetworkingV1NetworkPolicyPeer`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) GetToOk() (*[]IoK8sApiNetworkingV1NetworkPolicyPeer, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) SetTo(v []IoK8sApiNetworkingV1NetworkPolicyPeer)`

SetTo sets To field to given value.

### HasTo

`func (o *IoK8sApiNetworkingV1NetworkPolicyEgressRule) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


