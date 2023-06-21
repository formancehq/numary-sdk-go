/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type AccountsApi interface {

	/*
	AddMetadataToAccount Add metadata to an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ledger Name of the ledger.
	@param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
	@return ApiAddMetadataToAccountRequest
	*/
	AddMetadataToAccount(ctx context.Context, ledger string, address string) ApiAddMetadataToAccountRequest

	// AddMetadataToAccountExecute executes the request
	AddMetadataToAccountExecute(r ApiAddMetadataToAccountRequest) (*http.Response, error)

	/*
	CountAccounts Count the accounts from a ledger

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ledger Name of the ledger.
	@return ApiCountAccountsRequest
	*/
	CountAccounts(ctx context.Context, ledger string) ApiCountAccountsRequest

	// CountAccountsExecute executes the request
	CountAccountsExecute(r ApiCountAccountsRequest) (*http.Response, error)

	/*
	GetAccount Get account by its address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ledger Name of the ledger.
	@param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
	@return ApiGetAccountRequest
	*/
	GetAccount(ctx context.Context, ledger string, address string) ApiGetAccountRequest

	// GetAccountExecute executes the request
	//  @return AccountResponse
	GetAccountExecute(r ApiGetAccountRequest) (*AccountResponse, *http.Response, error)

	/*
	ListAccounts List accounts from a ledger

	List accounts from a ledger, sorted by address in descending order.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ledger Name of the ledger.
	@return ApiListAccountsRequest
	*/
	ListAccounts(ctx context.Context, ledger string) ApiListAccountsRequest

	// ListAccountsExecute executes the request
	//  @return AccountsCursorResponse
	ListAccountsExecute(r ApiListAccountsRequest) (*AccountsCursorResponse, *http.Response, error)
}

// AccountsApiService AccountsApi service
type AccountsApiService service

type ApiAddMetadataToAccountRequest struct {
	ctx context.Context
	ApiService AccountsApi
	ledger string
	address string
	requestBody *map[string]interface{}
}

// metadata
func (r ApiAddMetadataToAccountRequest) RequestBody(requestBody map[string]interface{}) ApiAddMetadataToAccountRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiAddMetadataToAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddMetadataToAccountExecute(r)
}

/*
AddMetadataToAccount Add metadata to an account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
 @return ApiAddMetadataToAccountRequest
*/
func (a *AccountsApiService) AddMetadataToAccount(ctx context.Context, ledger string, address string) ApiAddMetadataToAccountRequest {
	return ApiAddMetadataToAccountRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		address: address,
	}
}

// Execute executes the request
func (a *AccountsApiService) AddMetadataToAccountExecute(r ApiAddMetadataToAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.AddMetadataToAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts/{address}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", url.PathEscape(parameterValueToString(r.ledger, "ledger")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=utf-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCountAccountsRequest struct {
	ctx context.Context
	ApiService AccountsApi
	ledger string
	address *string
	metadata *map[string]interface{}
}

// Filter accounts by address pattern (regular expression placed between ^ and $).
func (r ApiCountAccountsRequest) Address(address string) ApiCountAccountsRequest {
	r.address = &address
	return r
}

// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
func (r ApiCountAccountsRequest) Metadata(metadata map[string]interface{}) ApiCountAccountsRequest {
	r.metadata = &metadata
	return r
}

func (r ApiCountAccountsRequest) Execute() (*http.Response, error) {
	return r.ApiService.CountAccountsExecute(r)
}

/*
CountAccounts Count the accounts from a ledger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiCountAccountsRequest
*/
func (a *AccountsApiService) CountAccounts(ctx context.Context, ledger string) ApiCountAccountsRequest {
	return ApiCountAccountsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
func (a *AccountsApiService) CountAccountsExecute(r ApiCountAccountsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.CountAccounts")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", url.PathEscape(parameterValueToString(r.ledger, "ledger")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.address != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
	}
	if r.metadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metadata", r.metadata, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAccountRequest struct {
	ctx context.Context
	ApiService AccountsApi
	ledger string
	address string
}

func (r ApiGetAccountRequest) Execute() (*AccountResponse, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Get account by its address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
 @return ApiGetAccountRequest
*/
func (a *AccountsApiService) GetAccount(ctx context.Context, ledger string, address string) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		address: address,
	}
}

// Execute executes the request
//  @return AccountResponse
func (a *AccountsApiService) GetAccountExecute(r ApiGetAccountRequest) (*AccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", url.PathEscape(parameterValueToString(r.ledger, "ledger")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterValueToString(r.address, "address")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAccountsRequest struct {
	ctx context.Context
	ApiService AccountsApi
	ledger string
	pageSize *int64
	pageSize2 *int64
	after *string
	address *string
	metadata *map[string]interface{}
	balance *int64
	balanceOperator *string
	balanceOperator2 *string
	cursor *string
	paginationToken *string
}

// The maximum number of results to return per page. 
func (r ApiListAccountsRequest) PageSize(pageSize int64) ApiListAccountsRequest {
	r.pageSize = &pageSize
	return r
}

// The maximum number of results to return per page. Deprecated, please use &#x60;pageSize&#x60; instead. 
// Deprecated
func (r ApiListAccountsRequest) PageSize2(pageSize2 int64) ApiListAccountsRequest {
	r.pageSize2 = &pageSize2
	return r
}

// Pagination cursor, will return accounts after given address, in descending order.
func (r ApiListAccountsRequest) After(after string) ApiListAccountsRequest {
	r.after = &after
	return r
}

// Filter accounts by address pattern (regular expression placed between ^ and $).
func (r ApiListAccountsRequest) Address(address string) ApiListAccountsRequest {
	r.address = &address
	return r
}

// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
func (r ApiListAccountsRequest) Metadata(metadata map[string]interface{}) ApiListAccountsRequest {
	r.metadata = &metadata
	return r
}

// Filter accounts by their balance (default operator is gte)
func (r ApiListAccountsRequest) Balance(balance int64) ApiListAccountsRequest {
	r.balance = &balance
	return r
}

// Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not. 
func (r ApiListAccountsRequest) BalanceOperator(balanceOperator string) ApiListAccountsRequest {
	r.balanceOperator = &balanceOperator
	return r
}

// Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not. Deprecated, please use &#x60;balanceOperator&#x60; instead. 
// Deprecated
func (r ApiListAccountsRequest) BalanceOperator2(balanceOperator2 string) ApiListAccountsRequest {
	r.balanceOperator2 = &balanceOperator2
	return r
}

// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. 
func (r ApiListAccountsRequest) Cursor(cursor string) ApiListAccountsRequest {
	r.cursor = &cursor
	return r
}

// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use &#x60;cursor&#x60; instead. 
// Deprecated
func (r ApiListAccountsRequest) PaginationToken(paginationToken string) ApiListAccountsRequest {
	r.paginationToken = &paginationToken
	return r
}

func (r ApiListAccountsRequest) Execute() (*AccountsCursorResponse, *http.Response, error) {
	return r.ApiService.ListAccountsExecute(r)
}

/*
ListAccounts List accounts from a ledger

List accounts from a ledger, sorted by address in descending order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiListAccountsRequest
*/
func (a *AccountsApiService) ListAccounts(ctx context.Context, ledger string) ApiListAccountsRequest {
	return ApiListAccountsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return AccountsCursorResponse
func (a *AccountsApiService) ListAccountsExecute(r ApiListAccountsRequest) (*AccountsCursorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountsCursorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.ListAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", url.PathEscape(parameterValueToString(r.ledger, "ledger")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.pageSize2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize2, "")
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.address != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
	}
	if r.metadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metadata", r.metadata, "")
	}
	if r.balance != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "balance", r.balance, "")
	}
	if r.balanceOperator != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "balanceOperator", r.balanceOperator, "")
	}
	if r.balanceOperator2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "balance_operator", r.balanceOperator2, "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.paginationToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination_token", r.paginationToken, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
