/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.8.1-rc.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type BalancesApi interface {

	/*
	GetBalances Get the balances from a ledger's account

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiGetBalancesRequest
	*/
	GetBalances(ctx _context.Context, ledger string) ApiGetBalancesRequest

	// GetBalancesExecute executes the request
	//  @return GetBalances200Response
	GetBalancesExecute(r ApiGetBalancesRequest) (GetBalances200Response, *_nethttp.Response, error)

	/*
	GetBalancesAggregated Get the aggregated balances from selected accounts

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiGetBalancesAggregatedRequest
	*/
	GetBalancesAggregated(ctx _context.Context, ledger string) ApiGetBalancesAggregatedRequest

	// GetBalancesAggregatedExecute executes the request
	//  @return GetBalancesAggregated200Response
	GetBalancesAggregatedExecute(r ApiGetBalancesAggregatedRequest) (GetBalancesAggregated200Response, *_nethttp.Response, error)
}

// BalancesApiService BalancesApi service
type BalancesApiService service

type ApiGetBalancesRequest struct {
	ctx _context.Context
	ApiService BalancesApi
	ledger string
	address *string
	after *string
	paginationToken *string
}

// Filter balances involving given account, either as source or destination.
func (r ApiGetBalancesRequest) Address(address string) ApiGetBalancesRequest {
	r.address = &address
	return r
}
// Pagination cursor, will return accounts after given address, in descending order.
func (r ApiGetBalancesRequest) After(after string) ApiGetBalancesRequest {
	r.after = &after
	return r
}
// Parameter used in pagination requests.  Set to the value of next for the next page of results.  Set to the value of previous for the previous page of results.
func (r ApiGetBalancesRequest) PaginationToken(paginationToken string) ApiGetBalancesRequest {
	r.paginationToken = &paginationToken
	return r
}

func (r ApiGetBalancesRequest) Execute() (GetBalances200Response, *_nethttp.Response, error) {
	return r.ApiService.GetBalancesExecute(r)
}

/*
GetBalances Get the balances from a ledger's account

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiGetBalancesRequest
*/
func (a *BalancesApiService) GetBalances(ctx _context.Context, ledger string) ApiGetBalancesRequest {
	return ApiGetBalancesRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return GetBalances200Response
func (a *BalancesApiService) GetBalancesExecute(r ApiGetBalancesRequest) (GetBalances200Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  GetBalances200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesApiService.GetBalances")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/balances"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.address != nil {
		localVarQueryParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.paginationToken != nil {
		localVarQueryParams.Add("pagination_token", parameterToString(*r.paginationToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ListAccounts400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBalancesAggregatedRequest struct {
	ctx _context.Context
	ApiService BalancesApi
	ledger string
	address *string
}

// Filter balances involving given account, either as source or destination.
func (r ApiGetBalancesAggregatedRequest) Address(address string) ApiGetBalancesAggregatedRequest {
	r.address = &address
	return r
}

func (r ApiGetBalancesAggregatedRequest) Execute() (GetBalancesAggregated200Response, *_nethttp.Response, error) {
	return r.ApiService.GetBalancesAggregatedExecute(r)
}

/*
GetBalancesAggregated Get the aggregated balances from selected accounts

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiGetBalancesAggregatedRequest
*/
func (a *BalancesApiService) GetBalancesAggregated(ctx _context.Context, ledger string) ApiGetBalancesAggregatedRequest {
	return ApiGetBalancesAggregatedRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return GetBalancesAggregated200Response
func (a *BalancesApiService) GetBalancesAggregatedExecute(r ApiGetBalancesAggregatedRequest) (GetBalancesAggregated200Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  GetBalancesAggregated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesApiService.GetBalancesAggregated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/aggregate/balances"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.address != nil {
		localVarQueryParams.Add("address", parameterToString(*r.address, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetBalancesAggregated400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
