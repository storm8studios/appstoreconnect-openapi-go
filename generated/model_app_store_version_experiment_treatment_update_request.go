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

// checks if the AppStoreVersionExperimentTreatmentUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentUpdateRequest{}

// AppStoreVersionExperimentTreatmentUpdateRequest struct for AppStoreVersionExperimentTreatmentUpdateRequest
type AppStoreVersionExperimentTreatmentUpdateRequest struct {
	Data AppStoreVersionExperimentTreatmentUpdateRequestData `json:"data"`
}

// NewAppStoreVersionExperimentTreatmentUpdateRequest instantiates a new AppStoreVersionExperimentTreatmentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentUpdateRequest(data AppStoreVersionExperimentTreatmentUpdateRequestData) *AppStoreVersionExperimentTreatmentUpdateRequest {
	this := AppStoreVersionExperimentTreatmentUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionExperimentTreatmentUpdateRequestWithDefaults instantiates a new AppStoreVersionExperimentTreatmentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentUpdateRequestWithDefaults() *AppStoreVersionExperimentTreatmentUpdateRequest {
	this := AppStoreVersionExperimentTreatmentUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionExperimentTreatmentUpdateRequest) GetData() AppStoreVersionExperimentTreatmentUpdateRequestData {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentUpdateRequest) GetDataOk() (*AppStoreVersionExperimentTreatmentUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionExperimentTreatmentUpdateRequest) SetData(v AppStoreVersionExperimentTreatmentUpdateRequestData) {
	o.Data = v
}

func (o AppStoreVersionExperimentTreatmentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentUpdateRequest struct {
	value *AppStoreVersionExperimentTreatmentUpdateRequest
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentUpdateRequest) Get() *AppStoreVersionExperimentTreatmentUpdateRequest {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentUpdateRequest) Set(val *AppStoreVersionExperimentTreatmentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentUpdateRequest(val *AppStoreVersionExperimentTreatmentUpdateRequest) *NullableAppStoreVersionExperimentTreatmentUpdateRequest {
	return &NullableAppStoreVersionExperimentTreatmentUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


