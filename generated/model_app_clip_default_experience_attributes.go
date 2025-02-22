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

// checks if the AppClipDefaultExperienceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceAttributes{}

// AppClipDefaultExperienceAttributes struct for AppClipDefaultExperienceAttributes
type AppClipDefaultExperienceAttributes struct {
	Action *AppClipAction `json:"action,omitempty"`
}

// NewAppClipDefaultExperienceAttributes instantiates a new AppClipDefaultExperienceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceAttributes() *AppClipDefaultExperienceAttributes {
	this := AppClipDefaultExperienceAttributes{}
	return &this
}

// NewAppClipDefaultExperienceAttributesWithDefaults instantiates a new AppClipDefaultExperienceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceAttributesWithDefaults() *AppClipDefaultExperienceAttributes {
	this := AppClipDefaultExperienceAttributes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceAttributes) GetAction() AppClipAction {
	if o == nil || IsNil(o.Action) {
		var ret AppClipAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceAttributes) GetActionOk() (*AppClipAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceAttributes) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given AppClipAction and assigns it to the Action field.
func (o *AppClipDefaultExperienceAttributes) SetAction(v AppClipAction) {
	o.Action = &v
}

func (o AppClipDefaultExperienceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceAttributes struct {
	value *AppClipDefaultExperienceAttributes
	isSet bool
}

func (v NullableAppClipDefaultExperienceAttributes) Get() *AppClipDefaultExperienceAttributes {
	return v.value
}

func (v *NullableAppClipDefaultExperienceAttributes) Set(val *AppClipDefaultExperienceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceAttributes(val *AppClipDefaultExperienceAttributes) *NullableAppClipDefaultExperienceAttributes {
	return &NullableAppClipDefaultExperienceAttributes{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


