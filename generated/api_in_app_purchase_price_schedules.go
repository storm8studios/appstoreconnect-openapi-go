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


// InAppPurchasePriceSchedulesApiService InAppPurchasePriceSchedulesApi service
type InAppPurchasePriceSchedulesApiService service

type InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *InAppPurchasePriceSchedulesApiService
	id string
	filterTerritory *[]string
	fieldsInAppPurchasePricePoints *[]string
	fieldsInAppPurchasePrices *[]string
	fieldsTerritories *[]string
	limit *int32
	include *[]string
}

// filter by id(s) of related &#39;territory&#39;
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) FilterTerritory(filterTerritory []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.filterTerritory = &filterTerritory
	return r
}

// the fields to include for returned resources of type inAppPurchasePricePoints
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.fieldsInAppPurchasePricePoints = &fieldsInAppPurchasePricePoints
	return r
}

// the fields to include for returned resources of type inAppPurchasePrices
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) FieldsInAppPurchasePrices(fieldsInAppPurchasePrices []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.fieldsInAppPurchasePrices = &fieldsInAppPurchasePrices
	return r
}

// the fields to include for returned resources of type territories
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) Limit(limit int32) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) Include(include []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) Execute() (*InAppPurchasePricesResponse, *http.Response, error) {
	return r.ApiService.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedExecute(r)
}

/*
InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated Method for InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest
*/
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated(ctx context.Context, id string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InAppPurchasePricesResponse
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedExecute(r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest) (*InAppPurchasePricesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchasePricesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchasePriceSchedulesApiService.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchasePriceSchedules/{id}/automaticPrices"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterTerritory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[territory]", r.filterTerritory, "csv")
	}
	if r.fieldsInAppPurchasePricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchasePricePoints]", r.fieldsInAppPurchasePricePoints, "csv")
	}
	if r.fieldsInAppPurchasePrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchasePrices]", r.fieldsInAppPurchasePrices, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
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

type InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *InAppPurchasePriceSchedulesApiService
	id string
	fieldsTerritories *[]string
}

// the fields to include for returned resources of type territories
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest) FieldsTerritories(fieldsTerritories []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest) Execute() (*TerritoryResponse, *http.Response, error) {
	return r.ApiService.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedExecute(r)
}

/*
InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated Method for InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest
*/
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated(ctx context.Context, id string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest {
	return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TerritoryResponse
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedExecute(r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest) (*TerritoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TerritoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchasePriceSchedulesApiService.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchasePriceSchedules/{id}/baseTerritory"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
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

type InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchasePriceSchedulesApiService
	inAppPurchasePriceScheduleCreateRequest *InAppPurchasePriceScheduleCreateRequest
}

// InAppPurchasePriceSchedule representation
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest) InAppPurchasePriceScheduleCreateRequest(inAppPurchasePriceScheduleCreateRequest InAppPurchasePriceScheduleCreateRequest) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest {
	r.inAppPurchasePriceScheduleCreateRequest = &inAppPurchasePriceScheduleCreateRequest
	return r
}

func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest) Execute() (*InAppPurchasePriceScheduleResponse, *http.Response, error) {
	return r.ApiService.InAppPurchasePriceSchedulesCreateInstanceExecute(r)
}

/*
InAppPurchasePriceSchedulesCreateInstance Method for InAppPurchasePriceSchedulesCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest
*/
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesCreateInstance(ctx context.Context) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest {
	return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InAppPurchasePriceScheduleResponse
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesCreateInstanceExecute(r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesCreateInstanceRequest) (*InAppPurchasePriceScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchasePriceScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchasePriceSchedulesApiService.InAppPurchasePriceSchedulesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchasePriceSchedules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inAppPurchasePriceScheduleCreateRequest == nil {
		return localVarReturnValue, nil, reportError("inAppPurchasePriceScheduleCreateRequest is required and must be specified")
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
	localVarPostBody = r.inAppPurchasePriceScheduleCreateRequest
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

type InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchasePriceSchedulesApiService
	id string
	fieldsInAppPurchasePriceSchedules *[]string
	include *[]string
	fieldsInAppPurchasePrices *[]string
	fieldsTerritories *[]string
	limitAutomaticPrices *int32
	limitManualPrices *int32
}

// the fields to include for returned resources of type inAppPurchasePriceSchedules
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	r.fieldsInAppPurchasePriceSchedules = &fieldsInAppPurchasePriceSchedules
	return r
}

// comma-separated list of relationships to include
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) Include(include []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type inAppPurchasePrices
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) FieldsInAppPurchasePrices(fieldsInAppPurchasePrices []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	r.fieldsInAppPurchasePrices = &fieldsInAppPurchasePrices
	return r
}

// the fields to include for returned resources of type territories
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) FieldsTerritories(fieldsTerritories []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum number of related automaticPrices returned (when they are included)
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) LimitAutomaticPrices(limitAutomaticPrices int32) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	r.limitAutomaticPrices = &limitAutomaticPrices
	return r
}

// maximum number of related manualPrices returned (when they are included)
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) LimitManualPrices(limitManualPrices int32) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	r.limitManualPrices = &limitManualPrices
	return r
}

func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) Execute() (*InAppPurchasePriceScheduleResponse, *http.Response, error) {
	return r.ApiService.InAppPurchasePriceSchedulesGetInstanceExecute(r)
}

/*
InAppPurchasePriceSchedulesGetInstance Method for InAppPurchasePriceSchedulesGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest
*/
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesGetInstance(ctx context.Context, id string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest {
	return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InAppPurchasePriceScheduleResponse
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesGetInstanceExecute(r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesGetInstanceRequest) (*InAppPurchasePriceScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchasePriceScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchasePriceSchedulesApiService.InAppPurchasePriceSchedulesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchasePriceSchedules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsInAppPurchasePriceSchedules != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchasePriceSchedules]", r.fieldsInAppPurchasePriceSchedules, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsInAppPurchasePrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchasePrices]", r.fieldsInAppPurchasePrices, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
	}
	if r.limitAutomaticPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[automaticPrices]", r.limitAutomaticPrices, "")
	}
	if r.limitManualPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[manualPrices]", r.limitManualPrices, "")
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

type InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *InAppPurchasePriceSchedulesApiService
	id string
	filterTerritory *[]string
	fieldsInAppPurchasePricePoints *[]string
	fieldsInAppPurchasePrices *[]string
	fieldsTerritories *[]string
	limit *int32
	include *[]string
}

// filter by id(s) of related &#39;territory&#39;
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) FilterTerritory(filterTerritory []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	r.filterTerritory = &filterTerritory
	return r
}

// the fields to include for returned resources of type inAppPurchasePricePoints
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	r.fieldsInAppPurchasePricePoints = &fieldsInAppPurchasePricePoints
	return r
}

// the fields to include for returned resources of type inAppPurchasePrices
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) FieldsInAppPurchasePrices(fieldsInAppPurchasePrices []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	r.fieldsInAppPurchasePrices = &fieldsInAppPurchasePrices
	return r
}

// the fields to include for returned resources of type territories
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) Limit(limit int32) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) Include(include []string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) Execute() (*InAppPurchasePricesResponse, *http.Response, error) {
	return r.ApiService.InAppPurchasePriceSchedulesManualPricesGetToManyRelatedExecute(r)
}

/*
InAppPurchasePriceSchedulesManualPricesGetToManyRelated Method for InAppPurchasePriceSchedulesManualPricesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest
*/
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesManualPricesGetToManyRelated(ctx context.Context, id string) InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest {
	return InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InAppPurchasePricesResponse
func (a *InAppPurchasePriceSchedulesApiService) InAppPurchasePriceSchedulesManualPricesGetToManyRelatedExecute(r InAppPurchasePriceSchedulesApiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest) (*InAppPurchasePricesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchasePricesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchasePriceSchedulesApiService.InAppPurchasePriceSchedulesManualPricesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchasePriceSchedules/{id}/manualPrices"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterTerritory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[territory]", r.filterTerritory, "csv")
	}
	if r.fieldsInAppPurchasePricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchasePricePoints]", r.fieldsInAppPurchasePricePoints, "csv")
	}
	if r.fieldsInAppPurchasePrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchasePrices]", r.fieldsInAppPurchasePrices, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
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
