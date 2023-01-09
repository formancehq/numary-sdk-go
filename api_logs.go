/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-rc.2
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

type LogsApi interface {

	/*
	ListLogs List the logs from a ledger.

	List the logs from a ledger, sorted by ID in descending order.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiListLogsRequest
	*/
	ListLogs(ctx _context.Context, ledger string) ApiListLogsRequest

	// ListLogsExecute executes the request
	//  @return LogsCursorResponse
	ListLogsExecute(r ApiListLogsRequest) (LogsCursorResponse, *_nethttp.Response, error)
}

// LogsApiService LogsApi service
type LogsApiService service

type ApiListLogsRequest struct {
	ctx _context.Context
	ApiService LogsApi
	ledger string
	pageSize *int64
	after *string
	startTime *string
	endTime *string
	paginationToken *string
}

// The maximum number of results to return per page
func (r ApiListLogsRequest) PageSize(pageSize int64) ApiListLogsRequest {
	r.pageSize = &pageSize
	return r
}
// Pagination cursor, will return the logs after a given ID. (in descending order).
func (r ApiListLogsRequest) After(after string) ApiListLogsRequest {
	r.after = &after
	return r
}
// Filter logs that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, 12:00:01 includes the first second of the minute). 
func (r ApiListLogsRequest) StartTime(startTime string) ApiListLogsRequest {
	r.startTime = &startTime
	return r
}
// Filter logs that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, 12:00:01 excludes the first second of the minute). 
func (r ApiListLogsRequest) EndTime(endTime string) ApiListLogsRequest {
	r.endTime = &endTime
	return r
}
// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set. 
func (r ApiListLogsRequest) PaginationToken(paginationToken string) ApiListLogsRequest {
	r.paginationToken = &paginationToken
	return r
}

func (r ApiListLogsRequest) Execute() (LogsCursorResponse, *_nethttp.Response, error) {
	return r.ApiService.ListLogsExecute(r)
}

/*
ListLogs List the logs from a ledger.

List the logs from a ledger, sorted by ID in descending order.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiListLogsRequest
*/
func (a *LogsApiService) ListLogs(ctx _context.Context, ledger string) ApiListLogsRequest {
	return ApiListLogsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return LogsCursorResponse
func (a *LogsApiService) ListLogsExecute(r ApiListLogsRequest) (LogsCursorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  LogsCursorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.ListLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/log"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.startTime != nil {
		localVarQueryParams.Add("start_time", parameterToString(*r.startTime, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("end_time", parameterToString(*r.endTime, ""))
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
