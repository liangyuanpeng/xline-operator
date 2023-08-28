# IoK8sApiStorageV1CSIStorageCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Capacity** | Pointer to **string** | Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.  The serialization format is:  &#x60;&#x60;&#x60; &lt;quantity&gt;        ::&#x3D; &lt;signedNumber&gt;&lt;suffix&gt;   (Note that &lt;suffix&gt; may be empty, from the \&quot;\&quot; case in &lt;decimalSI&gt;.)  &lt;digit&gt;           ::&#x3D; 0 | 1 | ... | 9 &lt;digits&gt;          ::&#x3D; &lt;digit&gt; | &lt;digit&gt;&lt;digits&gt; &lt;number&gt;          ::&#x3D; &lt;digits&gt; | &lt;digits&gt;.&lt;digits&gt; | &lt;digits&gt;. | .&lt;digits&gt; &lt;sign&gt;            ::&#x3D; \&quot;+\&quot; | \&quot;-\&quot; &lt;signedNumber&gt;    ::&#x3D; &lt;number&gt; | &lt;sign&gt;&lt;number&gt; &lt;suffix&gt;          ::&#x3D; &lt;binarySI&gt; | &lt;decimalExponent&gt; | &lt;decimalSI&gt; &lt;binarySI&gt;        ::&#x3D; Ki | Mi | Gi | Ti | Pi | Ei   (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)  &lt;decimalSI&gt;       ::&#x3D; m | \&quot;\&quot; | k | M | G | T | P | E   (Note that 1024 &#x3D; 1Ki but 1000 &#x3D; 1k; I didn&#39;t choose the capitalization.)  &lt;decimalExponent&gt; ::&#x3D; \&quot;e\&quot; &lt;signedNumber&gt; | \&quot;E\&quot; &lt;signedNumber&gt; &#x60;&#x60;&#x60;  No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.  When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.  Before serializing, Quantity will be put in \&quot;canonical form\&quot;. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:  - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.  The sign will be omitted unless the number is negative.  Examples:  - 1.5 will be serialized as \&quot;1500m\&quot; - 1.5Gi will be serialized as \&quot;1536Mi\&quot;  Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.  Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don&#39;t diff.)  This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation. | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**MaximumVolumeSize** | Pointer to **string** | Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.  The serialization format is:  &#x60;&#x60;&#x60; &lt;quantity&gt;        ::&#x3D; &lt;signedNumber&gt;&lt;suffix&gt;   (Note that &lt;suffix&gt; may be empty, from the \&quot;\&quot; case in &lt;decimalSI&gt;.)  &lt;digit&gt;           ::&#x3D; 0 | 1 | ... | 9 &lt;digits&gt;          ::&#x3D; &lt;digit&gt; | &lt;digit&gt;&lt;digits&gt; &lt;number&gt;          ::&#x3D; &lt;digits&gt; | &lt;digits&gt;.&lt;digits&gt; | &lt;digits&gt;. | .&lt;digits&gt; &lt;sign&gt;            ::&#x3D; \&quot;+\&quot; | \&quot;-\&quot; &lt;signedNumber&gt;    ::&#x3D; &lt;number&gt; | &lt;sign&gt;&lt;number&gt; &lt;suffix&gt;          ::&#x3D; &lt;binarySI&gt; | &lt;decimalExponent&gt; | &lt;decimalSI&gt; &lt;binarySI&gt;        ::&#x3D; Ki | Mi | Gi | Ti | Pi | Ei   (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)  &lt;decimalSI&gt;       ::&#x3D; m | \&quot;\&quot; | k | M | G | T | P | E   (Note that 1024 &#x3D; 1Ki but 1000 &#x3D; 1k; I didn&#39;t choose the capitalization.)  &lt;decimalExponent&gt; ::&#x3D; \&quot;e\&quot; &lt;signedNumber&gt; | \&quot;E\&quot; &lt;signedNumber&gt; &#x60;&#x60;&#x60;  No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.  When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.  Before serializing, Quantity will be put in \&quot;canonical form\&quot;. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:  - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible.  The sign will be omitted unless the number is negative.  Examples:  - 1.5 will be serialized as \&quot;1500m\&quot; - 1.5Gi will be serialized as \&quot;1536Mi\&quot;  Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.  Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don&#39;t diff.)  This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation. | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**NodeTopology** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**StorageClassName** | **string** | storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable. | 

## Methods

### NewIoK8sApiStorageV1CSIStorageCapacity

`func NewIoK8sApiStorageV1CSIStorageCapacity(storageClassName string, ) *IoK8sApiStorageV1CSIStorageCapacity`

NewIoK8sApiStorageV1CSIStorageCapacity instantiates a new IoK8sApiStorageV1CSIStorageCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiStorageV1CSIStorageCapacityWithDefaults

`func NewIoK8sApiStorageV1CSIStorageCapacityWithDefaults() *IoK8sApiStorageV1CSIStorageCapacity`

NewIoK8sApiStorageV1CSIStorageCapacityWithDefaults instantiates a new IoK8sApiStorageV1CSIStorageCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoK8sApiStorageV1CSIStorageCapacity) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCapacity

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *IoK8sApiStorageV1CSIStorageCapacity) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetKind

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoK8sApiStorageV1CSIStorageCapacity) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaximumVolumeSize

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetMaximumVolumeSize() string`

GetMaximumVolumeSize returns the MaximumVolumeSize field if non-nil, zero value otherwise.

### GetMaximumVolumeSizeOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetMaximumVolumeSizeOk() (*string, bool)`

GetMaximumVolumeSizeOk returns a tuple with the MaximumVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumVolumeSize

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetMaximumVolumeSize(v string)`

SetMaximumVolumeSize sets MaximumVolumeSize field to given value.

### HasMaximumVolumeSize

`func (o *IoK8sApiStorageV1CSIStorageCapacity) HasMaximumVolumeSize() bool`

HasMaximumVolumeSize returns a boolean if a field has been set.

### GetMetadata

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoK8sApiStorageV1CSIStorageCapacity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNodeTopology

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetNodeTopology() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetNodeTopology returns the NodeTopology field if non-nil, zero value otherwise.

### GetNodeTopologyOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetNodeTopologyOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetNodeTopologyOk returns a tuple with the NodeTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTopology

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetNodeTopology(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetNodeTopology sets NodeTopology field to given value.

### HasNodeTopology

`func (o *IoK8sApiStorageV1CSIStorageCapacity) HasNodeTopology() bool`

HasNodeTopology returns a boolean if a field has been set.

### GetStorageClassName

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *IoK8sApiStorageV1CSIStorageCapacity) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *IoK8sApiStorageV1CSIStorageCapacity) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


