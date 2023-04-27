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

// checks if the AppClipDefaultExperienceLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalizationCreateRequestDataAttributes{}

// AppClipDefaultExperienceLocalizationCreateRequestDataAttributes struct for AppClipDefaultExperienceLocalizationCreateRequestDataAttributes
type AppClipDefaultExperienceLocalizationCreateRequestDataAttributes struct {
	Locale string `json:"locale"`
	Subtitle *string `json:"subtitle,omitempty"`
}

// NewAppClipDefaultExperienceLocalizationCreateRequestDataAttributes instantiates a new AppClipDefaultExperienceLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalizationCreateRequestDataAttributes(locale string) *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes {
	this := AppClipDefaultExperienceLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewAppClipDefaultExperienceLocalizationCreateRequestDataAttributesWithDefaults instantiates a new AppClipDefaultExperienceLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationCreateRequestDataAttributesWithDefaults() *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes {
	this := AppClipDefaultExperienceLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

func (o AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locale"] = o.Locale
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes struct {
	value *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes) Get() *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes) Set(val *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes(val *AppClipDefaultExperienceLocalizationCreateRequestDataAttributes) *NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes {
	return &NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


