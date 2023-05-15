/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.10.0-rc.1
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
	"time"
)

// Linger please
var (
	_ _context.Context
)

type LogsAPI interface {

	/*
	ListLogs List the logs from a ledger

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

// LogsAPIService LogsAPI service
type LogsAPIService service

type ApiListLogsRequest struct {
	ctx _context.Context
	ApiService LogsAPI
	ledger string
	pageSize *int64
	pageSize2 *int64
	after *string
	startTime *time.Time
	startTime2 *time.Time
	endTime *time.Time
	endTime2 *time.Time
	cursor *string
	paginationToken *string
}

// The maximum number of results to return per page. 
func (r ApiListLogsRequest) PageSize(pageSize int64) ApiListLogsRequest {
	r.pageSize = &pageSize
	return r
}
// The maximum number of results to return per page. Deprecated, please use &#x60;pageSize&#x60; instead. 
// Deprecated
func (r ApiListLogsRequest) PageSize2(pageSize2 int64) ApiListLogsRequest {
	r.pageSize2 = &pageSize2
	return r
}
// Pagination cursor, will return the logs after a given ID. (in descending order).
func (r ApiListLogsRequest) After(after string) ApiListLogsRequest {
	r.after = &after
	return r
}
// Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute). 
func (r ApiListLogsRequest) StartTime(startTime time.Time) ApiListLogsRequest {
	r.startTime = &startTime
	return r
}
// Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; includes the first second of 4th minute). Deprecated, please use &#x60;startTime&#x60; instead. 
// Deprecated
func (r ApiListLogsRequest) StartTime2(startTime2 time.Time) ApiListLogsRequest {
	r.startTime2 = &startTime2
	return r
}
// Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute). 
func (r ApiListLogsRequest) EndTime(endTime time.Time) ApiListLogsRequest {
	r.endTime = &endTime
	return r
}
// Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, \&quot;2023-01-02T15:04:01Z\&quot; excludes the first second of 4th minute). Deprecated, please use &#x60;endTime&#x60; instead. 
// Deprecated
func (r ApiListLogsRequest) EndTime2(endTime2 time.Time) ApiListLogsRequest {
	r.endTime2 = &endTime2
	return r
}
// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. 
func (r ApiListLogsRequest) Cursor(cursor string) ApiListLogsRequest {
	r.cursor = &cursor
	return r
}
// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set. Deprecated, please use &#x60;cursor&#x60; instead. 
// Deprecated
func (r ApiListLogsRequest) PaginationToken(paginationToken string) ApiListLogsRequest {
	r.paginationToken = &paginationToken
	return r
}

func (r ApiListLogsRequest) Execute() (LogsCursorResponse, *_nethttp.Response, error) {
	return r.ApiService.ListLogsExecute(r)
}

/*
ListLogs List the logs from a ledger

List the logs from a ledger, sorted by ID in descending order.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiListLogsRequest
*/
func (a *LogsAPIService) ListLogs(ctx _context.Context, ledger string) ApiListLogsRequest {
	return ApiListLogsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return LogsCursorResponse
func (a *LogsAPIService) ListLogsExecute(r ApiListLogsRequest) (LogsCursorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  LogsCursorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsAPIService.ListLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageSize2 != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize2, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.startTime != nil {
		localVarQueryParams.Add("startTime", parameterToString(*r.startTime, ""))
	}
	if r.startTime2 != nil {
		localVarQueryParams.Add("start_time", parameterToString(*r.startTime2, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.endTime2 != nil {
		localVarQueryParams.Add("end_time", parameterToString(*r.endTime2, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
