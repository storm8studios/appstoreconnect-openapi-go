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

// checks if the AppPricePointV2RelationshipsPriceTierData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPricePointV2RelationshipsPriceTierData{}

// AppPricePointV2RelationshipsPriceTierData struct for AppPricePointV2RelationshipsPriceTierData
type AppPricePointV2RelationshipsPriceTierData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppPricePointV2RelationshipsPriceTierData instantiates a new AppPricePointV2RelationshipsPriceTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointV2RelationshipsPriceTierData(type_ string, id string) *AppPricePointV2RelationshipsPriceTierData {
	this := AppPricePointV2RelationshipsPriceTierData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPricePointV2RelationshipsPriceTierDataWithDefaults instantiates a new AppPricePointV2RelationshipsPriceTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointV2RelationshipsPriceTierDataWithDefaults() *AppPricePointV2RelationshipsPriceTierData {
	this := AppPricePointV2RelationshipsPriceTierData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPricePointV2RelationshipsPriceTierData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPricePointV2RelationshipsPriceTierData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPricePointV2RelationshipsPriceTierData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPricePointV2RelationshipsPriceTierData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPricePointV2RelationshipsPriceTierData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPricePointV2RelationshipsPriceTierData) SetId(v string) {
	o.Id = v
}

func (o AppPricePointV2RelationshipsPriceTierData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPricePointV2RelationshipsPriceTierData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppPricePointV2RelationshipsPriceTierData struct {
	value *AppPricePointV2RelationshipsPriceTierData
	isSet bool
}

func (v NullableAppPricePointV2RelationshipsPriceTierData) Get() *AppPricePointV2RelationshipsPriceTierData {
	return v.value
}

func (v *NullableAppPricePointV2RelationshipsPriceTierData) Set(val *AppPricePointV2RelationshipsPriceTierData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointV2RelationshipsPriceTierData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointV2RelationshipsPriceTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointV2RelationshipsPriceTierData(val *AppPricePointV2RelationshipsPriceTierData) *NullableAppPricePointV2RelationshipsPriceTierData {
	return &NullableAppPricePointV2RelationshipsPriceTierData{value: val, isSet: true}
}

func (v NullableAppPricePointV2RelationshipsPriceTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointV2RelationshipsPriceTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


