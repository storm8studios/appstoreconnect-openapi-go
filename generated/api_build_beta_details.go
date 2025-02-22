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


// BuildBetaDetailsApiService BuildBetaDetailsApi service
type BuildBetaDetailsApiService service

type BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsApiService
	id string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type builds
func (r BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest) FieldsBuilds(fieldsBuilds []string) BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest) Execute() (*BuildResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsBuildGetToOneRelatedExecute(r)
}

/*
BuildBetaDetailsBuildGetToOneRelated Method for BuildBetaDetailsBuildGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest
*/
func (a *BuildBetaDetailsApiService) BuildBetaDetailsBuildGetToOneRelated(ctx context.Context, id string) BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest {
	return BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildResponse
func (a *BuildBetaDetailsApiService) BuildBetaDetailsBuildGetToOneRelatedExecute(r BuildBetaDetailsApiBuildBetaDetailsBuildGetToOneRelatedRequest) (*BuildResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsApiService.BuildBetaDetailsBuildGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails/{id}/build"
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

type BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsApiService
	filterBuild *[]string
	filterId *[]string
	fieldsBuildBetaDetails *[]string
	limit *int32
	include *[]string
	fieldsBuilds *[]string
}

// filter by id(s) of related &#39;build&#39;
func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) FilterBuild(filterBuild []string) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	r.filterBuild = &filterBuild
	return r
}

// filter by id(s)
func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) FilterId(filterId []string) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	r.filterId = &filterId
	return r
}

// the fields to include for returned resources of type buildBetaDetails
func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) FieldsBuildBetaDetails(fieldsBuildBetaDetails []string) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	r.fieldsBuildBetaDetails = &fieldsBuildBetaDetails
	return r
}

// maximum resources per page
func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) Limit(limit int32) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) Include(include []string) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type builds
func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) FieldsBuilds(fieldsBuilds []string) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) Execute() (*BuildBetaDetailsResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsGetCollectionExecute(r)
}

/*
BuildBetaDetailsGetCollection Method for BuildBetaDetailsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest
*/
func (a *BuildBetaDetailsApiService) BuildBetaDetailsGetCollection(ctx context.Context) BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest {
	return BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BuildBetaDetailsResponse
func (a *BuildBetaDetailsApiService) BuildBetaDetailsGetCollectionExecute(r BuildBetaDetailsApiBuildBetaDetailsGetCollectionRequest) (*BuildBetaDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsApiService.BuildBetaDetailsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBuild != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[build]", r.filterBuild, "csv")
	}
	if r.filterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[id]", r.filterId, "csv")
	}
	if r.fieldsBuildBetaDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[buildBetaDetails]", r.fieldsBuildBetaDetails, "csv")
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

type BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsApiService
	id string
	fieldsBuildBetaDetails *[]string
	include *[]string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type buildBetaDetails
func (r BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest) FieldsBuildBetaDetails(fieldsBuildBetaDetails []string) BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest {
	r.fieldsBuildBetaDetails = &fieldsBuildBetaDetails
	return r
}

// comma-separated list of relationships to include
func (r BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest) Include(include []string) BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type builds
func (r BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest) FieldsBuilds(fieldsBuilds []string) BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest) Execute() (*BuildBetaDetailResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsGetInstanceExecute(r)
}

/*
BuildBetaDetailsGetInstance Method for BuildBetaDetailsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest
*/
func (a *BuildBetaDetailsApiService) BuildBetaDetailsGetInstance(ctx context.Context, id string) BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest {
	return BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildBetaDetailResponse
func (a *BuildBetaDetailsApiService) BuildBetaDetailsGetInstanceExecute(r BuildBetaDetailsApiBuildBetaDetailsGetInstanceRequest) (*BuildBetaDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsApiService.BuildBetaDetailsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBuildBetaDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[buildBetaDetails]", r.fieldsBuildBetaDetails, "csv")
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

type BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsApiService
	id string
	buildBetaDetailUpdateRequest *BuildBetaDetailUpdateRequest
}

// BuildBetaDetail representation
func (r BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest) BuildBetaDetailUpdateRequest(buildBetaDetailUpdateRequest BuildBetaDetailUpdateRequest) BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest {
	r.buildBetaDetailUpdateRequest = &buildBetaDetailUpdateRequest
	return r
}

func (r BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest) Execute() (*BuildBetaDetailResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsUpdateInstanceExecute(r)
}

/*
BuildBetaDetailsUpdateInstance Method for BuildBetaDetailsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest
*/
func (a *BuildBetaDetailsApiService) BuildBetaDetailsUpdateInstance(ctx context.Context, id string) BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest {
	return BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildBetaDetailResponse
func (a *BuildBetaDetailsApiService) BuildBetaDetailsUpdateInstanceExecute(r BuildBetaDetailsApiBuildBetaDetailsUpdateInstanceRequest) (*BuildBetaDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsApiService.BuildBetaDetailsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.buildBetaDetailUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("buildBetaDetailUpdateRequest is required and must be specified")
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
	localVarPostBody = r.buildBetaDetailUpdateRequest
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
