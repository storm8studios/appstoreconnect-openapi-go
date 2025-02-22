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

// checks if the AppCustomProductPageVersionInlineCreateRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionInlineCreateRelationships{}

// AppCustomProductPageVersionInlineCreateRelationships struct for AppCustomProductPageVersionInlineCreateRelationships
type AppCustomProductPageVersionInlineCreateRelationships struct {
	AppCustomProductPage *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage `json:"appCustomProductPage,omitempty"`
	AppCustomProductPageLocalizations *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations `json:"appCustomProductPageLocalizations,omitempty"`
}

// NewAppCustomProductPageVersionInlineCreateRelationships instantiates a new AppCustomProductPageVersionInlineCreateRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionInlineCreateRelationships() *AppCustomProductPageVersionInlineCreateRelationships {
	this := AppCustomProductPageVersionInlineCreateRelationships{}
	return &this
}

// NewAppCustomProductPageVersionInlineCreateRelationshipsWithDefaults instantiates a new AppCustomProductPageVersionInlineCreateRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionInlineCreateRelationshipsWithDefaults() *AppCustomProductPageVersionInlineCreateRelationships {
	this := AppCustomProductPageVersionInlineCreateRelationships{}
	return &this
}

// GetAppCustomProductPage returns the AppCustomProductPage field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionInlineCreateRelationships) GetAppCustomProductPage() AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage {
	if o == nil || IsNil(o.AppCustomProductPage) {
		var ret AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage
		return ret
	}
	return *o.AppCustomProductPage
}

// GetAppCustomProductPageOk returns a tuple with the AppCustomProductPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationships) GetAppCustomProductPageOk() (*AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage, bool) {
	if o == nil || IsNil(o.AppCustomProductPage) {
		return nil, false
	}
	return o.AppCustomProductPage, true
}

// HasAppCustomProductPage returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationships) HasAppCustomProductPage() bool {
	if o != nil && !IsNil(o.AppCustomProductPage) {
		return true
	}

	return false
}

// SetAppCustomProductPage gets a reference to the given AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage and assigns it to the AppCustomProductPage field.
func (o *AppCustomProductPageVersionInlineCreateRelationships) SetAppCustomProductPage(v AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) {
	o.AppCustomProductPage = &v
}

// GetAppCustomProductPageLocalizations returns the AppCustomProductPageLocalizations field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionInlineCreateRelationships) GetAppCustomProductPageLocalizations() AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations {
	if o == nil || IsNil(o.AppCustomProductPageLocalizations) {
		var ret AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations
		return ret
	}
	return *o.AppCustomProductPageLocalizations
}

// GetAppCustomProductPageLocalizationsOk returns a tuple with the AppCustomProductPageLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationships) GetAppCustomProductPageLocalizationsOk() (*AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations, bool) {
	if o == nil || IsNil(o.AppCustomProductPageLocalizations) {
		return nil, false
	}
	return o.AppCustomProductPageLocalizations, true
}

// HasAppCustomProductPageLocalizations returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationships) HasAppCustomProductPageLocalizations() bool {
	if o != nil && !IsNil(o.AppCustomProductPageLocalizations) {
		return true
	}

	return false
}

// SetAppCustomProductPageLocalizations gets a reference to the given AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations and assigns it to the AppCustomProductPageLocalizations field.
func (o *AppCustomProductPageVersionInlineCreateRelationships) SetAppCustomProductPageLocalizations(v AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) {
	o.AppCustomProductPageLocalizations = &v
}

func (o AppCustomProductPageVersionInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionInlineCreateRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppCustomProductPage) {
		toSerialize["appCustomProductPage"] = o.AppCustomProductPage
	}
	if !IsNil(o.AppCustomProductPageLocalizations) {
		toSerialize["appCustomProductPageLocalizations"] = o.AppCustomProductPageLocalizations
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionInlineCreateRelationships struct {
	value *AppCustomProductPageVersionInlineCreateRelationships
	isSet bool
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationships) Get() *AppCustomProductPageVersionInlineCreateRelationships {
	return v.value
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationships) Set(val *AppCustomProductPageVersionInlineCreateRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionInlineCreateRelationships(val *AppCustomProductPageVersionInlineCreateRelationships) *NullableAppCustomProductPageVersionInlineCreateRelationships {
	return &NullableAppCustomProductPageVersionInlineCreateRelationships{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


