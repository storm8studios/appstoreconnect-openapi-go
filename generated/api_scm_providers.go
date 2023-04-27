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


// ScmProvidersApiService ScmProvidersApi service
type ScmProvidersApiService service

type ScmProvidersApiScmProvidersGetCollectionRequest struct {
	ctx context.Context
	ApiService *ScmProvidersApiService
	fieldsScmProviders *[]string
	limit *int32
	fieldsScmRepositories *[]string
}

// the fields to include for returned resources of type scmProviders
func (r ScmProvidersApiScmProvidersGetCollectionRequest) FieldsScmProviders(fieldsScmProviders []string) ScmProvidersApiScmProvidersGetCollectionRequest {
	r.fieldsScmProviders = &fieldsScmProviders
	return r
}

// maximum resources per page
func (r ScmProvidersApiScmProvidersGetCollectionRequest) Limit(limit int32) ScmProvidersApiScmProvidersGetCollectionRequest {
	r.limit = &limit
	return r
}

// the fields to include for returned resources of type scmRepositories
func (r ScmProvidersApiScmProvidersGetCollectionRequest) FieldsScmRepositories(fieldsScmRepositories []string) ScmProvidersApiScmProvidersGetCollectionRequest {
	r.fieldsScmRepositories = &fieldsScmRepositories
	return r
}

func (r ScmProvidersApiScmProvidersGetCollectionRequest) Execute() (*ScmProvidersResponse, *http.Response, error) {
	return r.ApiService.ScmProvidersGetCollectionExecute(r)
}

/*
ScmProvidersGetCollection Method for ScmProvidersGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ScmProvidersApiScmProvidersGetCollectionRequest
*/
func (a *ScmProvidersApiService) ScmProvidersGetCollection(ctx context.Context) ScmProvidersApiScmProvidersGetCollectionRequest {
	return ScmProvidersApiScmProvidersGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ScmProvidersResponse
func (a *ScmProvidersApiService) ScmProvidersGetCollectionExecute(r ScmProvidersApiScmProvidersGetCollectionRequest) (*ScmProvidersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScmProvidersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScmProvidersApiService.ScmProvidersGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/scmProviders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsScmProviders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmProviders]", r.fieldsScmProviders, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.fieldsScmRepositories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmRepositories]", r.fieldsScmRepositories, "csv")
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

type ScmProvidersApiScmProvidersGetInstanceRequest struct {
	ctx context.Context
	ApiService *ScmProvidersApiService
	id string
	fieldsScmProviders *[]string
	fieldsScmRepositories *[]string
}

// the fields to include for returned resources of type scmProviders
func (r ScmProvidersApiScmProvidersGetInstanceRequest) FieldsScmProviders(fieldsScmProviders []string) ScmProvidersApiScmProvidersGetInstanceRequest {
	r.fieldsScmProviders = &fieldsScmProviders
	return r
}

// the fields to include for returned resources of type scmRepositories
func (r ScmProvidersApiScmProvidersGetInstanceRequest) FieldsScmRepositories(fieldsScmRepositories []string) ScmProvidersApiScmProvidersGetInstanceRequest {
	r.fieldsScmRepositories = &fieldsScmRepositories
	return r
}

func (r ScmProvidersApiScmProvidersGetInstanceRequest) Execute() (*ScmProviderResponse, *http.Response, error) {
	return r.ApiService.ScmProvidersGetInstanceExecute(r)
}

/*
ScmProvidersGetInstance Method for ScmProvidersGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ScmProvidersApiScmProvidersGetInstanceRequest
*/
func (a *ScmProvidersApiService) ScmProvidersGetInstance(ctx context.Context, id string) ScmProvidersApiScmProvidersGetInstanceRequest {
	return ScmProvidersApiScmProvidersGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ScmProviderResponse
func (a *ScmProvidersApiService) ScmProvidersGetInstanceExecute(r ScmProvidersApiScmProvidersGetInstanceRequest) (*ScmProviderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScmProviderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScmProvidersApiService.ScmProvidersGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/scmProviders/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsScmProviders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmProviders]", r.fieldsScmProviders, "csv")
	}
	if r.fieldsScmRepositories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmRepositories]", r.fieldsScmRepositories, "csv")
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

type ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *ScmProvidersApiService
	id string
	filterId *[]string
	fieldsScmGitReferences *[]string
	fieldsScmProviders *[]string
	fieldsScmRepositories *[]string
	limit *int32
	include *[]string
}

// filter by id(s)
func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) FilterId(filterId []string) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	r.filterId = &filterId
	return r
}

// the fields to include for returned resources of type scmGitReferences
func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) FieldsScmGitReferences(fieldsScmGitReferences []string) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	r.fieldsScmGitReferences = &fieldsScmGitReferences
	return r
}

// the fields to include for returned resources of type scmProviders
func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) FieldsScmProviders(fieldsScmProviders []string) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	r.fieldsScmProviders = &fieldsScmProviders
	return r
}

// the fields to include for returned resources of type scmRepositories
func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) FieldsScmRepositories(fieldsScmRepositories []string) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	r.fieldsScmRepositories = &fieldsScmRepositories
	return r
}

// maximum resources per page
func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) Limit(limit int32) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) Include(include []string) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) Execute() (*ScmRepositoriesResponse, *http.Response, error) {
	return r.ApiService.ScmProvidersRepositoriesGetToManyRelatedExecute(r)
}

/*
ScmProvidersRepositoriesGetToManyRelated Method for ScmProvidersRepositoriesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest
*/
func (a *ScmProvidersApiService) ScmProvidersRepositoriesGetToManyRelated(ctx context.Context, id string) ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest {
	return ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ScmRepositoriesResponse
func (a *ScmProvidersApiService) ScmProvidersRepositoriesGetToManyRelatedExecute(r ScmProvidersApiScmProvidersRepositoriesGetToManyRelatedRequest) (*ScmRepositoriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScmRepositoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScmProvidersApiService.ScmProvidersRepositoriesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/scmProviders/{id}/repositories"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[id]", r.filterId, "csv")
	}
	if r.fieldsScmGitReferences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmGitReferences]", r.fieldsScmGitReferences, "csv")
	}
	if r.fieldsScmProviders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmProviders]", r.fieldsScmProviders, "csv")
	}
	if r.fieldsScmRepositories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmRepositories]", r.fieldsScmRepositories, "csv")
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
