# IoK8sApimachineryPkgApisMetaV1WatchEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **map[string]interface{}** | RawExtension is used to hold extensions in external versions.  To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types.  // Internal package:   type MyAPIObject struct {   runtime.TypeMeta &#x60;json:\&quot;,inline\&quot;&#x60;   MyPlugin runtime.Object &#x60;json:\&quot;myPlugin\&quot;&#x60;  }   type PluginA struct {   AOption string &#x60;json:\&quot;aOption\&quot;&#x60;  }  // External package:   type MyAPIObject struct {   runtime.TypeMeta &#x60;json:\&quot;,inline\&quot;&#x60;   MyPlugin runtime.RawExtension &#x60;json:\&quot;myPlugin\&quot;&#x60;  }   type PluginA struct {   AOption string &#x60;json:\&quot;aOption\&quot;&#x60;  }  // On the wire, the JSON will look something like this:   {   \&quot;kind\&quot;:\&quot;MyAPIObject\&quot;,   \&quot;apiVersion\&quot;:\&quot;v1\&quot;,   \&quot;myPlugin\&quot;: {    \&quot;kind\&quot;:\&quot;PluginA\&quot;,    \&quot;aOption\&quot;:\&quot;foo\&quot;,   },  }  So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package&#39;s DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.) | 
**Type** | **string** |  | 

## Methods

### NewIoK8sApimachineryPkgApisMetaV1WatchEvent

`func NewIoK8sApimachineryPkgApisMetaV1WatchEvent(object map[string]interface{}, type_ string, ) *IoK8sApimachineryPkgApisMetaV1WatchEvent`

NewIoK8sApimachineryPkgApisMetaV1WatchEvent instantiates a new IoK8sApimachineryPkgApisMetaV1WatchEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApimachineryPkgApisMetaV1WatchEventWithDefaults

`func NewIoK8sApimachineryPkgApisMetaV1WatchEventWithDefaults() *IoK8sApimachineryPkgApisMetaV1WatchEvent`

NewIoK8sApimachineryPkgApisMetaV1WatchEventWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1WatchEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *IoK8sApimachineryPkgApisMetaV1WatchEvent) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IoK8sApimachineryPkgApisMetaV1WatchEvent) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IoK8sApimachineryPkgApisMetaV1WatchEvent) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.


### GetType

`func (o *IoK8sApimachineryPkgApisMetaV1WatchEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoK8sApimachineryPkgApisMetaV1WatchEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoK8sApimachineryPkgApisMetaV1WatchEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


