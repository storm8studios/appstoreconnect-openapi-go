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

// checks if the AppInfoLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationCreateRequestData{}

// AppInfoLocalizationCreateRequestData struct for AppInfoLocalizationCreateRequestData
type AppInfoLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppInfoLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppInfoLocalizationCreateRequestDataRelationships `json:"relationships"`
}

// NewAppInfoLocalizationCreateRequestData instantiates a new AppInfoLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationCreateRequestData(type_ string, attributes AppInfoLocalizationCreateRequestDataAttributes, relationships AppInfoLocalizationCreateRequestDataRelationships) *AppInfoLocalizationCreateRequestData {
	this := AppInfoLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppInfoLocalizationCreateRequestDataWithDefaults instantiates a new AppInfoLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationCreateRequestDataWithDefaults() *AppInfoLocalizationCreateRequestData {
	this := AppInfoLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppInfoLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppInfoLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppInfoLocalizationCreateRequestData) GetAttributes() AppInfoLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret AppInfoLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestData) GetAttributesOk() (*AppInfoLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppInfoLocalizationCreateRequestData) SetAttributes(v AppInfoLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppInfoLocalizationCreateRequestData) GetRelationships() AppInfoLocalizationCreateRequestDataRelationships {
	if o == nil {
		var ret AppInfoLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestData) GetRelationshipsOk() (*AppInfoLocalizationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppInfoLocalizationCreateRequestData) SetRelationships(v AppInfoLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppInfoLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppInfoLocalizationCreateRequestData struct {
	value *AppInfoLocalizationCreateRequestData
	isSet bool
}

func (v NullableAppInfoLocalizationCreateRequestData) Get() *AppInfoLocalizationCreateRequestData {
	return v.value
}

func (v *NullableAppInfoLocalizationCreateRequestData) Set(val *AppInfoLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationCreateRequestData(val *AppInfoLocalizationCreateRequestData) *NullableAppInfoLocalizationCreateRequestData {
	return &NullableAppInfoLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


