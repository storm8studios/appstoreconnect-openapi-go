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

// checks if the AppClipDefaultExperienceLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalizationAttributes{}

// AppClipDefaultExperienceLocalizationAttributes struct for AppClipDefaultExperienceLocalizationAttributes
type AppClipDefaultExperienceLocalizationAttributes struct {
	Locale *string `json:"locale,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
}

// NewAppClipDefaultExperienceLocalizationAttributes instantiates a new AppClipDefaultExperienceLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalizationAttributes() *AppClipDefaultExperienceLocalizationAttributes {
	this := AppClipDefaultExperienceLocalizationAttributes{}
	return &this
}

// NewAppClipDefaultExperienceLocalizationAttributesWithDefaults instantiates a new AppClipDefaultExperienceLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationAttributesWithDefaults() *AppClipDefaultExperienceLocalizationAttributes {
	this := AppClipDefaultExperienceLocalizationAttributes{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AppClipDefaultExperienceLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceLocalizationAttributes) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceLocalizationAttributes) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppClipDefaultExperienceLocalizationAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

func (o AppClipDefaultExperienceLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceLocalizationAttributes struct {
	value *AppClipDefaultExperienceLocalizationAttributes
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationAttributes) Get() *AppClipDefaultExperienceLocalizationAttributes {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationAttributes) Set(val *AppClipDefaultExperienceLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationAttributes(val *AppClipDefaultExperienceLocalizationAttributes) *NullableAppClipDefaultExperienceLocalizationAttributes {
	return &NullableAppClipDefaultExperienceLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


