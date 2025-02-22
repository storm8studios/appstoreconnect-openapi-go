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

// checks if the AppScreenshotCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotCreateRequest{}

// AppScreenshotCreateRequest struct for AppScreenshotCreateRequest
type AppScreenshotCreateRequest struct {
	Data AppScreenshotCreateRequestData `json:"data"`
}

// NewAppScreenshotCreateRequest instantiates a new AppScreenshotCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotCreateRequest(data AppScreenshotCreateRequestData) *AppScreenshotCreateRequest {
	this := AppScreenshotCreateRequest{}
	this.Data = data
	return &this
}

// NewAppScreenshotCreateRequestWithDefaults instantiates a new AppScreenshotCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotCreateRequestWithDefaults() *AppScreenshotCreateRequest {
	this := AppScreenshotCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppScreenshotCreateRequest) GetData() AppScreenshotCreateRequestData {
	if o == nil {
		var ret AppScreenshotCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotCreateRequest) GetDataOk() (*AppScreenshotCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppScreenshotCreateRequest) SetData(v AppScreenshotCreateRequestData) {
	o.Data = v
}

func (o AppScreenshotCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppScreenshotCreateRequest struct {
	value *AppScreenshotCreateRequest
	isSet bool
}

func (v NullableAppScreenshotCreateRequest) Get() *AppScreenshotCreateRequest {
	return v.value
}

func (v *NullableAppScreenshotCreateRequest) Set(val *AppScreenshotCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotCreateRequest(val *AppScreenshotCreateRequest) *NullableAppScreenshotCreateRequest {
	return &NullableAppScreenshotCreateRequest{value: val, isSet: true}
}

func (v NullableAppScreenshotCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


