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
	"strings"
)


// AppScreenshotSetsApiService AppScreenshotSetsApi service
type AppScreenshotSetsApiService service

type AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppScreenshotSetsApiService
	id string
	fieldsAppScreenshotSets *[]string
	fieldsAppScreenshots *[]string
	limit *int32
	include *[]string
}

// the fields to include for returned resources of type appScreenshotSets
func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest) FieldsAppScreenshotSets(fieldsAppScreenshotSets []string) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest {
	r.fieldsAppScreenshotSets = &fieldsAppScreenshotSets
	return r
}

// the fields to include for returned resources of type appScreenshots
func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest) FieldsAppScreenshots(fieldsAppScreenshots []string) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest {
	r.fieldsAppScreenshots = &fieldsAppScreenshots
	return r
}

// maximum resources per page
func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest) Limit(limit int32) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest) Include(include []string) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest) Execute() (*AppScreenshotsResponse, *http.Response, error) {
	return r.ApiService.AppScreenshotSetsAppScreenshotsGetToManyRelatedExecute(r)
}

/*
AppScreenshotSetsAppScreenshotsGetToManyRelated Method for AppScreenshotSetsAppScreenshotsGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest
*/
func (a *AppScreenshotSetsApiService) AppScreenshotSetsAppScreenshotsGetToManyRelated(ctx context.Context, id string) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest {
	return AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppScreenshotsResponse
func (a *AppScreenshotSetsApiService) AppScreenshotSetsAppScreenshotsGetToManyRelatedExecute(r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest) (*AppScreenshotsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppScreenshotsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppScreenshotSetsApiService.AppScreenshotSetsAppScreenshotsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appScreenshotSets/{id}/appScreenshots"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppScreenshotSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appScreenshotSets]", r.fieldsAppScreenshotSets, "csv")
	}
	if r.fieldsAppScreenshots != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appScreenshots]", r.fieldsAppScreenshots, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest struct {
	ctx context.Context
	ApiService *AppScreenshotSetsApiService
	id string
	limit *int32
}

// maximum resources per page
func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest) Limit(limit int32) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest {
	r.limit = &limit
	return r
}

func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest) Execute() (*AppScreenshotSetAppScreenshotsLinkagesResponse, *http.Response, error) {
	return r.ApiService.AppScreenshotSetsAppScreenshotsGetToManyRelationshipExecute(r)
}

/*
AppScreenshotSetsAppScreenshotsGetToManyRelationship Method for AppScreenshotSetsAppScreenshotsGetToManyRelationship

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest
*/
func (a *AppScreenshotSetsApiService) AppScreenshotSetsAppScreenshotsGetToManyRelationship(ctx context.Context, id string) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest {
	return AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppScreenshotSetAppScreenshotsLinkagesResponse
func (a *AppScreenshotSetsApiService) AppScreenshotSetsAppScreenshotsGetToManyRelationshipExecute(r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest) (*AppScreenshotSetAppScreenshotsLinkagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppScreenshotSetAppScreenshotsLinkagesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppScreenshotSetsApiService.AppScreenshotSetsAppScreenshotsGetToManyRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appScreenshotSets/{id}/relationships/appScreenshots"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest struct {
	ctx context.Context
	ApiService *AppScreenshotSetsApiService
	id string
	appScreenshotSetAppScreenshotsLinkagesRequest *AppScreenshotSetAppScreenshotsLinkagesRequest
}

// List of related linkages
func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest) AppScreenshotSetAppScreenshotsLinkagesRequest(appScreenshotSetAppScreenshotsLinkagesRequest AppScreenshotSetAppScreenshotsLinkagesRequest) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest {
	r.appScreenshotSetAppScreenshotsLinkagesRequest = &appScreenshotSetAppScreenshotsLinkagesRequest
	return r
}

func (r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppScreenshotSetsAppScreenshotsReplaceToManyRelationshipExecute(r)
}

/*
AppScreenshotSetsAppScreenshotsReplaceToManyRelationship Method for AppScreenshotSetsAppScreenshotsReplaceToManyRelationship

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest
*/
func (a *AppScreenshotSetsApiService) AppScreenshotSetsAppScreenshotsReplaceToManyRelationship(ctx context.Context, id string) AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest {
	return AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AppScreenshotSetsApiService) AppScreenshotSetsAppScreenshotsReplaceToManyRelationshipExecute(r AppScreenshotSetsApiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppScreenshotSetsApiService.AppScreenshotSetsAppScreenshotsReplaceToManyRelationship")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appScreenshotSets/{id}/relationships/appScreenshots"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appScreenshotSetAppScreenshotsLinkagesRequest == nil {
		return nil, reportError("appScreenshotSetAppScreenshotsLinkagesRequest is required and must be specified")
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
	localVarPostBody = r.appScreenshotSetAppScreenshotsLinkagesRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppScreenshotSetsApiService
	appScreenshotSetCreateRequest *AppScreenshotSetCreateRequest
}

// AppScreenshotSet representation
func (r AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest) AppScreenshotSetCreateRequest(appScreenshotSetCreateRequest AppScreenshotSetCreateRequest) AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest {
	r.appScreenshotSetCreateRequest = &appScreenshotSetCreateRequest
	return r
}

func (r AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest) Execute() (*AppScreenshotSetResponse, *http.Response, error) {
	return r.ApiService.AppScreenshotSetsCreateInstanceExecute(r)
}

/*
AppScreenshotSetsCreateInstance Method for AppScreenshotSetsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest
*/
func (a *AppScreenshotSetsApiService) AppScreenshotSetsCreateInstance(ctx context.Context) AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest {
	return AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppScreenshotSetResponse
func (a *AppScreenshotSetsApiService) AppScreenshotSetsCreateInstanceExecute(r AppScreenshotSetsApiAppScreenshotSetsCreateInstanceRequest) (*AppScreenshotSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppScreenshotSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppScreenshotSetsApiService.AppScreenshotSetsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appScreenshotSets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appScreenshotSetCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appScreenshotSetCreateRequest is required and must be specified")
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
	localVarPostBody = r.appScreenshotSetCreateRequest
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

type AppScreenshotSetsApiAppScreenshotSetsDeleteInstanceRequest struct {
	ctx context.Context
	ApiService *AppScreenshotSetsApiService
	id string
}

func (r AppScreenshotSetsApiAppScreenshotSetsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppScreenshotSetsDeleteInstanceExecute(r)
}

/*
AppScreenshotSetsDeleteInstance Method for AppScreenshotSetsDeleteInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppScreenshotSetsApiAppScreenshotSetsDeleteInstanceRequest
*/
func (a *AppScreenshotSetsApiService) AppScreenshotSetsDeleteInstance(ctx context.Context, id string) AppScreenshotSetsApiAppScreenshotSetsDeleteInstanceRequest {
	return AppScreenshotSetsApiAppScreenshotSetsDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AppScreenshotSetsApiService) AppScreenshotSetsDeleteInstanceExecute(r AppScreenshotSetsApiAppScreenshotSetsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppScreenshotSetsApiService.AppScreenshotSetsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appScreenshotSets/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest struct {
	ctx context.Context
	ApiService *AppScreenshotSetsApiService
	id string
	fieldsAppScreenshotSets *[]string
	include *[]string
	fieldsAppScreenshots *[]string
	limitAppScreenshots *int32
}

// the fields to include for returned resources of type appScreenshotSets
func (r AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest) FieldsAppScreenshotSets(fieldsAppScreenshotSets []string) AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest {
	r.fieldsAppScreenshotSets = &fieldsAppScreenshotSets
	return r
}

// comma-separated list of relationships to include
func (r AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest) Include(include []string) AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appScreenshots
func (r AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest) FieldsAppScreenshots(fieldsAppScreenshots []string) AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest {
	r.fieldsAppScreenshots = &fieldsAppScreenshots
	return r
}

// maximum number of related appScreenshots returned (when they are included)
func (r AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest) LimitAppScreenshots(limitAppScreenshots int32) AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest {
	r.limitAppScreenshots = &limitAppScreenshots
	return r
}

func (r AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest) Execute() (*AppScreenshotSetResponse, *http.Response, error) {
	return r.ApiService.AppScreenshotSetsGetInstanceExecute(r)
}

/*
AppScreenshotSetsGetInstance Method for AppScreenshotSetsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest
*/
func (a *AppScreenshotSetsApiService) AppScreenshotSetsGetInstance(ctx context.Context, id string) AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest {
	return AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppScreenshotSetResponse
func (a *AppScreenshotSetsApiService) AppScreenshotSetsGetInstanceExecute(r AppScreenshotSetsApiAppScreenshotSetsGetInstanceRequest) (*AppScreenshotSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppScreenshotSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppScreenshotSetsApiService.AppScreenshotSetsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appScreenshotSets/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppScreenshotSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appScreenshotSets]", r.fieldsAppScreenshotSets, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppScreenshots != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appScreenshots]", r.fieldsAppScreenshots, "csv")
	}
	if r.limitAppScreenshots != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appScreenshots]", r.limitAppScreenshots, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
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
