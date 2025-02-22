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


// AppAvailabilitiesApiService AppAvailabilitiesApi service
type AppAvailabilitiesApiService service

type AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppAvailabilitiesApiService
	id string
	fieldsTerritories *[]string
	limit *int32
}

// the fields to include for returned resources of type territories
func (r AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) Limit(limit int32) AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) Execute() (*TerritoriesResponse, *http.Response, error) {
	return r.ApiService.AppAvailabilitiesAvailableTerritoriesGetToManyRelatedExecute(r)
}

/*
AppAvailabilitiesAvailableTerritoriesGetToManyRelated Method for AppAvailabilitiesAvailableTerritoriesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest
*/
func (a *AppAvailabilitiesApiService) AppAvailabilitiesAvailableTerritoriesGetToManyRelated(ctx context.Context, id string) AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	return AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TerritoriesResponse
func (a *AppAvailabilitiesApiService) AppAvailabilitiesAvailableTerritoriesGetToManyRelatedExecute(r AppAvailabilitiesApiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) (*TerritoriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TerritoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAvailabilitiesApiService.AppAvailabilitiesAvailableTerritoriesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appAvailabilities/{id}/availableTerritories"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
	}
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

type AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppAvailabilitiesApiService
	appAvailabilityCreateRequest *AppAvailabilityCreateRequest
}

// AppAvailability representation
func (r AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest) AppAvailabilityCreateRequest(appAvailabilityCreateRequest AppAvailabilityCreateRequest) AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest {
	r.appAvailabilityCreateRequest = &appAvailabilityCreateRequest
	return r
}

func (r AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest) Execute() (*AppAvailabilityResponse, *http.Response, error) {
	return r.ApiService.AppAvailabilitiesCreateInstanceExecute(r)
}

/*
AppAvailabilitiesCreateInstance Method for AppAvailabilitiesCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest
*/
func (a *AppAvailabilitiesApiService) AppAvailabilitiesCreateInstance(ctx context.Context) AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest {
	return AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppAvailabilityResponse
func (a *AppAvailabilitiesApiService) AppAvailabilitiesCreateInstanceExecute(r AppAvailabilitiesApiAppAvailabilitiesCreateInstanceRequest) (*AppAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAvailabilitiesApiService.AppAvailabilitiesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appAvailabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appAvailabilityCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appAvailabilityCreateRequest is required and must be specified")
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
	localVarPostBody = r.appAvailabilityCreateRequest
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

type AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest struct {
	ctx context.Context
	ApiService *AppAvailabilitiesApiService
	id string
	fieldsAppAvailabilities *[]string
	include *[]string
	fieldsTerritories *[]string
	limitAvailableTerritories *int32
}

// the fields to include for returned resources of type appAvailabilities
func (r AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest) FieldsAppAvailabilities(fieldsAppAvailabilities []string) AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest {
	r.fieldsAppAvailabilities = &fieldsAppAvailabilities
	return r
}

// comma-separated list of relationships to include
func (r AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest) Include(include []string) AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type territories
func (r AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest) FieldsTerritories(fieldsTerritories []string) AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum number of related availableTerritories returned (when they are included)
func (r AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest) LimitAvailableTerritories(limitAvailableTerritories int32) AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest {
	r.limitAvailableTerritories = &limitAvailableTerritories
	return r
}

func (r AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest) Execute() (*AppAvailabilityResponse, *http.Response, error) {
	return r.ApiService.AppAvailabilitiesGetInstanceExecute(r)
}

/*
AppAvailabilitiesGetInstance Method for AppAvailabilitiesGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest
*/
func (a *AppAvailabilitiesApiService) AppAvailabilitiesGetInstance(ctx context.Context, id string) AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest {
	return AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppAvailabilityResponse
func (a *AppAvailabilitiesApiService) AppAvailabilitiesGetInstanceExecute(r AppAvailabilitiesApiAppAvailabilitiesGetInstanceRequest) (*AppAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAvailabilitiesApiService.AppAvailabilitiesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appAvailabilities/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppAvailabilities != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appAvailabilities]", r.fieldsAppAvailabilities, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
	}
	if r.limitAvailableTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[availableTerritories]", r.limitAvailableTerritories, "")
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
