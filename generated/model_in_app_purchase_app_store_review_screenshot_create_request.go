/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
)

// checks if the InAppPurchaseAppStoreReviewScreenshotCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAppStoreReviewScreenshotCreateRequest{}

// InAppPurchaseAppStoreReviewScreenshotCreateRequest struct for InAppPurchaseAppStoreReviewScreenshotCreateRequest
type InAppPurchaseAppStoreReviewScreenshotCreateRequest struct {
	Data InAppPurchaseAppStoreReviewScreenshotCreateRequestData `json:"data"`
}

// NewInAppPurchaseAppStoreReviewScreenshotCreateRequest instantiates a new InAppPurchaseAppStoreReviewScreenshotCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAppStoreReviewScreenshotCreateRequest(data InAppPurchaseAppStoreReviewScreenshotCreateRequestData) *InAppPurchaseAppStoreReviewScreenshotCreateRequest {
	this := InAppPurchaseAppStoreReviewScreenshotCreateRequest{}
	this.Data = data
	return &this
}

// NewInAppPurchaseAppStoreReviewScreenshotCreateRequestWithDefaults instantiates a new InAppPurchaseAppStoreReviewScreenshotCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAppStoreReviewScreenshotCreateRequestWithDefaults() *InAppPurchaseAppStoreReviewScreenshotCreateRequest {
	this := InAppPurchaseAppStoreReviewScreenshotCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequest) GetData() InAppPurchaseAppStoreReviewScreenshotCreateRequestData {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshotCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequest) GetDataOk() (*InAppPurchaseAppStoreReviewScreenshotCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequest) SetData(v InAppPurchaseAppStoreReviewScreenshotCreateRequestData) {
	o.Data = v
}

func (o InAppPurchaseAppStoreReviewScreenshotCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAppStoreReviewScreenshotCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest struct {
	value *InAppPurchaseAppStoreReviewScreenshotCreateRequest
	isSet bool
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest) Get() *InAppPurchaseAppStoreReviewScreenshotCreateRequest {
	return v.value
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest) Set(val *InAppPurchaseAppStoreReviewScreenshotCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAppStoreReviewScreenshotCreateRequest(val *InAppPurchaseAppStoreReviewScreenshotCreateRequest) *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest {
	return &NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest{value: val, isSet: true}
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


