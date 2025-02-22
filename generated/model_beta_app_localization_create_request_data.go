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

// checks if the BetaAppLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppLocalizationCreateRequestData{}

// BetaAppLocalizationCreateRequestData struct for BetaAppLocalizationCreateRequestData
type BetaAppLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes BetaAppLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppEventCreateRequestDataRelationships `json:"relationships"`
}

// NewBetaAppLocalizationCreateRequestData instantiates a new BetaAppLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationCreateRequestData(type_ string, attributes BetaAppLocalizationCreateRequestDataAttributes, relationships AppEventCreateRequestDataRelationships) *BetaAppLocalizationCreateRequestData {
	this := BetaAppLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewBetaAppLocalizationCreateRequestDataWithDefaults instantiates a new BetaAppLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationCreateRequestDataWithDefaults() *BetaAppLocalizationCreateRequestData {
	this := BetaAppLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaAppLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaAppLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BetaAppLocalizationCreateRequestData) GetAttributes() BetaAppLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret BetaAppLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequestData) GetAttributesOk() (*BetaAppLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BetaAppLocalizationCreateRequestData) SetAttributes(v BetaAppLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *BetaAppLocalizationCreateRequestData) GetRelationships() AppEventCreateRequestDataRelationships {
	if o == nil {
		var ret AppEventCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequestData) GetRelationshipsOk() (*AppEventCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BetaAppLocalizationCreateRequestData) SetRelationships(v AppEventCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BetaAppLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableBetaAppLocalizationCreateRequestData struct {
	value *BetaAppLocalizationCreateRequestData
	isSet bool
}

func (v NullableBetaAppLocalizationCreateRequestData) Get() *BetaAppLocalizationCreateRequestData {
	return v.value
}

func (v *NullableBetaAppLocalizationCreateRequestData) Set(val *BetaAppLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationCreateRequestData(val *BetaAppLocalizationCreateRequestData) *NullableBetaAppLocalizationCreateRequestData {
	return &NullableBetaAppLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


