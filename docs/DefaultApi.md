# \DefaultApi

All URIs are relative to *https://api.tailscale.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TailnetTailnetNameAclGet**](DefaultApi.md#TailnetTailnetNameAclGet) | **Get** /tailnet/{tailnet_name}/acl | 
[**TailnetTailnetNameAclPost**](DefaultApi.md#TailnetTailnetNameAclPost) | **Post** /tailnet/{tailnet_name}/acl | 



## TailnetTailnetNameAclGet

> TailnetACL TailnetTailnetNameAclGet(ctx, tailnetName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tailnetName := "tailnetName_example" // string | The name of the Tailnet

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TailnetTailnetNameAclGet(context.Background(), tailnetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TailnetTailnetNameAclGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailnetTailnetNameAclGet`: TailnetACL
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TailnetTailnetNameAclGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tailnetName** | **string** | The name of the Tailnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailnetTailnetNameAclGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TailnetACL**](TailnetACL.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TailnetTailnetNameAclPost

> TailnetACL TailnetTailnetNameAclPost(ctx, tailnetName).IfMatch(ifMatch).TailnetACL(tailnetACL).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tailnetName := "tailnetName_example" // string | The name of the Tailnet
    ifMatch := "ifMatch_example" // string | Set this value to the ETag header provided in an ACL GET request to avoid missed updates. (optional)
    tailnetACL := *openapiclient.NewTailnetACL() // TailnetACL |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TailnetTailnetNameAclPost(context.Background(), tailnetName).IfMatch(ifMatch).TailnetACL(tailnetACL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TailnetTailnetNameAclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TailnetTailnetNameAclPost`: TailnetACL
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TailnetTailnetNameAclPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tailnetName** | **string** | The name of the Tailnet | 

### Other Parameters

Other parameters are passed through a pointer to a apiTailnetTailnetNameAclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Set this value to the ETag header provided in an ACL GET request to avoid missed updates. | 
 **tailnetACL** | [**TailnetACL**](TailnetACL.md) |  | 

### Return type

[**TailnetACL**](TailnetACL.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

