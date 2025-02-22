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

// checks if the AppCustomProductPageAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageAttributes{}

// AppCustomProductPageAttributes struct for AppCustomProductPageAttributes
type AppCustomProductPageAttributes struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Visible *bool `json:"visible,omitempty"`
}

// NewAppCustomProductPageAttributes instantiates a new AppCustomProductPageAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageAttributes() *AppCustomProductPageAttributes {
	this := AppCustomProductPageAttributes{}
	return &this
}

// NewAppCustomProductPageAttributesWithDefaults instantiates a new AppCustomProductPageAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageAttributesWithDefaults() *AppCustomProductPageAttributes {
	this := AppCustomProductPageAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppCustomProductPageAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppCustomProductPageAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppCustomProductPageAttributes) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AppCustomProductPageAttributes) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageAttributes) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AppCustomProductPageAttributes) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AppCustomProductPageAttributes) SetUrl(v string) {
	o.Url = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *AppCustomProductPageAttributes) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageAttributes) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *AppCustomProductPageAttributes) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *AppCustomProductPageAttributes) SetVisible(v bool) {
	o.Visible = &v
}

func (o AppCustomProductPageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageAttributes struct {
	value *AppCustomProductPageAttributes
	isSet bool
}

func (v NullableAppCustomProductPageAttributes) Get() *AppCustomProductPageAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageAttributes) Set(val *AppCustomProductPageAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageAttributes(val *AppCustomProductPageAttributes) *NullableAppCustomProductPageAttributes {
	return &NullableAppCustomProductPageAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


