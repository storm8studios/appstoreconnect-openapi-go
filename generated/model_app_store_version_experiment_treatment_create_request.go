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

// checks if the AppStoreVersionExperimentTreatmentCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentCreateRequest{}

// AppStoreVersionExperimentTreatmentCreateRequest struct for AppStoreVersionExperimentTreatmentCreateRequest
type AppStoreVersionExperimentTreatmentCreateRequest struct {
	Data AppStoreVersionExperimentTreatmentCreateRequestData `json:"data"`
}

// NewAppStoreVersionExperimentTreatmentCreateRequest instantiates a new AppStoreVersionExperimentTreatmentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentCreateRequest(data AppStoreVersionExperimentTreatmentCreateRequestData) *AppStoreVersionExperimentTreatmentCreateRequest {
	this := AppStoreVersionExperimentTreatmentCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionExperimentTreatmentCreateRequestWithDefaults instantiates a new AppStoreVersionExperimentTreatmentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentCreateRequestWithDefaults() *AppStoreVersionExperimentTreatmentCreateRequest {
	this := AppStoreVersionExperimentTreatmentCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionExperimentTreatmentCreateRequest) GetData() AppStoreVersionExperimentTreatmentCreateRequestData {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequest) GetDataOk() (*AppStoreVersionExperimentTreatmentCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionExperimentTreatmentCreateRequest) SetData(v AppStoreVersionExperimentTreatmentCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionExperimentTreatmentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentCreateRequest struct {
	value *AppStoreVersionExperimentTreatmentCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequest) Get() *AppStoreVersionExperimentTreatmentCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequest) Set(val *AppStoreVersionExperimentTreatmentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentCreateRequest(val *AppStoreVersionExperimentTreatmentCreateRequest) *NullableAppStoreVersionExperimentTreatmentCreateRequest {
	return &NullableAppStoreVersionExperimentTreatmentCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


