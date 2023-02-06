/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.0-rc.9
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

type MappingApi interface {

	/*
	GetMapping Get the mapping of a ledger

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiGetMappingRequest
	*/
	GetMapping(ctx _context.Context, ledger string) ApiGetMappingRequest

	// GetMappingExecute executes the request
	//  @return MappingResponse
	GetMappingExecute(r ApiGetMappingRequest) (MappingResponse, *_nethttp.Response, error)

	/*
	UpdateMapping Update the mapping of a ledger

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiUpdateMappingRequest
	*/
	UpdateMapping(ctx _context.Context, ledger string) ApiUpdateMappingRequest

	// UpdateMappingExecute executes the request
	//  @return MappingResponse
	UpdateMappingExecute(r ApiUpdateMappingRequest) (MappingResponse, *_nethttp.Response, error)
}

// MappingApiService MappingApi service
type MappingApiService service

type ApiGetMappingRequest struct {
	ctx _context.Context
	ApiService MappingApi
	ledger string
}


func (r ApiGetMappingRequest) Execute() (MappingResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMappingExecute(r)
}

/*
GetMapping Get the mapping of a ledger

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiGetMappingRequest
*/
func (a *MappingApiService) GetMapping(ctx _context.Context, ledger string) ApiGetMappingRequest {
	return ApiGetMappingRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return MappingResponse
func (a *MappingApiService) GetMappingExecute(r ApiGetMappingRequest) (MappingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MappingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingApiService.GetMapping")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/mapping"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiUpdateMappingRequest struct {
	ctx _context.Context
	ApiService MappingApi
	ledger string
	mapping *Mapping
}

func (r ApiUpdateMappingRequest) Mapping(mapping Mapping) ApiUpdateMappingRequest {
	r.mapping = &mapping
	return r
}

func (r ApiUpdateMappingRequest) Execute() (MappingResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateMappingExecute(r)
}

/*
UpdateMapping Update the mapping of a ledger

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiUpdateMappingRequest
*/
func (a *MappingApiService) UpdateMapping(ctx _context.Context, ledger string) ApiUpdateMappingRequest {
	return ApiUpdateMappingRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return MappingResponse
func (a *MappingApiService) UpdateMappingExecute(r ApiUpdateMappingRequest) (MappingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MappingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MappingApiService.UpdateMapping")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/mapping"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.mapping == nil {
		return localVarReturnValue, nil, reportError("mapping is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.mapping
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
