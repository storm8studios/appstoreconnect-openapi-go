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

// checks if the AppInfoLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationAttributes{}

// AppInfoLocalizationAttributes struct for AppInfoLocalizationAttributes
type AppInfoLocalizationAttributes struct {
	Locale *string `json:"locale,omitempty"`
	Name *string `json:"name,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	PrivacyPolicyUrl *string `json:"privacyPolicyUrl,omitempty"`
	PrivacyChoicesUrl *string `json:"privacyChoicesUrl,omitempty"`
	PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
}

// NewAppInfoLocalizationAttributes instantiates a new AppInfoLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationAttributes() *AppInfoLocalizationAttributes {
	this := AppInfoLocalizationAttributes{}
	return &this
}

// NewAppInfoLocalizationAttributesWithDefaults instantiates a new AppInfoLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationAttributesWithDefaults() *AppInfoLocalizationAttributes {
	this := AppInfoLocalizationAttributes{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AppInfoLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AppInfoLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AppInfoLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppInfoLocalizationAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppInfoLocalizationAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppInfoLocalizationAttributes) SetName(v string) {
	o.Name = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppInfoLocalizationAttributes) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppInfoLocalizationAttributes) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppInfoLocalizationAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *AppInfoLocalizationAttributes) GetPrivacyPolicyUrl() string {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationAttributes) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *AppInfoLocalizationAttributes) HasPrivacyPolicyUrl() bool {
	if o != nil && !IsNil(o.PrivacyPolicyUrl) {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given string and assigns it to the PrivacyPolicyUrl field.
func (o *AppInfoLocalizationAttributes) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl = &v
}

// GetPrivacyChoicesUrl returns the PrivacyChoicesUrl field value if set, zero value otherwise.
func (o *AppInfoLocalizationAttributes) GetPrivacyChoicesUrl() string {
	if o == nil || IsNil(o.PrivacyChoicesUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyChoicesUrl
}

// GetPrivacyChoicesUrlOk returns a tuple with the PrivacyChoicesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationAttributes) GetPrivacyChoicesUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyChoicesUrl) {
		return nil, false
	}
	return o.PrivacyChoicesUrl, true
}

// HasPrivacyChoicesUrl returns a boolean if a field has been set.
func (o *AppInfoLocalizationAttributes) HasPrivacyChoicesUrl() bool {
	if o != nil && !IsNil(o.PrivacyChoicesUrl) {
		return true
	}

	return false
}

// SetPrivacyChoicesUrl gets a reference to the given string and assigns it to the PrivacyChoicesUrl field.
func (o *AppInfoLocalizationAttributes) SetPrivacyChoicesUrl(v string) {
	o.PrivacyChoicesUrl = &v
}

// GetPrivacyPolicyText returns the PrivacyPolicyText field value if set, zero value otherwise.
func (o *AppInfoLocalizationAttributes) GetPrivacyPolicyText() string {
	if o == nil || IsNil(o.PrivacyPolicyText) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyText
}

// GetPrivacyPolicyTextOk returns a tuple with the PrivacyPolicyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationAttributes) GetPrivacyPolicyTextOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyText) {
		return nil, false
	}
	return o.PrivacyPolicyText, true
}

// HasPrivacyPolicyText returns a boolean if a field has been set.
func (o *AppInfoLocalizationAttributes) HasPrivacyPolicyText() bool {
	if o != nil && !IsNil(o.PrivacyPolicyText) {
		return true
	}

	return false
}

// SetPrivacyPolicyText gets a reference to the given string and assigns it to the PrivacyPolicyText field.
func (o *AppInfoLocalizationAttributes) SetPrivacyPolicyText(v string) {
	o.PrivacyPolicyText = &v
}

func (o AppInfoLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.PrivacyPolicyUrl) {
		toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl
	}
	if !IsNil(o.PrivacyChoicesUrl) {
		toSerialize["privacyChoicesUrl"] = o.PrivacyChoicesUrl
	}
	if !IsNil(o.PrivacyPolicyText) {
		toSerialize["privacyPolicyText"] = o.PrivacyPolicyText
	}
	return toSerialize, nil
}

type NullableAppInfoLocalizationAttributes struct {
	value *AppInfoLocalizationAttributes
	isSet bool
}

func (v NullableAppInfoLocalizationAttributes) Get() *AppInfoLocalizationAttributes {
	return v.value
}

func (v *NullableAppInfoLocalizationAttributes) Set(val *AppInfoLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationAttributes(val *AppInfoLocalizationAttributes) *NullableAppInfoLocalizationAttributes {
	return &NullableAppInfoLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


