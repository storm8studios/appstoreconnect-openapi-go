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


// BetaBuildLocalizationsApiService BetaBuildLocalizationsApi service
type BetaBuildLocalizationsApiService service

type BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *BetaBuildLocalizationsApiService
	id string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type builds
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest) FieldsBuilds(fieldsBuilds []string) BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest) Execute() (*BuildResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsBuildGetToOneRelatedExecute(r)
}

/*
BetaBuildLocalizationsBuildGetToOneRelated Method for BetaBuildLocalizationsBuildGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest
*/
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsBuildGetToOneRelated(ctx context.Context, id string) BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest {
	return BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildResponse
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsBuildGetToOneRelatedExecute(r BetaBuildLocalizationsApiBetaBuildLocalizationsBuildGetToOneRelatedRequest) (*BuildResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsApiService.BetaBuildLocalizationsBuildGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}/build"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
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

type BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *BetaBuildLocalizationsApiService
	betaBuildLocalizationCreateRequest *BetaBuildLocalizationCreateRequest
}

// BetaBuildLocalization representation
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest) BetaBuildLocalizationCreateRequest(betaBuildLocalizationCreateRequest BetaBuildLocalizationCreateRequest) BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest {
	r.betaBuildLocalizationCreateRequest = &betaBuildLocalizationCreateRequest
	return r
}

func (r BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest) Execute() (*BetaBuildLocalizationResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsCreateInstanceExecute(r)
}

/*
BetaBuildLocalizationsCreateInstance Method for BetaBuildLocalizationsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest
*/
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsCreateInstance(ctx context.Context) BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest {
	return BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BetaBuildLocalizationResponse
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsCreateInstanceExecute(r BetaBuildLocalizationsApiBetaBuildLocalizationsCreateInstanceRequest) (*BetaBuildLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaBuildLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsApiService.BetaBuildLocalizationsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaBuildLocalizationCreateRequest == nil {
		return localVarReturnValue, nil, reportError("betaBuildLocalizationCreateRequest is required and must be specified")
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
	localVarPostBody = r.betaBuildLocalizationCreateRequest
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

type BetaBuildLocalizationsApiBetaBuildLocalizationsDeleteInstanceRequest struct {
	ctx context.Context
	ApiService *BetaBuildLocalizationsApiService
	id string
}

func (r BetaBuildLocalizationsApiBetaBuildLocalizationsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsDeleteInstanceExecute(r)
}

/*
BetaBuildLocalizationsDeleteInstance Method for BetaBuildLocalizationsDeleteInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BetaBuildLocalizationsApiBetaBuildLocalizationsDeleteInstanceRequest
*/
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsDeleteInstance(ctx context.Context, id string) BetaBuildLocalizationsApiBetaBuildLocalizationsDeleteInstanceRequest {
	return BetaBuildLocalizationsApiBetaBuildLocalizationsDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsDeleteInstanceExecute(r BetaBuildLocalizationsApiBetaBuildLocalizationsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsApiService.BetaBuildLocalizationsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}"
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

type BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest struct {
	ctx context.Context
	ApiService *BetaBuildLocalizationsApiService
	filterLocale *[]string
	filterBuild *[]string
	fieldsBetaBuildLocalizations *[]string
	limit *int32
	include *[]string
	fieldsBuilds *[]string
}

// filter by attribute &#39;locale&#39;
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) FilterLocale(filterLocale []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	r.filterLocale = &filterLocale
	return r
}

// filter by id(s) of related &#39;build&#39;
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) FilterBuild(filterBuild []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	r.filterBuild = &filterBuild
	return r
}

// the fields to include for returned resources of type betaBuildLocalizations
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	r.fieldsBetaBuildLocalizations = &fieldsBetaBuildLocalizations
	return r
}

// maximum resources per page
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) Limit(limit int32) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) Include(include []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type builds
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) FieldsBuilds(fieldsBuilds []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) Execute() (*BetaBuildLocalizationsResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsGetCollectionExecute(r)
}

/*
BetaBuildLocalizationsGetCollection Method for BetaBuildLocalizationsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest
*/
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsGetCollection(ctx context.Context) BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest {
	return BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BetaBuildLocalizationsResponse
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsGetCollectionExecute(r BetaBuildLocalizationsApiBetaBuildLocalizationsGetCollectionRequest) (*BetaBuildLocalizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaBuildLocalizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsApiService.BetaBuildLocalizationsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[locale]", r.filterLocale, "csv")
	}
	if r.filterBuild != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[build]", r.filterBuild, "csv")
	}
	if r.fieldsBetaBuildLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaBuildLocalizations]", r.fieldsBetaBuildLocalizations, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
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

type BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest struct {
	ctx context.Context
	ApiService *BetaBuildLocalizationsApiService
	id string
	fieldsBetaBuildLocalizations *[]string
	include *[]string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type betaBuildLocalizations
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest) FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest {
	r.fieldsBetaBuildLocalizations = &fieldsBetaBuildLocalizations
	return r
}

// comma-separated list of relationships to include
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest) Include(include []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type builds
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest) FieldsBuilds(fieldsBuilds []string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest) Execute() (*BetaBuildLocalizationResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsGetInstanceExecute(r)
}

/*
BetaBuildLocalizationsGetInstance Method for BetaBuildLocalizationsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest
*/
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsGetInstance(ctx context.Context, id string) BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest {
	return BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaBuildLocalizationResponse
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsGetInstanceExecute(r BetaBuildLocalizationsApiBetaBuildLocalizationsGetInstanceRequest) (*BetaBuildLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaBuildLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsApiService.BetaBuildLocalizationsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBetaBuildLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaBuildLocalizations]", r.fieldsBetaBuildLocalizations, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
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

type BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *BetaBuildLocalizationsApiService
	id string
	betaBuildLocalizationUpdateRequest *BetaBuildLocalizationUpdateRequest
}

// BetaBuildLocalization representation
func (r BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest) BetaBuildLocalizationUpdateRequest(betaBuildLocalizationUpdateRequest BetaBuildLocalizationUpdateRequest) BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest {
	r.betaBuildLocalizationUpdateRequest = &betaBuildLocalizationUpdateRequest
	return r
}

func (r BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest) Execute() (*BetaBuildLocalizationResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsUpdateInstanceExecute(r)
}

/*
BetaBuildLocalizationsUpdateInstance Method for BetaBuildLocalizationsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest
*/
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsUpdateInstance(ctx context.Context, id string) BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest {
	return BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaBuildLocalizationResponse
func (a *BetaBuildLocalizationsApiService) BetaBuildLocalizationsUpdateInstanceExecute(r BetaBuildLocalizationsApiBetaBuildLocalizationsUpdateInstanceRequest) (*BetaBuildLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaBuildLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsApiService.BetaBuildLocalizationsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaBuildLocalizationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("betaBuildLocalizationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.betaBuildLocalizationUpdateRequest
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
