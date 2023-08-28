# \CoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCoreAPIVersions**](CoreApi.md#GetCoreAPIVersions) | **Get** /api/ | 



## GetCoreAPIVersions

> IoK8sApimachineryPkgApisMetaV1APIVersions GetCoreAPIVersions(ctx).Execute()





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
    resp, r, err := apiClient.CoreApi.GetCoreAPIVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreApi.GetCoreAPIVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCoreAPIVersions`: IoK8sApimachineryPkgApisMetaV1APIVersions
    fmt.Fprintf(os.Stdout, "Response from `CoreApi.GetCoreAPIVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoreAPIVersionsRequest struct via the builder pattern


### Return type

[**IoK8sApimachineryPkgApisMetaV1APIVersions**](IoK8sApimachineryPkgApisMetaV1APIVersions.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

