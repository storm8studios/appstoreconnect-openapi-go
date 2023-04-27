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

// checks if the AppEventLocalizationRelationshipsAppEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationRelationshipsAppEventData{}

// AppEventLocalizationRelationshipsAppEventData struct for AppEventLocalizationRelationshipsAppEventData
type AppEventLocalizationRelationshipsAppEventData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppEventLocalizationRelationshipsAppEventData instantiates a new AppEventLocalizationRelationshipsAppEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationRelationshipsAppEventData(type_ string, id string) *AppEventLocalizationRelationshipsAppEventData {
	this := AppEventLocalizationRelationshipsAppEventData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEventLocalizationRelationshipsAppEventDataWithDefaults instantiates a new AppEventLocalizationRelationshipsAppEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationRelationshipsAppEventDataWithDefaults() *AppEventLocalizationRelationshipsAppEventData {
	this := AppEventLocalizationRelationshipsAppEventData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventLocalizationRelationshipsAppEventData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventLocalizationRelationshipsAppEventData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventLocalizationRelationshipsAppEventData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventLocalizationRelationshipsAppEventData) SetId(v string) {
	o.Id = v
}

func (o AppEventLocalizationRelationshipsAppEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationRelationshipsAppEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppEventLocalizationRelationshipsAppEventData struct {
	value *AppEventLocalizationRelationshipsAppEventData
	isSet bool
}

func (v NullableAppEventLocalizationRelationshipsAppEventData) Get() *AppEventLocalizationRelationshipsAppEventData {
	return v.value
}

func (v *NullableAppEventLocalizationRelationshipsAppEventData) Set(val *AppEventLocalizationRelationshipsAppEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationRelationshipsAppEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationRelationshipsAppEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationRelationshipsAppEventData(val *AppEventLocalizationRelationshipsAppEventData) *NullableAppEventLocalizationRelationshipsAppEventData {
	return &NullableAppEventLocalizationRelationshipsAppEventData{value: val, isSet: true}
}

func (v NullableAppEventLocalizationRelationshipsAppEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationRelationshipsAppEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


