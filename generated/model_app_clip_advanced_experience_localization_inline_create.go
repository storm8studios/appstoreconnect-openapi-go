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

// checks if the AppClipAdvancedExperienceLocalizationInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceLocalizationInlineCreate{}

// AppClipAdvancedExperienceLocalizationInlineCreate struct for AppClipAdvancedExperienceLocalizationInlineCreate
type AppClipAdvancedExperienceLocalizationInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
	Attributes *AppClipAdvancedExperienceLocalizationAttributes `json:"attributes,omitempty"`
}

// NewAppClipAdvancedExperienceLocalizationInlineCreate instantiates a new AppClipAdvancedExperienceLocalizationInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceLocalizationInlineCreate(type_ string) *AppClipAdvancedExperienceLocalizationInlineCreate {
	this := AppClipAdvancedExperienceLocalizationInlineCreate{}
	this.Type = type_
	return &this
}

// NewAppClipAdvancedExperienceLocalizationInlineCreateWithDefaults instantiates a new AppClipAdvancedExperienceLocalizationInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceLocalizationInlineCreateWithDefaults() *AppClipAdvancedExperienceLocalizationInlineCreate {
	this := AppClipAdvancedExperienceLocalizationInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) GetAttributes() AppClipAdvancedExperienceLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) GetAttributesOk() (*AppClipAdvancedExperienceLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceLocalizationAttributes and assigns it to the Attributes field.
func (o *AppClipAdvancedExperienceLocalizationInlineCreate) SetAttributes(v AppClipAdvancedExperienceLocalizationAttributes) {
	o.Attributes = &v
}

func (o AppClipAdvancedExperienceLocalizationInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceLocalizationInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceLocalizationInlineCreate struct {
	value *AppClipAdvancedExperienceLocalizationInlineCreate
	isSet bool
}

func (v NullableAppClipAdvancedExperienceLocalizationInlineCreate) Get() *AppClipAdvancedExperienceLocalizationInlineCreate {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceLocalizationInlineCreate) Set(val *AppClipAdvancedExperienceLocalizationInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceLocalizationInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceLocalizationInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceLocalizationInlineCreate(val *AppClipAdvancedExperienceLocalizationInlineCreate) *NullableAppClipAdvancedExperienceLocalizationInlineCreate {
	return &NullableAppClipAdvancedExperienceLocalizationInlineCreate{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceLocalizationInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceLocalizationInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


