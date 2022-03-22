# \TransactionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMetadataOnTransaction**](TransactionsApi.md#AddMetadataOnTransaction) | **Post** /{ledger}/transactions/{txid}/metadata | Set Transaction Metadata
[**CreateTransaction**](TransactionsApi.md#CreateTransaction) | **Post** /{ledger}/transactions | Create Transaction
[**CreateTransactions**](TransactionsApi.md#CreateTransactions) | **Post** /{ledger}/transactions/batch | Create Transactions Batch
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /{ledger}/transactions/{txid} | Get Transaction
[**ListTransactions**](TransactionsApi.md#ListTransactions) | **Get** /{ledger}/transactions | Get all Transactions
[**RevertTransaction**](TransactionsApi.md#RevertTransaction) | **Post** /{ledger}/transactions/{txid}/revert | Revert Transaction



## AddMetadataOnTransaction

> AddMetadataOnTransaction(ctx, ledger, txid).RequestBody(requestBody).Execute()

Set Transaction Metadata



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
    txid := int32(56) // int32 | txid
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | metadata (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.AddMetadataOnTransaction(context.Background(), ledger, txid).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.AddMetadataOnTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 
**txid** | **int32** | txid | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataOnTransactionRequest struct via the builder pattern


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


## CreateTransaction

> CreateTransactionResponse CreateTransaction(ctx, ledger).TransactionData(transactionData).Preview(preview).Execute()

Create Transaction



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
    transactionData := *client.NewTransactionData([]client.Posting{*client.NewPosting(int32(123), "Asset_example", "Destination_example", "Source_example")}) // TransactionData | transaction
    preview := true // bool | Preview mode (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CreateTransaction(context.Background(), ledger).TransactionData(transactionData).Preview(preview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransaction`: CreateTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionData** | [**TransactionData**](TransactionData.md) | transaction | 
 **preview** | **bool** | Preview mode | 

### Return type

[**CreateTransactionResponse**](CreateTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransactions

> TransactionListResponse CreateTransactions(ctx, ledger).Transactions(transactions).Execute()

Create Transactions Batch



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
    transactions := *client.NewTransactions([]client.TransactionData{*client.NewTransactionData([]client.Posting{*client.NewPosting(int32(123), "Asset_example", "Destination_example", "Source_example")})}) // Transactions | transactions

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CreateTransactions(context.Background(), ledger).Transactions(transactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransactions`: TransactionListResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactions** | [**Transactions**](Transactions.md) | transactions | 

### Return type

[**TransactionListResponse**](TransactionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> TransactionResponse GetTransaction(ctx, ledger, txid).Execute()

Get Transaction



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
    txid := int32(56) // int32 | txid

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetTransaction(context.Background(), ledger, txid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: TransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 
**txid** | **int32** | txid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionCursorResponse ListTransactions(ctx, ledger).After(after).Reference(reference).Account(account).Execute()

Get all Transactions



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
    after := "after_example" // string | pagination cursor, will return transactions after given txid (in descending order) (optional)
    reference := "reference_example" // string | find transactions by reference field (optional)
    account := "account_example" // string | find transactions with postings involving given account, either as source or destination (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.ListTransactions(context.Background(), ledger).After(after).Reference(reference).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: TransactionCursorResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | pagination cursor, will return transactions after given txid (in descending order) | 
 **reference** | **string** | find transactions by reference field | 
 **account** | **string** | find transactions with postings involving given account, either as source or destination | 

### Return type

[**TransactionCursorResponse**](TransactionCursorResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertTransaction

> TransactionResponse RevertTransaction(ctx, ledger, txid).Execute()

Revert Transaction



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
    txid := int32(56) // int32 | txid

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.RevertTransaction(context.Background(), ledger, txid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.RevertTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertTransaction`: TransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.RevertTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 
**txid** | **int32** | txid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

