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

// checks if the AppClipAdvancedExperienceAttributesPlaceDisplayPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceAttributesPlaceDisplayPoint{}

// AppClipAdvancedExperienceAttributesPlaceDisplayPoint struct for AppClipAdvancedExperienceAttributesPlaceDisplayPoint
type AppClipAdvancedExperienceAttributesPlaceDisplayPoint struct {
	Coordinates *AppClipAdvancedExperienceAttributesPlaceDisplayPointCoordinates `json:"coordinates,omitempty"`
	Source *string `json:"source,omitempty"`
}

// NewAppClipAdvancedExperienceAttributesPlaceDisplayPoint instantiates a new AppClipAdvancedExperienceAttributesPlaceDisplayPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceAttributesPlaceDisplayPoint() *AppClipAdvancedExperienceAttributesPlaceDisplayPoint {
	this := AppClipAdvancedExperienceAttributesPlaceDisplayPoint{}
	return &this
}

// NewAppClipAdvancedExperienceAttributesPlaceDisplayPointWithDefaults instantiates a new AppClipAdvancedExperienceAttributesPlaceDisplayPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceAttributesPlaceDisplayPointWithDefaults() *AppClipAdvancedExperienceAttributesPlaceDisplayPoint {
	this := AppClipAdvancedExperienceAttributesPlaceDisplayPoint{}
	return &this
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) GetCoordinates() AppClipAdvancedExperienceAttributesPlaceDisplayPointCoordinates {
	if o == nil || IsNil(o.Coordinates) {
		var ret AppClipAdvancedExperienceAttributesPlaceDisplayPointCoordinates
		return ret
	}
	return *o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) GetCoordinatesOk() (*AppClipAdvancedExperienceAttributesPlaceDisplayPointCoordinates, bool) {
	if o == nil || IsNil(o.Coordinates) {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) HasCoordinates() bool {
	if o != nil && !IsNil(o.Coordinates) {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given AppClipAdvancedExperienceAttributesPlaceDisplayPointCoordinates and assigns it to the Coordinates field.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) SetCoordinates(v AppClipAdvancedExperienceAttributesPlaceDisplayPointCoordinates) {
	o.Coordinates = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) SetSource(v string) {
	o.Source = &v
}

func (o AppClipAdvancedExperienceAttributesPlaceDisplayPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceAttributesPlaceDisplayPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coordinates) {
		toSerialize["coordinates"] = o.Coordinates
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint struct {
	value *AppClipAdvancedExperienceAttributesPlaceDisplayPoint
	isSet bool
}

func (v NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint) Get() *AppClipAdvancedExperienceAttributesPlaceDisplayPoint {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint) Set(val *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint(val *AppClipAdvancedExperienceAttributesPlaceDisplayPoint) *NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint {
	return &NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceAttributesPlaceDisplayPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


