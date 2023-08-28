# \ApisApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIVersions**](ApisApi.md#GetAPIVersions) | **Get** /apis/ | 



## GetAPIVersions

> IoK8sApimachineryPkgApisMetaV1APIGroupList GetAPIVersions(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApisApi.GetAPIVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApisApi.GetAPIVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPIVersions`: IoK8sApimachineryPkgApisMetaV1APIGroupList
    fmt.Fprintf(os.Stdout, "Response from `ApisApi.GetAPIVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIVersionsRequest struct via the builder pattern


### Return type

[**IoK8sApimachineryPkgApisMetaV1APIGroupList**](IoK8sApimachineryPkgApisMetaV1APIGroupList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

