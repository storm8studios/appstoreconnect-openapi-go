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

// checks if the AppScreenshotUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotUpdateRequest{}

// AppScreenshotUpdateRequest struct for AppScreenshotUpdateRequest
type AppScreenshotUpdateRequest struct {
	Data AppScreenshotUpdateRequestData `json:"data"`
}

// NewAppScreenshotUpdateRequest instantiates a new AppScreenshotUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotUpdateRequest(data AppScreenshotUpdateRequestData) *AppScreenshotUpdateRequest {
	this := AppScreenshotUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppScreenshotUpdateRequestWithDefaults instantiates a new AppScreenshotUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotUpdateRequestWithDefaults() *AppScreenshotUpdateRequest {
	this := AppScreenshotUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppScreenshotUpdateRequest) GetData() AppScreenshotUpdateRequestData {
	if o == nil {
		var ret AppScreenshotUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotUpdateRequest) GetDataOk() (*AppScreenshotUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppScreenshotUpdateRequest) SetData(v AppScreenshotUpdateRequestData) {
	o.Data = v
}

func (o AppScreenshotUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppScreenshotUpdateRequest struct {
	value *AppScreenshotUpdateRequest
	isSet bool
}

func (v NullableAppScreenshotUpdateRequest) Get() *AppScreenshotUpdateRequest {
	return v.value
}

func (v *NullableAppScreenshotUpdateRequest) Set(val *AppScreenshotUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotUpdateRequest(val *AppScreenshotUpdateRequest) *NullableAppScreenshotUpdateRequest {
	return &NullableAppScreenshotUpdateRequest{value: val, isSet: true}
}

func (v NullableAppScreenshotUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


