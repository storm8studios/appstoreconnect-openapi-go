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

// checks if the AppCustomProductPageLocalizationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationCreateRequestDataRelationships{}

// AppCustomProductPageLocalizationCreateRequestDataRelationships struct for AppCustomProductPageLocalizationCreateRequestDataRelationships
type AppCustomProductPageLocalizationCreateRequestDataRelationships struct {
	AppCustomProductPageVersion AppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion `json:"appCustomProductPageVersion"`
}

// NewAppCustomProductPageLocalizationCreateRequestDataRelationships instantiates a new AppCustomProductPageLocalizationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationCreateRequestDataRelationships(appCustomProductPageVersion AppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion) *AppCustomProductPageLocalizationCreateRequestDataRelationships {
	this := AppCustomProductPageLocalizationCreateRequestDataRelationships{}
	this.AppCustomProductPageVersion = appCustomProductPageVersion
	return &this
}

// NewAppCustomProductPageLocalizationCreateRequestDataRelationshipsWithDefaults instantiates a new AppCustomProductPageLocalizationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationCreateRequestDataRelationshipsWithDefaults() *AppCustomProductPageLocalizationCreateRequestDataRelationships {
	this := AppCustomProductPageLocalizationCreateRequestDataRelationships{}
	return &this
}

// GetAppCustomProductPageVersion returns the AppCustomProductPageVersion field value
func (o *AppCustomProductPageLocalizationCreateRequestDataRelationships) GetAppCustomProductPageVersion() AppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion {
	if o == nil {
		var ret AppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion
		return ret
	}

	return o.AppCustomProductPageVersion
}

// GetAppCustomProductPageVersionOk returns a tuple with the AppCustomProductPageVersion field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationCreateRequestDataRelationships) GetAppCustomProductPageVersionOk() (*AppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppCustomProductPageVersion, true
}

// SetAppCustomProductPageVersion sets field value
func (o *AppCustomProductPageLocalizationCreateRequestDataRelationships) SetAppCustomProductPageVersion(v AppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion) {
	o.AppCustomProductPageVersion = v
}

func (o AppCustomProductPageLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appCustomProductPageVersion"] = o.AppCustomProductPageVersion
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationCreateRequestDataRelationships struct {
	value *AppCustomProductPageLocalizationCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationCreateRequestDataRelationships) Get() *AppCustomProductPageLocalizationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationCreateRequestDataRelationships) Set(val *AppCustomProductPageLocalizationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationCreateRequestDataRelationships(val *AppCustomProductPageLocalizationCreateRequestDataRelationships) *NullableAppCustomProductPageLocalizationCreateRequestDataRelationships {
	return &NullableAppCustomProductPageLocalizationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


