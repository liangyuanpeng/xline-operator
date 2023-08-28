# IoK8sApiCoreV1ClientIPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutSeconds** | Pointer to **int32** | timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be &gt;0 &amp;&amp; &lt;&#x3D;86400(for 1 day) if ServiceAffinity &#x3D;&#x3D; \&quot;ClientIP\&quot;. Default value is 10800(for 3 hours). | [optional] 

## Methods

### NewIoK8sApiCoreV1ClientIPConfig

`func NewIoK8sApiCoreV1ClientIPConfig() *IoK8sApiCoreV1ClientIPConfig`

NewIoK8sApiCoreV1ClientIPConfig instantiates a new IoK8sApiCoreV1ClientIPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ClientIPConfigWithDefaults

`func NewIoK8sApiCoreV1ClientIPConfigWithDefaults() *IoK8sApiCoreV1ClientIPConfig`

NewIoK8sApiCoreV1ClientIPConfigWithDefaults instantiates a new IoK8sApiCoreV1ClientIPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutSeconds

`func (o *IoK8sApiCoreV1ClientIPConfig) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *IoK8sApiCoreV1ClientIPConfig) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *IoK8sApiCoreV1ClientIPConfig) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *IoK8sApiCoreV1ClientIPConfig) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


