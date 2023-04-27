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

// checks if the AppClipDefaultExperienceUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceUpdateRequest{}

// AppClipDefaultExperienceUpdateRequest struct for AppClipDefaultExperienceUpdateRequest
type AppClipDefaultExperienceUpdateRequest struct {
	Data AppClipDefaultExperienceUpdateRequestData `json:"data"`
}

// NewAppClipDefaultExperienceUpdateRequest instantiates a new AppClipDefaultExperienceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceUpdateRequest(data AppClipDefaultExperienceUpdateRequestData) *AppClipDefaultExperienceUpdateRequest {
	this := AppClipDefaultExperienceUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppClipDefaultExperienceUpdateRequestWithDefaults instantiates a new AppClipDefaultExperienceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceUpdateRequestWithDefaults() *AppClipDefaultExperienceUpdateRequest {
	this := AppClipDefaultExperienceUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipDefaultExperienceUpdateRequest) GetData() AppClipDefaultExperienceUpdateRequestData {
	if o == nil {
		var ret AppClipDefaultExperienceUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceUpdateRequest) GetDataOk() (*AppClipDefaultExperienceUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipDefaultExperienceUpdateRequest) SetData(v AppClipDefaultExperienceUpdateRequestData) {
	o.Data = v
}

func (o AppClipDefaultExperienceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceUpdateRequest struct {
	value *AppClipDefaultExperienceUpdateRequest
	isSet bool
}

func (v NullableAppClipDefaultExperienceUpdateRequest) Get() *AppClipDefaultExperienceUpdateRequest {
	return v.value
}

func (v *NullableAppClipDefaultExperienceUpdateRequest) Set(val *AppClipDefaultExperienceUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceUpdateRequest(val *AppClipDefaultExperienceUpdateRequest) *NullableAppClipDefaultExperienceUpdateRequest {
	return &NullableAppClipDefaultExperienceUpdateRequest{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


