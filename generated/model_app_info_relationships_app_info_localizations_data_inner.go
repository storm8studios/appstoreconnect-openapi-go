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

// checks if the AppInfoRelationshipsAppInfoLocalizationsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoRelationshipsAppInfoLocalizationsDataInner{}

// AppInfoRelationshipsAppInfoLocalizationsDataInner struct for AppInfoRelationshipsAppInfoLocalizationsDataInner
type AppInfoRelationshipsAppInfoLocalizationsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppInfoRelationshipsAppInfoLocalizationsDataInner instantiates a new AppInfoRelationshipsAppInfoLocalizationsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoRelationshipsAppInfoLocalizationsDataInner(type_ string, id string) *AppInfoRelationshipsAppInfoLocalizationsDataInner {
	this := AppInfoRelationshipsAppInfoLocalizationsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppInfoRelationshipsAppInfoLocalizationsDataInnerWithDefaults instantiates a new AppInfoRelationshipsAppInfoLocalizationsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoRelationshipsAppInfoLocalizationsDataInnerWithDefaults() *AppInfoRelationshipsAppInfoLocalizationsDataInner {
	this := AppInfoRelationshipsAppInfoLocalizationsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppInfoRelationshipsAppInfoLocalizationsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppInfoRelationshipsAppInfoLocalizationsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppInfoRelationshipsAppInfoLocalizationsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppInfoRelationshipsAppInfoLocalizationsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppInfoRelationshipsAppInfoLocalizationsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppInfoRelationshipsAppInfoLocalizationsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppInfoRelationshipsAppInfoLocalizationsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoRelationshipsAppInfoLocalizationsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppInfoRelationshipsAppInfoLocalizationsDataInner struct {
	value *AppInfoRelationshipsAppInfoLocalizationsDataInner
	isSet bool
}

func (v NullableAppInfoRelationshipsAppInfoLocalizationsDataInner) Get() *AppInfoRelationshipsAppInfoLocalizationsDataInner {
	return v.value
}

func (v *NullableAppInfoRelationshipsAppInfoLocalizationsDataInner) Set(val *AppInfoRelationshipsAppInfoLocalizationsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoRelationshipsAppInfoLocalizationsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoRelationshipsAppInfoLocalizationsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoRelationshipsAppInfoLocalizationsDataInner(val *AppInfoRelationshipsAppInfoLocalizationsDataInner) *NullableAppInfoRelationshipsAppInfoLocalizationsDataInner {
	return &NullableAppInfoRelationshipsAppInfoLocalizationsDataInner{value: val, isSet: true}
}

func (v NullableAppInfoRelationshipsAppInfoLocalizationsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoRelationshipsAppInfoLocalizationsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


