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

// checks if the AppStoreVersionPhasedReleaseUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPhasedReleaseUpdateRequest{}

// AppStoreVersionPhasedReleaseUpdateRequest struct for AppStoreVersionPhasedReleaseUpdateRequest
type AppStoreVersionPhasedReleaseUpdateRequest struct {
	Data AppStoreVersionPhasedReleaseUpdateRequestData `json:"data"`
}

// NewAppStoreVersionPhasedReleaseUpdateRequest instantiates a new AppStoreVersionPhasedReleaseUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPhasedReleaseUpdateRequest(data AppStoreVersionPhasedReleaseUpdateRequestData) *AppStoreVersionPhasedReleaseUpdateRequest {
	this := AppStoreVersionPhasedReleaseUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionPhasedReleaseUpdateRequestWithDefaults instantiates a new AppStoreVersionPhasedReleaseUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPhasedReleaseUpdateRequestWithDefaults() *AppStoreVersionPhasedReleaseUpdateRequest {
	this := AppStoreVersionPhasedReleaseUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionPhasedReleaseUpdateRequest) GetData() AppStoreVersionPhasedReleaseUpdateRequestData {
	if o == nil {
		var ret AppStoreVersionPhasedReleaseUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseUpdateRequest) GetDataOk() (*AppStoreVersionPhasedReleaseUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionPhasedReleaseUpdateRequest) SetData(v AppStoreVersionPhasedReleaseUpdateRequestData) {
	o.Data = v
}

func (o AppStoreVersionPhasedReleaseUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPhasedReleaseUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionPhasedReleaseUpdateRequest struct {
	value *AppStoreVersionPhasedReleaseUpdateRequest
	isSet bool
}

func (v NullableAppStoreVersionPhasedReleaseUpdateRequest) Get() *AppStoreVersionPhasedReleaseUpdateRequest {
	return v.value
}

func (v *NullableAppStoreVersionPhasedReleaseUpdateRequest) Set(val *AppStoreVersionPhasedReleaseUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPhasedReleaseUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPhasedReleaseUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPhasedReleaseUpdateRequest(val *AppStoreVersionPhasedReleaseUpdateRequest) *NullableAppStoreVersionPhasedReleaseUpdateRequest {
	return &NullableAppStoreVersionPhasedReleaseUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionPhasedReleaseUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPhasedReleaseUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


