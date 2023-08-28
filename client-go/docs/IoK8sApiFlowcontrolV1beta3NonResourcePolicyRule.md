# IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NonResourceURLs** | **[]string** | &#x60;nonResourceURLs&#x60; is a set of url prefixes that a user should have access to and may not be empty. For example:   - \&quot;/healthz\&quot; is legal   - \&quot;/hea*\&quot; is illegal   - \&quot;/hea\&quot; is legal but matches nothing   - \&quot;/hea/_*\&quot; also matches nothing   - \&quot;/healthz/_*\&quot; matches all per-component health checks. \&quot;*\&quot; matches all non-resource urls. if it is present, it must be the only entry. Required. | 
**Verbs** | **[]string** | &#x60;verbs&#x60; is a list of matching verbs and may not be empty. \&quot;*\&quot; matches all verbs. If it is present, it must be the only entry. Required. | 

## Methods

### NewIoK8sApiFlowcontrolV1beta3NonResourcePolicyRule

`func NewIoK8sApiFlowcontrolV1beta3NonResourcePolicyRule(nonResourceURLs []string, verbs []string, ) *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule`

NewIoK8sApiFlowcontrolV1beta3NonResourcePolicyRule instantiates a new IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiFlowcontrolV1beta3NonResourcePolicyRuleWithDefaults

`func NewIoK8sApiFlowcontrolV1beta3NonResourcePolicyRuleWithDefaults() *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule`

NewIoK8sApiFlowcontrolV1beta3NonResourcePolicyRuleWithDefaults instantiates a new IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonResourceURLs

`func (o *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule) GetNonResourceURLs() []string`

GetNonResourceURLs returns the NonResourceURLs field if non-nil, zero value otherwise.

### GetNonResourceURLsOk

`func (o *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule) GetNonResourceURLsOk() (*[]string, bool)`

GetNonResourceURLsOk returns a tuple with the NonResourceURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceURLs

`func (o *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule) SetNonResourceURLs(v []string)`

SetNonResourceURLs sets NonResourceURLs field to given value.


### GetVerbs

`func (o *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *IoK8sApiFlowcontrolV1beta3NonResourcePolicyRule) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


