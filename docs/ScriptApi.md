# \ScriptApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunScript**](ScriptApi.md#RunScript) | **Post** /{ledger}/script | Execute Numscript



## RunScript

> ScriptResult RunScript(ctx, ledger).Script(script).Execute()

Execute Numscript



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
    script := *client.NewScript("Plain_example") // Script | script

    configuration := client.NewConfiguration()
    api_client := client.NewAPIClient(configuration)
    resp, r, err := api_client.ScriptApi.RunScript(context.Background(), ledger).Script(script).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScriptApi.RunScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunScript`: ScriptResult
    fmt.Fprintf(os.Stdout, "Response from `ScriptApi.RunScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | ledger | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **script** | [**Script**](Script.md) | script | 

### Return type

[**ScriptResult**](ScriptResult.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cloudToken](../README.md#cloudToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

