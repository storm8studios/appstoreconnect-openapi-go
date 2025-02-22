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

// checks if the AppStoreVersionRelationshipsAgeRatingDeclaration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationshipsAgeRatingDeclaration{}

// AppStoreVersionRelationshipsAgeRatingDeclaration struct for AppStoreVersionRelationshipsAgeRatingDeclaration
type AppStoreVersionRelationshipsAgeRatingDeclaration struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppInfoRelationshipsAgeRatingDeclarationData `json:"data,omitempty"`
}

// NewAppStoreVersionRelationshipsAgeRatingDeclaration instantiates a new AppStoreVersionRelationshipsAgeRatingDeclaration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsAgeRatingDeclaration() *AppStoreVersionRelationshipsAgeRatingDeclaration {
	this := AppStoreVersionRelationshipsAgeRatingDeclaration{}
	return &this
}

// NewAppStoreVersionRelationshipsAgeRatingDeclarationWithDefaults instantiates a new AppStoreVersionRelationshipsAgeRatingDeclaration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsAgeRatingDeclarationWithDefaults() *AppStoreVersionRelationshipsAgeRatingDeclaration {
	this := AppStoreVersionRelationshipsAgeRatingDeclaration{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) GetData() AppInfoRelationshipsAgeRatingDeclarationData {
	if o == nil || IsNil(o.Data) {
		var ret AppInfoRelationshipsAgeRatingDeclarationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) GetDataOk() (*AppInfoRelationshipsAgeRatingDeclarationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppInfoRelationshipsAgeRatingDeclarationData and assigns it to the Data field.
func (o *AppStoreVersionRelationshipsAgeRatingDeclaration) SetData(v AppInfoRelationshipsAgeRatingDeclarationData) {
	o.Data = &v
}

func (o AppStoreVersionRelationshipsAgeRatingDeclaration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationshipsAgeRatingDeclaration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionRelationshipsAgeRatingDeclaration struct {
	value *AppStoreVersionRelationshipsAgeRatingDeclaration
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsAgeRatingDeclaration) Get() *AppStoreVersionRelationshipsAgeRatingDeclaration {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsAgeRatingDeclaration) Set(val *AppStoreVersionRelationshipsAgeRatingDeclaration) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsAgeRatingDeclaration) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsAgeRatingDeclaration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsAgeRatingDeclaration(val *AppStoreVersionRelationshipsAgeRatingDeclaration) *NullableAppStoreVersionRelationshipsAgeRatingDeclaration {
	return &NullableAppStoreVersionRelationshipsAgeRatingDeclaration{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsAgeRatingDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsAgeRatingDeclaration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


