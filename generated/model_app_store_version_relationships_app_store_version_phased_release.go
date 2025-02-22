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

// checks if the AppStoreVersionRelationshipsAppStoreVersionPhasedRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationshipsAppStoreVersionPhasedRelease{}

// AppStoreVersionRelationshipsAppStoreVersionPhasedRelease struct for AppStoreVersionRelationshipsAppStoreVersionPhasedRelease
type AppStoreVersionRelationshipsAppStoreVersionPhasedRelease struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData `json:"data,omitempty"`
}

// NewAppStoreVersionRelationshipsAppStoreVersionPhasedRelease instantiates a new AppStoreVersionRelationshipsAppStoreVersionPhasedRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsAppStoreVersionPhasedRelease() *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease {
	this := AppStoreVersionRelationshipsAppStoreVersionPhasedRelease{}
	return &this
}

// NewAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseWithDefaults instantiates a new AppStoreVersionRelationshipsAppStoreVersionPhasedRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseWithDefaults() *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease {
	this := AppStoreVersionRelationshipsAppStoreVersionPhasedRelease{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) GetData() AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData {
	if o == nil || IsNil(o.Data) {
		var ret AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) GetDataOk() (*AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData and assigns it to the Data field.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) SetData(v AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) {
	o.Data = &v
}

func (o AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease struct {
	value *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease) Get() *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease) Set(val *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease(val *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease {
	return &NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


