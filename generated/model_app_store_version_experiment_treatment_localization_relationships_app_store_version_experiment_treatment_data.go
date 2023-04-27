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

// checks if the AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData{}

// AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData struct for AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData
type AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData instantiates a new AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData(type_ string, id string) *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData {
	this := AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentDataWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentDataWithDefaults() *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData {
	this := AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) SetId(v string) {
	o.Id = v
}

func (o AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData struct {
	value *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) Get() *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) Set(val *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData(val *AppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) *NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData {
	return &NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


