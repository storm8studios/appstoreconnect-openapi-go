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

// checks if the AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData{}

// AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData struct for AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData
type AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData instantiates a new AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData(type_ string, id string) *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData {
	this := AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailDataWithDefaults instantiates a new AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailDataWithDefaults() *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData {
	this := AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) SetId(v string) {
	o.Id = v
}

func (o AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData struct {
	value *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData
	isSet bool
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) Get() *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData {
	return v.value
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) Set(val *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData(val *AppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) *NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData {
	return &NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipAppStoreReviewDetailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


