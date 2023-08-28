# IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the version of this resource that this field set applies to. The format is \&quot;group/version\&quot; just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted. | [optional] 
**FieldsType** | Pointer to **string** | FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \&quot;FieldsV1\&quot; | [optional] 
**FieldsV1** | Pointer to **map[string]interface{}** | FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.  Each key is either a &#39;.&#39; representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: &#39;f:&lt;name&gt;&#39;, where &lt;name&gt; is the name of a field in a struct, or key in a map &#39;v:&lt;value&gt;&#39;, where &lt;value&gt; is the exact json formatted value of a list item &#39;i:&lt;index&gt;&#39;, where &lt;index&gt; is position of a item in a list &#39;k:&lt;keys&gt;&#39;, where &lt;keys&gt; is a map of  a list item&#39;s key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.  The exact format is defined in sigs.k8s.io/structured-merge-diff | [optional] 
**Manager** | Pointer to **string** | Manager is an identifier of the workflow managing these fields. | [optional] 
**Operation** | Pointer to **string** | Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are &#39;Apply&#39; and &#39;Update&#39;. | [optional] 
**Subresource** | Pointer to **string** | Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource. | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry

`func NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry() *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry`

NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry instantiates a new IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntryWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntryWithDefaults() *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry`

NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntryWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetFieldsType

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsType() string`

GetFieldsType returns the FieldsType field if non-nil, zero value otherwise.

### GetFieldsTypeOk

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsTypeOk() (*string, bool)`

GetFieldsTypeOk returns a tuple with the FieldsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsType

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetFieldsType(v string)`

SetFieldsType sets FieldsType field to given value.

### HasFieldsType

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasFieldsType() bool`

HasFieldsType returns a boolean if a field has been set.

### GetFieldsV1

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsV1() map[string]interface{}`

GetFieldsV1 returns the FieldsV1 field if non-nil, zero value otherwise.

### GetFieldsV1Ok

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsV1Ok() (*map[string]interface{}, bool)`

GetFieldsV1Ok returns a tuple with the FieldsV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsV1

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetFieldsV1(v map[string]interface{})`

SetFieldsV1 sets FieldsV1 field to given value.

### HasFieldsV1

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasFieldsV1() bool`

HasFieldsV1 returns a boolean if a field has been set.

### GetManager

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetOperation

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSubresource

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetSubresource() string`

GetSubresource returns the Subresource field if non-nil, zero value otherwise.

### GetSubresourceOk

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetSubresourceOk() (*string, bool)`

GetSubresourceOk returns a tuple with the Subresource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresource

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetSubresource(v string)`

SetSubresource sets Subresource field to given value.

### HasSubresource

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasSubresource() bool`

HasSubresource returns a boolean if a field has been set.

### GetTime

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


