# IoK8sApiCoreV1PodDNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameservers** | Pointer to **[]string** | A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed. | [optional] 
**Options** | Pointer to [**[]IoK8sApiCoreV1PodDNSConfigOption**](IoK8sApiCoreV1PodDNSConfigOption.md) | A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. | [optional] 
**Searches** | Pointer to **[]string** | A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed. | [optional] 

## Methods

### NewIoK8sApiCoreV1PodDNSConfig

`func NewIoK8sApiCoreV1PodDNSConfig() *IoK8sApiCoreV1PodDNSConfig`

NewIoK8sApiCoreV1PodDNSConfig instantiates a new IoK8sApiCoreV1PodDNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PodDNSConfigWithDefaults

`func NewIoK8sApiCoreV1PodDNSConfigWithDefaults() *IoK8sApiCoreV1PodDNSConfig`

NewIoK8sApiCoreV1PodDNSConfigWithDefaults instantiates a new IoK8sApiCoreV1PodDNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameservers

`func (o *IoK8sApiCoreV1PodDNSConfig) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *IoK8sApiCoreV1PodDNSConfig) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *IoK8sApiCoreV1PodDNSConfig) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *IoK8sApiCoreV1PodDNSConfig) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetOptions

`func (o *IoK8sApiCoreV1PodDNSConfig) GetOptions() []IoK8sApiCoreV1PodDNSConfigOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IoK8sApiCoreV1PodDNSConfig) GetOptionsOk() (*[]IoK8sApiCoreV1PodDNSConfigOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IoK8sApiCoreV1PodDNSConfig) SetOptions(v []IoK8sApiCoreV1PodDNSConfigOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IoK8sApiCoreV1PodDNSConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSearches

`func (o *IoK8sApiCoreV1PodDNSConfig) GetSearches() []string`

GetSearches returns the Searches field if non-nil, zero value otherwise.

### GetSearchesOk

`func (o *IoK8sApiCoreV1PodDNSConfig) GetSearchesOk() (*[]string, bool)`

GetSearchesOk returns a tuple with the Searches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearches

`func (o *IoK8sApiCoreV1PodDNSConfig) SetSearches(v []string)`

SetSearches sets Searches field to given value.

### HasSearches

`func (o *IoK8sApiCoreV1PodDNSConfig) HasSearches() bool`

HasSearches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


