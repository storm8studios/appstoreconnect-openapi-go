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

// checks if the AppStoreVersionExperimentTreatmentRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentRelationships{}

// AppStoreVersionExperimentTreatmentRelationships struct for AppStoreVersionExperimentTreatmentRelationships
type AppStoreVersionExperimentTreatmentRelationships struct {
	AppStoreVersionExperiment *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperiment,omitempty"`
	AppStoreVersionExperimentTreatmentLocalizations *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentTreatmentLocalizations `json:"appStoreVersionExperimentTreatmentLocalizations,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentRelationships instantiates a new AppStoreVersionExperimentTreatmentRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentRelationships() *AppStoreVersionExperimentTreatmentRelationships {
	this := AppStoreVersionExperimentTreatmentRelationships{}
	return &this
}

// NewAppStoreVersionExperimentTreatmentRelationshipsWithDefaults instantiates a new AppStoreVersionExperimentTreatmentRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentRelationshipsWithDefaults() *AppStoreVersionExperimentTreatmentRelationships {
	this := AppStoreVersionExperimentTreatmentRelationships{}
	return &this
}

// GetAppStoreVersionExperiment returns the AppStoreVersionExperiment field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentRelationships) GetAppStoreVersionExperiment() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperiment
}

// GetAppStoreVersionExperimentOk returns a tuple with the AppStoreVersionExperiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentRelationships) GetAppStoreVersionExperimentOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		return nil, false
	}
	return o.AppStoreVersionExperiment, true
}

// HasAppStoreVersionExperiment returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentRelationships) HasAppStoreVersionExperiment() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperiment) {
		return true
	}

	return false
}

// SetAppStoreVersionExperiment gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperiment field.
func (o *AppStoreVersionExperimentTreatmentRelationships) SetAppStoreVersionExperiment(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperiment = &v
}

// GetAppStoreVersionExperimentTreatmentLocalizations returns the AppStoreVersionExperimentTreatmentLocalizations field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentRelationships) GetAppStoreVersionExperimentTreatmentLocalizations() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentTreatmentLocalizations {
	if o == nil || IsNil(o.AppStoreVersionExperimentTreatmentLocalizations) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentTreatmentLocalizations
		return ret
	}
	return *o.AppStoreVersionExperimentTreatmentLocalizations
}

// GetAppStoreVersionExperimentTreatmentLocalizationsOk returns a tuple with the AppStoreVersionExperimentTreatmentLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentRelationships) GetAppStoreVersionExperimentTreatmentLocalizationsOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentTreatmentLocalizations, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperimentTreatmentLocalizations) {
		return nil, false
	}
	return o.AppStoreVersionExperimentTreatmentLocalizations, true
}

// HasAppStoreVersionExperimentTreatmentLocalizations returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentRelationships) HasAppStoreVersionExperimentTreatmentLocalizations() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperimentTreatmentLocalizations) {
		return true
	}

	return false
}

// SetAppStoreVersionExperimentTreatmentLocalizations gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentTreatmentLocalizations and assigns it to the AppStoreVersionExperimentTreatmentLocalizations field.
func (o *AppStoreVersionExperimentTreatmentRelationships) SetAppStoreVersionExperimentTreatmentLocalizations(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentTreatmentLocalizations) {
	o.AppStoreVersionExperimentTreatmentLocalizations = &v
}

func (o AppStoreVersionExperimentTreatmentRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreVersionExperiment) {
		toSerialize["appStoreVersionExperiment"] = o.AppStoreVersionExperiment
	}
	if !IsNil(o.AppStoreVersionExperimentTreatmentLocalizations) {
		toSerialize["appStoreVersionExperimentTreatmentLocalizations"] = o.AppStoreVersionExperimentTreatmentLocalizations
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentRelationships struct {
	value *AppStoreVersionExperimentTreatmentRelationships
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentRelationships) Get() *AppStoreVersionExperimentTreatmentRelationships {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationships) Set(val *AppStoreVersionExperimentTreatmentRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentRelationships(val *AppStoreVersionExperimentTreatmentRelationships) *NullableAppStoreVersionExperimentTreatmentRelationships {
	return &NullableAppStoreVersionExperimentTreatmentRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


