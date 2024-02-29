# Formance Go SDK

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This SDK was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate SDK.

- API version: v1.11.1
- Package version: v1.11.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

```shell
go get github.com/formancehq/numary-sdk-go
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ledgerclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**AddMetadataToAccount**](docs/AccountsApi.md#addmetadatatoaccount) | **Post** /{ledger}/accounts/{address}/metadata | Add metadata to an account
*AccountsApi* | [**CountAccounts**](docs/AccountsApi.md#countaccounts) | **Head** /{ledger}/accounts | Count the accounts from a ledger
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /{ledger}/accounts/{address} | Get account by its address
*AccountsApi* | [**ListAccounts**](docs/AccountsApi.md#listaccounts) | **Get** /{ledger}/accounts | List accounts from a ledger
*BalancesApi* | [**GetBalances**](docs/BalancesApi.md#getbalances) | **Get** /{ledger}/balances | Get the balances from a ledger&#39;s account
*BalancesApi* | [**GetBalancesAggregated**](docs/BalancesApi.md#getbalancesaggregated) | **Get** /{ledger}/aggregate/balances | Get the aggregated balances from selected accounts
*LedgerApi* | [**AddMetadataOnTransaction**](docs/LedgerApi.md#addmetadataontransaction) | **Post** /{ledger}/transactions/{txid}/metadata | Set the metadata of a transaction by its ID
*LedgerApi* | [**AddMetadataToAccount**](docs/LedgerApi.md#addmetadatatoaccount) | **Post** /{ledger}/accounts/{address}/metadata | Add metadata to an account
*LedgerApi* | [**CountAccounts**](docs/LedgerApi.md#countaccounts) | **Head** /{ledger}/accounts | Count the accounts from a ledger
*LedgerApi* | [**CountTransactions**](docs/LedgerApi.md#counttransactions) | **Head** /{ledger}/transactions | Count the transactions from a ledger
*LedgerApi* | [**CreateTransaction**](docs/LedgerApi.md#createtransaction) | **Post** /{ledger}/transactions | Create a new transaction to a ledger
*LedgerApi* | [**CreateTransactions**](docs/LedgerApi.md#createtransactions) | **Post** /{ledger}/transactions/batch | Create a new batch of transactions to a ledger
*LedgerApi* | [**GetAccount**](docs/LedgerApi.md#getaccount) | **Get** /{ledger}/accounts/{address} | Get account by its address
*LedgerApi* | [**GetBalances**](docs/LedgerApi.md#getbalances) | **Get** /{ledger}/balances | Get the balances from a ledger&#39;s account
*LedgerApi* | [**GetBalancesAggregated**](docs/LedgerApi.md#getbalancesaggregated) | **Get** /{ledger}/aggregate/balances | Get the aggregated balances from selected accounts
*LedgerApi* | [**GetInfo**](docs/LedgerApi.md#getinfo) | **Get** /_info | Show server information
*LedgerApi* | [**GetLedgerInfo**](docs/LedgerApi.md#getledgerinfo) | **Get** /{ledger}/_info | Get information about a ledger
*LedgerApi* | [**GetMapping**](docs/LedgerApi.md#getmapping) | **Get** /{ledger}/mapping | Get the mapping of a ledger
*LedgerApi* | [**GetTransaction**](docs/LedgerApi.md#gettransaction) | **Get** /{ledger}/transactions/{txid} | Get transaction from a ledger by its ID
*LedgerApi* | [**ListAccounts**](docs/LedgerApi.md#listaccounts) | **Get** /{ledger}/accounts | List accounts from a ledger
*LedgerApi* | [**ListLogs**](docs/LedgerApi.md#listlogs) | **Get** /{ledger}/logs | List the logs from a ledger
*LedgerApi* | [**ListTransactions**](docs/LedgerApi.md#listtransactions) | **Get** /{ledger}/transactions | List transactions from a ledger
*LedgerApi* | [**ReadStats**](docs/LedgerApi.md#readstats) | **Get** /{ledger}/stats | Get statistics from a ledger
*LedgerApi* | [**RevertTransaction**](docs/LedgerApi.md#reverttransaction) | **Post** /{ledger}/transactions/{txid}/revert | Revert a ledger transaction by its ID
*LedgerApi* | [**RunScript**](docs/LedgerApi.md#runscript) | **Post** /{ledger}/script | Execute a Numscript
*LedgerApi* | [**UpdateMapping**](docs/LedgerApi.md#updatemapping) | **Put** /{ledger}/mapping | Update the mapping of a ledger
*LogsApi* | [**ListLogs**](docs/LogsApi.md#listlogs) | **Get** /{ledger}/logs | List the logs from a ledger
*MappingApi* | [**GetMapping**](docs/MappingApi.md#getmapping) | **Get** /{ledger}/mapping | Get the mapping of a ledger
*MappingApi* | [**UpdateMapping**](docs/MappingApi.md#updatemapping) | **Put** /{ledger}/mapping | Update the mapping of a ledger
*ScriptApi* | [**RunScript**](docs/ScriptApi.md#runscript) | **Post** /{ledger}/script | Execute a Numscript
*ServerApi* | [**GetInfo**](docs/ServerApi.md#getinfo) | **Get** /_info | Show server information
*StatsApi* | [**ReadStats**](docs/StatsApi.md#readstats) | **Get** /{ledger}/stats | Get statistics from a ledger
*TransactionsApi* | [**AddMetadataOnTransaction**](docs/TransactionsApi.md#addmetadataontransaction) | **Post** /{ledger}/transactions/{txid}/metadata | Set the metadata of a transaction by its ID
*TransactionsApi* | [**CountTransactions**](docs/TransactionsApi.md#counttransactions) | **Head** /{ledger}/transactions | Count the transactions from a ledger
*TransactionsApi* | [**CreateTransaction**](docs/TransactionsApi.md#createtransaction) | **Post** /{ledger}/transactions | Create a new transaction to a ledger
*TransactionsApi* | [**CreateTransactions**](docs/TransactionsApi.md#createtransactions) | **Post** /{ledger}/transactions/batch | Create a new batch of transactions to a ledger
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /{ledger}/transactions/{txid} | Get transaction from a ledger by its ID
*TransactionsApi* | [**ListTransactions**](docs/TransactionsApi.md#listtransactions) | **Get** /{ledger}/transactions | List transactions from a ledger
*TransactionsApi* | [**RevertTransaction**](docs/TransactionsApi.md#reverttransaction) | **Post** /{ledger}/transactions/{txid}/revert | Revert a ledger transaction by its ID


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountResponse](docs/AccountResponse.md)
 - [AccountWithVolumesAndBalances](docs/AccountWithVolumesAndBalances.md)
 - [AccountsCursorResponse](docs/AccountsCursorResponse.md)
 - [AccountsCursorResponseCursor](docs/AccountsCursorResponseCursor.md)
 - [AggregateBalancesResponse](docs/AggregateBalancesResponse.md)
 - [BalancesCursorResponse](docs/BalancesCursorResponse.md)
 - [BalancesCursorResponseCursor](docs/BalancesCursorResponseCursor.md)
 - [Config](docs/Config.md)
 - [ConfigInfo](docs/ConfigInfo.md)
 - [ConfigInfoResponse](docs/ConfigInfoResponse.md)
 - [Contract](docs/Contract.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorsEnum](docs/ErrorsEnum.md)
 - [LedgerInfo](docs/LedgerInfo.md)
 - [LedgerInfoResponse](docs/LedgerInfoResponse.md)
 - [LedgerInfoStorage](docs/LedgerInfoStorage.md)
 - [LedgerStorage](docs/LedgerStorage.md)
 - [Log](docs/Log.md)
 - [LogsCursorResponse](docs/LogsCursorResponse.md)
 - [LogsCursorResponseCursor](docs/LogsCursorResponseCursor.md)
 - [Mapping](docs/Mapping.md)
 - [MappingResponse](docs/MappingResponse.md)
 - [MigrationInfo](docs/MigrationInfo.md)
 - [PostTransaction](docs/PostTransaction.md)
 - [PostTransactionScript](docs/PostTransactionScript.md)
 - [Posting](docs/Posting.md)
 - [Script](docs/Script.md)
 - [ScriptResponse](docs/ScriptResponse.md)
 - [Stats](docs/Stats.md)
 - [StatsResponse](docs/StatsResponse.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionData](docs/TransactionData.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [Transactions](docs/Transactions.md)
 - [TransactionsCursorResponse](docs/TransactionsCursorResponse.md)
 - [TransactionsCursorResponseCursor](docs/TransactionsCursorResponseCursor.md)
 - [TransactionsResponse](docs/TransactionsResponse.md)
 - [Volume](docs/Volume.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



