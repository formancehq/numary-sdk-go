# \MappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMapping**](MappingApi.md#GetMapping) | **Get** /{ledger}/mapping | Get mapping
[**UpdateMapping**](MappingApi.md#UpdateMapping) | **Put** /{ledger}/mapping | Put mapping



## GetMapping

> MappingResponse GetMapping(ctx, ledger).Execute()

Get mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "github.com/numary/numary-go"
)

func main() {
    ledger := "ledger_example" // string | ledger

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.MappingApi.GetMapping(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappingApi.GetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapping`: MappingResponse
    fmt.Fprintf(os.Stdout, "Response from `MappingApi.GetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MappingResponse**](MappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMapping

> MappingResponse UpdateMapping(ctx, ledger).Mapping(mapping).Execute()

Put mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "github.com/numary/numary-go"
)

func main() {
    ledger := "ledger_example" // string | ledger
    mapping := *client.NewMapping([]client.Contract{*client.NewContract(map[string]interface{}(123))}) // Mapping | mapping

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.MappingApi.UpdateMapping(context.Background(), ledger).Mapping(mapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappingApi.UpdateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMapping`: MappingResponse
    fmt.Fprintf(os.Stdout, "Response from `MappingApi.UpdateMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapping** | [**Mapping**](Mapping.md) | mapping | 

### Return type

[**MappingResponse**](MappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

