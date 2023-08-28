# \AuthenticationV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationV1TokenReview**](AuthenticationV1Api.md#CreateAuthenticationV1TokenReview) | **Post** /apis/authentication.k8s.io/v1/tokenreviews | 
[**GetAuthenticationV1APIResources**](AuthenticationV1Api.md#GetAuthenticationV1APIResources) | **Get** /apis/authentication.k8s.io/v1/ | 



## CreateAuthenticationV1TokenReview

> IoK8sApiAuthenticationV1TokenReview CreateAuthenticationV1TokenReview(ctx).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()





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
    body := *openapiclient.NewIoK8sApiAuthenticationV1TokenReview(*openapiclient.NewIoK8sApiAuthenticationV1TokenReviewSpec()) // IoK8sApiAuthenticationV1TokenReview | 
    dryRun := "dryRun_example" // string | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed (optional)
    fieldManager := "fieldManager_example" // string | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. (optional)
    fieldValidation := "fieldValidation_example" // string | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. (optional)
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationV1Api.CreateAuthenticationV1TokenReview(context.Background()).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationV1Api.CreateAuthenticationV1TokenReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthenticationV1TokenReview`: IoK8sApiAuthenticationV1TokenReview
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationV1Api.CreateAuthenticationV1TokenReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationV1TokenReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IoK8sApiAuthenticationV1TokenReview**](IoK8sApiAuthenticationV1TokenReview.md) |  | 
 **dryRun** | **string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **fieldManager** | **string** | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. | 
 **fieldValidation** | **string** | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**IoK8sApiAuthenticationV1TokenReview**](IoK8sApiAuthenticationV1TokenReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationV1APIResources

> IoK8sApimachineryPkgApisMetaV1APIResourceList GetAuthenticationV1APIResources(ctx).Execute()





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
    resp, r, err := apiClient.AuthenticationV1Api.GetAuthenticationV1APIResources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationV1Api.GetAuthenticationV1APIResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationV1APIResources`: IoK8sApimachineryPkgApisMetaV1APIResourceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationV1Api.GetAuthenticationV1APIResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationV1APIResourcesRequest struct via the builder pattern


### Return type

[**IoK8sApimachineryPkgApisMetaV1APIResourceList**](IoK8sApimachineryPkgApisMetaV1APIResourceList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

