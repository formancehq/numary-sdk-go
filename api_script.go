/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.9.1
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


type ScriptApi interface {

	/*
	RunScript Execute a Numscript

	This route is deprecated, and has been merged into `POST /{ledger}/transactions`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ledger Name of the ledger.
	@return ApiRunScriptRequest

	Deprecated
	*/
	RunScript(ctx context.Context, ledger string) ApiRunScriptRequest

	// RunScriptExecute executes the request
	//  @return ScriptResponse
	// Deprecated
	RunScriptExecute(r ApiRunScriptRequest) (*ScriptResponse, *http.Response, error)
}

// ScriptApiService ScriptApi service
type ScriptApiService service

type ApiRunScriptRequest struct {
	ctx context.Context
	ApiService ScriptApi
	ledger string
	script *Script
	preview *bool
}

func (r ApiRunScriptRequest) Script(script Script) ApiRunScriptRequest {
	r.script = &script
	return r
}

// Set the preview mode. Preview mode doesn&#39;t add the logs to the database or publish a message to the message broker.
func (r ApiRunScriptRequest) Preview(preview bool) ApiRunScriptRequest {
	r.preview = &preview
	return r
}

func (r ApiRunScriptRequest) Execute() (*ScriptResponse, *http.Response, error) {
	return r.ApiService.RunScriptExecute(r)
}

/*
RunScript Execute a Numscript

This route is deprecated, and has been merged into `POST /{ledger}/transactions`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiRunScriptRequest

Deprecated
*/
func (a *ScriptApiService) RunScript(ctx context.Context, ledger string) ApiRunScriptRequest {
	return ApiRunScriptRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return ScriptResponse
// Deprecated
func (a *ScriptApiService) RunScriptExecute(r ApiRunScriptRequest) (*ScriptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScriptResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScriptApiService.RunScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/script"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", url.PathEscape(parameterValueToString(r.ledger, "ledger")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.script == nil {
		return localVarReturnValue, nil, reportError("script is required and must be specified")
	}

	if r.preview != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "preview", r.preview, "")
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
	localVarPostBody = r.script
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
