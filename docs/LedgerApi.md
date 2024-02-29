# \LedgerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMetadataOnTransaction**](LedgerApi.md#AddMetadataOnTransaction) | **Post** /{ledger}/transactions/{txid}/metadata | Set the metadata of a transaction by its ID
[**AddMetadataToAccount**](LedgerApi.md#AddMetadataToAccount) | **Post** /{ledger}/accounts/{address}/metadata | Add metadata to an account
[**CountAccounts**](LedgerApi.md#CountAccounts) | **Head** /{ledger}/accounts | Count the accounts from a ledger
[**CountTransactions**](LedgerApi.md#CountTransactions) | **Head** /{ledger}/transactions | Count the transactions from a ledger
[**CreateTransaction**](LedgerApi.md#CreateTransaction) | **Post** /{ledger}/transactions | Create a new transaction to a ledger
[**CreateTransactions**](LedgerApi.md#CreateTransactions) | **Post** /{ledger}/transactions/batch | Create a new batch of transactions to a ledger
[**GetAccount**](LedgerApi.md#GetAccount) | **Get** /{ledger}/accounts/{address} | Get account by its address
[**GetBalances**](LedgerApi.md#GetBalances) | **Get** /{ledger}/balances | Get the balances from a ledger&#39;s account
[**GetBalancesAggregated**](LedgerApi.md#GetBalancesAggregated) | **Get** /{ledger}/aggregate/balances | Get the aggregated balances from selected accounts
[**GetInfo**](LedgerApi.md#GetInfo) | **Get** /_info | Show server information
[**GetLedgerInfo**](LedgerApi.md#GetLedgerInfo) | **Get** /{ledger}/_info | Get information about a ledger
[**GetMapping**](LedgerApi.md#GetMapping) | **Get** /{ledger}/mapping | Get the mapping of a ledger
[**GetTransaction**](LedgerApi.md#GetTransaction) | **Get** /{ledger}/transactions/{txid} | Get transaction from a ledger by its ID
[**ListAccounts**](LedgerApi.md#ListAccounts) | **Get** /{ledger}/accounts | List accounts from a ledger
[**ListLogs**](LedgerApi.md#ListLogs) | **Get** /{ledger}/logs | List the logs from a ledger
[**ListTransactions**](LedgerApi.md#ListTransactions) | **Get** /{ledger}/transactions | List transactions from a ledger
[**ReadStats**](LedgerApi.md#ReadStats) | **Get** /{ledger}/stats | Get statistics from a ledger
[**RevertTransaction**](LedgerApi.md#RevertTransaction) | **Post** /{ledger}/transactions/{txid}/revert | Revert a ledger transaction by its ID
[**RunScript**](LedgerApi.md#RunScript) | **Post** /{ledger}/script | Execute a Numscript
[**UpdateMapping**](LedgerApi.md#UpdateMapping) | **Put** /{ledger}/mapping | Update the mapping of a ledger



## AddMetadataOnTransaction

> AddMetadataOnTransaction(ctx, ledger, txid).RequestBody(requestBody).Execute()

Set the metadata of a transaction by its ID

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
    ledger := "ledger001" // string | Name of the ledger.
    txid := int64(1234) // int64 | Transaction ID.
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | metadata (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.AddMetadataOnTransaction(context.Background(), ledger, txid).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.AddMetadataOnTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 
**txid** | **int64** | Transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataOnTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** | metadata | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMetadataToAccount

> AddMetadataToAccount(ctx, ledger, address).RequestBody(requestBody).Execute()

Add metadata to an account

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
    ledger := "ledger001" // string | Name of the ledger.
    address := "users:001" // string | Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | metadata

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.AddMetadataToAccount(context.Background(), ledger, address).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.AddMetadataToAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 
**address** | **string** | Exact address of the account. It must match the following regular expressions pattern: &#x60;&#x60;&#x60; ^\\w+(:\\w+)*$ &#x60;&#x60;&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataToAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** | metadata | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountAccounts

> CountAccounts(ctx, ledger).Address(address).Metadata(metadata).Execute()

Count the accounts from a ledger

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
    ledger := "ledger001" // string | Name of the ledger.
    address := "users:.+" // string | Filter accounts by address pattern (regular expression placed between ^ and $). (optional)
    metadata := map[string]interface{}{"key": map[string]interface{}(123)} // map[string]interface{} | Filter accounts by metadata key value pairs. The filter can be used like this metadata[key]=value1&metadata[a.nested.key]=value2 (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.CountAccounts(context.Background(), ledger).Address(address).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.CountAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **string** | Filter accounts by address pattern (regular expression placed between ^ and $). | 
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Filter accounts by metadata key value pairs. The filter can be used like this metadata[key]&#x3D;value1&amp;metadata[a.nested.key]&#x3D;value2 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountTransactions

> CountTransactions(ctx, ledger).Reference(reference).Account(account).Source(source).Destination(destination).StartTime(startTime).StartTime2(startTime2).EndTime(endTime).EndTime2(endTime2).Metadata(metadata).Execute()

Count the transactions from a ledger

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    client "github.com/numary/numary-go"
)

func main() {
    ledger := "ledger001" // string | Name of the ledger.
    reference := "ref:001" // string | Filter transactions by reference field. (optional)
    account := "users:001" // string | Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $). (optional)
    source := "users:001" // string | Filter transactions with postings involving given account at source (regular expression placed between ^ and $). (optional)
    destination := "users:001" // string | Filter transactions with postings involving given account at destination (regular expression placed between ^ and $). (optional)
    startTime := time.Now() // time.Time | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \"2023-01-02T15:04:01Z\" includes the first second of 4th minute).  (optional)
    startTime2 := time.Now() // time.Time | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \"2023-01-02T15:04:01Z\" includes the first second of 4th minute). Deprecated, please use `startTime` instead.  (optional)
    endTime := time.Now() // time.Time | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \"2023-01-02T15:04:01Z\" excludes the first second of 4th minute).  (optional)
    endTime2 := time.Now() // time.Time | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \"2023-01-02T15:04:01Z\" excludes the first second of 4th minute). Deprecated, please use `endTime` instead.  (optional)
    metadata := map[string]interface{}{"key": map[string]interface{}(123)} // map[string]interface{} | Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below. (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.CountTransactions(context.Background(), ledger).Reference(reference).Account(account).Source(source).Destination(destination).StartTime(startTime).StartTime2(startTime2).EndTime(endTime).EndTime2(endTime2).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.CountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reference** | **string** | Filter transactions by reference field. | 
 **account** | **string** | Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $). | 
 **source** | **string** | Filter transactions with postings involving given account at source (regular expression placed between ^ and $). | 
 **destination** | **string** | Filter transactions with postings involving given account at destination (regular expression placed between ^ and $). | 
 **startTime** | **time.Time** | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute).  | 
 **startTime2** | **time.Time** | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute). Deprecated, please use &#x60;startTime&#x60; instead.  | 
 **endTime** | **time.Time** | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute).  | 
 **endTime2** | **time.Time** | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute). Deprecated, please use &#x60;endTime&#x60; instead.  | 
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransaction

> TransactionsResponse CreateTransaction(ctx, ledger).PostTransaction(postTransaction).Preview(preview).Execute()

Create a new transaction to a ledger

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
    ledger := "ledger001" // string | Name of the ledger.
    postTransaction := *client.NewPostTransaction() // PostTransaction | The request body must contain at least one of the following objects:   - `postings`: suitable for simple transactions   - `script`: enabling more complex transactions with Numscript 
    preview := true // bool | Set the preview mode. Preview mode doesn't add the logs to the database or publish a message to the message broker. (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.CreateTransaction(context.Background(), ledger).PostTransaction(postTransaction).Preview(preview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.CreateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransaction`: TransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.CreateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postTransaction** | [**PostTransaction**](PostTransaction.md) | The request body must contain at least one of the following objects:   - &#x60;postings&#x60;: suitable for simple transactions   - &#x60;script&#x60;: enabling more complex transactions with Numscript  | 
 **preview** | **bool** | Set the preview mode. Preview mode doesn&#39;t add the logs to the database or publish a message to the message broker. | 

### Return type

[**TransactionsResponse**](TransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransactions

> TransactionsResponse CreateTransactions(ctx, ledger).Transactions(transactions).Execute()

Create a new batch of transactions to a ledger

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
    ledger := "ledger001" // string | Name of the ledger.
    transactions := *client.NewTransactions([]client.TransactionData{*client.NewTransactionData([]client.Posting{*client.NewPosting(int64(100), "COIN", "users:002", "users:001")})}) // Transactions | 

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.CreateTransactions(context.Background(), ledger).Transactions(transactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.CreateTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransactions`: TransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.CreateTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactions** | [**Transactions**](Transactions.md) |  | 

### Return type

[**TransactionsResponse**](TransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> AccountResponse GetAccount(ctx, ledger, address).Execute()

Get account by its address

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
    ledger := "ledger001" // string | Name of the ledger.
    address := "users:001" // string | Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetAccount(context.Background(), ledger, address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 
**address** | **string** | Exact address of the account. It must match the following regular expressions pattern: &#x60;&#x60;&#x60; ^\\w+(:\\w+)*$ &#x60;&#x60;&#x60;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalances

> BalancesCursorResponse GetBalances(ctx, ledger).Address(address).After(after).Cursor(cursor).PaginationToken(paginationToken).Execute()

Get the balances from a ledger's account

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
    ledger := "ledger001" // string | Name of the ledger.
    address := "users:001" // string | Filter balances involving given account, either as source or destination. (optional)
    after := "users:003" // string | Pagination cursor, will return accounts after given address, in descending order. (optional)
    cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  (optional)
    paginationToken := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. Deprecated, please use `cursor` instead. (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetBalances(context.Background(), ledger).Address(address).After(after).Cursor(cursor).PaginationToken(paginationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalances`: BalancesCursorResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **string** | Filter balances involving given account, either as source or destination. | 
 **after** | **string** | Pagination cursor, will return accounts after given address, in descending order. | 
 **cursor** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  | 
 **paginationToken** | **string** | Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. Deprecated, please use &#x60;cursor&#x60; instead. | 

### Return type

[**BalancesCursorResponse**](BalancesCursorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancesAggregated

> AggregateBalancesResponse GetBalancesAggregated(ctx, ledger).Address(address).Execute()

Get the aggregated balances from selected accounts

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
    ledger := "ledger001" // string | Name of the ledger.
    address := "users:001" // string | Filter balances involving given account, either as source or destination. (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetBalancesAggregated(context.Background(), ledger).Address(address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetBalancesAggregated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancesAggregated`: AggregateBalancesResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetBalancesAggregated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancesAggregatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **string** | Filter balances involving given account, either as source or destination. | 

### Return type

[**AggregateBalancesResponse**](AggregateBalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> ConfigInfoResponse GetInfo(ctx).Execute()

Show server information

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

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfo`: ConfigInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoRequest struct via the builder pattern


### Return type

[**ConfigInfoResponse**](ConfigInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLedgerInfo

> LedgerInfoResponse GetLedgerInfo(ctx, ledger).Execute()

Get information about a ledger

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
    ledger := "ledger001" // string | Name of the ledger.

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetLedgerInfo(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetLedgerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLedgerInfo`: LedgerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetLedgerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LedgerInfoResponse**](LedgerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapping

> MappingResponse GetMapping(ctx, ledger).Execute()

Get the mapping of a ledger

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
    ledger := "ledger001" // string | Name of the ledger.

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetMapping(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapping`: MappingResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MappingResponse**](MappingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> TransactionResponse GetTransaction(ctx, ledger, txid).Execute()

Get transaction from a ledger by its ID

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
    ledger := "ledger001" // string | Name of the ledger.
    txid := int64(1234) // int64 | Transaction ID.

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.GetTransaction(context.Background(), ledger, txid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: TransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 
**txid** | **int64** | Transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountsCursorResponse ListAccounts(ctx, ledger).PageSize(pageSize).PageSize2(pageSize2).After(after).Address(address).Metadata(metadata).Balance(balance).BalanceAsset(balanceAsset).BalanceOperator(balanceOperator).BalanceOperator2(balanceOperator2).Cursor(cursor).PaginationToken(paginationToken).Execute()

List accounts from a ledger



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
    ledger := "ledger001" // string | Name of the ledger.
    pageSize := int64(100) // int64 | The maximum number of results to return per page.  (optional) (default to 15)
    pageSize2 := int64(100) // int64 | The maximum number of results to return per page. Deprecated, please use `pageSize` instead.  (optional) (default to 15)
    after := "users:003" // string | Pagination cursor, will return accounts after given address, in descending order. (optional)
    address := "users:.+" // string | Filter accounts by address pattern (regular expression placed between ^ and $). (optional)
    metadata := map[string]interface{}{"key": map[string]interface{}(123)} // map[string]interface{} | Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below. (optional)
    balance := int64(2400) // int64 | Filter accounts by their balance (default operator is gte) (optional)
    balanceAsset := "balanceAsset_example" // string | Filter accounts by their balance asset (optional)
    balanceOperator := "gte" // string | Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.  (optional)
    balanceOperator2 := "gte" // string | Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not. Deprecated, please use `balanceOperator` instead.  (optional)
    cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  (optional)
    paginationToken := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use `cursor` instead.  (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.ListAccounts(context.Background(), ledger).PageSize(pageSize).PageSize2(pageSize2).After(after).Address(address).Metadata(metadata).Balance(balance).BalanceAsset(balanceAsset).BalanceOperator(balanceOperator).BalanceOperator2(balanceOperator2).Cursor(cursor).PaginationToken(paginationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountsCursorResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The maximum number of results to return per page.  | [default to 15]
 **pageSize2** | **int64** | The maximum number of results to return per page. Deprecated, please use &#x60;pageSize&#x60; instead.  | [default to 15]
 **after** | **string** | Pagination cursor, will return accounts after given address, in descending order. | 
 **address** | **string** | Filter accounts by address pattern (regular expression placed between ^ and $). | 
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below. | 
 **balance** | **int64** | Filter accounts by their balance (default operator is gte) | 
 **balanceAsset** | **string** | Filter accounts by their balance asset | 
 **balanceOperator** | **string** | Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.  | 
 **balanceOperator2** | **string** | Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not. Deprecated, please use &#x60;balanceOperator&#x60; instead.  | 
 **cursor** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  | 
 **paginationToken** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use &#x60;cursor&#x60; instead.  | 

### Return type

[**AccountsCursorResponse**](AccountsCursorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogs

> LogsCursorResponse ListLogs(ctx, ledger).PageSize(pageSize).PageSize2(pageSize2).After(after).StartTime(startTime).StartTime2(startTime2).EndTime(endTime).EndTime2(endTime2).Cursor(cursor).PaginationToken(paginationToken).Execute()

List the logs from a ledger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    client "github.com/numary/numary-go"
)

func main() {
    ledger := "ledger001" // string | Name of the ledger.
    pageSize := int64(100) // int64 | The maximum number of results to return per page.  (optional) (default to 15)
    pageSize2 := int64(100) // int64 | The maximum number of results to return per page. Deprecated, please use `pageSize` instead.  (optional) (default to 15)
    after := "1234" // string | Pagination cursor, will return the logs after a given ID. (in descending order). (optional)
    startTime := time.Now() // time.Time | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \"2023-01-02T15:04:01Z\" includes the first second of 4th minute).  (optional)
    startTime2 := time.Now() // time.Time | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \"2023-01-02T15:04:01Z\" includes the first second of 4th minute). Deprecated, please use `startTime` instead.  (optional)
    endTime := time.Now() // time.Time | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \"2023-01-02T15:04:01Z\" excludes the first second of 4th minute).  (optional)
    endTime2 := time.Now() // time.Time | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \"2023-01-02T15:04:01Z\" excludes the first second of 4th minute). Deprecated, please use `endTime` instead.  (optional)
    cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  (optional)
    paginationToken := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use `cursor` instead.  (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.ListLogs(context.Background(), ledger).PageSize(pageSize).PageSize2(pageSize2).After(after).StartTime(startTime).StartTime2(startTime2).EndTime(endTime).EndTime2(endTime2).Cursor(cursor).PaginationToken(paginationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.ListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: LogsCursorResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.ListLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The maximum number of results to return per page.  | [default to 15]
 **pageSize2** | **int64** | The maximum number of results to return per page. Deprecated, please use &#x60;pageSize&#x60; instead.  | [default to 15]
 **after** | **string** | Pagination cursor, will return the logs after a given ID. (in descending order). | 
 **startTime** | **time.Time** | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute).  | 
 **startTime2** | **time.Time** | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute). Deprecated, please use &#x60;startTime&#x60; instead.  | 
 **endTime** | **time.Time** | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute).  | 
 **endTime2** | **time.Time** | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute). Deprecated, please use &#x60;endTime&#x60; instead.  | 
 **cursor** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  | 
 **paginationToken** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use &#x60;cursor&#x60; instead.  | 

### Return type

[**LogsCursorResponse**](LogsCursorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionsCursorResponse ListTransactions(ctx, ledger).PageSize(pageSize).PageSize2(pageSize2).After(after).Reference(reference).Account(account).Source(source).Destination(destination).StartTime(startTime).StartTime2(startTime2).EndTime(endTime).EndTime2(endTime2).Cursor(cursor).PaginationToken(paginationToken).Metadata(metadata).Execute()

List transactions from a ledger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    client "github.com/numary/numary-go"
)

func main() {
    ledger := "ledger001" // string | Name of the ledger.
    pageSize := int64(100) // int64 | The maximum number of results to return per page.  (optional) (default to 15)
    pageSize2 := int64(100) // int64 | The maximum number of results to return per page. Deprecated, please use `pageSize` instead.  (optional) (default to 15)
    after := "1234" // string | Pagination cursor, will return transactions after given txid (in descending order). (optional)
    reference := "ref:001" // string | Find transactions by reference field. (optional)
    account := "users:001" // string | Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $). (optional)
    source := "users:001" // string | Filter transactions with postings involving given account at source (regular expression placed between ^ and $). (optional)
    destination := "users:001" // string | Filter transactions with postings involving given account at destination (regular expression placed between ^ and $). (optional)
    startTime := time.Now() // time.Time | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \"2023-01-02T15:04:01Z\" includes the first second of 4th minute).  (optional)
    startTime2 := time.Now() // time.Time | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \"2023-01-02T15:04:01Z\" includes the first second of 4th minute). Deprecated, please use `startTime` instead.  (optional)
    endTime := time.Now() // time.Time | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \"2023-01-02T15:04:01Z\" excludes the first second of 4th minute).  (optional)
    endTime2 := time.Now() // time.Time | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \"2023-01-02T15:04:01Z\" excludes the first second of 4th minute). Deprecated, please use `endTime` instead.  (optional)
    cursor := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  (optional)
    paginationToken := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use `cursor` instead.  (optional)
    metadata := map[string]interface{}{"key": map[string]interface{}(123)} // map[string]interface{} | Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below. (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.ListTransactions(context.Background(), ledger).PageSize(pageSize).PageSize2(pageSize2).After(after).Reference(reference).Account(account).Source(source).Destination(destination).StartTime(startTime).StartTime2(startTime2).EndTime(endTime).EndTime2(endTime2).Cursor(cursor).PaginationToken(paginationToken).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: TransactionsCursorResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.ListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The maximum number of results to return per page.  | [default to 15]
 **pageSize2** | **int64** | The maximum number of results to return per page. Deprecated, please use &#x60;pageSize&#x60; instead.  | [default to 15]
 **after** | **string** | Pagination cursor, will return transactions after given txid (in descending order). | 
 **reference** | **string** | Find transactions by reference field. | 
 **account** | **string** | Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $). | 
 **source** | **string** | Filter transactions with postings involving given account at source (regular expression placed between ^ and $). | 
 **destination** | **string** | Filter transactions with postings involving given account at destination (regular expression placed between ^ and $). | 
 **startTime** | **time.Time** | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute).  | 
 **startTime2** | **time.Time** | Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute). Deprecated, please use &#x60;startTime&#x60; instead.  | 
 **endTime** | **time.Time** | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute).  | 
 **endTime2** | **time.Time** | Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute). Deprecated, please use &#x60;endTime&#x60; instead.  | 
 **cursor** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.  | 
 **paginationToken** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use &#x60;cursor&#x60; instead.  | 
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below. | 

### Return type

[**TransactionsCursorResponse**](TransactionsCursorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadStats

> StatsResponse ReadStats(ctx, ledger).Execute()

Get statistics from a ledger



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
    ledger := "ledger001" // string | name of the ledger

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.ReadStats(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.ReadStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadStats`: StatsResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.ReadStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | name of the ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatsResponse**](StatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertTransaction

> TransactionResponse RevertTransaction(ctx, ledger, txid).DisableChecks(disableChecks).Execute()

Revert a ledger transaction by its ID

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
    ledger := "ledger001" // string | Name of the ledger.
    txid := int64(1234) // int64 | Transaction ID.
    disableChecks := true // bool | Allow to disable balances checks (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.RevertTransaction(context.Background(), ledger, txid).DisableChecks(disableChecks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.RevertTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertTransaction`: TransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.RevertTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 
**txid** | **int64** | Transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableChecks** | **bool** | Allow to disable balances checks | 

### Return type

[**TransactionResponse**](TransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunScript

> ScriptResponse RunScript(ctx, ledger).Script(script).Preview(preview).Execute()

Execute a Numscript



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
    ledger := "ledger001" // string | Name of the ledger.
    script := *client.NewScript("vars {
account $user
}
send [COIN 10] (
	source = @world
	destination = $user
)
") // Script | 
    preview := true // bool | Set the preview mode. Preview mode doesn't add the logs to the database or publish a message to the message broker. (optional)

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.RunScript(context.Background(), ledger).Script(script).Preview(preview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.RunScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunScript`: ScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.RunScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **script** | [**Script**](Script.md) |  | 
 **preview** | **bool** | Set the preview mode. Preview mode doesn&#39;t add the logs to the database or publish a message to the message broker. | 

### Return type

[**ScriptResponse**](ScriptResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMapping

> MappingResponse UpdateMapping(ctx, ledger).Mapping(mapping).Execute()

Update the mapping of a ledger

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
    ledger := "ledger001" // string | Name of the ledger.
    mapping := *client.NewMapping([]client.Contract{*client.NewContract(map[string]interface{}(123))}) // Mapping | 

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.UpdateMapping(context.Background(), ledger).Mapping(mapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.UpdateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMapping`: MappingResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.UpdateMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Name of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapping** | [**Mapping**](Mapping.md) |  | 

### Return type

[**MappingResponse**](MappingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

