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

// checks if the AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience{}

// AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience struct for AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience
type AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience struct {
	Data AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData `json:"data"`
}

// NewAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience instantiates a new AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience(data AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData) *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience {
	this := AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience{}
	this.Data = data
	return &this
}

// NewAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperienceWithDefaults instantiates a new AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperienceWithDefaults() *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience {
	this := AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) GetData() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData {
	if o == nil {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) GetDataOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) SetData(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData) {
	o.Data = v
}

func (o AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience struct {
	value *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) Get() *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) Set(val *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience(val *AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience {
	return &NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


