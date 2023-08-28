# IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedNames** | Pointer to [**IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames.md) |  | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition**](IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition.md) | conditions indicate state for particular aspects of a CustomResourceDefinition | [optional] 
**StoredVersions** | Pointer to **[]string** | storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list. Versions may not be removed from &#x60;spec.versions&#x60; while they exist in this list. | [optional] 

## Methods

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatusWithDefaults

`func NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatusWithDefaults() *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus`

NewIoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatusWithDefaults instantiates a new IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) GetAcceptedNames() IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames`

GetAcceptedNames returns the AcceptedNames field if non-nil, zero value otherwise.

### GetAcceptedNamesOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) GetAcceptedNamesOk() (*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames, bool)`

GetAcceptedNamesOk returns a tuple with the AcceptedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) SetAcceptedNames(v IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames)`

SetAcceptedNames sets AcceptedNames field to given value.

### HasAcceptedNames

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) HasAcceptedNames() bool`

HasAcceptedNames returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) GetConditions() []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) GetConditionsOk() (*[]IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) SetConditions(v []IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetStoredVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) GetStoredVersions() []string`

GetStoredVersions returns the StoredVersions field if non-nil, zero value otherwise.

### GetStoredVersionsOk

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) GetStoredVersionsOk() (*[]string, bool)`

GetStoredVersionsOk returns a tuple with the StoredVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) SetStoredVersions(v []string)`

SetStoredVersions sets StoredVersions field to given value.

### HasStoredVersions

`func (o *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) HasStoredVersions() bool`

HasStoredVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


