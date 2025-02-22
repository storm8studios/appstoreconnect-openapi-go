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

// checks if the AppEventLocalizationRelationshipsAppEventScreenshotsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationRelationshipsAppEventScreenshotsDataInner{}

// AppEventLocalizationRelationshipsAppEventScreenshotsDataInner struct for AppEventLocalizationRelationshipsAppEventScreenshotsDataInner
type AppEventLocalizationRelationshipsAppEventScreenshotsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppEventLocalizationRelationshipsAppEventScreenshotsDataInner instantiates a new AppEventLocalizationRelationshipsAppEventScreenshotsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationRelationshipsAppEventScreenshotsDataInner(type_ string, id string) *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner {
	this := AppEventLocalizationRelationshipsAppEventScreenshotsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEventLocalizationRelationshipsAppEventScreenshotsDataInnerWithDefaults instantiates a new AppEventLocalizationRelationshipsAppEventScreenshotsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationRelationshipsAppEventScreenshotsDataInnerWithDefaults() *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner {
	this := AppEventLocalizationRelationshipsAppEventScreenshotsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner struct {
	value *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner
	isSet bool
}

func (v NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner) Get() *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner {
	return v.value
}

func (v *NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner) Set(val *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner(val *AppEventLocalizationRelationshipsAppEventScreenshotsDataInner) *NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner {
	return &NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner{value: val, isSet: true}
}

func (v NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationRelationshipsAppEventScreenshotsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


