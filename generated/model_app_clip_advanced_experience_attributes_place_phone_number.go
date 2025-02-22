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

// checks if the AppClipAdvancedExperienceAttributesPlacePhoneNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceAttributesPlacePhoneNumber{}

// AppClipAdvancedExperienceAttributesPlacePhoneNumber struct for AppClipAdvancedExperienceAttributesPlacePhoneNumber
type AppClipAdvancedExperienceAttributesPlacePhoneNumber struct {
	Number *string `json:"number,omitempty"`
	Type *string `json:"type,omitempty"`
	Intent *string `json:"intent,omitempty"`
}

// NewAppClipAdvancedExperienceAttributesPlacePhoneNumber instantiates a new AppClipAdvancedExperienceAttributesPlacePhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceAttributesPlacePhoneNumber() *AppClipAdvancedExperienceAttributesPlacePhoneNumber {
	this := AppClipAdvancedExperienceAttributesPlacePhoneNumber{}
	return &this
}

// NewAppClipAdvancedExperienceAttributesPlacePhoneNumberWithDefaults instantiates a new AppClipAdvancedExperienceAttributesPlacePhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceAttributesPlacePhoneNumberWithDefaults() *AppClipAdvancedExperienceAttributesPlacePhoneNumber {
	this := AppClipAdvancedExperienceAttributesPlacePhoneNumber{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) SetNumber(v string) {
	o.Number = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) SetType(v string) {
	o.Type = &v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) GetIntent() string {
	if o == nil || IsNil(o.Intent) {
		var ret string
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) GetIntentOk() (*string, bool) {
	if o == nil || IsNil(o.Intent) {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) HasIntent() bool {
	if o != nil && !IsNil(o.Intent) {
		return true
	}

	return false
}

// SetIntent gets a reference to the given string and assigns it to the Intent field.
func (o *AppClipAdvancedExperienceAttributesPlacePhoneNumber) SetIntent(v string) {
	o.Intent = &v
}

func (o AppClipAdvancedExperienceAttributesPlacePhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceAttributesPlacePhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Intent) {
		toSerialize["intent"] = o.Intent
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber struct {
	value *AppClipAdvancedExperienceAttributesPlacePhoneNumber
	isSet bool
}

func (v NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber) Get() *AppClipAdvancedExperienceAttributesPlacePhoneNumber {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber) Set(val *AppClipAdvancedExperienceAttributesPlacePhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceAttributesPlacePhoneNumber(val *AppClipAdvancedExperienceAttributesPlacePhoneNumber) *NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber {
	return &NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceAttributesPlacePhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


