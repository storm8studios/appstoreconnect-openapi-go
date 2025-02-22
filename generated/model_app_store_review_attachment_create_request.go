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

// checks if the AppStoreReviewAttachmentCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentCreateRequest{}

// AppStoreReviewAttachmentCreateRequest struct for AppStoreReviewAttachmentCreateRequest
type AppStoreReviewAttachmentCreateRequest struct {
	Data AppStoreReviewAttachmentCreateRequestData `json:"data"`
}

// NewAppStoreReviewAttachmentCreateRequest instantiates a new AppStoreReviewAttachmentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentCreateRequest(data AppStoreReviewAttachmentCreateRequestData) *AppStoreReviewAttachmentCreateRequest {
	this := AppStoreReviewAttachmentCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreReviewAttachmentCreateRequestWithDefaults instantiates a new AppStoreReviewAttachmentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentCreateRequestWithDefaults() *AppStoreReviewAttachmentCreateRequest {
	this := AppStoreReviewAttachmentCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreReviewAttachmentCreateRequest) GetData() AppStoreReviewAttachmentCreateRequestData {
	if o == nil {
		var ret AppStoreReviewAttachmentCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequest) GetDataOk() (*AppStoreReviewAttachmentCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreReviewAttachmentCreateRequest) SetData(v AppStoreReviewAttachmentCreateRequestData) {
	o.Data = v
}

func (o AppStoreReviewAttachmentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreReviewAttachmentCreateRequest struct {
	value *AppStoreReviewAttachmentCreateRequest
	isSet bool
}

func (v NullableAppStoreReviewAttachmentCreateRequest) Get() *AppStoreReviewAttachmentCreateRequest {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentCreateRequest) Set(val *AppStoreReviewAttachmentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentCreateRequest(val *AppStoreReviewAttachmentCreateRequest) *NullableAppStoreReviewAttachmentCreateRequest {
	return &NullableAppStoreReviewAttachmentCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


