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

// checks if the AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion{}

// AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion struct for AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion
type AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion struct {
	Data *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData `json:"data,omitempty"`
}

// NewAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion instantiates a new AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion() *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion {
	this := AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion{}
	return &this
}

// NewAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersionWithDefaults instantiates a new AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersionWithDefaults() *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion {
	this := AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) GetData() AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	if o == nil || IsNil(o.Data) {
		var ret AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) GetDataOk() (*AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData and assigns it to the Data field.
func (o *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) SetData(v AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) {
	o.Data = &v
}

func (o AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion struct {
	value *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) Get() *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) Set(val *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion(val *AppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) *NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion {
	return &NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationInlineCreateRelationshipsAppCustomProductPageVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


