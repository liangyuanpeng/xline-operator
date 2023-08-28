# IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaBundle** | Pointer to **string** | CABundle is a PEM encoded CA bundle which will be used to validate an API server&#39;s serving certificate. If unspecified, system trust roots on the apiserver are used. | [optional] 
**Group** | Pointer to **string** | Group is the API group name this server hosts | [optional] 
**GroupPriorityMinimum** | **int32** | GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We&#39;d recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s | 
**InsecureSkipTLSVerify** | Pointer to **bool** | InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead. | [optional] 
**Service** | Pointer to [**IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference**](IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference.md) |  | [optional] 
**Version** | Pointer to **string** | Version is the API version this server hosts.  For example, \&quot;v1\&quot; | [optional] 
**VersionPriority** | **int32** | VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it&#39;s inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is \&quot;kube-like\&quot;, it will sort above non \&quot;kube-like\&quot; version strings, which are ordered lexicographically. \&quot;Kube-like\&quot; versions start with a \&quot;v\&quot;, then are followed by a number (the major version), then optionally the string \&quot;alpha\&quot; or \&quot;beta\&quot; and another number (the minor version). These are sorted first by GA &gt; beta &gt; alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10. | 

## Methods

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec(groupPriorityMinimum int32, versionPriority int32, ) *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpecWithDefaults

`func NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpecWithDefaults() *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec`

NewIoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpecWithDefaults instantiates a new IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaBundle

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.

### HasCaBundle

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) HasCaBundle() bool`

HasCaBundle returns a boolean if a field has been set.

### GetGroup

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupPriorityMinimum

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetGroupPriorityMinimum() int32`

GetGroupPriorityMinimum returns the GroupPriorityMinimum field if non-nil, zero value otherwise.

### GetGroupPriorityMinimumOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetGroupPriorityMinimumOk() (*int32, bool)`

GetGroupPriorityMinimumOk returns a tuple with the GroupPriorityMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPriorityMinimum

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetGroupPriorityMinimum(v int32)`

SetGroupPriorityMinimum sets GroupPriorityMinimum field to given value.


### GetInsecureSkipTLSVerify

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetInsecureSkipTLSVerify() bool`

GetInsecureSkipTLSVerify returns the InsecureSkipTLSVerify field if non-nil, zero value otherwise.

### GetInsecureSkipTLSVerifyOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetInsecureSkipTLSVerifyOk() (*bool, bool)`

GetInsecureSkipTLSVerifyOk returns a tuple with the InsecureSkipTLSVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipTLSVerify

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetInsecureSkipTLSVerify(v bool)`

SetInsecureSkipTLSVerify sets InsecureSkipTLSVerify field to given value.

### HasInsecureSkipTLSVerify

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) HasInsecureSkipTLSVerify() bool`

HasInsecureSkipTLSVerify returns a boolean if a field has been set.

### GetService

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetService() IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetServiceOk() (*IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetService(v IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference)`

SetService sets Service field to given value.

### HasService

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) HasService() bool`

HasService returns a boolean if a field has been set.

### GetVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionPriority

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetVersionPriority() int32`

GetVersionPriority returns the VersionPriority field if non-nil, zero value otherwise.

### GetVersionPriorityOk

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) GetVersionPriorityOk() (*int32, bool)`

GetVersionPriorityOk returns a tuple with the VersionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPriority

`func (o *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) SetVersionPriority(v int32)`

SetVersionPriority sets VersionPriority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


