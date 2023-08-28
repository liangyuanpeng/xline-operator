# \AuthorizationV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorizationV1NamespacedLocalSubjectAccessReview**](AuthorizationV1Api.md#CreateAuthorizationV1NamespacedLocalSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1/namespaces/{namespace}/localsubjectaccessreviews | 
[**CreateAuthorizationV1SelfSubjectAccessReview**](AuthorizationV1Api.md#CreateAuthorizationV1SelfSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1/selfsubjectaccessreviews | 
[**CreateAuthorizationV1SelfSubjectRulesReview**](AuthorizationV1Api.md#CreateAuthorizationV1SelfSubjectRulesReview) | **Post** /apis/authorization.k8s.io/v1/selfsubjectrulesreviews | 
[**CreateAuthorizationV1SubjectAccessReview**](AuthorizationV1Api.md#CreateAuthorizationV1SubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1/subjectaccessreviews | 
[**GetAuthorizationV1APIResources**](AuthorizationV1Api.md#GetAuthorizationV1APIResources) | **Get** /apis/authorization.k8s.io/v1/ | 



## CreateAuthorizationV1NamespacedLocalSubjectAccessReview

> IoK8sApiAuthorizationV1LocalSubjectAccessReview CreateAuthorizationV1NamespacedLocalSubjectAccessReview(ctx, namespace).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()





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
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    body := *openapiclient.NewIoK8sApiAuthorizationV1LocalSubjectAccessReview(*openapiclient.NewIoK8sApiAuthorizationV1SubjectAccessReviewSpec()) // IoK8sApiAuthorizationV1LocalSubjectAccessReview | 
    dryRun := "dryRun_example" // string | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed (optional)
    fieldManager := "fieldManager_example" // string | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. (optional)
    fieldValidation := "fieldValidation_example" // string | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. (optional)
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationV1Api.CreateAuthorizationV1NamespacedLocalSubjectAccessReview(context.Background(), namespace).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationV1Api.CreateAuthorizationV1NamespacedLocalSubjectAccessReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationV1NamespacedLocalSubjectAccessReview`: IoK8sApiAuthorizationV1LocalSubjectAccessReview
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationV1Api.CreateAuthorizationV1NamespacedLocalSubjectAccessReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationV1NamespacedLocalSubjectAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoK8sApiAuthorizationV1LocalSubjectAccessReview**](IoK8sApiAuthorizationV1LocalSubjectAccessReview.md) |  | 
 **dryRun** | **string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **fieldManager** | **string** | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. | 
 **fieldValidation** | **string** | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**IoK8sApiAuthorizationV1LocalSubjectAccessReview**](IoK8sApiAuthorizationV1LocalSubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationV1SelfSubjectAccessReview

> IoK8sApiAuthorizationV1SelfSubjectAccessReview CreateAuthorizationV1SelfSubjectAccessReview(ctx).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()





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
    body := *openapiclient.NewIoK8sApiAuthorizationV1SelfSubjectAccessReview(*openapiclient.NewIoK8sApiAuthorizationV1SelfSubjectAccessReviewSpec()) // IoK8sApiAuthorizationV1SelfSubjectAccessReview | 
    dryRun := "dryRun_example" // string | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed (optional)
    fieldManager := "fieldManager_example" // string | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. (optional)
    fieldValidation := "fieldValidation_example" // string | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. (optional)
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationV1Api.CreateAuthorizationV1SelfSubjectAccessReview(context.Background()).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationV1Api.CreateAuthorizationV1SelfSubjectAccessReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationV1SelfSubjectAccessReview`: IoK8sApiAuthorizationV1SelfSubjectAccessReview
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationV1Api.CreateAuthorizationV1SelfSubjectAccessReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationV1SelfSubjectAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IoK8sApiAuthorizationV1SelfSubjectAccessReview**](IoK8sApiAuthorizationV1SelfSubjectAccessReview.md) |  | 
 **dryRun** | **string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **fieldManager** | **string** | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. | 
 **fieldValidation** | **string** | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**IoK8sApiAuthorizationV1SelfSubjectAccessReview**](IoK8sApiAuthorizationV1SelfSubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationV1SelfSubjectRulesReview

> IoK8sApiAuthorizationV1SelfSubjectRulesReview CreateAuthorizationV1SelfSubjectRulesReview(ctx).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()





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
    body := *openapiclient.NewIoK8sApiAuthorizationV1SelfSubjectRulesReview(*openapiclient.NewIoK8sApiAuthorizationV1SelfSubjectRulesReviewSpec()) // IoK8sApiAuthorizationV1SelfSubjectRulesReview | 
    dryRun := "dryRun_example" // string | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed (optional)
    fieldManager := "fieldManager_example" // string | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. (optional)
    fieldValidation := "fieldValidation_example" // string | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. (optional)
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationV1Api.CreateAuthorizationV1SelfSubjectRulesReview(context.Background()).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationV1Api.CreateAuthorizationV1SelfSubjectRulesReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationV1SelfSubjectRulesReview`: IoK8sApiAuthorizationV1SelfSubjectRulesReview
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationV1Api.CreateAuthorizationV1SelfSubjectRulesReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationV1SelfSubjectRulesReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IoK8sApiAuthorizationV1SelfSubjectRulesReview**](IoK8sApiAuthorizationV1SelfSubjectRulesReview.md) |  | 
 **dryRun** | **string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **fieldManager** | **string** | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. | 
 **fieldValidation** | **string** | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**IoK8sApiAuthorizationV1SelfSubjectRulesReview**](IoK8sApiAuthorizationV1SelfSubjectRulesReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationV1SubjectAccessReview

> IoK8sApiAuthorizationV1SubjectAccessReview CreateAuthorizationV1SubjectAccessReview(ctx).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()





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
    body := *openapiclient.NewIoK8sApiAuthorizationV1SubjectAccessReview(*openapiclient.NewIoK8sApiAuthorizationV1SubjectAccessReviewSpec()) // IoK8sApiAuthorizationV1SubjectAccessReview | 
    dryRun := "dryRun_example" // string | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed (optional)
    fieldManager := "fieldManager_example" // string | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. (optional)
    fieldValidation := "fieldValidation_example" // string | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. (optional)
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationV1Api.CreateAuthorizationV1SubjectAccessReview(context.Background()).Body(body).DryRun(dryRun).FieldManager(fieldManager).FieldValidation(fieldValidation).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationV1Api.CreateAuthorizationV1SubjectAccessReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationV1SubjectAccessReview`: IoK8sApiAuthorizationV1SubjectAccessReview
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationV1Api.CreateAuthorizationV1SubjectAccessReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationV1SubjectAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IoK8sApiAuthorizationV1SubjectAccessReview**](IoK8sApiAuthorizationV1SubjectAccessReview.md) |  | 
 **dryRun** | **string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **fieldManager** | **string** | fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. | 
 **fieldValidation** | **string** | fieldValidation instructs the server on how to handle objects in the request (POST/PUT/PATCH) containing unknown or duplicate fields. Valid values are: - Ignore: This will ignore any unknown fields that are silently dropped from the object, and will ignore all but the last duplicate field that the decoder encounters. This is the default behavior prior to v1.23. - Warn: This will send a warning via the standard warning response header for each unknown field that is dropped from the object, and for each duplicate field that is encountered. The request will still succeed if there are no other errors, and will only persist the last of any duplicate fields. This is the default in v1.23+ - Strict: This will fail the request with a BadRequest error if any unknown fields would be dropped from the object, or if any duplicate fields are present. The error returned from the server will contain all unknown and duplicate fields encountered. | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**IoK8sApiAuthorizationV1SubjectAccessReview**](IoK8sApiAuthorizationV1SubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationV1APIResources

> IoK8sApimachineryPkgApisMetaV1APIResourceList GetAuthorizationV1APIResources(ctx).Execute()





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
    resp, r, err := apiClient.AuthorizationV1Api.GetAuthorizationV1APIResources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationV1Api.GetAuthorizationV1APIResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationV1APIResources`: IoK8sApimachineryPkgApisMetaV1APIResourceList
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationV1Api.GetAuthorizationV1APIResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationV1APIResourcesRequest struct via the builder pattern


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

