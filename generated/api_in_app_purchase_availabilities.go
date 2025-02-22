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


// InAppPurchaseAvailabilitiesApiService InAppPurchaseAvailabilitiesApi service
type InAppPurchaseAvailabilitiesApiService service

type InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAvailabilitiesApiService
	id string
	fieldsTerritories *[]string
	limit *int32
}

// the fields to include for returned resources of type territories
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) Limit(limit int32) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) Execute() (*TerritoriesResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedExecute(r)
}

/*
InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelated Method for InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest
*/
func (a *InAppPurchaseAvailabilitiesApiService) InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelated(ctx context.Context, id string) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	return InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TerritoriesResponse
func (a *InAppPurchaseAvailabilitiesApiService) InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedExecute(r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) (*TerritoriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TerritoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAvailabilitiesApiService.InAppPurchaseAvailabilitiesAvailableTerritoriesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAvailabilities/{id}/availableTerritories"
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

type InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAvailabilitiesApiService
	inAppPurchaseAvailabilityCreateRequest *InAppPurchaseAvailabilityCreateRequest
}

// InAppPurchaseAvailability representation
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest) InAppPurchaseAvailabilityCreateRequest(inAppPurchaseAvailabilityCreateRequest InAppPurchaseAvailabilityCreateRequest) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest {
	r.inAppPurchaseAvailabilityCreateRequest = &inAppPurchaseAvailabilityCreateRequest
	return r
}

func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest) Execute() (*InAppPurchaseAvailabilityResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseAvailabilitiesCreateInstanceExecute(r)
}

/*
InAppPurchaseAvailabilitiesCreateInstance Method for InAppPurchaseAvailabilitiesCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest
*/
func (a *InAppPurchaseAvailabilitiesApiService) InAppPurchaseAvailabilitiesCreateInstance(ctx context.Context) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest {
	return InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InAppPurchaseAvailabilityResponse
func (a *InAppPurchaseAvailabilitiesApiService) InAppPurchaseAvailabilitiesCreateInstanceExecute(r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesCreateInstanceRequest) (*InAppPurchaseAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchaseAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAvailabilitiesApiService.InAppPurchaseAvailabilitiesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAvailabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inAppPurchaseAvailabilityCreateRequest == nil {
		return localVarReturnValue, nil, reportError("inAppPurchaseAvailabilityCreateRequest is required and must be specified")
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
	localVarPostBody = r.inAppPurchaseAvailabilityCreateRequest
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

type InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAvailabilitiesApiService
	id string
	fieldsInAppPurchaseAvailabilities *[]string
	include *[]string
	fieldsTerritories *[]string
	limitAvailableTerritories *int32
}

// the fields to include for returned resources of type inAppPurchaseAvailabilities
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest) FieldsInAppPurchaseAvailabilities(fieldsInAppPurchaseAvailabilities []string) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest {
	r.fieldsInAppPurchaseAvailabilities = &fieldsInAppPurchaseAvailabilities
	return r
}

// comma-separated list of relationships to include
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest) Include(include []string) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type territories
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest) FieldsTerritories(fieldsTerritories []string) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum number of related availableTerritories returned (when they are included)
func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest) LimitAvailableTerritories(limitAvailableTerritories int32) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest {
	r.limitAvailableTerritories = &limitAvailableTerritories
	return r
}

func (r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest) Execute() (*InAppPurchaseAvailabilityResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseAvailabilitiesGetInstanceExecute(r)
}

/*
InAppPurchaseAvailabilitiesGetInstance Method for InAppPurchaseAvailabilitiesGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest
*/
func (a *InAppPurchaseAvailabilitiesApiService) InAppPurchaseAvailabilitiesGetInstance(ctx context.Context, id string) InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest {
	return InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InAppPurchaseAvailabilityResponse
func (a *InAppPurchaseAvailabilitiesApiService) InAppPurchaseAvailabilitiesGetInstanceExecute(r InAppPurchaseAvailabilitiesApiInAppPurchaseAvailabilitiesGetInstanceRequest) (*InAppPurchaseAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchaseAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAvailabilitiesApiService.InAppPurchaseAvailabilitiesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAvailabilities/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsInAppPurchaseAvailabilities != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchaseAvailabilities]", r.fieldsInAppPurchaseAvailabilities, "csv")
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
