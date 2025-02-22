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

// checks if the AppPreviewUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewUpdateRequest{}

// AppPreviewUpdateRequest struct for AppPreviewUpdateRequest
type AppPreviewUpdateRequest struct {
	Data AppPreviewUpdateRequestData `json:"data"`
}

// NewAppPreviewUpdateRequest instantiates a new AppPreviewUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewUpdateRequest(data AppPreviewUpdateRequestData) *AppPreviewUpdateRequest {
	this := AppPreviewUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppPreviewUpdateRequestWithDefaults instantiates a new AppPreviewUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewUpdateRequestWithDefaults() *AppPreviewUpdateRequest {
	this := AppPreviewUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppPreviewUpdateRequest) GetData() AppPreviewUpdateRequestData {
	if o == nil {
		var ret AppPreviewUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequest) GetDataOk() (*AppPreviewUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPreviewUpdateRequest) SetData(v AppPreviewUpdateRequestData) {
	o.Data = v
}

func (o AppPreviewUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppPreviewUpdateRequest struct {
	value *AppPreviewUpdateRequest
	isSet bool
}

func (v NullableAppPreviewUpdateRequest) Get() *AppPreviewUpdateRequest {
	return v.value
}

func (v *NullableAppPreviewUpdateRequest) Set(val *AppPreviewUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewUpdateRequest(val *AppPreviewUpdateRequest) *NullableAppPreviewUpdateRequest {
	return &NullableAppPreviewUpdateRequest{value: val, isSet: true}
}

func (v NullableAppPreviewUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


