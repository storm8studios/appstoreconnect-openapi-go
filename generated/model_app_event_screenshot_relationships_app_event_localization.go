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

// checks if the AppEventScreenshotRelationshipsAppEventLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotRelationshipsAppEventLocalization{}

// AppEventScreenshotRelationshipsAppEventLocalization struct for AppEventScreenshotRelationshipsAppEventLocalization
type AppEventScreenshotRelationshipsAppEventLocalization struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppEventScreenshotRelationshipsAppEventLocalizationData `json:"data,omitempty"`
}

// NewAppEventScreenshotRelationshipsAppEventLocalization instantiates a new AppEventScreenshotRelationshipsAppEventLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotRelationshipsAppEventLocalization() *AppEventScreenshotRelationshipsAppEventLocalization {
	this := AppEventScreenshotRelationshipsAppEventLocalization{}
	return &this
}

// NewAppEventScreenshotRelationshipsAppEventLocalizationWithDefaults instantiates a new AppEventScreenshotRelationshipsAppEventLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotRelationshipsAppEventLocalizationWithDefaults() *AppEventScreenshotRelationshipsAppEventLocalization {
	this := AppEventScreenshotRelationshipsAppEventLocalization{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) GetData() AppEventScreenshotRelationshipsAppEventLocalizationData {
	if o == nil || IsNil(o.Data) {
		var ret AppEventScreenshotRelationshipsAppEventLocalizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) GetDataOk() (*AppEventScreenshotRelationshipsAppEventLocalizationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppEventScreenshotRelationshipsAppEventLocalizationData and assigns it to the Data field.
func (o *AppEventScreenshotRelationshipsAppEventLocalization) SetData(v AppEventScreenshotRelationshipsAppEventLocalizationData) {
	o.Data = &v
}

func (o AppEventScreenshotRelationshipsAppEventLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotRelationshipsAppEventLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppEventScreenshotRelationshipsAppEventLocalization struct {
	value *AppEventScreenshotRelationshipsAppEventLocalization
	isSet bool
}

func (v NullableAppEventScreenshotRelationshipsAppEventLocalization) Get() *AppEventScreenshotRelationshipsAppEventLocalization {
	return v.value
}

func (v *NullableAppEventScreenshotRelationshipsAppEventLocalization) Set(val *AppEventScreenshotRelationshipsAppEventLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotRelationshipsAppEventLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotRelationshipsAppEventLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotRelationshipsAppEventLocalization(val *AppEventScreenshotRelationshipsAppEventLocalization) *NullableAppEventScreenshotRelationshipsAppEventLocalization {
	return &NullableAppEventScreenshotRelationshipsAppEventLocalization{value: val, isSet: true}
}

func (v NullableAppEventScreenshotRelationshipsAppEventLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotRelationshipsAppEventLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


