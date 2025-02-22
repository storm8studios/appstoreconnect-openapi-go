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

// checks if the AppPreOrderAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreOrderAttributes{}

// AppPreOrderAttributes struct for AppPreOrderAttributes
type AppPreOrderAttributes struct {
	PreOrderAvailableDate *string `json:"preOrderAvailableDate,omitempty"`
	AppReleaseDate *string `json:"appReleaseDate,omitempty"`
}

// NewAppPreOrderAttributes instantiates a new AppPreOrderAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreOrderAttributes() *AppPreOrderAttributes {
	this := AppPreOrderAttributes{}
	return &this
}

// NewAppPreOrderAttributesWithDefaults instantiates a new AppPreOrderAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreOrderAttributesWithDefaults() *AppPreOrderAttributes {
	this := AppPreOrderAttributes{}
	return &this
}

// GetPreOrderAvailableDate returns the PreOrderAvailableDate field value if set, zero value otherwise.
func (o *AppPreOrderAttributes) GetPreOrderAvailableDate() string {
	if o == nil || IsNil(o.PreOrderAvailableDate) {
		var ret string
		return ret
	}
	return *o.PreOrderAvailableDate
}

// GetPreOrderAvailableDateOk returns a tuple with the PreOrderAvailableDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrderAttributes) GetPreOrderAvailableDateOk() (*string, bool) {
	if o == nil || IsNil(o.PreOrderAvailableDate) {
		return nil, false
	}
	return o.PreOrderAvailableDate, true
}

// HasPreOrderAvailableDate returns a boolean if a field has been set.
func (o *AppPreOrderAttributes) HasPreOrderAvailableDate() bool {
	if o != nil && !IsNil(o.PreOrderAvailableDate) {
		return true
	}

	return false
}

// SetPreOrderAvailableDate gets a reference to the given string and assigns it to the PreOrderAvailableDate field.
func (o *AppPreOrderAttributes) SetPreOrderAvailableDate(v string) {
	o.PreOrderAvailableDate = &v
}

// GetAppReleaseDate returns the AppReleaseDate field value if set, zero value otherwise.
func (o *AppPreOrderAttributes) GetAppReleaseDate() string {
	if o == nil || IsNil(o.AppReleaseDate) {
		var ret string
		return ret
	}
	return *o.AppReleaseDate
}

// GetAppReleaseDateOk returns a tuple with the AppReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrderAttributes) GetAppReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.AppReleaseDate) {
		return nil, false
	}
	return o.AppReleaseDate, true
}

// HasAppReleaseDate returns a boolean if a field has been set.
func (o *AppPreOrderAttributes) HasAppReleaseDate() bool {
	if o != nil && !IsNil(o.AppReleaseDate) {
		return true
	}

	return false
}

// SetAppReleaseDate gets a reference to the given string and assigns it to the AppReleaseDate field.
func (o *AppPreOrderAttributes) SetAppReleaseDate(v string) {
	o.AppReleaseDate = &v
}

func (o AppPreOrderAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreOrderAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreOrderAvailableDate) {
		toSerialize["preOrderAvailableDate"] = o.PreOrderAvailableDate
	}
	if !IsNil(o.AppReleaseDate) {
		toSerialize["appReleaseDate"] = o.AppReleaseDate
	}
	return toSerialize, nil
}

type NullableAppPreOrderAttributes struct {
	value *AppPreOrderAttributes
	isSet bool
}

func (v NullableAppPreOrderAttributes) Get() *AppPreOrderAttributes {
	return v.value
}

func (v *NullableAppPreOrderAttributes) Set(val *AppPreOrderAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreOrderAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreOrderAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreOrderAttributes(val *AppPreOrderAttributes) *NullableAppPreOrderAttributes {
	return &NullableAppPreOrderAttributes{value: val, isSet: true}
}

func (v NullableAppPreOrderAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreOrderAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


