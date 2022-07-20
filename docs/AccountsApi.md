# \AccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMetadataToAccount**](AccountsApi.md#AddMetadataToAccount) | **Post** /{ledger}/accounts/{accountId}/metadata | Add metadata to account
[**CountAccounts**](AccountsApi.md#CountAccounts) | **Head** /{ledger}/accounts | Count accounts
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /{ledger}/accounts/{accountId} | Get account by address
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /{ledger}/accounts | List all accounts



## AddMetadataToAccount

> AddMetadataToAccount(ctx, ledger, accountId).RequestBody(requestBody).Execute()

Add metadata to account

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
    accountId := "accountId_example" // string | accountId
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | metadata

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AddMetadataToAccount(context.Background(), ledger, accountId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AddMetadataToAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 
**accountId** | **string** | accountId | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataToAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** | metadata | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountAccounts

> CountAccounts(ctx, ledger).After(after).Address(address).Metadata(metadata).Execute()

Count accounts

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
    after := "after_example" // string | pagination cursor, will return accounts after given address (in descending order) (optional)
    address := "address_example" // string | account address (optional)
    metadata := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | metadata (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CountAccounts(context.Background(), ledger).After(after).Address(address).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CountAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | pagination cursor, will return accounts after given address (in descending order) | 
 **address** | **string** | account address | 
 **metadata** | **map[string]map[string]string** | metadata | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> AccountResponse GetAccount(ctx, ledger, accountId).Execute()

Get account by address

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
    accountId := "accountId_example" // string | accountId

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), ledger, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 
**accountId** | **string** | accountId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountCursorResponse ListAccounts(ctx, ledger).After(after).Address(address).Metadata(metadata).Execute()

List all accounts

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
    after := "after_example" // string | pagination cursor, will return accounts after given address (in descending order) (optional)
    address := "address_example" // string | account address (optional)
    metadata := map[string]string{"key": map[string]string{"key": "Inner_example"}} // map[string]string | account address (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccounts(context.Background(), ledger).After(after).Address(address).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountCursorResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | pagination cursor, will return accounts after given address (in descending order) | 
 **address** | **string** | account address | 
 **metadata** | **map[string]map[string]string** | account address | 

### Return type

[**AccountCursorResponse**](AccountCursorResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

