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

// checks if the AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner{}

// AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner struct for AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner
type AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner instantiates a new AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner(type_ string, id string) *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner {
	this := AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInnerWithDefaults instantiates a new AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInnerWithDefaults() *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner {
	this := AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner struct {
	value *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) Get() *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) Set(val *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner(val *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) *NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner {
	return &NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


