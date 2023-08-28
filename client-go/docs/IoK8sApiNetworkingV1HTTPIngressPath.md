# IoK8sApiNetworkingV1HTTPIngressPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | [**IoK8sApiNetworkingV1IngressBackend**](IoK8sApiNetworkingV1IngressBackend.md) |  | 
**Path** | Pointer to **string** | path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \&quot;path\&quot; part of a URL as defined by RFC 3986. Paths must begin with a &#39;/&#39; and must be present when using PathType with value \&quot;Exact\&quot; or \&quot;Prefix\&quot;. | [optional] 
**PathType** | **string** | pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by &#39;/&#39;. Matching is   done on a path element by element basis. A path element refers is the   list of labels in the path split by the &#39;/&#39; separator. A request is a   match for path p if every p is an element-wise prefix of p of the   request path. Note that if the last element of the path is a substring   of the last element in request path, it is not a match (e.g. /foo/bar   matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to   the IngressClass. Implementations can treat this as a separate PathType   or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.  Possible enum values:  - &#x60;\&quot;Exact\&quot;&#x60; matches the URL path exactly and with case sensitivity.  - &#x60;\&quot;ImplementationSpecific\&quot;&#x60; matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types.  - &#x60;\&quot;Prefix\&quot;&#x60; matches based on a URL path prefix split by &#39;/&#39;. Matching is case sensitive and done on a path element by element basis. A path element refers to the list of labels in the path split by the &#39;/&#39; separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). If multiple matching paths exist in an Ingress spec, the longest matching path is given priority. Examples: - /foo/bar does not match requests to /foo/barbaz - /foo/bar matches request to /foo/bar and /foo/bar/baz - /foo and /foo/ both match requests to /foo and /foo/. If both paths are present in an Ingress spec, the longest matching path (/foo/) is given priority. | 

## Methods

### NewIoK8sApiNetworkingV1HTTPIngressPath

`func NewIoK8sApiNetworkingV1HTTPIngressPath(backend IoK8sApiNetworkingV1IngressBackend, pathType string, ) *IoK8sApiNetworkingV1HTTPIngressPath`

NewIoK8sApiNetworkingV1HTTPIngressPath instantiates a new IoK8sApiNetworkingV1HTTPIngressPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiNetworkingV1HTTPIngressPathWithDefaults

`func NewIoK8sApiNetworkingV1HTTPIngressPathWithDefaults() *IoK8sApiNetworkingV1HTTPIngressPath`

NewIoK8sApiNetworkingV1HTTPIngressPathWithDefaults instantiates a new IoK8sApiNetworkingV1HTTPIngressPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) GetBackend() IoK8sApiNetworkingV1IngressBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) GetBackendOk() (*IoK8sApiNetworkingV1IngressBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) SetBackend(v IoK8sApiNetworkingV1IngressBackend)`

SetBackend sets Backend field to given value.


### GetPath

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPathType

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) GetPathType() string`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) GetPathTypeOk() (*string, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *IoK8sApiNetworkingV1HTTPIngressPath) SetPathType(v string)`

SetPathType sets PathType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


