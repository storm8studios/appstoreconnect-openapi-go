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

// checks if the AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes{}

// AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes struct for AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes
type AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes struct {
	Subtitle *string `json:"subtitle,omitempty"`
}

// NewAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes instantiates a new AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes() *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes {
	this := AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewAppClipDefaultExperienceLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationUpdateRequestDataAttributesWithDefaults() *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes {
	this := AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

func (o AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes struct {
	value *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) Get() *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) Set(val *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes(val *AppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) *NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes {
	return &NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


