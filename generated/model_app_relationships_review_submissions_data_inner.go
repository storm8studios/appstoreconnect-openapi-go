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

// checks if the AppRelationshipsReviewSubmissionsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsReviewSubmissionsDataInner{}

// AppRelationshipsReviewSubmissionsDataInner struct for AppRelationshipsReviewSubmissionsDataInner
type AppRelationshipsReviewSubmissionsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppRelationshipsReviewSubmissionsDataInner instantiates a new AppRelationshipsReviewSubmissionsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsReviewSubmissionsDataInner(type_ string, id string) *AppRelationshipsReviewSubmissionsDataInner {
	this := AppRelationshipsReviewSubmissionsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsReviewSubmissionsDataInnerWithDefaults instantiates a new AppRelationshipsReviewSubmissionsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsReviewSubmissionsDataInnerWithDefaults() *AppRelationshipsReviewSubmissionsDataInner {
	this := AppRelationshipsReviewSubmissionsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsReviewSubmissionsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsReviewSubmissionsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsReviewSubmissionsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsReviewSubmissionsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsReviewSubmissionsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsReviewSubmissionsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsReviewSubmissionsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsReviewSubmissionsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppRelationshipsReviewSubmissionsDataInner struct {
	value *AppRelationshipsReviewSubmissionsDataInner
	isSet bool
}

func (v NullableAppRelationshipsReviewSubmissionsDataInner) Get() *AppRelationshipsReviewSubmissionsDataInner {
	return v.value
}

func (v *NullableAppRelationshipsReviewSubmissionsDataInner) Set(val *AppRelationshipsReviewSubmissionsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsReviewSubmissionsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsReviewSubmissionsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsReviewSubmissionsDataInner(val *AppRelationshipsReviewSubmissionsDataInner) *NullableAppRelationshipsReviewSubmissionsDataInner {
	return &NullableAppRelationshipsReviewSubmissionsDataInner{value: val, isSet: true}
}

func (v NullableAppRelationshipsReviewSubmissionsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsReviewSubmissionsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


