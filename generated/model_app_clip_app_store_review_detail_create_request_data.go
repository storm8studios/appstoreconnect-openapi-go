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

// checks if the AppClipAppStoreReviewDetailCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailCreateRequestData{}

// AppClipAppStoreReviewDetailCreateRequestData struct for AppClipAppStoreReviewDetailCreateRequestData
type AppClipAppStoreReviewDetailCreateRequestData struct {
	Type string `json:"type"`
	Attributes *AppClipAppStoreReviewDetailAttributes `json:"attributes,omitempty"`
	Relationships AppClipAppStoreReviewDetailCreateRequestDataRelationships `json:"relationships"`
}

// NewAppClipAppStoreReviewDetailCreateRequestData instantiates a new AppClipAppStoreReviewDetailCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailCreateRequestData(type_ string, relationships AppClipAppStoreReviewDetailCreateRequestDataRelationships) *AppClipAppStoreReviewDetailCreateRequestData {
	this := AppClipAppStoreReviewDetailCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppClipAppStoreReviewDetailCreateRequestDataWithDefaults instantiates a new AppClipAppStoreReviewDetailCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailCreateRequestDataWithDefaults() *AppClipAppStoreReviewDetailCreateRequestData {
	this := AppClipAppStoreReviewDetailCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAppStoreReviewDetailCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAppStoreReviewDetailCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipAppStoreReviewDetailCreateRequestData) GetAttributes() AppClipAppStoreReviewDetailAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAppStoreReviewDetailAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailCreateRequestData) GetAttributesOk() (*AppClipAppStoreReviewDetailAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipAppStoreReviewDetailCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAppStoreReviewDetailAttributes and assigns it to the Attributes field.
func (o *AppClipAppStoreReviewDetailCreateRequestData) SetAttributes(v AppClipAppStoreReviewDetailAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value
func (o *AppClipAppStoreReviewDetailCreateRequestData) GetRelationships() AppClipAppStoreReviewDetailCreateRequestDataRelationships {
	if o == nil {
		var ret AppClipAppStoreReviewDetailCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailCreateRequestData) GetRelationshipsOk() (*AppClipAppStoreReviewDetailCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppClipAppStoreReviewDetailCreateRequestData) SetRelationships(v AppClipAppStoreReviewDetailCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppClipAppStoreReviewDetailCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailCreateRequestData struct {
	value *AppClipAppStoreReviewDetailCreateRequestData
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestData) Get() *AppClipAppStoreReviewDetailCreateRequestData {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestData) Set(val *AppClipAppStoreReviewDetailCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailCreateRequestData(val *AppClipAppStoreReviewDetailCreateRequestData) *NullableAppClipAppStoreReviewDetailCreateRequestData {
	return &NullableAppClipAppStoreReviewDetailCreateRequestData{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


