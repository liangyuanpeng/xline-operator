# \OpenidApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceAccountIssuerOpenIDKeyset**](OpenidApi.md#GetServiceAccountIssuerOpenIDKeyset) | **Get** /openid/v1/jwks/ | 



## GetServiceAccountIssuerOpenIDKeyset

> string GetServiceAccountIssuerOpenIDKeyset(ctx).Execute()





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
    resp, r, err := apiClient.OpenidApi.GetServiceAccountIssuerOpenIDKeyset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenidApi.GetServiceAccountIssuerOpenIDKeyset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccountIssuerOpenIDKeyset`: string
    fmt.Fprintf(os.Stdout, "Response from `OpenidApi.GetServiceAccountIssuerOpenIDKeyset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountIssuerOpenIDKeysetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/jwk-set+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

