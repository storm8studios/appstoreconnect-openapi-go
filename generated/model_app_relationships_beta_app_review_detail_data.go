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

// checks if the AppRelationshipsBetaAppReviewDetailData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsBetaAppReviewDetailData{}

// AppRelationshipsBetaAppReviewDetailData struct for AppRelationshipsBetaAppReviewDetailData
type AppRelationshipsBetaAppReviewDetailData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppRelationshipsBetaAppReviewDetailData instantiates a new AppRelationshipsBetaAppReviewDetailData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaAppReviewDetailData(type_ string, id string) *AppRelationshipsBetaAppReviewDetailData {
	this := AppRelationshipsBetaAppReviewDetailData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsBetaAppReviewDetailDataWithDefaults instantiates a new AppRelationshipsBetaAppReviewDetailData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaAppReviewDetailDataWithDefaults() *AppRelationshipsBetaAppReviewDetailData {
	this := AppRelationshipsBetaAppReviewDetailData{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsBetaAppReviewDetailData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppReviewDetailData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsBetaAppReviewDetailData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsBetaAppReviewDetailData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppReviewDetailData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsBetaAppReviewDetailData) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsBetaAppReviewDetailData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsBetaAppReviewDetailData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppRelationshipsBetaAppReviewDetailData struct {
	value *AppRelationshipsBetaAppReviewDetailData
	isSet bool
}

func (v NullableAppRelationshipsBetaAppReviewDetailData) Get() *AppRelationshipsBetaAppReviewDetailData {
	return v.value
}

func (v *NullableAppRelationshipsBetaAppReviewDetailData) Set(val *AppRelationshipsBetaAppReviewDetailData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaAppReviewDetailData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaAppReviewDetailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaAppReviewDetailData(val *AppRelationshipsBetaAppReviewDetailData) *NullableAppRelationshipsBetaAppReviewDetailData {
	return &NullableAppRelationshipsBetaAppReviewDetailData{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaAppReviewDetailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaAppReviewDetailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


