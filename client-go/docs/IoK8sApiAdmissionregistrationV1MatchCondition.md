# IoK8sApiAdmissionregistrationV1MatchCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:  &#39;object&#39; - The object from the incoming request. The value is null for DELETE requests. &#39;oldObject&#39; - The existing object. The value is null for CREATE requests. &#39;request&#39; - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). &#39;authorizer&#39; - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.   See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz &#39;authorizer.requestResource&#39; - A CEL ResourceCheck constructed from the &#39;authorizer&#39; and configured with the   request resource. Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/  Required. | 
**Name** | **string** | Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, &#39;-&#39;, &#39;_&#39; or &#39;.&#39;, and must start and end with an alphanumeric character (e.g. &#39;MyName&#39;,  or &#39;my.name&#39;,  or &#39;123-abc&#39;, regex used for validation is &#39;([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]&#39;) with an optional DNS subdomain prefix and &#39;/&#39; (e.g. &#39;example.com/MyName&#39;)  Required. | 

## Methods

### NewIoK8sApiAdmissionregistrationV1MatchCondition

`func NewIoK8sApiAdmissionregistrationV1MatchCondition(expression string, name string, ) *IoK8sApiAdmissionregistrationV1MatchCondition`

NewIoK8sApiAdmissionregistrationV1MatchCondition instantiates a new IoK8sApiAdmissionregistrationV1MatchCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiAdmissionregistrationV1MatchConditionWithDefaults

`func NewIoK8sApiAdmissionregistrationV1MatchConditionWithDefaults() *IoK8sApiAdmissionregistrationV1MatchCondition`

NewIoK8sApiAdmissionregistrationV1MatchConditionWithDefaults instantiates a new IoK8sApiAdmissionregistrationV1MatchCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *IoK8sApiAdmissionregistrationV1MatchCondition) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *IoK8sApiAdmissionregistrationV1MatchCondition) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *IoK8sApiAdmissionregistrationV1MatchCondition) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetName

`func (o *IoK8sApiAdmissionregistrationV1MatchCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiAdmissionregistrationV1MatchCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiAdmissionregistrationV1MatchCondition) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


