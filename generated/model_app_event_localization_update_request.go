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

// checks if the AppEventLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationUpdateRequest{}

// AppEventLocalizationUpdateRequest struct for AppEventLocalizationUpdateRequest
type AppEventLocalizationUpdateRequest struct {
	Data AppEventLocalizationUpdateRequestData `json:"data"`
}

// NewAppEventLocalizationUpdateRequest instantiates a new AppEventLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationUpdateRequest(data AppEventLocalizationUpdateRequestData) *AppEventLocalizationUpdateRequest {
	this := AppEventLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppEventLocalizationUpdateRequestWithDefaults instantiates a new AppEventLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationUpdateRequestWithDefaults() *AppEventLocalizationUpdateRequest {
	this := AppEventLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventLocalizationUpdateRequest) GetData() AppEventLocalizationUpdateRequestData {
	if o == nil {
		var ret AppEventLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationUpdateRequest) GetDataOk() (*AppEventLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEventLocalizationUpdateRequest) SetData(v AppEventLocalizationUpdateRequestData) {
	o.Data = v
}

func (o AppEventLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppEventLocalizationUpdateRequest struct {
	value *AppEventLocalizationUpdateRequest
	isSet bool
}

func (v NullableAppEventLocalizationUpdateRequest) Get() *AppEventLocalizationUpdateRequest {
	return v.value
}

func (v *NullableAppEventLocalizationUpdateRequest) Set(val *AppEventLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationUpdateRequest(val *AppEventLocalizationUpdateRequest) *NullableAppEventLocalizationUpdateRequest {
	return &NullableAppEventLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableAppEventLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


