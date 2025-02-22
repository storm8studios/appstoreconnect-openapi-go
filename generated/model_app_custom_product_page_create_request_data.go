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

// checks if the AppCustomProductPageCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageCreateRequestData{}

// AppCustomProductPageCreateRequestData struct for AppCustomProductPageCreateRequestData
type AppCustomProductPageCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppCustomProductPageCreateRequestDataAttributes `json:"attributes"`
	Relationships AppCustomProductPageCreateRequestDataRelationships `json:"relationships"`
}

// NewAppCustomProductPageCreateRequestData instantiates a new AppCustomProductPageCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageCreateRequestData(type_ string, attributes AppCustomProductPageCreateRequestDataAttributes, relationships AppCustomProductPageCreateRequestDataRelationships) *AppCustomProductPageCreateRequestData {
	this := AppCustomProductPageCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppCustomProductPageCreateRequestDataWithDefaults instantiates a new AppCustomProductPageCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageCreateRequestDataWithDefaults() *AppCustomProductPageCreateRequestData {
	this := AppCustomProductPageCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppCustomProductPageCreateRequestData) GetAttributes() AppCustomProductPageCreateRequestDataAttributes {
	if o == nil {
		var ret AppCustomProductPageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestData) GetAttributesOk() (*AppCustomProductPageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppCustomProductPageCreateRequestData) SetAttributes(v AppCustomProductPageCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppCustomProductPageCreateRequestData) GetRelationships() AppCustomProductPageCreateRequestDataRelationships {
	if o == nil {
		var ret AppCustomProductPageCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestData) GetRelationshipsOk() (*AppCustomProductPageCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppCustomProductPageCreateRequestData) SetRelationships(v AppCustomProductPageCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppCustomProductPageCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppCustomProductPageCreateRequestData struct {
	value *AppCustomProductPageCreateRequestData
	isSet bool
}

func (v NullableAppCustomProductPageCreateRequestData) Get() *AppCustomProductPageCreateRequestData {
	return v.value
}

func (v *NullableAppCustomProductPageCreateRequestData) Set(val *AppCustomProductPageCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageCreateRequestData(val *AppCustomProductPageCreateRequestData) *NullableAppCustomProductPageCreateRequestData {
	return &NullableAppCustomProductPageCreateRequestData{value: val, isSet: true}
}

func (v NullableAppCustomProductPageCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


