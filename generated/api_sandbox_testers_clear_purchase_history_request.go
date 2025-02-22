/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SandboxTestersClearPurchaseHistoryRequestApiService SandboxTestersClearPurchaseHistoryRequestApi service
type SandboxTestersClearPurchaseHistoryRequestApiService service

type SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest struct {
	ctx context.Context
	ApiService *SandboxTestersClearPurchaseHistoryRequestApiService
	sandboxTestersClearPurchaseHistoryRequestV2CreateRequest *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest
}

// SandboxTestersClearPurchaseHistoryRequest representation
func (r SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest) SandboxTestersClearPurchaseHistoryRequestV2CreateRequest(sandboxTestersClearPurchaseHistoryRequestV2CreateRequest SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest {
	r.sandboxTestersClearPurchaseHistoryRequestV2CreateRequest = &sandboxTestersClearPurchaseHistoryRequestV2CreateRequest
	return r
}

func (r SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest) Execute() (*SandboxTestersClearPurchaseHistoryRequestV2Response, *http.Response, error) {
	return r.ApiService.SandboxTestersClearPurchaseHistoryRequestCreateInstanceExecute(r)
}

/*
SandboxTestersClearPurchaseHistoryRequestCreateInstance Method for SandboxTestersClearPurchaseHistoryRequestCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest
*/
func (a *SandboxTestersClearPurchaseHistoryRequestApiService) SandboxTestersClearPurchaseHistoryRequestCreateInstance(ctx context.Context) SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest {
	return SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SandboxTestersClearPurchaseHistoryRequestV2Response
func (a *SandboxTestersClearPurchaseHistoryRequestApiService) SandboxTestersClearPurchaseHistoryRequestCreateInstanceExecute(r SandboxTestersClearPurchaseHistoryRequestApiSandboxTestersClearPurchaseHistoryRequestCreateInstanceRequest) (*SandboxTestersClearPurchaseHistoryRequestV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SandboxTestersClearPurchaseHistoryRequestV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SandboxTestersClearPurchaseHistoryRequestApiService.SandboxTestersClearPurchaseHistoryRequestCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/sandboxTestersClearPurchaseHistoryRequest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sandboxTestersClearPurchaseHistoryRequestV2CreateRequest == nil {
		return localVarReturnValue, nil, reportError("sandboxTestersClearPurchaseHistoryRequestV2CreateRequest is required and must be specified")
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
	localVarPostBody = r.sandboxTestersClearPurchaseHistoryRequestV2CreateRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
