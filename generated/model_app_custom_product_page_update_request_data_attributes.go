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

// checks if the AppCustomProductPageUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageUpdateRequestDataAttributes{}

// AppCustomProductPageUpdateRequestDataAttributes struct for AppCustomProductPageUpdateRequestDataAttributes
type AppCustomProductPageUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	Visible *bool `json:"visible,omitempty"`
}

// NewAppCustomProductPageUpdateRequestDataAttributes instantiates a new AppCustomProductPageUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageUpdateRequestDataAttributes() *AppCustomProductPageUpdateRequestDataAttributes {
	this := AppCustomProductPageUpdateRequestDataAttributes{}
	return &this
}

// NewAppCustomProductPageUpdateRequestDataAttributesWithDefaults instantiates a new AppCustomProductPageUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageUpdateRequestDataAttributesWithDefaults() *AppCustomProductPageUpdateRequestDataAttributes {
	this := AppCustomProductPageUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppCustomProductPageUpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppCustomProductPageUpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppCustomProductPageUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *AppCustomProductPageUpdateRequestDataAttributes) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageUpdateRequestDataAttributes) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *AppCustomProductPageUpdateRequestDataAttributes) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *AppCustomProductPageUpdateRequestDataAttributes) SetVisible(v bool) {
	o.Visible = &v
}

func (o AppCustomProductPageUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageUpdateRequestDataAttributes struct {
	value *AppCustomProductPageUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppCustomProductPageUpdateRequestDataAttributes) Get() *AppCustomProductPageUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageUpdateRequestDataAttributes) Set(val *AppCustomProductPageUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageUpdateRequestDataAttributes(val *AppCustomProductPageUpdateRequestDataAttributes) *NullableAppCustomProductPageUpdateRequestDataAttributes {
	return &NullableAppCustomProductPageUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


